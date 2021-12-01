package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
	"github.com/rs/zerolog/log"
)

// Verify and Parse uses init event and received bytes to return bool for finished, messages.Event
// the parsed bytes or an error
type VerifyAndParse func(messages.Event, []byte, messages.Message) (bool, messages.Event, error)

// OnPrevious uses the output of a program to create a new init for the current
type OnPrevious func(messages.Event) messages.Event

// State represents the state of a received message
type State int8

const (
	None    = -1 // Not relevant for the program
	Failure = 0  // Identified as a failure message
	Success = 1  // Identified as a success message
)

// Received is used to identify downstream handler about success or failure
type Received struct {
	State State
	Event messages.Event
}

// Program is a predefined message handler
type Program struct {
	sync.RWMutex
	context           string
	init              messages.Event
	verifySuccess     VerifyAndParse
	verifyFailure     VerifyAndParse
	onPreviousSuccess OnPrevious
	onPreviousFailure OnPrevious
	success           messages.Event
	failure           messages.Event
	next              *Program
	previous          *Program
	send              chan<- *connection.SendResponse
	receive           <-chan *connection.TopicData
	received          chan<- *Received // used to inform downstream about failure or success; mostly useful on multiple success or fialure states (e.g. a scan)
	timeout           time.Duration    // Timeout of trying to receive a response; a timeout is mandataroy ohterwise it could block
	retries           int
	retryInterval     time.Duration
	finish            bool
}

func (p *Program) onMessage(td *connection.TopicData) {
	var maym messages.Message
	if err := json.Unmarshal(td.Message, &maym); err != nil {
		log.Trace().
			Err(err).
			Msgf("Skipping message (%s) on %s because it is not parseable to Message", string(td.Message), td.Topic)
		return
	}
	finish, msg, err := p.verifySuccess(p.init, td.Message, maym)
	var received *Received
	p.Lock()
	defer p.Unlock()
	if err != nil {
		if finish, msg, e := p.verifyFailure(p.init, td.Message, maym); e == nil {
			p.failure = msg
			p.finish = finish
			received = &Received{
				State: Failure,
				Event: msg,
			}
		} else {
			log.Trace().Err(e).Msgf("Ignoring message %s", maym.Type)
		}
	} else {
		p.success = msg
		p.finish = finish
		received = &Received{
			State: Success,
			Event: msg,
		}
	}

	if received != nil && p.received != nil {
		p.received <- received
	}
}

// Next adds a step to the execution chain
func (p *Program) Next(
	onFailure, onSuccess OnPrevious,
	verifyFailure, verifySuccess VerifyAndParse,
) *Program {
	np := &Program{
		context:           p.context,
		verifySuccess:     verifySuccess,
		verifyFailure:     verifyFailure,
		onPreviousSuccess: onSuccess,
		onPreviousFailure: onFailure,
		send:              p.send,
		receive:           p.receive,
		previous:          p,
		received:          p.received,
		timeout:           p.timeout,
	}
	p.next = np
	return np
}

// First finds the first step within a Program
func (p *Program) First() *Program {
	p.RLock()
	defer p.RUnlock()
	var result *Program = p
	for result.previous != nil {
		result = result.previous
	}
	return result
}

// Start identifies the First step and runs it
func (p *Program) Start() (success interface{}, failure interface{}, err error) {
	return p.First().Run()
}

// Run runs the current and next steps
func (p *Program) Run() (success interface{}, failure interface{}, err error) {
	// it is not a startpoint and depends on a previous program
	if p.init == nil {
		if p.previous == nil {
			return nil, nil, errors.New("unable to calculate init event without previous program")
		}
		if p.previous.success == nil && p.previous.failure == nil {
			return nil, nil, errors.New("previous program did not run yet")
		}
		if p.previous.failure != nil && p.onPreviousFailure != nil {
			p.init = p.onPreviousFailure(p.previous.failure)
		} else if p.previous.success != nil && p.onPreviousSuccess != nil {
			p.init = p.onPreviousSuccess(p.previous.success)
		}
	}
	if p.init == nil {
		return nil, nil, errors.New("no initial send found")
	}
	timeout := func(on string) {
		log.Info().Msgf("Timeout after %s while %s", p.timeout, on)
		err = fmt.Errorf("timeout after %s", p.timeout)
		time.Sleep(p.retryInterval)
	}
retryloop:
	for i := 0; i < p.retries+1; i++ {
		log.Trace().Msgf("[%d][%s] run %T", i, p.init.GetMessage().GroupID, p.init)
		err = nil
		p.failure = nil
		failure = nil
		p.finish = false
		select {
		case p.send <- messages.EventToResponse(p.context, p.init):
		case <-time.After(p.timeout):
			timeout("sending message")
			continue retryloop
		}
		for !p.finish {
			select {
			case td, open := <-p.receive:
				p.onMessage(td)
				if !open {
					return nil, nil, errors.New("channel in closed")
				}
			case <-time.After(p.timeout):
				timeout("receiving")
				continue retryloop
			}
		}
		success = p.success
		failure = p.failure
		if success != nil {
			break
		}
		time.Sleep(p.retryInterval)
	}

	if success == nil && err == nil {
		err = errors.New("program failed")
	}
	if p.next != nil {
		log.Trace().Msg("Running next Program")
		success, failure, err = p.next.Run()
	}
	return

}

// ModifyBasedOnGetID is used to create a modify message based on the previous message
func ModifyBasedOnGetID(aggregate string, destination string, values func(messages.GetID) map[string]interface{}) func(messages.Event) messages.Event {
	return func(e messages.Event) messages.Event {
		if v, ok := e.(messages.GetID); ok {

			return cmds.NewModify(aggregate, v.GetID(), values(v), destination, v.GetMessage().GroupID)
		}
		return nil

	}
}

// StartBasedOnGetID  used to create a start message based on the previous message
func StartBasedOnGetID(aggregate string, destination string) func(messages.Event) messages.Event {
	return func(e messages.Event) messages.Event {
		if v, ok := e.(messages.GetID); ok {
			return cmds.NewStart(aggregate, v.GetID(), destination, v.GetMessage().GroupID)
		}
		return nil
	}
}

// DeleteBasedOnGetID  used to create a delete message based on the previous message
func DeleteBasedOnGetID(aggregate string, destination string) func(messages.Event) messages.Event {
	return func(e messages.Event) messages.Event {
		if v, ok := e.(messages.GetID); ok {
			return cmds.NewDelete(aggregate, v.GetID(), destination, v.GetMessage().GroupID)
		}
		return nil
	}
}

// DefaultVerifier creates a messages.Event verifier based on given parser
func DefaultVerifier(to func([]byte) (string, messages.Event, error)) VerifyAndParse {
	return func(e messages.Event, b []byte, m messages.Message) (bool, messages.Event, error) {

		mmt := m.MessageType()
		emt := e.MessageType()

		if m.GroupID == e.GetMessage().GroupID && mmt.Aggregate == emt.Aggregate {
			ff, r, err := to(b)
			if err != nil {
				return true, nil, err
			}
			if ff == mmt.Function {
				return true, r, nil
			} else {
				return true, nil, fmt.Errorf("wrong function %s (expected %s)", mmt.Function, ff)
			}
		}
		return true, nil, errors.New("incorrect aggregate or group")
	}

}

// OpenvasScanSuccess is a success verifier for start scan message for openvas
func OpenvasScanSuccess(e messages.Event, b []byte, m messages.Message) (bool, messages.Event, error) {
	basedOn, ok := e.(messages.GetID)
	if !ok {
		return false, nil, errors.New("unable parse message to messages.GetID")
	}
	mt := m.MessageType()
	switch strings.ToLower(mt.Function) {
	case "status":
		var status info.Status
		if err := json.Unmarshal(b, &status); err != nil {
			return false, nil, err
		}
		if status.ID != basedOn.GetID() {
			return false, nil, fmt.Errorf("status ID (%s) does not match scan id (%s)", status.ID, basedOn.GetID())
		}
		switch status.Status {
		case info.REQUESTED, info.QUEUED, info.INIT, info.RUNNING, info.STOPPING:
			return false, status, nil
		case info.FINISHED, info.STOPPED:
			return true, status, nil
		default:
			return false, nil, fmt.Errorf("status (%s) is not a success case", status.Status)
		}
	case "got":
		if mt.Aggregate != "result" {
			return false, nil, fmt.Errorf("aggregate (%s) does not match result", mt.Aggregate)
		}
		var result models.GotResult
		if err := json.Unmarshal(b, &result); err != nil {
			return false, nil, err
		}
		if result.ID != basedOn.GetID() {
			return false, nil, fmt.Errorf("id (%s) does not match expected ID (%s)", result.ID, basedOn.GetID())
		}
		return false, result, nil
	default:
		return false, nil, fmt.Errorf("invalid function (%s)", mt.Function)
	}

}

// OpenvasScanFailure is a failure verifier for start scan message for openvas
func OpenvasScanFailure(e messages.Event, b []byte, m messages.Message) (bool, messages.Event, error) {
	basedOn, ok := e.(messages.GetID)
	if !ok {
		return false, nil, errors.New("unable cast message to messages.GetID")
	}
	mt := m.MessageType()
	switch strings.ToLower(mt.Function) {
	case "status":
		var status info.Status
		if err := json.Unmarshal(b, &status); err != nil {
			return false, nil, err
		}
		if status.ID != basedOn.GetID() {
			return false, nil, fmt.Errorf("status ID (%s) does not match scan id (%s)", status.ID, basedOn.GetID())
		}
		switch status.Status {
		case info.FAILED, info.INTERRUPTED, info.STOPPED:
			return true, status, nil
		default:
			return false, nil, fmt.Errorf("status (%s) is not a fail case", status.Status)
		}
	case "failure":
		var c info.Failure
		if err := json.Unmarshal(b, &c); err != nil {
			return false, nil, err
		}
		return true, c, nil

	default:
		return false, nil, fmt.Errorf("invalid function (%s)", mt.Function)
	}

}

func CreatedParser(b []byte) (string, messages.Event, error) {
	var c info.Created
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "created", c, nil
}

func ModifiedParser(b []byte) (string, messages.Event, error) {
	var c info.Modified
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "modified", c, nil
}
func DeletedParser(b []byte) (string, messages.Event, error) {
	var c info.Deleted
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "deleted", c, nil
}
func GotTargetParser(b []byte) (string, messages.Event, error) {
	var c models.GotTarget
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "got", c, nil
}
func GotScanParser(b []byte) (string, messages.Event, error) {
	var c models.GotScan
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "got", c, nil
}
func GotResultParser(b []byte) (string, messages.Event, error) {
	var c models.GotResult
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "got", c, nil
}
func GotVTParser(b []byte) (string, messages.Event, error) {
	var c models.GotVT
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "got", c, nil
}
func GotSensorParser(b []byte) (string, messages.Event, error) {
	var c models.GotSensor
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "got", c, nil
}

func FailureParser(b []byte) (string, messages.Event, error) {
	var c info.Failure
	if err := json.Unmarshal(b, &c); err != nil {
		return "", nil, err
	}
	return "failure", c, nil
}

// ToValues transforms a given interface to map[string]interface{}
//
// It is mainly used when creating modify events.
func ToValues(i interface{}) (map[string]interface{}, error) {
	var vals map[string]interface{}
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &vals); err != nil {
		return nil, err
	}
	return vals, nil
}

// Configuration is used to configure programs
type Configuration struct {
	Context       string                          // is the context of the topic
	Out           chan<- *connection.SendResponse // required channel to sendout messages
	In            <-chan *connection.TopicData    // required channel of incoming messages
	DownStream    chan<- *Received                // Optional channel to inform about identified success or failure messages
	Timeout       time.Duration                   // Optional, default 5 minutes; timeout duration for Out, In or Downstream channel operations
	Retries       int                             // Optional amount of retries on timeout or failure; can only be set on the first step
	RetryInterval time.Duration                   // Optional sleep duration before anoter try gets started
}

// From creates the first step of a Program based on given configuration and message
func From(
	c Configuration,
	msg messages.Event) (*Program, error) {
	if c.Timeout < 1 {
		log.Info().Msg("No timeout specified; setting it to 5 minutes per message")
		c.Timeout = 5 * time.Minute
	}
	if msg.GetMessage().GroupID == "" {
		return nil, errors.New("a program needs to have group id for identification of belonging messages")
	}

	var vs VerifyAndParse
	var vf VerifyAndParse
	switch v := msg.(type) {
	case cmds.Create:
		vs = DefaultVerifier(CreatedParser)
		vf = DefaultVerifier(FailureParser)
	case cmds.Get:
		var parser func([]byte) (string, messages.Event, error)
		switch v.MessageType().Aggregate {
		case "target":
			parser = GotTargetParser
		case "sensor":
			parser = GotSensorParser
		case "scan":
			parser = GotScanParser
		case "result":
			parser = GotResultParser
		case "vt":
			parser = GotVTParser
		default:
			return nil, fmt.Errorf("no known parser for %s", v.MessageType().Aggregate)
		}

		vs = DefaultVerifier(parser)
		vf = DefaultVerifier(FailureParser)
	case cmds.Delete:
		vs = DefaultVerifier(DeletedParser)
		vf = DefaultVerifier(FailureParser)
	case cmds.Modify:
		vs = DefaultVerifier(ModifiedParser)
		vf = DefaultVerifier(FailureParser)
	default:
		return nil, errors.New("unable to create Program from: %v")
	}
	return New(c, msg, vs, vf), nil
}

// Nes creates the first step of a Program based on given configuration, message and verifier
func New(c Configuration, init messages.Event, vs, vf VerifyAndParse) *Program {
	return &Program{
		context:       c.Context,
		init:          init,
		verifySuccess: vs,
		verifyFailure: vf,
		send:          c.Out,
		receive:       c.In,
		received:      c.DownStream,
		timeout:       c.Timeout,
		retries:       c.Retries,
		retryInterval: c.RetryInterval,
	}
}
