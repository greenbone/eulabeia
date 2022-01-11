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

package handler

import (
	"encoding/json"
	"fmt"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/rs/zerolog/log"
)

type Register struct {
	context      string
	onMessage    map[string]connection.OnMessage
	container    map[string]Container
	preprocessor []connection.Preprocessor
	publisher    []connection.Publisher
	out          <-chan *connection.SendResponse
	in           <-chan *connection.TopicData
	republish    func(connection.TopicData) *connection.SendResponse
}

// NewSingleMessageHandler creates a MessageHandler with DefaultOut based on
// PubSub
func NewDefaultMessageHandler(
	context string,
	container []Container,
	onMessages map[string]connection.OnMessage,
	pubsub ...connection.PubSub,
) Register {
	pubs := make([]connection.Publisher, len(pubsub))
	topics := make([]string, len(container)+len(onMessages))
	i := 0
	for _, c := range container {
		topics[i] = c.Topic
		i = i + 1
	}

	for k := range onMessages {
		topics[i] = k
		i = i + 1
	}
	for i, p := range pubsub {
		pubs[i] = p
		p.Subscribe(topics)
	}

	return NewRegister(
		RegisterConfiguration{
			context,
			container,
			onMessages,
			connection.NoOpPreprocessor,
			pubs,
			connection.MergePubSubIn(pubsub),
			connection.DefaultOut,
			nil,
		},
	)
}

// RegisterConfiguration is used to configure a register
// TODO translate Container to OnMessages instead of providing both directly
type RegisterConfiguration struct {
	Context      string                                              // The context of mqtt usually scanner
	Container    []Container                                         // Message processor preferred by abstract systems (e.g. director)
	OnMessages   map[string]connection.OnMessage                     // Message processor preferred by systems that require a direct form of control (e.g. sensor)
	Preprocessor []connection.Preprocessor                           // Message preprocessor to manipulate incoming message before message processor can handle them
	Publisher    []connection.Publisher                              // Are used to publish outgoing messages (informed via Out channel)
	In           <-chan *connection.TopicData                        // Channel to inform about incoming messages
	Out          <-chan *connection.SendResponse                     // Channel to inform about outgoing messages
	RePublish    func(connection.TopicData) *connection.SendResponse // Closure to decide if a message should be published again; nil if not supported
}

// NewRegister creates a new register for message handler
//
// A register does delegate incoming messages to registered handler
// and may publish outgoing data to registered publisher.
func NewRegister(
	config RegisterConfiguration,
) Register {
	handler := make(map[string]Container)
	for _, c := range config.Container {
		handler[c.Topic] = c
	}
	return Register{
		config.Context,
		config.OnMessages,
		handler,
		config.Preprocessor,
		config.Publisher,
		config.Out,
		config.In,
		config.RePublish,
	}
}

func (s *Register) preprocess(
	topic string,
	message []byte,
) ([]connection.TopicData, bool) {
	var td []connection.TopicData
	handled := false
	for _, p := range s.preprocessor {
		if r, ok := p.Preprocess(topic, message); ok {
			handled = true
			td = append(td, r...)
		}
	}
	return td, handled
}

func (om *Register) elevate(message []byte) (*connection.SendResponse, error) {
	var msg messages.Message
	if err := json.Unmarshal(message, &msg); err != nil {
		return messages.EventToResponse(om.context, info.Failure{
			Identifier: messages.Identifier{
				Message: messages.NewMessage("failure", "", ""),
			},
			Error: fmt.Sprintf("%s", err),
		}), nil
	}
	mt := msg.MessageType()
	if h, ok := om.container[mt.Aggregate]; ok {
		use, fuse := ContainerMethod(h, mt.Function)
		if use == nil && fuse == nil {
			return nil, nil
		}
		if e := json.Unmarshal(message, use); e != nil {
			return messages.EventToResponse(om.context, info.Failure{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("failure", "", ""),
				},
				Error: fmt.Sprintf("unable to parse %s: %s", mt, e),
			}), nil
		}
		r, f, e := fuse()
		if e != nil {
			return nil, e
		}
		if f != nil {
			return messages.EventToResponse(om.context, f), e
		}
		return messages.EventToResponse(om.context, r), e

	}
	log.Printf("unable to identify entity %s", mt)
	return nil, nil

}

func (s *Register) send(resp *connection.SendResponse) {
	for i, p := range s.publisher {
		if err := p.Publish(resp.Topic, resp.MSG); err != nil {
			log.Error().
				Err(err).
				Msgf("Error occured while publishing data on topic %s on publisher %d", resp.Topic, i)
		}
	}

}
func (s *Register) Check() bool {
	select {
	case in, open := <-s.in:
		if in != nil {
			log.Trace().Msgf("Received msg (%s) on %s", in.Message, in.Topic)
			var tds []connection.TopicData
			if t, ok := s.preprocess(in.Topic, in.Message); ok {
				tds = t
			} else {
				tds = []connection.TopicData{*in}
			}
			log.Trace().Msgf("Calling %d handler", len(tds))
			for _, t := range tds {
				var resp *connection.SendResponse
				var err error
				if s.republish != nil {
					if msg := s.republish(t); msg != nil {
						s.send(msg)
					}
				}
				if h, ok := s.onMessage[t.Topic]; ok {
					resp, err = h.On(t.Topic, t.Message)
				} else {
					resp, err = s.elevate(t.Message)
					// Extend republish to verify for topics
					if resp == nil && err == nil && s.republish != nil {
						log.Trace().Msgf("Sending message to %s", in.Topic)
						s.send(&connection.SendResponse{
							Topic: in.Topic,
							MSG:   in.Message,
						})

						log.Trace().Msgf("Sent message to %s", in.Topic)
					}
				}
				if err != nil {
					log.Error().
						Err(err).
						Msgf("Error occured while processing message for %s", t.Topic)
				} else if resp != nil {
					log.Debug().Msgf("Sending resp to %s", resp.Topic)
					s.send(resp)
				}

				log.Trace().Msgf("Finished %s handler", t.Topic)
			}
		}
		return open

	case out, open := <-s.out:
		if out != nil {
			log.Trace().Msgf("Sending message (%v) to %s", out.MSG, out.Topic)
			s.send(out)
		}
		return open
	}

}

func (s *Register) Start() {
	go func() {
		for s.Check() {
			log.Trace().Msg("checking")
		}
		log.Debug().Msg("MessageHander stopped")
	}()
}
