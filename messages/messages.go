// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package messages

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection"
)

// EventType is used to identify a message
type EventType string

const (
	CMD  EventType = "cmd"  // Event is a cmd
	INFO EventType = "info" // Event is a info
)

type GetMessage interface {
	GetMessage() Message
}

type Event interface {
	GetMessage
	Event() EventType
	MessageType() MessageType
}

// EventToResponse is using given context and Event to calculate the topic
func EventToResponse(context string, e Event) *connection.SendResponse {
	if e == nil {
		return nil
	}
	mt := e.MessageType()
	topic := fmt.Sprintf("%s/%s/%s", context, mt.Aggregate, e.Event())
	if mt.Destination != "" {
		topic = fmt.Sprintf("%s/%s", topic, mt.Destination)
	}
	return &connection.SendResponse{
		Topic: topic,
		MSG:   e,
	}
}

// Message contains the meta data for each sent message.
// It should be embedded into all messages send to or received by eulabeia.
type Message struct {
	Created   int    `json:"message_created"` // Timestamp when this message was created
	Type      string `json:"message_type"`    // Identifier what this message actually contains
	MessageID string `json:"message_id"`      // The ID of a message, responses will have the same ID
	GroupID   string `json:"group_id"`        // The ID of a group of messages, responses will have the same ID
}

func (m Message) GetMessage() Message {
	return m
}

func (m Message) MessageType() MessageType {
	result, err := ParseMessageType(m.Type)
	if err != nil {
		panic(fmt.Errorf("unable to parse MessageType: %s", err))
	}
	return *result
}

type GetID interface {
	Event
	GetID() string
}

// Identifier is an ID based cmd it contains an ID for
// messages.Message.MessageType
type Identifier struct {
	ID string `json:"id"`
	Message
}

func (s Identifier) GetID() string {
	return s.ID
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

func ParseMessageType(mt string) (*MessageType, error) {
	smt := strings.Split(mt, ".")
	if len(smt) < 1 {
		return nil, fmt.Errorf("unable to parse %s to MessageType", mt)
	}
	s := 0
	if smt[0] == "failure" {
		s = 1
	}
	result := MessageType{
		Function: smt[0],
	}
	if len(smt) > s+1 {
		result.Aggregate = smt[s+1]
	}
	if len(smt) > s+2 {
		result.Destination = smt[s+2]
	}
	return &result, nil
}

// NewMessage creates a new message; if messageID oder groupID are empty a new
// uuid will be used instead.
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
