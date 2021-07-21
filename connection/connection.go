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

// Package connection contains interfaces for generic message handling
package connection

import (
	"io"
)

// Send Response is used to indicate that a response message should be send
type SendResponse struct {
	MSG   interface{}
	Topic string
}

// OnMessage is the interface that wraps the basic On method.
//
// The behavior of On is that the interface{} is a response and should be send
// back to the same topic as the initial message and errors should be handled
// by the user of OnMessage.
//
// If the initial message is incorrect or if it can only be fixed by the sender
// the implementation of OnMessage should response with  a response
// (e.g. messages.Failure) instead of an error.
type OnMessage interface {
	On(topic string, message []byte) (*SendResponse, error)
}

// Publisher is the interface that wraps the basic Publish method.
//
// Publish sends a message to the given topic.
type Publisher interface {
	Publish(topic string, message interface{}) error
}

// Subscriber the interface that wraps the basic Subscribe method.
//
// Subscribe iterates through each handler and registers each OnMessage to a
// topic.
type Subscriber interface {
	Subscribe(handler map[string]OnMessage) error
}

// Connecter is the interface that wraps the basic Connect method.
//
// Connect connectes to a broker if necessary.
type Connecter interface {
	Connect() error
}

// PubSub is the interface that contains the methods needed to simulate a broker
//
// The typical call order of PubSub is:
// - Connect
// - Subscribe
// - 1..n Publish
// - Close
type PubSub interface {
	io.Closer
	Connecter
	Publisher
	Subscriber
}
