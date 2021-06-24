// Package handler implements various handler for events and messages
package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/tidwall/gjson"
)

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

// Getter is the interface that wraps the basic Get method.
//
// Gets a entity of a given type via messages.Message.MessageType and ID.
// It responds with interface{} on success and messages.Failure when not found.
type Getter interface {
	Get(messages.Get) (interface{}, *messages.Failure, error)
}

// Aggregate is the interface to handle Aggregate messages
type Aggregate interface {
	Creater
	Modifier
	Getter
}

// onAggregate is a struct containing aggregates for registered types.
//
// The messages.MessageType is normalized like what.on e.g. create.target
// onAggregate tries to parse the given messages to messages.Create,
// messages.Modify, messages.Get then tries to find via MessageType the
// Aggregate via handler.
type onAggregate struct {
	handler map[string]Aggregate
}

func (mh onAggregate) On(message []byte) (interface{}, error) {
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
	if h, ok := mh.handler[smt[1]]; ok {
		switch m := smt[0]; m {
		case "create":
			var create messages.Create
			if e := json.Unmarshal(message, &create); e != nil {
				return messages.Failure{
					Message: messages.NewMessage("failure", "", ""),
					Error:   fmt.Sprintf("unable to parse %s: %s", m, e),
				}, nil
			}
			return h.Create(create)
		case "modify":
			var modify messages.Modify
			if e := json.Unmarshal(message, &modify); e != nil {
				return messages.Failure{
					Message: messages.NewMessage("failure", "", ""),
					Error:   fmt.Sprintf("unable to parse %s: %s", m, e),
				}, nil
			}
			r, f, e := h.Modify(modify)
			if f != nil {
				return f, e
			}
			return r, e
		case "get":
			var get messages.Get
			if e := json.Unmarshal(message, &get); e != nil {
				return messages.Failure{
					Message: messages.NewMessage("failure", "", ""),
					Error:   fmt.Sprintf("unable to parse %s: %s", m, e),
				}, nil
			}
			r, f, e := h.Get(get)
			if f != nil {
				return f, e
			}
			return r, e
		default:
			return &messages.Failure{
				Message: messages.NewMessage("failure", "", ""),
				Error:   fmt.Sprintf("unable to identify method %s", m),
			}, nil
		}
	}
	return &messages.Failure{
		Message: messages.NewMessage("failure", "", ""),
		Error:   fmt.Sprintf("unable to find handler for %s", smt[1]),
	}, nil
}

func New(handler map[string]Aggregate) connection.OnMessage {
	return onAggregate{handler: handler}
}
