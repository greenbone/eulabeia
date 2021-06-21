// Package handler implements various handler
// for events and messages
package handler

import (
	"fmt"
	"github.com/greenbone/eulabeia/connection"
	"github.com/tidwall/gjson"
)

// OnEvent are typically used to handle messages that change/create a data structure
type OnEvent interface {
	// On will be called by a connection.OnMessage implementation.
	// Returns a response as interface{} if the broker implementation should reply with a message.
	On(messageType string, message []byte) (interface{}, error)
}

type Generic struct {
	changeEventHandler []OnEvent
}

func (mh Generic) On(message []byte) (result interface{}, err error) {
	messageType := gjson.GetBytes(message, "message_type")
	if messageType.Type == gjson.Null {
		err = fmt.Errorf("message: %s does not contain messageType", string(message))
		return
	}
	for _, i := range mh.changeEventHandler {
		result, err = i.On(messageType.String(), message)
		if result != nil || err != nil {
			return
		}
	}
	return
}

func New(changeEventHandler []OnEvent) (connection.OnMessage, error) {
	return Generic{changeEventHandler: changeEventHandler}, nil
}
