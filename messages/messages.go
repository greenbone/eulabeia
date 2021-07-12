package messages

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// EventType is used to identify a message
type EventType string

const (
	CMD  EventType = "cmd"  // Event is a cmd
	INFO EventType = "info" // Event is a info
)

type Event interface {
	Event() EventType
	MessageType() MessageType
}

// Message contains the meta data for each sent message.
// It should be embedded into all messages send to or received by eulabeia.
type Message struct {
	Created   int    `json:"created"`      // Timestamp when this message was created
	Type      string `json:"message_type"` // Identifier what this message actually contains
	MessageID string `json:"message_id"`   // The ID of a message, responses will have the same ID
	GroupID   string `json:"group_id"`     // The ID of a group of messages, responses will have the same ID
}

func (m Message) MessageType() MessageType {
	result, err := ParseMessageType(m.Type)
	if err != nil {
		panic(fmt.Errorf("unable to parse MessageType: %s", err))
	}
	return *result
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

func (m MessageType) String() string {
	result := fmt.Sprintf("%s.%s", m.Function, m.Aggregate)
	if m.Destination != "" {
		result = fmt.Sprintf("%s.%s", result, m.Destination)
	}
	return result
}

func ParseMessageType(typ string) (*MessageType, error) {
	smt := strings.Split(typ, ".")
	if len(smt) < 1 {
		return nil, fmt.Errorf("unable to parse %s to MessageType", typ)
	}
	result := MessageType{
		Function: smt[0],
	}
	if len(smt) > 1 {
		result.Aggregate = smt[1]
	}
	if len(smt) > 2 {
		result.Destination = smt[2]
	}
	return &result, nil
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
		Created:   time.Now().Nanosecond(),
		Type:      messageType,
		MessageID: messageID,
		GroupID:   groupID,
	}
}

// Command is used by the director to run a command on a sensor. Possible
// commands are:
//  - start
//  - stop
//  - version
//  - loadvts
type Command struct {
	ID  string `json:"id"`
	Cmd string `json:"cmd"`
	Message
}

type ScanInfo struct {
	ID       string `json:"id"`
	InfoType string `json:"type"`
	Info     string `json:"info"`
	Message
}
