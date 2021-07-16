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

// Package handler contains various message handler for sensors and initializes MQTT connection
package handler

import (
	"encoding/json"
	"log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

type StartStop struct {
	StartFunc func(string) error
	StopFunc  func(string) error
}

func (handler StartStop) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg cmds.IDCMD
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	if mt.Aggregate == "scan" {
		switch mt.Function {
		case "start":
			if err := handler.StartFunc(msg.ID); err != nil {
				log.Printf("Unable to start scan: %s", err)
			}
		case "stop":
			if err := handler.StartFunc(msg.ID); err != nil {
				log.Printf("Unable to stop scan: %s", err)
			}
		}
	}
	return nil, nil
}

type Registered struct {
	RegChan chan struct{}
	ID      string
}

func (handler Registered) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg info.Created
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	if msg.ID == handler.ID {
		handler.RegChan <- struct{}{}
	}
	return nil, nil
}

type Status struct {
	RunFunc func(string) error
	FinFunc func(string) error
}

func (handler Status) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg info.Status
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	switch msg.Status {
	case "running":
		if err := handler.RunFunc(msg.ID); err != nil {
			log.Printf("Unable to set status to running: %s", err)
		}
	case "stopped", "finished", "interrupted":
		if err := handler.FinFunc(msg.ID); err != nil {
			log.Printf("Unable to set status to running: %s", err)
		}
	}
	return nil, nil
}

type LoadVTs struct {
	VtsFunc func()
}

func (handler LoadVTs) On(topic string, message []byte) (*connection.SendResponse, error) {
	handler.VtsFunc()
	return nil, nil
}
