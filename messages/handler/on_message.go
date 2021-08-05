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

package handler

import (
	"errors"
	"fmt"
	"log"

	"encoding/json"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
	"github.com/tidwall/gjson"
)

// InterfaceArrayToStringArray is a conenvience function to transform []interface{} to []string
//
// It is usually used on modify due to the map[string]interface{} within a modify message
func InterfaceArrayToStringArray(v interface{}) []string {
	if cv, ok := v.([]interface{}); ok {
		strings := make([]string, len(cv), cap(cv))
		for i, j := range cv {
			if s, ok := j.(string); ok {
				strings[i] = s
			}
		}
		return strings
	}
	return nil
}

func InterfaceToPlugins(v interface{}) models.VTsList {
	return models.VTsList{}
}

// ParseMessageType tries to parse the messages.MessageType based on a []byte message
func ParseMessageType(message []byte) (*messages.MessageType, error) {
	messageType := gjson.GetBytes(message, "message_type")
	if messageType.Type == gjson.Null {
		return nil, errors.New("unable to find message_type")
	}
	mt, err := messages.ParseMessageType(messageType.String())
	if err != nil {
		return nil, fmt.Errorf("incorrect message_type %s", messageType.String())
	}
	return mt, nil
}

// onMessage is a struct containing aggregates for registered types.
//
// The messages.MessageType is normalized like what.on e.g. create.target
// onMessage tries to parse the given messages to messages.Create,
// messages.Modify, messages.Get then tries to find via MessageType the
// Aggregate via handler.
type onMessage struct {
	lookup  map[string]Container
	context string
}

func (om onMessage) On(topic string, message []byte) (*connection.SendResponse, error) {
	mt, err := ParseMessageType(message)
	if err != nil {
		return messages.EventToResponse(om.context, info.Failure{
			Identifier: messages.Identifier{Message: messages.NewMessage("failure", "", "")},
			Error:      fmt.Sprintf("%s", err),
		}), nil
	}
	if h, ok := om.lookup[mt.Aggregate]; ok {
		use, fuse := containerClosure(h, mt.Function)
		if e := json.Unmarshal(message, use); e != nil {
			return messages.EventToResponse(om.context, info.Failure{
				Identifier: messages.Identifier{Message: messages.NewMessage("failure", "", "")},
				Error:      fmt.Sprintf("unable to parse %s: %s", mt, e),
			}), nil
		}
		r, f, e := fuse()
		if e != nil {
			return nil, e
		}
		if f != nil {
			return messages.EventToResponse(om.context, f), e
		}
		return messages.EventToResponse(om.context, r), e

	}
	log.Printf("unable to identify entity %s", mt)
	return nil, nil
}

// New returns a new connection.OnMessage handler
func New(context string, container ...Container) connection.OnMessage {
	lookup := map[string]Container{}

	for _, c := range container {
		lookup[c.Topic] = c
	}
	return onMessage{
		lookup:  lookup,
		context: context,
	}
}
