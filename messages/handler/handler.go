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
	// Start should send a get.memory for a sensor and react on
	// got.memory of given sensor; when sufficient memory is
	// available then it should send start scan event to
	// sensor and return started event
	Start(messages.Start) (*messages.Started, *messages.Failure, error)
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
	creater  map[string]Creater
	modifier map[string]Modifier
	getter   map[string]Getter
	deleter  map[string]Deleter
}

// enhance for topic information
func (mh onMessage) On(message []byte) (interface{}, error) {
	messageType := gjson.GetBytes(message, "message_type")
	if messageType.Type == gjson.Null {
		return messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   "unable to find message_type",
		}, nil
	}
	smt := strings.Split(messageType.String(), ".")
	if len(smt) < 2 {
		return messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   fmt.Sprintf("incorrect message_type %s", messageType.String()),
		}, nil
	}
	switch m := smt[0]; m {
	case "delete":
		var del messages.Delete
		if e := json.Unmarshal(message, &del); e != nil {
			return messages.Failure{
				Message: messages.NewMessage("failure", "", ""),
				Error:   fmt.Sprintf("unable to parse %s: %s", m, e),
			}, nil
		}
		if h, ok := mh.deleter[smt[1]]; ok {
			r, f, e := h.Delete(del)
			if f != nil {
				return f, e
			}
			return r, e
		}
		return &messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   fmt.Sprintf("unable to find handler for %s", smt[1]),
		}, nil
	case "create":
		var create messages.Create
		if e := json.Unmarshal(message, &create); e != nil {
			return messages.Failure{
				Message: messages.NewMessage("failure", "", ""),
				Error:   fmt.Sprintf("unable to parse %s: %s", m, e),
			}, nil
		}
		if h, ok := mh.creater[smt[1]]; ok {
			return h.Create(create)
		}
		return &messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   fmt.Sprintf("unable to find handler for %s", smt[1]),
		}, nil
	case "modify":
		var modify messages.Modify
		if e := json.Unmarshal(message, &modify); e != nil {
			return messages.Failure{
				Message: messages.NewMessage("failure", "", ""),
				Error:   fmt.Sprintf("unable to parse %s: %s", m, e),
			}, nil
		}

		if h, ok := mh.modifier[smt[1]]; ok {
			r, f, e := h.Modify(modify)
			if f != nil {
				return f, e
			}
			return r, e
		}
		return &messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   fmt.Sprintf("unable to find handler for %s", smt[1]),
		}, nil
	case "get":
		var get messages.Get
		if e := json.Unmarshal(message, &get); e != nil {
			return messages.Failure{
				Message: messages.NewMessage("failure", "", ""),
				Error:   fmt.Sprintf("unable to parse %s: %s", m, e),
			}, nil
		}
		if h, ok := mh.getter[smt[1]]; ok {
			r, f, e := h.Get(get)
			if f != nil {
				return f, e
			}
			return r, e
		}
		return &messages.Failure{
			Message: messages.NewMessage("failure", "", ""),
			Error:   fmt.Sprintf("unable to find handler for %s", smt[1]),
		}, nil
	default:
		log.Printf("unable to identify method %s", m)
		return nil, nil
	}
}

// Holder contains interfaces needed for OnMessage
type Holder struct {
	Topic    string
	Creater  Creater
	Modifier Modifier
	Getter   Getter
	Deleter  Deleter
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
	creater := map[string]Creater{}
	modifier := map[string]Modifier{}
	getter := map[string]Getter{}
	deleter := map[string]Deleter{}

	for _, h := range holder {
		if h.Creater != nil {
			creater[h.Topic] = h.Creater
		}
		if h.Modifier != nil {
			modifier[h.Topic] = h.Modifier
		}
		if h.Getter != nil {
			getter[h.Topic] = h.Getter
		}
		if h.Deleter != nil {
			deleter[h.Topic] = h.Deleter
		}
	}
	return onMessage{
		creater:  creater,
		modifier: modifier,
		getter:   getter,
		deleter:  deleter,
	}
}
