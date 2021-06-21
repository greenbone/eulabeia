// Package connection contains interfaces for generic message handling
package connection

import (
	"io"
)

type OnMessage interface {
	On(message []byte) (interface{}, error)
}

type Publisher interface {
	Publish(topic string, message interface{}) error
}

type Subscriber interface {
	Subscribe(handler map[string]OnMessage) error
}

type Connect interface {
	Connect() error
}

type PubSub interface {
	io.Closer
	Connect
	Publisher
	Subscriber
}
