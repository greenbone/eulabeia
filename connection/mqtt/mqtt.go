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

// package mqtt contains the mqtt implementation of connection
package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/eclipse/paho.golang/paho"
	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/rs/zerolog/log"
)

// MQTT is connection type for
type MQTT struct {
	client            *paho.Client
	connectProperties *paho.Connect
	qos               byte
	in                chan *connection.TopicData // Is used to send respons messages of a handler downwards
}

func (m MQTT) In() <-chan *connection.TopicData {
	return m.in
}

func (m MQTT) Close() error {
	return m.client.Disconnect(&paho.Disconnect{ReasonCode: 0})
}

func (m MQTT) register(topic string) error {

	m.client.Router.RegisterHandler(topic, func(p *paho.Publish) {
		// verif that it's not sent by this client; this can happen although
		// NoLocal is on when
		// tunneling
		if m.client.ClientID != "" &&
			strings.HasPrefix(
				p.Properties.User.Get("sender"),
				m.client.ClientID,
			) {
			log.Printf(
				"ignoring message on %s due to same clientID (%s)",
				p.Topic,
				m.client.ClientID,
			)
			return
		}
		m.in <- &connection.TopicData{Topic: topic, Message: p.Payload}

	})

	_, err := m.client.Subscribe(context.Background(), &paho.Subscribe{
		// we need NoLocal otherwise we would consum our own messages
		// again and ack them.
		Subscriptions: map[string]paho.SubscribeOptions{
			topic: {QoS: m.qos, NoLocal: true},
		},
	},
	)
	return err
}

func (m MQTT) Subscribe(topics []string) error {
	for _, t := range topics {
		if err := m.register(t); err != nil {
			return err
		}
	}
	return nil
}

func (m MQTT) Publish(topic string, message interface{}) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	props := &paho.PublishProperties{}
	pb := &paho.Publish{
		Topic:      topic,
		QoS:        m.qos,
		Payload:    b,
		Properties: props,
	}
	_, err = m.client.Publish(context.Background(), pb)
	return err
}

func (m MQTT) Connect() error {
	ca, err := m.client.Connect(context.Background(), m.connectProperties)
	if err != nil {
		return err
	}
	if ca.ReasonCode != 0 {
		return fmt.Errorf(
			"failed to connect to %s : %d - %s",
			m.client.Conn.RemoteAddr().String(),
			ca.ReasonCode,
			ca.Properties.ReasonString,
		)
	}
	return nil
}

type LastWillMessage struct {
	Topic string
	MSG   interface{}
}

func (lwm LastWillMessage) asBytes() ([]byte, error) {
	return json.Marshal(lwm.MSG)
}

// Configuration holds information for MQTT
type Configuration struct {
	ClientID      string           // The ID to be used when connecting to a broker
	Username      string           // Username to be used as authentication; empty for anonymous
	Password      string           // Password to be used as authentication with Username
	LWM           *LastWillMessage // LastWillMessage to be send when disconnecting
	CleanStart    bool             // CleanStart when false and SessionExpiry set to > 1 it will reuse a session
	SessionExpiry uint64           // Amount of seconds a session is valid; WARNING when set to 0 it is effectively a cleanstart.
	QOS           byte
	KeepAlive     uint16
	Inflight      uint
}

func FromConfiguration(clientID string, lwm *LastWillMessage, c *config.Connection) (connection.PubSub, error) {
	if !c.CleanStart && clientID == "" {
		log.Warn().Msg("Setting clean start to false requires a clientID; setting it to true due to missing clientID")
		c.CleanStart = true
	}
	if !c.CleanStart && c.SessionExpiry == 0 {
		log.Trace().Msg("Inactive clean start requires a session expiry; using default one day")
		c.SessionExpiry = uint64((24 * time.Hour).Seconds())
	}
	if !c.CleanStart && c.QOS < 1 {
		log.Trace().Msgf("Setting QOS from %d to 1 based on CleanStart false", c.QOS)
		c.QOS = 1
	}

	mqttConfig := Configuration{
		ClientID:      clientID,
		Username:      c.Username,
		Password:      c.Password,
		LWM:           lwm,
		CleanStart:    c.CleanStart,
		SessionExpiry: c.SessionExpiry,
		QOS:           c.QOS,
		KeepAlive:     30,
		Inflight:      1,
	}
	conn, err := net.Dial("tcp", c.Server)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("MQTT: client ID %s, username %s, clean start %v, session expiry %d, qos %d",
		mqttConfig.ClientID,
		mqttConfig.Username,
		mqttConfig.CleanStart,
		mqttConfig.SessionExpiry,
		mqttConfig.QOS,
	)
	return New(conn, mqttConfig)
}

func New(conn net.Conn,
	cfg Configuration,
) (connection.PubSub, error) {

	c := paho.NewClient(paho.ClientConfig{
		Router: paho.NewStandardRouter(),
		Conn:   conn,
	})

	cp := &paho.Connect{
		KeepAlive:  cfg.KeepAlive,
		ClientID:   cfg.ClientID,
		CleanStart: cfg.CleanStart,
		Username:   cfg.Username,
		Password:   []byte(cfg.Password),
	}
	if cfg.LWM != nil {
		if b, err := cfg.LWM.asBytes(); err != nil {
			return nil, err
		} else {
			cp.WillMessage = &paho.WillMessage{
				Topic:   cfg.LWM.Topic,
				QoS:     cfg.QOS,
				Payload: b,
			}
		}
	}
	if cfg.Username != "" {
		cp.UsernameFlag = true
	}
	if cfg.Password != "" {
		cp.PasswordFlag = true
	}

	return &MQTT{
		client:            c,
		connectProperties: cp,
		qos:               cfg.QOS,
		in:                make(chan *connection.TopicData, cfg.Inflight),
	}, nil
}
