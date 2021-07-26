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

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/director/scan"
	"github.com/greenbone/eulabeia/director/sensor"
	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/process"
	"github.com/greenbone/eulabeia/storage"
)

func main() {
	clientid := flag.String("clientid", "eulabeia-director", "A clientid for the connection")
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	configuration, err := config.New(*configPath, "eulabeia")
	if err != nil {
		panic(err)
	}
	config.OverrideViaENV(configuration)
	server := configuration.Connection.Server

	log.Println("Starting director")
	client, err := mqtt.New(server, *clientid, "", "", nil, []connection.Preprocessor{scan.StartMegaScan{}})
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = client.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	device := storage.File{Dir: configuration.Director.StoragePath}
	context := configuration.Context
	err = client.Subscribe(map[string]connection.OnMessage{
		"eulabeia/sensor/cmd/director": handler.New(context, sensor.New(device)),
		"eulabeia/target/cmd/director": handler.New(context, target.New(device)),
		"eulabeia/scan/cmd/director":   handler.New(context, scan.New(device)),
	})
	if err != nil {
		panic(err)
	}

	process.Block(client)
}
