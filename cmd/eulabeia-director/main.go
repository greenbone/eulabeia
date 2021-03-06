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
	"fmt"

	_ "github.com/greenbone/eulabeia/logging/configuration"
	"github.com/rs/zerolog/log"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/director/scan"
	"github.com/greenbone/eulabeia/director/sensor"
	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/director/vt"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/process"
	"github.com/greenbone/eulabeia/storage"
)

func main() {
	clientid := flag.String(
		"clientid",
		"eulabeia-director",
		"A clientid for the connection",
	)
	configPath := flag.String(
		"config",
		"",
		"Path to config file, default: search for config file in TODO",
	)
	flag.Parse()
	configuration, err := config.New(*configPath, "eulabeia")
	if err != nil {
		panic(err)
	}
	config.OverrideViaENV(configuration)

	prepare_topic := func(aggregate_name string) string {
		return fmt.Sprintf(
			"%s/%s/cmd/director",
			configuration.Context,
			aggregate_name,
		)
	}
	log.Info().Msgf("Starting director with context %s", configuration.Context)
	client, err := mqtt.FromConfiguration(*clientid, nil, &configuration.Connection)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create MQTT client.")
	}
	err = client.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MQTT.")
	}
	crypt, err := storage.NewRSACrypt(*configuration)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create RSA-Key.")
	}
	device, err := storage.New(configuration.Director.StoragePath, crypt)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create storage.")
	}
	container := []handler.Container{
		sensor.New(device),
		target.New(device),
		scan.New(device),
		vt.New(configuration.Director.VTSensor),
	}

	err = client.Subscribe([]string{
		prepare_topic("sensor"),
		prepare_topic("target"),
		prepare_topic("scan"),
		prepare_topic("vt"),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to subscribe.")
	}
	preoprocessor := []connection.Preprocessor{
		scan.ScanPreprocessor{Context: configuration.Context},
	}

	r := handler.NewRegister(
		configuration.Context,
		container,
		nil,
		preoprocessor,
		[]connection.Publisher{client},
		client.In(),
		connection.DefaultOut,
	)
	r.Start()

	process.Block(client)
}
