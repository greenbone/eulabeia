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
