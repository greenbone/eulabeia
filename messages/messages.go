package messages

import (
	"github.com/google/uuid"
	"time"
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
