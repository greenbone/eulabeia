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

// Package handler implements various handler for events and messages
package handler

import (
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

// Starter is the interface that wraps the basic Start method.
//
// Start is used to start an new event chain (e.g. start.scan)
type Starter interface {
	Start(cmds.Start) (messages.Event, *info.Failure, error)
}

// Creater is the interface that wraps the basic Create method.
//
// Create is used on aggregate handler to handle messages.Create.
// Creates a new entity of a given type via messages.Message.MessageType.
// It responds with info.Created which contains the id of the entity.
type Creater interface {
	Create(cmds.Create) (*info.Created, error)
}

// Modifier is the interface that wraps the basic Modify method.
//
// Modifies a entity of a given type via messages.Message.MessageType and ID.
// It responds with info.Modified on successful alteration
// info.Failure on incorrect Values
type Modifier interface {
	Modify(cmds.Modify) (*info.Modified, *info.Failure, error)
}

// Getter is the interface that wraps the basic Get method.
//
// Gets a entity of a given type via messages.Message.MessageType and ID.
// It responds with interface{} on success and info.Failure when not found.
type Getter interface {
	Get(cmds.Get) (messages.Event, *info.Failure, error)
}

// Deleter is the interface that wraps the basic Delete method.
//
// Delets a entity of a given type found via messages.Message,MessageType and ID.
// It responds with info.Deleted on successful removal
type Deleter interface {
	Delete(cmds.Delete) (*info.Deleted, *info.Failure, error)
}

// Aggregate is the interface to handle Aggregate messages
type Aggregate interface {
	Creater
	Modifier
	Deleter
	Getter
}
