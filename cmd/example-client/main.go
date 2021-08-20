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

// This example will create and modify
// 1. a target
// 2. a scan
// when a scan has been modified it starts a scan.
package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/director/scan"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
)

const MEGA_ID = "mega_scan_123"
const context = "eulabeia"
const topic = context + "/+/info"

const (
	CREATED_TARGET  = "created.target"
	MODIFIED_TARGET = "modified.target"
	MODIFIED_SCAN   = "modified.scan"
)

// ExampleHandler parses the message and calls corresponding function of MessageType within do map.
type ExampleHandler struct {
	sync.RWMutex
	do      map[string]func(info.IDInfo, []byte) *connection.SendResponse
	handled []string
	exit    chan os.Signal
}

func (e *ExampleHandler) On(topic string, msg []byte) (*connection.SendResponse, error) {
	e.Lock()
	defer e.Unlock()
	mt, err := handler.ParseMessageType(msg)
	if err != nil {
		// In this example we end the program on a unexpected message so that we can
		// reuse it as a smoke test.
		// However in a production environment you want to either log and ignore or
		// just ignore unparseable messages.
		panic(err)
	}
	log.Printf("Got message: %s", mt)
	var infoMSG info.IDInfo
	if err := json.Unmarshal(msg, &infoMSG); err != nil {
		log.Panicf("Unable to parse %s to info.IDInfo (%s)", msg, err)
	}
	f, ok := e.do[infoMSG.ID]
	if !ok {
		f, ok = e.do[mt.String()]
		if !ok {
			log.Panicf("No handler for %s found", mt)
		}
	}
	response := f(infoMSG, msg)
	// On a response without a topic we assume ignored
	if response != nil && response.Topic == "" {
		return nil, nil
	}
	e.handled = append(e.handled, mt.String())
	e.handled = append(e.handled, infoMSG.ID)
	// We assume that if there is no response message that the test scenario is finished
	if response == nil {
		e.exit <- syscall.SIGCONT
		// on response.Topic ignore SendResponse
	}
	return response, nil
}

func ModifyTarget(msg info.IDInfo, _ []byte) *connection.SendResponse {
	modify := cmds.NewModify(
		"target",
		msg.ID,
		map[string]interface{}{"sensor": "openvas"},
		"director",
		msg.GroupID)
	return messages.EventToResponse(context, modify)
}

func CreateScan(msg info.IDInfo, _ []byte) *connection.SendResponse {
	// We use the principle modify over create to directly create a scan with a target ID.
	// Otherwise we need to store the target ID and reuse it on created.scan.
	modify := cmds.NewModify(
		"scan",
		uuid.NewString(),
		map[string]interface{}{"target": msg.ID},
		"director",
		msg.GroupID)
	return messages.EventToResponse(context, modify)
}

func MegaScan(msg info.IDInfo, _ []byte) *connection.SendResponse {
	mega := scan.StartMegaScan{
		Message: messages.NewMessage("start.scan.director", "", ""),
		Scan: models.Scan{
			ID: MEGA_ID,
			Target: models.Target{
				ID:    MEGA_ID,
				Hosts: []string{"localhost"},
				Ports: []string{"80"},
				Plugins: models.VTsList{
					Single: []models.SingleVT{
						{
							OID: "1.3.6.1.4.1.25623.1.0.90022",
							PrefsByID: map[int]interface{}{
								0: "test1",
								1: 2,
							},
							PrefsByName: map[string]interface{}{
								"pref1": "test2",
								"pref2": true,
							},
						},
					},
					Group: map[string]string{
						"family": "my test family",
					},
				},
				Exclude:  []string{"exclude1"},
				Sensor:   "localhorst",
				Alive:    true,
				Parallel: true,
				Credentials: map[string]map[string]string{
					"ssh": {
						"private_key": "denkste",
					},
				},
			},
			Finished: []string{"hosts2"},
		},
	}
	return messages.EventToResponse(context, mega)
}

func VerifyForScanStatus(i info.IDInfo, b []byte) *connection.SendResponse {
	if i.Type == "status.scan" {
		var status info.Status
		if json.Unmarshal(b, &status) != nil {
			log.Panicf("Unable to parse: %s", string(b))
		}
		if status.Status == "interrupted" {
			return nil
		}
	}
	return &connection.SendResponse{}
}

func Done(_ info.IDInfo, _ []byte) *connection.SendResponse {
	return nil
}

func Verify(eh *ExampleHandler) {
	var difference []string
	for k := range eh.do {
		found := false
		for _, v := range eh.handled {
			if k == v {
				found = true
				break
			}
		}
		if !found {
			difference = append(difference, k)
		}
	}
	if len(difference) > 0 {
		log.Fatalf("FAILURE: %s were not handled.", difference)
	} else {
		log.Println("SUCCESS")
	}
}

func main() {
	clientid := flag.String("clientid", "", "A clientid for the connection")
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	configuration, err := config.New(*configPath, "eulabeia")
	if err != nil {
		panic(err)
	}
	config.OverrideViaENV(configuration)
	server := configuration.Connection.Server

	log.Println("Starting example client")
	c, err := mqtt.New(server, *clientid+uuid.NewString(), "", "", nil, nil)
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	defer c.Close()
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	ic := make(chan os.Signal, 1)
	defer close(ic)
	mh := ExampleHandler{
		do: map[string]func(info.IDInfo, []byte) *connection.SendResponse{
			CREATED_TARGET:  ModifyTarget,
			MODIFIED_TARGET: CreateScan,
			MODIFIED_SCAN:   MegaScan,
			MEGA_ID:         VerifyForScanStatus,
		},
		exit: ic,
	}
	defer Verify(&mh)
	err = c.Subscribe(map[string]connection.OnMessage{topic: &mh})
	if err != nil {
		panic(err)
	}
	timer := time.NewTimer(1 * time.Minute)
	defer timer.Stop()
	go func() {
		<-timer.C
		ic <- syscall.SIGABRT
	}()
	signal.Notify(ic, os.Interrupt, syscall.SIGTERM)
	err = c.Publish("eulabeia/target/cmd/director", cmds.NewCreate("target", "director", ""))
	if err != nil {
		log.Panicf("Failed to publish: %s", err)
	}
	<-ic
	log.Printf("After handling %s it is time to say good bye", mh.handled)
}
