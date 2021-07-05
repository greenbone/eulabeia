// Package handler implements various handler for events and messages
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
	"github.com/tidwall/gjson"
)

// Starter is the interface that wraps the basic Start method.
//
// Start is used to start an new event chain (e.g. start.scan)
type Starter interface {
	Start(messages.Start) (interface{}, *messages.Failure, error)
}

// Creater is the interface that wraps the basic Create method.
//
// Create is used on aggregate handler to handle messages.Create.
// Creates a new entity of a given type via messages.Message.MessageType.
// It responds with messages.Created which contains the id of the entity.
type Creater interface {
	Create(messages.Create) (*messages.Created, error)
}

// Modifier is the interface that wraps the basic Modify method.
//
// Modifies a entity of a given type via messages.Message.MessageType and ID.
// It responds with messages.Modified on successful alteration
// messages.Failure on incorrect Values
type Modifier interface {
	Modify(messages.Modify) (*messages.Modified, *messages.Failure, error)
}

// ModifySetValueOf is a conenvience function to set values of Modify to target
//
// Modifies a given target by trying to normalize the key of Values within Modify to
// match the naming scheme within models and then apply the given value to that.
// If it fails to apply the given value directly it calls the given function
// apply to try it via own handling mechanismn. If apply is nil or apply fails as well
// an messages.Failure is returned.
func ModifySetValueOf(target interface{},
	m messages.Modify,
	apply func(string, interface{}) error) *messages.Failure {
	for k, v := range m.Values {
		// normalize field name
		nk := strings.Title(k)
		var failure error
		// due to map[string]interface{} []string can be detected as []interface{} instead
		switch cv := v.(type) {
		case []interface{}:
			strings := make([]string, len(cv), cap(cv))
			for i, j := range cv {
				if s, ok := j.(string); ok {
					strings[i] = s
				}
			}
			failure = models.SetValueOf(target, nk, strings)
		default:
			failure = models.SetValueOf(target, nk, cv)
		}
		if failure != nil && apply != nil {
			failure = apply(k, v)
		}
		if failure != nil {
			log.Printf("Failure while processing field %v: %v", nk, failure)
			return &messages.Failure{
				Error:   fmt.Sprintf("Unable to set %s on target to %s: %v", nk, v, failure),
				Message: messages.NewMessage("failure."+m.MessageType, m.MessageID, m.GroupID),
			}
		}
	}
	return nil
}

// Getter is the interface that wraps the basic Get method.
//
// Gets a entity of a given type via messages.Message.MessageType and ID.
// It responds with interface{} on success and messages.Failure when not found.
type Getter interface {
	Get(messages.Get) (interface{}, *messages.Failure, error)
}

type Deleter interface {
	Delete(messages.Delete) (*messages.Deleted, *messages.Failure, error)
}

// Aggregate is the interface to handle Aggregate messages
type Aggregate interface {
	Creater
	Modifier
	Deleter
	Getter
}

// onMessage is a struct containing aggregates for registered types.
//
// The messages.MessageType is normalized like what.on e.g. create.target
// onMessage tries to parse the given messages to messages.Create,
// messages.Modify, messages.Get then tries to find via MessageType the
// Aggregate via handler.
type onMessage struct {
	lookup map[string]Holder
}

func asResponse(t string, d interface{}) *connection.SendResponse {
	return &connection.SendResponse{
		Topic: t,
		MSG:   d,
	}
}

func (mh onMessage) On(topic string, message []byte) (*connection.SendResponse, error) {
	messageType := gjson.GetBytes(message, "message_type")
	if messageType.Type == gjson.Null {
		return asResponse(topic, messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   "unable to find message_type",
		}), nil
	}
	smt := strings.Split(messageType.String(), ".")
	if len(smt) < 2 {
		return asResponse(topic, messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   fmt.Sprintf("incorrect message_type %s", messageType.String()),
		}), nil
	}
	if h, ok := mh.lookup[smt[1]]; ok {
		var del messages.Delete
		var create messages.Create
		var modify messages.Modify
		var get messages.Get
		var start messages.Start
		var use interface{}
		var fuse func() (interface{}, *messages.Failure, error)
		if m := smt[0]; m == "delete" && h.Deleter != nil {
			use = &del
			fuse = func() (interface{}, *messages.Failure, error) {
				return h.Deleter.Delete(del)
			}

		} else if m == "create" && h.Creater != nil {
			use = &create
			fuse = func() (interface{}, *messages.Failure, error) {
				r, e := h.Creater.Create(create)
				return r, nil, e
			}
		} else if m == "start" && h.Starter != nil {
			use = &start
			fuse = func() (interface{}, *messages.Failure, error) {
				return h.Starter.Start(start)
			}
		} else if m == "modify" && h.Modifier != nil {
			use = &modify
			fuse = func() (interface{}, *messages.Failure, error) {
				return h.Modifier.Modify(modify)
			}
		} else if m == "get" && h.Getter != nil {
			use = &get
			fuse = func() (interface{}, *messages.Failure, error) {
				return h.Getter.Get(get)
			}
		} else {
			log.Printf("unable to identify method %s", m)
			return nil, nil
		}
		if e := json.Unmarshal(message, use); e != nil {
			return asResponse(topic, messages.Failure{
				Message: messages.NewMessage("failure", "", ""),
				Error:   fmt.Sprintf("unable to parse %s: %s", smt[0], e),
			}), nil
		}
		r, f, e := fuse()
		if f != nil {
			return asResponse(topic, f), e
		}
		if r, ok := r.(*connection.SendResponse); ok {
			return r, e
		}
		return asResponse(topic, r), e

	}
	log.Printf("unable to identify entity %s", smt[1])
	return nil, nil
}

// Holder contains interfaces needed for OnMessage
type Holder struct {
	Topic    string
	Creater  Creater
	Modifier Modifier
	Getter   Getter
	Deleter  Deleter
	Starter  Starter
}

// FromAggregate is a convencience method to create specialized lookup maps for connection.OnMessage
func FromAggregate(topic string, a Aggregate) Holder {
	return Holder{
		Topic:    topic,
		Creater:  a,
		Modifier: a,
		Getter:   a,
		Deleter:  a,
	}
}

// FromGetter is a convencience method to create specialized lookup maps for connection.OnMessage
func FromGetter(topic string, a Getter) Holder {
	return Holder{
		Topic:  topic,
		Getter: a,
	}
}

// New returns a new connection.OnMessage handler
func New(holder ...Holder) connection.OnMessage {
	lookup := map[string]Holder{}

	for _, h := range holder {
		lookup[h.Topic] = h
	}
	return onMessage{
		lookup: lookup,
	}
}
