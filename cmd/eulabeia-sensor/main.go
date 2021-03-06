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
	"time"

	"github.com/greenbone/eulabeia/client"
	"github.com/greenbone/eulabeia/connection"
	_ "github.com/greenbone/eulabeia/logging/configuration"
	"github.com/rs/zerolog/log"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/feedservice"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/process"
	"github.com/greenbone/eulabeia/sensor"
)

func waitForRegistration(context string, clientID string, ps connection.PubSub) {
	exit := false
	defer func() {
		exit = true
	}()
	go func() {
		for !exit {
			select {
			case send := <-connection.DefaultOut:
				ps.Publish(send.Topic, send.MSG)
			default:
			}
		}
	}()
	ps.Subscribe([]string{
		fmt.Sprintf("%s/sensor/info", context),
	})
	clientc := client.Configuration{
		Context:       context,
		Out:           connection.DefaultOut,
		In:            ps.In(),
		Timeout:       5 * time.Second,
		Retries:       10,
		RetryInterval: 1 * time.Second,
	}
	reg, err := client.From(clientc, cmds.NewModify("sensor", clientID, nil, "director", clientID))
	if err != nil {
		log.Panic().Err(err).Msgf("Unable to create register program")
	}
	_, f, err := reg.Run()
	if err != nil || f != nil {
		log.Panic().Err(err).Msgf("Unable to register %s: %+v", clientID, f)
	}
}

func main() {
	configPath := flag.String(
		"config",
		"",
		"Path to config file, default: search for config file in TODO",
	)
	clientid := flag.String(
		"clientid",
		"eulabeia-sensor",
		"A clientid for the connection",
	)

	flag.Parse()
	configuration, err := config.New(*configPath, "eulabeia")
	if err != nil {
		panic(err)
	}

	config.OverrideViaENV(configuration)
	log.Info().
		Msgf("Starting sensor (%s) on context (%s)\n", *clientid, configuration.Context)
	lwm := &mqtt.LastWillMessage{
		Topic: "scanner/sensor/cmd/director",
		MSG: cmds.Delete{
			EventType: cmds.EventType{},
			Identifier: messages.Identifier{
				Message: messages.NewMessage("delete.sensor", "", ""),
				ID:      *clientid,
			},
		}}
	mqttc, err := mqtt.FromConfiguration(*clientid, lwm, &configuration.Connection)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create MQTT")
	}
	err = mqttc.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}

	waitForRegistration(configuration.Context, *clientid, mqttc)

	feed := feedservice.NewFeed(
		configuration.Context,
		*clientid,
		configuration.Feedservice.RedisDbAddress,
	)
	sens := sensor.NewScheduler(
		connection.DefaultOut,
		*clientid,
		configuration.ScannerPreferences,
		configuration.Context,
	)
	h := connection.CombineHandler(
		feed.Handler(),
		sens.Handler(),
	)
	mhm := handler.NewDefaultMessageHandler(configuration.Context, nil, h, mqttc)
	mhm.Start()
	sens.Start(mhm)
	log.Debug().Msg("Starting MessageListener")
	process.Block(mqttc, sens, feed)
}
