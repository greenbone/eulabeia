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
	_ "github.com/greenbone/eulabeia/logging/configuration"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/feedservice"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/process"
	"github.com/greenbone/eulabeia/sensor"
)

func main() {
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

	log.Info().Msgf("Starting sensor (%s) on context (%s)\n", configuration.Sensor.Id, configuration.Context)
	client, err := mqtt.New(server, configuration.Sensor.Id, "", "",
		&mqtt.LastWillMessage{
			Topic: "scanner/sensor/cmd/director",
			MSG: cmds.Delete{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("delete.sensor", "", ""),
					ID:      configuration.Sensor.Id,
				},
			}},
		nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create MQTT")
	}
	err = client.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}
	feed := feedservice.NewFeed(client, configuration.Context, configuration.Sensor.Id, configuration.Feedservice.RedisDbAddress)
	log.Printf("Starting Feed Service\n")
	feed.Start()
	sens := sensor.NewScheduler(client, configuration.Sensor.Id, configuration.ScannerPreferences, configuration.Context)
	log.Printf("Starting Scheduler\n")
	sens.Start()
	process.Block(client, sens, feed)
}
