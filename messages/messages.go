package messages

import (
	"github.com/google/uuid"
	"time"
)

// Message contains the meta data for each sent message.
// It should be embedded into all messages send to or received by eulabia.
type Message struct {
	Created     int    `json:"created"`      // Timestamp when this message was created
	MessageType string `json:"message_type"` // Identifier what this message actually contains
	MessageID   string `json:"message_id"`   // The ID of a message, responses will have the same ID
	GroupID     string `json:"group_id"`     // The ID of a group of messages, responses will have the same ID
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

// Create indicates that a new entity should be created.
// The type of of entity is indicated by `message_type`
// e.g. "message_type": "create.target" creates a target.
type Create struct {
	Message
}

// Created is returned by a create event and contains the `id` as an identifier for the created entity.
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

/*
Modify indicates that a entity should be modified.
The type of of entity is indicated by `message_type`
e.g. "message_type": "modify.target" modifies a target.
Values contains the fields as well as the data to be set; e.g.:

will override the fields:
- hosts
- plugins
*/
type Modify struct {
	Message
	ID     string                 `json:"id"`
	Values map[string]interface{} `json:"values"`
}

// Modified is returned by a modify event and contains the `id` as an identifier for the modified entity.
// The type of entity is indicated by `message_type`.
type Modified struct {
	ID string `json:"id"`
	Message
}

// Get is used by a client to get the latest snapshort of an aggregate.
// The response for Get is usually the aggragte with Message information and can be found within a model.
type Get struct {
	ID string `json:"id"`
	Message
}
