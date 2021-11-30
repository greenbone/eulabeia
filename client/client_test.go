package client

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection"
	_ "github.com/greenbone/eulabeia/logging/configuration"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
	"github.com/rs/zerolog/log"
)

var NO_VALUES = func(gi messages.GetID) map[string]interface{} { return nil }
var MODIFY_SCAN_VALUES = func(gi messages.GetID) map[string]interface{} { return map[string]interface{}{"target_id": gi.GetID()} }

func createDelegationFake(gid string, response ...messages.Event) (chan *connection.SendResponse, chan *connection.TopicData) {
	out := make(chan *connection.SendResponse, 1)
	in := make(chan *connection.TopicData, 1)
	go func(out <-chan *connection.SendResponse, in chan<- *connection.TopicData, gid string, respond []messages.Event) {
		i := 0
		for sr := range out {
			log.Trace().Msgf("Got message %+v", sr)
			switch respond[i].(type) {
			case info.Status:
				for ; i < len(response); i++ {
					b, _ := json.Marshal(&respond[i])
					log.Trace().Msgf("[%d] sending %s", i, string(b))
					in <- &connection.TopicData{Topic: "dontcare", Message: b}
				}
			default:
				b, _ := json.Marshal(&respond[i])
				log.Trace().Msgf("[%d] sending %s", i, string(b))
				in <- &connection.TopicData{Topic: "dontcare", Message: b}
				i = i + 1
			}
		}
		log.Trace().Msg("BYE BYE BYE BYE BYE")

	}(out, in, gid, response)
	return out, in
}

func TestCreateScanProgram(t *testing.T) {
	created := info.Created{
		Identifier: messages.Identifier{
			ID:      uuid.NewString(),
			Message: messages.NewMessage("created.scan", "", "tcsp"),
		},
	}
	out, in := createDelegationFake(created.GroupID, created)
	defer close(in)
	defer close(out)
	c := Configuration{
		Context: "scanner",
		Out:     out,
		In:      in,
	}

	p, err := From(c, cmds.NewCreate("scan", "director", created.GroupID))
	if err != nil {
		t.Fatalf("Expected to create program: %s", err)

	}
	s, f, err := p.Run()
	if err != nil {
		t.Fatalf("Expected success but got err: %s (%+v)", err, f)
	}
	switch v := s.(type) {
	case info.Created:
	default:
		t.Errorf("Expected info.Created as response but got %T", v)
	}
}

func TestFailCreateScanProgram(t *testing.T) {
	failed := info.FailureResponse(
		messages.NewMessage("create.scan", "", "gid"),
		"1",
		"Unable to create 1",
	)
	out, in := createDelegationFake(failed.GroupID, failed)
	defer close(in)
	defer close(out)
	c := Configuration{
		Context: "scanner",
		Out:     out,
		In:      in,
	}
	p, err := From(c, cmds.NewCreate("scan", "director", failed.GroupID))
	if err != nil {
		t.Fatalf("Expected to create program: %s", err)

	}
	s, f, err := p.Run()
	if err == nil {
		t.Fatalf("Expected err but got none")
	}
	if s != nil {
		t.Fatalf("Expected success to be nil but it is %+v", s)
	}
	switch v := f.(type) {
	case info.Failure:
	default:
		t.Errorf("Expected info.Created as response but got %T", v)
	}
}

func TestCreateThenModifyScan(t *testing.T) {
	created := info.Created{
		Identifier: messages.Identifier{
			ID:      uuid.NewString(),
			Message: messages.NewMessage("created.scan", "", "tcsp"),
		},
	}
	modified := info.Modified{
		Identifier: messages.Identifier{
			ID:      "test",
			Message: messages.NewMessage("modified.scan", "", created.GroupID),
		},
	}
	out, in := createDelegationFake(created.GroupID, created, modified)
	defer close(in)
	defer close(out)
	create := cmds.NewCreate("scan", "director", created.GroupID)
	c := Configuration{
		Context: "scanner",
		Out:     out,
		In:      in,
	}
	p, _ := From(c, create)
	tm := ModifyBasedOnGetID("scan", "director", NO_VALUES)
	p.Next(nil, tm, DefaultVerifier(FailureParser), DefaultVerifier(ModifiedParser))
	s, f, err := p.Start()
	if err != nil {
		t.Fatalf("Expected success but got err: %s (%+v)", err, f)
	}
	switch v := s.(type) {
	case info.Modified:
	default:
		t.Errorf("Expected info.Modified as response but got %T", v)
	}
}

func TestCreateTargetModifyTargetCreateScanModifyScanStartScan(t *testing.T) {
	groupID := "scantest"
	sID := "scanme"
	cstatus := func(status string) info.Status {
		return info.Status{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("status.scan", "", groupID),
			},
			Status: status,
		}
	}
	responses := []messages.Event{
		info.Created{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("created.target", "", groupID),
			},
		},
		info.Modified{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("modified.target", "", groupID),
			},
		},
		info.Modified{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("modified.scan", "", groupID),
			},
		},
		cstatus(info.REQUESTED),
		cstatus(info.QUEUED),
		cstatus(info.INIT),
		cstatus(info.RUNNING),
		models.GotResult{
			Message: messages.NewMessage("got.result", "", groupID),
			Result: models.Result{
				ID:    sID,
				OID:   "0.0.0.0.0.0.0.0.0.0.0.1",
				Type:  models.LOG,
				Value: "something",
			},
		},
		cstatus(info.FINISHED),
	}
	create := cmds.NewCreate("target", "director", groupID)
	cmds := []func(messages.Event) messages.Event{
		ModifyBasedOnGetID("target", "director", NO_VALUES),
		// we reuse the target id as a scan id
		ModifyBasedOnGetID("scan", "director", MODIFY_SCAN_VALUES),
		StartBasedOnGetID("scan", "director"),
	}
	s_verification := []VerifyAndParse{
		DefaultVerifier(ModifiedParser),
		DefaultVerifier(ModifiedParser),
		OpenvasScanSuccess,
	}
	f_verification := []VerifyAndParse{
		DefaultVerifier(FailureParser),
		DefaultVerifier(FailureParser),
		OpenvasScanFailure,
	}

	out, in := createDelegationFake(groupID, responses...)
	defer close(in)
	defer close(out)
	downstream := make(chan *Received, len(responses))

	c := Configuration{
		Context:    "scanner",
		DownStream: downstream,
		Out:        out,
		In:         in,
	}
	p, _ := From(c, create)
	for i, sv := range s_verification {
		fv := f_verification[i]
		op := cmds[i]
		p = p.Next(nil, op, fv, sv)
	}
	s, f, e := p.Start()
	if e != nil || f != nil {
		t.Fatalf("Expected success but got error (%s) or failure (%+v)", e, f)
	}
	countds := 0
	for countds < len(responses) {
		select {
		case <-downstream:
			countds = countds + 1
		case <-time.After(1 * time.Minute):
			log.Fatal().Msg("Timeout after a minute.")
		}
	}
	if countds != len(responses) {
		log.Fatal().Msgf("expected countds (%d) to be %d", countds, len(responses))
	}
	switch v := s.(type) {
	case info.Status:
		if v.Status != info.FINISHED {
			log.Fatal().Msgf("Expected status (%s) to be %s", v.Status, info.FINISHED)
		}
	default:
		t.Errorf("Expected info.Status as last response but got %T", v)
	}
}

func TestRetry(t *testing.T) {
	groupID := "retry_test"
	sID := "test"
	responses := []messages.Event{
		info.Failure{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("failure.create.target", "", groupID),
			},
		},
		info.Failure{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("failure.create.target", "", groupID),
			},
		},
		info.Failure{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("failure.create.target", "", groupID),
			},
		},
		info.Created{
			Identifier: messages.Identifier{
				ID:      sID,
				Message: messages.NewMessage("created.target", "", groupID),
			},
		},
	}
	out, in := createDelegationFake(groupID, responses...)
	defer close(in)
	defer close(out)
	downstream := make(chan *Received, len(responses))

	c := Configuration{
		Context:    "scanner",
		DownStream: downstream,
		Out:        out,
		In:         in,
		Retries:    len(responses),
		Timeout:    30 * time.Second,
	}
	create := cmds.NewCreate("target", "", groupID)
	p, _ := From(c, create)
	s, f, e := p.Run()
	if e != nil || f != nil {
		t.Fatalf("Expected success but got error (%s) or failure (%+v)", e, f)
	}
	countds := 0
	for countds < len(responses) {
		select {
		case <-downstream:
			countds = countds + 1
		case <-time.After(10 * time.Second):
			log.Fatal().Msgf("Timeout after a minute; counted: %d.", countds)
		}
	}
	switch v := s.(type) {
	case info.Created:
	default:
		t.Errorf("Expected info.Status as last response but got %T", v)
	}
}

func TestTimeoutOnNoResponse(t *testing.T) {
	groupID := "groupie"
	out, in := createDelegationFake(groupID)
	defer close(in)
	defer close(out)
	c := Configuration{
		Out:     out,
		In:      in,
		Timeout: 2,
	}
	p, _ := From(c, cmds.NewCreate("scan", "", ""))
	_, _, err := p.Run()
	if err == nil {
		log.Fatal().Msg("expected timeout error but got nil")
	}
}
