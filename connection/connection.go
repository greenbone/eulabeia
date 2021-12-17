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

// Package connection contains interfaces for generic message handling
package connection

import (
	"io"
	"sync"
)

// Send Response is used to indicate that a response message should be send
type SendResponse struct {
	MSG   interface{} // The actual messages, will be transformed to json
	Topic string      // The topic to send the message to
}

// OnMessage is the interface that wraps the basic On method.
//
// The behavior of On is that the interface{} is a response and should be send
// back to the same topic as the initial message and errors should be handled
// by the user of OnMessage.
//
// If the initial message is incorrect or if it can only be fixed by the sender
// the implementation of OnMessage should response with  a response
// (e.g. messages.Failure) instead of an error.
type OnMessage interface {
	On(topic string, message []byte) (*SendResponse, error)
}

// Publisher is the interface that wraps the basic Publish method.
//
// Publish sends a message to the given topic.
type Publisher interface {
	Publish(topic string, message interface{}) error
}

// TopicData is a tuple for incoming messages
type TopicData struct {
	Topic   string   // The topic this message was send to
	Message []byte   // The actual content
	Sender  []string // The sender this messages has been sent from
}

// Preprocessor is the interface the wrapt the basic Preprocess method.
//
// Preprocess is called to allow client specific complex events to be split up
// in mutiple smaller events. This allows essentially shortcuts like creating
// a huge start.scan event containing all the target data directly instead of
// follwing the actual design pattern.
// It will be called before the subscriber handling there it needs to follow
// the basic string, []byte principle instead of complex structs.
// Returns a list of []TopicData and true when it was handled.
// Returns nil or an empty list and false when it was not handled.
type Preprocessor interface {
	Preprocess(topic string, message []byte) ([]TopicData, bool)
}

// Subscriber the interface that wraps the basic Subscribe method.
//
// Subscribe iterates through each handler and registers each OnMessage to a
// topic and will send each incoming data to the in channel
type Subscriber interface {
	Subscribe([]string) error
	In() <-chan *TopicData
}

// Connecter is the interface that wraps the basic Connect method.
//
// Connect connectes to a broker if necessary.
type Connecter interface {
	Connect() error
}

// PubSub is the interface that contains the methods needed to simulate a broker
//
// The typical call order of PubSub is:
// - Connect
// - Subscribe
// - 1..n Publish
// - Close
type PubSub interface {
	io.Closer
	Connecter
	Publisher
	Subscriber
}

// CombineHandler is a method to combine multiple topic handler
func CombineHandler(mh ...map[string]OnMessage) map[string]OnMessage {
	result := map[string]OnMessage{}
	for _, h := range mh {
		for k, v := range h {
			result[k] = v
		}
	}
	return result
}

// ClosureOnMessage is struct for simple OnMesage implementation that don't
// require a own struct
type ClosureOnMessage struct {
	Closure func(td TopicData) (*SendResponse, error)
}

func (a ClosureOnMessage) On(
	topic string,
	message []byte,
) (*SendResponse, error) {
	return a.Closure(TopicData{topic, message, []string{}})
}

// ClosurePublisher is struct for simple Publish implementation that don't
// require a own struct
type ClosurePublisher struct {
	Closure func(string, interface{}) error
}

func (s ClosurePublisher) Publish(topic string, message interface{}) error {
	return s.Closure(topic, message)
}

// ClosurePublisher is struct for a simple preprocess implementation that don't
// require a own struct
type ClosurePreprocessor struct {
	Closure func(TopicData) ([]TopicData, bool)
}

func (s ClosurePreprocessor) Preprocess(
	topic string,
	message []byte,
) ([]TopicData, bool) {
	return s.Closure(TopicData{topic, message, []string{}})
}

var NoOpPreprocessor = []Preprocessor{}

// Defaulting to three in flight messages
var DefaultOut = make(chan *SendResponse, 3)

func MergePubSubIn(ps []PubSub) <-chan *TopicData {
	out := make(chan *TopicData, 3)
	var wg sync.WaitGroup
	wg.Add(len(ps))
	for _, p := range ps {
		go func(c <-chan *TopicData) {
			defer wg.Done()
			for in := range c {
				out <- in
			}
		}(p.In())
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}
