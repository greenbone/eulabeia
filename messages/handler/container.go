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

package handler

import (
	"log"

	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

// Container contains interfaces needed for OnMessage
//
// It is a convencience struct for handler so that it can be registered and chosen transparently.
type Container struct {
	Topic    string
	Creater  Creater
	Modifier Modifier
	Getter   Getter
	Deleter  Deleter
	Starter  Starter
}

// FromAggregate is a convencience method to create specialized lookup maps for connection.OnMessage
func FromAggregate(topic string, a Aggregate) Container {
	return Container{
		Topic:    topic,
		Creater:  a,
		Modifier: a,
		Getter:   a,
		Deleter:  a,
	}
}

// FromGetter is a convencience method to create specialized lookup maps for connection.OnMessage
func FromGetter(topic string, a Getter) Container {
	return Container{
		Topic:  topic,
		Getter: a,
	}
}

// containerClosure returns a pointer of a struct to unmarshall the json into as well as a closure to call the
// actual downstream handler.
func containerClosure(h Container, method string) (messages.Event, func() (messages.Event, *info.Failure, error)) {
	var del cmds.Delete
	var create cmds.Create
	var modify cmds.Modify
	var get cmds.Get
	var start cmds.Start
	if method == "delete" && h.Deleter != nil {
		return &del, func() (messages.Event, *info.Failure, error) {
			return h.Deleter.Delete(del)
		}

	} else if method == "create" && h.Creater != nil {
		return &create, func() (messages.Event, *info.Failure, error) {
			r, e := h.Creater.Create(create)
			return r, nil, e
		}
	} else if method == "start" && h.Starter != nil {
		return &start, func() (messages.Event, *info.Failure, error) {
			return h.Starter.Start(start)
		}
	} else if method == "modify" && h.Modifier != nil {
		return &modify, func() (messages.Event, *info.Failure, error) {
			return h.Modifier.Modify(modify)
		}
	} else if method == "get" && h.Getter != nil {
		return &get, func() (messages.Event, *info.Failure, error) {
			return h.Getter.Get(get)
		}
	} else {
		log.Printf("unable to identify method %s", method)
		return nil, nil
	}
}
