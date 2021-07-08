package info

import (
	"fmt"
	"github.com/greenbone/eulabeia/messages"
)

// Started is returned by a start event and contains the `id` as an identifier for the scan entity.
type Started struct {
	messages.Identifier
}

// Created is returned by a create event and contains the `id` as an identifier for the created entity.
//
// The type of entity is indicated by `message_type`.
// e.g. on "message_type": "created.target" the `id` is a identifier for a target.
type Created struct {
	messages.Identifier
}

// Failure is returned when an error occured while processing a message
type Failure struct {
	Error string `json:"error"`
	messages.Message
}

// Modified is returned by a modify event and contains the `id` as an identifier for the modified entity.
type Modified struct {
	messages.Identifier
}

// Deleted is the success response of Delete.
type Deleted struct {
	messages.Identifier
}

// DeleteFailureResponse is a conenvience method to return a Failure as Unable to delete
func DeleteFailureResponse(msg messages.Message, prefix string, id string) *Failure {
	return &Failure{
		Message: messages.NewMessage(fmt.Sprintf("failure.%s", msg.MessageType), msg.MessageID, msg.GroupID),
		Error:   fmt.Sprintf("Unable to delete %s %s.", prefix, id),
	}
}

// GetFailureResponse is a conenvience method to return a Failure as NotFound
func GetFailureResponse(msg messages.Message, prefix string, id string) *Failure {
	return &Failure{
		Message: messages.NewMessage(fmt.Sprintf("failure.%s", msg.MessageType), msg.MessageID, msg.GroupID),
		Error:   fmt.Sprintf("%s %s not found.", prefix, id),
	}
}
