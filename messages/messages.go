package messages

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Message contains the meta data for each sent message.
// It should be embedded into all messages send to or received by eulabeia.
type Message struct {
	Created     int    `json:"created"`      // Timestamp when this message was created
	MessageType string `json:"message_type"` // Identifier what this message actually contains
	MessageID   string `json:"message_id"`   // The ID of a message, responses will have the same ID
	GroupID     string `json:"group_id"`     // The ID of a group of messages, responses will have the same ID
}

// Identifier is an ID based cmd it contains an ID for messages.Message.MessageType
type Identifier struct {
	ID string `json:"id"`
	Message
}

type MessageType struct {
	Function    string // Function indicates if it is a cmd or info (e.g. create, created)
	Aggregate   string // Aggregate defines to which aggregate this message belonds to (e.g. target)
	Destination string // Destination is an optinal parameter to indicate if this message is deicated for a special consumer
}

// NewMessage creates a new message; if messageID oder groupID are empty a new uuid will be used instead.
func NewMessage(messageType string, messageID string, groupID string) Message {
	if messageID == "" {
		messageID = uuid.NewString()
	}
	if groupID == "" {
		groupID = uuid.NewString()
	}
	return Message{
		Created:     time.Now().Nanosecond(),
		MessageType: messageType,
		MessageID:   messageID,
		GroupID:     groupID,
	}
}

// DeleteFailureResponse is a conenvience method to return a Failure as Unable to delete
func DeleteFailureResponse(msg Message, prefix string, id string) *Failure {
	return &Failure{
		Message: NewMessage(fmt.Sprintf("failure.%s", msg.MessageType), msg.MessageID, msg.GroupID),
		Error:   fmt.Sprintf("Unable to delete %s %s.", prefix, id),
	}
}

// GetFailureResponse is a conenvience method to return a Failure as NotFound
func GetFailureResponse(msg Message, prefix string, id string) *Failure {
	return &Failure{
		Message: NewMessage(fmt.Sprintf("failure.%s", msg.MessageType), msg.MessageID, msg.GroupID),
		Error:   fmt.Sprintf("%s %s not found.", prefix, id),
	}
}

// Started is returned by a start event and contains the `id` as an identifier for the scan entity.
type Started struct {
	ID string `json:"id"`
	Message
}

// Created is returned by a create event and contains the `id` as an identifier for the created entity.
//
// The type of entity is indicated by `message_type`.
// e.g. on "message_type": "created.target" the `id` is a identifier for a target.
type Created struct {
	ID string `json:"id"`
	Message
}

// Failure is returned when an error occured while processing a message
type Failure struct {
	Error string `json:"error"`
	Message
}

// Modified is returned by a modify event and contains the `id` as an identifier for the modified entity.
// The type of entity is indicated by `message_type`.
type Modified struct {
	ID string `json:"id"`
	Message
}

// Deleted is the success response of Delete.
type Deleted struct {
	ID string `json:"id"`
	Message
}
