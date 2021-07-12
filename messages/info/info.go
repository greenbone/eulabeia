package info

import (
	"fmt"
	"github.com/greenbone/eulabeia/messages"
)

type EventType struct{}

func (EventType) Event() messages.EventType {
	return messages.INFO
}

// IDInfo is an Info message with just an entity ID
type IDInfo struct {
	EventType
	messages.Identifier
}

// Started is the success response of a cmd.Start
type Started IDInfo

// Created is the success response of a cmd.Create
type Created IDInfo

// Modified is the success response of a cmd.Modify
type Modified IDInfo

// Deleted is the success response of a cmd.Delete
type Deleted IDInfo

// Failure is returned when an error occured while processing a message
type Failure struct {
	EventType
	messages.Message
	Error string `json:"error"`
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
