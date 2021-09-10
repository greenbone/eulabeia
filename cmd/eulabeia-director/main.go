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
	"github.com/greenbone/eulabeia/topic"
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

	log.Printf("Starting director with context %s\n", configuration.Context)
	client, err := mqtt.New(server, *clientid, "", "", nil, []connection.Preprocessor{
		scan.ScanPreprocessor{Context: configuration.Context}})
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = client.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	crypt, err := storage.NewRSACrypt(*configuration)
	if err != nil {
		log.Panicf("Failed create RSA: %s", err)
	}
	device, err := storage.New(configuration.Director.StoragePath, crypt)
	if err != nil {
		log.Panicf("Failed to create storage: %s", err)
	}
	err = client.Subscribe(map[string]connection.OnMessage{
		topic.NewCmd(configuration.Context, "sensor", "director"): handler.New(configuration.Context, sensor.New(device)),
		topic.NewCmd(configuration.Context, "target", "director"): handler.New(configuration.Context, target.New(device)),
		topic.NewCmd(configuration.Context, "scan", "director"):   handler.New(configuration.Context, scan.New(device)),
	})
	if err != nil {
		log.Panicf("Subscribing failed: %s", err)
	}

	process.Block(client)
}
