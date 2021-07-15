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

package main

import (
	"flag"
	"log"
	"os"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/process"
	"github.com/greenbone/eulabeia/sensor"
)

func main() {
	// topic := "eulabeia/+/+/sensor"
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	configuration, err := config.New(*configPath, "eulabeia")
	if err != nil {
		panic(err)
	}

	config.OverrideViaENV(configuration)
	server := configuration.Connection.Server
	if configuration.Sensor.Id == "" {
		sensor_id, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		configuration.Sensor.Id = sensor_id
	}

	log.Println("Starting sensor")
	client, err := mqtt.New(server, configuration.Sensor.Id, "", "",
		&mqtt.LastWillMessage{
			Topic: "eulabeia/sensor/cmd/director",
			MSG: cmds.Delete{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("delete.sensor", "", ""),
					ID:      configuration.Sensor.Id,
				},
			}})
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = client.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	sens := sensor.NewScheduler(client, configuration.Sensor.Id, configuration.ScannerPreferences)
	log.Printf("Starting Scheduler")
	sens.Start()
	process.Block(client)
	log.Printf("Stopping Scheduler")
	term := sens.Stop()
	log.Printf("Wait for all OpenVAS processes to end.")
	<-term
}
