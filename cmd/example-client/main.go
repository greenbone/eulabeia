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
	"flag"
	"time"

	"github.com/greenbone/eulabeia/client"
	"github.com/greenbone/eulabeia/connection"
	_ "github.com/greenbone/eulabeia/logging/configuration"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/rs/zerolog/log"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/models"
)

var target = models.Target{
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
		Group: []models.VTFilter{
			{
				Key:   "family",
				Value: "my test family",
			},
		},
	},
	Exclude: []string{"exclude1"},
	Sensor:  "localhorst",
	AliveTest: models.AliveTest{
		Test_alive_hosts_only: true,
		Methods:               2,
		Ports:                 []int{80, 137, 587, 3128, 8081},
	},
	Parallel: true,
	Credentials: map[string]map[string]string{
		"ssh": {
			"private_key": "denkste",
		},
	},
}

func verifyGetVT(cc client.Configuration) {
	p, err := client.From(cc, cmds.NewGet("vt", "0.0.0.0.0.0.0.0.0.1", "director", "getvts"))
	if err != nil {
		log.Panic().Err(err).Msg("Unable create program for get.vts")
	}
	_, f, err := p.Run()
	if err != nil {
		log.Panic().Err(err).Msgf("get.vt failed %+v", f)
	}
}

func verifyStartScan(cc client.Configuration) messages.GetID {

	cc.Retries = 0
	cc.RetryInterval = 0

	p, err := client.From(cc, cmds.NewCreate("target", "director", "scantest"))
	if err != nil {
		log.Panic().Err(err).Msg("Unable to create scantest program")
	}
	modifyTarget := client.ModifyBasedOnGetID("target", "director", func(gi messages.GetID) map[string]interface{} {
		v, err := client.ToValues(target)
		if err != nil {
			log.Panic().Err(err).Msgf("%T to values", target)
		}
		return v
	})
	modifyScan := client.ModifyBasedOnGetID("scan", "director", func(gi messages.GetID) map[string]interface{} {
		return map[string]interface{}{
			"target_id": gi.GetID(),
		}
	})

	verifyFailure := client.DefaultVerifier(client.FailureParser)
	p = p.Next(nil, modifyTarget, verifyFailure, client.DefaultVerifier(client.ModifiedParser))
	p = p.Next(nil, modifyScan, verifyFailure, client.DefaultVerifier(client.ModifiedParser))
	p = p.Next(nil, client.StartBasedOnGetID("scan", "director"), client.OpenvasScanFailure, client.OpenvasScanSuccess)
	s, f, err := p.Start()
	if f != nil {
		log.Panic().Msgf("Failure while waiting for sensor: %+v", f)
	}
	if err != nil {
		log.Panic().Err(err).Msg("Start scan failed")
	}

	log.Info().Msgf("Finished with %+v", s)
	return s.(messages.GetID)
}

func main() {
	log.Info().Msg("Starting example client")
	clientid := flag.String("clientid", "", "A clientid for the connection")
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
	c, err := mqtt.FromConfiguration(*clientid, nil, &configuration.Connection)
	if err != nil {
		log.Fatal().Msgf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Fatal().Msgf("Failed to connect: %s", err)
	}
	defer c.Close()
	c.Subscribe([]string{
		"#",
	})
	out := make(chan *connection.SendResponse, 1)
	go func(pub connection.Publisher) {
		for o := range out {
			log.Trace().Msgf("Sending message to %s", o.Topic)
			pub.Publish(o.Topic, o.MSG)
		}
	}(c)
	cc := client.Configuration{
		Context:       "scanner",
		Out:           out,
		In:            c.In(),
		Timeout:       5 * time.Second,
		Retries:       10,
		RetryInterval: 1 * time.Second,
	}
	p, err := client.From(cc, cmds.NewGet("sensor", "localhorst", "director", "initial"))
	if err != nil {
		log.Panic().Err(err).Msg("unable to create wait program")
	}
	_, f, err := p.Start()
	if f != nil {
		log.Panic().Msgf("Failure while waiting for sensor: %+v", f)

	}
	if err != nil {
		log.Panic().Err(err).Msg("error while waiting for sensor")
	}
	received := make(chan *client.Received)
	go func(r chan *client.Received) {
		for m := range r {
			log.Debug().Msgf("[%d][%T] %+v", m.State, m.Event, m.Event)
		}
	}(received)
	cc.DownStream = received
	_ = verifyStartScan(cc)
	verifyGetVT(cc)
}
