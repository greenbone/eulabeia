package info

import (
	"fmt"

	"github.com/greenbone/eulabeia/messages"
)

type EventType struct{}

func (EventType) Event() messages.EventType {
	return messages.INFO
}

// Started is returned by a start event and contains the `id` as an identifier for the scan entity.
type Started struct {
	messages.Identifier
}

// Created is returned by a create event and contains the `id` as an identifier for the created entity.
//
// The type of entity is indicated by `message_type`.
// e.g. on "message_type": "created.target" the `id` is a identifier for a target.
type Created struct {
	EventType
	messages.Identifier
}

// Failure is returned when an error occured while processing a message
type Failure struct {
	EventType
	messages.Message
	Error string `json:"error"`
}

// Modified is returned by a modify event and contains the `id` as an identifier for the modified entity.
type Modified struct {
	EventType
	messages.Identifier
}

// Deleted is the success response of Delete.
type Deleted struct {
	EventType
	messages.Identifier
}

// Contains the status of a scan
type Status struct {
	EventType
	messages.Identifier
	Status string `json:"status"`
}

type Response struct {
	EventType
	messages.Message
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type Version struct {
	EventType
	messages.Identifier
	Version string `json:"version"`
}

// DeleteFailureResponse is a conenvience method to return a Failure as Unable to delete
func DeleteFailureResponse(msg messages.Message, prefix string, id string) *Failure {
	return &Failure{
		Message: messages.NewMessage(fmt.Sprintf("failure.%s", msg.Type), msg.MessageID, msg.GroupID),
		Error:   fmt.Sprintf("Unable to delete %s %s.", prefix, id),
	}
}

// GetFailureResponse is a conenvience method to return a Failure as NotFound
func GetFailureResponse(msg messages.Message, prefix string, id string) *Failure {
	return &Failure{
		Message: messages.NewMessage(fmt.Sprintf("failure.%s", msg.Type), msg.MessageID, msg.GroupID),
		Error:   fmt.Sprintf("%s %s not found.", prefix, id),
	}
}
