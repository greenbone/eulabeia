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
package sensor

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
)

type ScanCmd struct {
	Context string
	Sensor  string
	Stop    func(scanID string) error              // Function to Stop a scan
	Get     func(string) (models.ScanPrefs, error) // Function to get scan prefs corresponding to its ID
}

func (handler ScanCmd) On(topic string, message []byte) (*connection.SendResponse, error) {
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
			return &connection.SendResponse{
				MSG:   cmds.NewGet("scan", msg.ID, "director", msg.GroupID),
				Topic: fmt.Sprintf("%s/%s/%s/%s", handler.Context, "scan", "cmd", "director"),
			}, nil

		case "stop":
			if err := handler.Stop(msg.ID); err != nil {
				log.Printf("Unable to stop scan: %s", err)
			}
		case "get":
			if sp, err := handler.Get(msg.ID); err != nil {
				return nil, err
			} else {
				return &connection.SendResponse{
					Topic: fmt.Sprintf("%s/%s/%s", handler.Context, "scan", "info"),
					MSG: models.GotScanPrefs{
						Message:   messages.NewMessage("got.scan", "", msg.GroupID),
						ScanPrefs: sp,
					},
				}, nil
			}
		}
	}
	return nil, nil
}

type Registered struct {
	Register chan struct{} // Channel to signal succesful registration
	ID       string        // SensorID to compare registered ID with own
}

func (handler Registered) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg info.Created
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	if msg.ID == handler.ID && mt.Function == "modified" && mt.Aggregate == "sensor" {
		handler.Register <- struct{}{}
	}
	return nil, nil
}

type ScanInfo struct {
	Context string
	Sensor  string
	Run     func(string) error      // Function to mark a scan as running
	Fin     func(string) error      // Function to mark a scan as finished
	Start   func(models.Scan) error // Function to Start a scan event chain
}

func (handler ScanInfo) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg info.Status
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}

	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}

	switch mt.Function {
	case "got":
		var msg models.GotScan
		if err := json.Unmarshal(message, &msg); err != nil {
			return nil, err
		}

		go func() {
			if err := handler.Start(msg.Scan); err != nil {
				log.Printf("Unable to start scan: %s", err)
			}
		}()

		return nil, nil

	case "status":
		switch msg.Status {
		case "running":
			if err := handler.Run(msg.ID); err != nil {
				log.Printf("Unable to set status to running: %s", err)
			}
		case "finished":
			if err := handler.Fin(msg.ID); err != nil {
				log.Printf("Unable to set status to finished: %s", err)
			}
		}
	}

	return nil, nil
}

type LoadVTs struct {
	VtsLoad func() // Function to start LoadingVTs (into redis by openvas)
}

func (handler LoadVTs) On(topic string, message []byte) (*connection.SendResponse, error) {
	handler.VtsLoad()
	return nil, nil
}
