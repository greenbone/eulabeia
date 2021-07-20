// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmds

import (
	"fmt"

	"github.com/greenbone/eulabeia/messages"
)

type eventType struct{}

func (eventType) Event() messages.EventType {
	return messages.CMD
}

// IDCMD is a command with just an ID to identify a specific entity
type IDCMD struct {
	eventType
	messages.Identifier
}

// Create indicates that a new entity should be created.
// The type of of entity is indicated by `message_type`
// e.g. "message_type": "create.target" creates a target.
type Create struct {
	eventType
	messages.Message
}

func buildMessageType(function string, aggregate string, destination string) string {
	result := fmt.Sprintf("%s.%s", function, aggregate)
	if destination != "" {
		result = fmt.Sprintf("%s.%s", result, destination)
	}
	return result
}

// NewCreate creates a new Create cmd
func NewCreate(aggregate string, destination string, groupID string) Create {
	return Create{
		Message: messages.NewMessage(buildMessageType("create", aggregate, destination), "", groupID),
	}
}

// Get is used by a client to get the latest snapshot of an aggregate.
//
// The response for Get is usually the aggragte with Message information and can be found within a model.
type Get IDCMD

// NewGet creates a new Get cmd
func NewGet(aggregate string, id string, destination string, groupID string) Get {
	return Get{
		Identifier: messages.Identifier{
			ID:      id,
			Message: messages.NewMessage(buildMessageType("get", aggregate, destination), "", groupID),
		},
	}
}

// Delete is used by a client to delete the latest snapshot of an aggregate.
//
// The response of Delete is Deleted.
type Delete IDCMD

// NewDelete creates a new Delete cmd
func NewDelete(aggregate string, id string, destination string, groupID string) Delete {
	return Delete{
		Identifier: messages.Identifier{
			ID:      id,
			Message: messages.NewMessage(buildMessageType("delete", aggregate, destination), "", groupID),
		},
	}
}

// Start indicates that something with the ID should be started.
//
// As an example an event with the type start.scan with the id 1 would start scan id 1
type Start IDCMD

// NewStart creates a new Start cmd
func NewStart(aggregate string, id string, destination string, groupID string) Start {
	return Start{
		Identifier: messages.Identifier{
			ID:      id,
			Message: messages.NewMessage(buildMessageType("start", aggregate, destination), "", groupID),
		},
	}
}

// Stop indicates that something with the ID should be stopped
type Stop struct {
	eventType
	messages.Identifier
}

// Modify indicates that a entity should be modified.
//
// The values to modify are within Values, they typically match the names of the aggregate
// but with lower case starting.
type Modify struct {
	eventType
	messages.Identifier
	Values map[string]interface{} `json:"values"`
}

// Register indicates that something with the ID should be registered
type Register struct {
	eventType
	messages.Identifier
}

// LoadVTs signals that all sensors should update their VTs
type LoadVTs struct {
	eventType
	messages.Message
}

// NewModify creates a new Modify cmd
func NewModify(aggregate string, id string, values map[string]interface{}, destination string, groupID string) Modify {
	return Modify{
		Identifier: messages.Identifier{
			ID:      id,
			Message: messages.NewMessage(buildMessageType("modify", aggregate, destination), "", groupID),
		},
		Values: values,
	}
}
