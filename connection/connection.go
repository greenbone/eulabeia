// Package connection contains interfaces for generic message handling
package connection

import (
	"io"
)

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
	On(message []byte) (interface{}, error)
}

// Publisher is the interface that wraps the basic Publish method.
//
// Publish sends a message to the given topic.
type Publisher interface {
	Publish(topic string, message interface{}) error
}

// Subscriber the interface that wraps the basic Subscribe method.
//
// Subscribe iterates through each handler and registers each OnMessage to a
// topic.
type Subscriber interface {
	Subscribe(handler map[string]OnMessage) error
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
