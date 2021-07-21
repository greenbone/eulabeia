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

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

type StartStop struct {
	StartChan chan string
	StopChan  chan string
}

func (handler StartStop) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg cmds.IDCMD
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	switch msg.Type {
	case "scan.start":
		handler.StartChan <- msg.ID
	case "scan.stop":
		handler.StopChan <- msg.ID
	}
	return nil, nil
}

type Registered struct {
	RegChan chan struct{}
}

func (handler Registered) On(topic string, message []byte) (*connection.SendResponse, error) {
	handler.RegChan <- struct{}{}
	return nil, nil
}

type Status struct {
	RunChan chan string
	FinChan chan string
}

func (handler Status) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg info.Status
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	switch msg.Status {
	case "running":
		handler.RunChan <- msg.ID
	case "stopped", "finished", "interrupted":
		handler.FinChan <- msg.ID
	}
	return nil, nil
}

type LoadVTs struct {
	VtsChan chan struct{}
}

func (handler LoadVTs) On(topic string, message []byte) (*connection.SendResponse, error) {
	handler.VtsChan <- struct{}{}
	return nil, nil
}
