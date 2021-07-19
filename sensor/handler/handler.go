// Package handler contains various message handler for sensors and initializes MQTT connection
package handler

import (
	"encoding/json"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

type StartStop struct {
	StartChan chan string
	StopChan  chan string
}

func (handler StartStop) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg cmds.IDCMD
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	switch msg.Type {
	case "scan.start":
		handler.StartChan <- msg.ID
	case "scan.stop":
		handler.StopChan <- msg.ID
	}
	return nil, nil
}

type Registered struct {
	RegChan chan struct{}
}

func (handler Registered) On(topic string, message []byte) (*connection.SendResponse, error) {
	handler.RegChan <- struct{}{}
	return nil, nil
}

type Status struct {
	RunChan chan string
	FinChan chan string
}

func (handler Status) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg info.Status
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	switch msg.Status {
	case "running":
		handler.RunChan <- msg.ID
	case "stopped", "finished", "interrupted":
		handler.FinChan <- msg.ID
	}
	return nil, nil
}

type LoadVTs struct {
	VtsChan chan struct{}
}

func (handler LoadVTs) On(topic string, message []byte) (*connection.SendResponse, error) {
	handler.VtsChan <- struct{}{}
	return nil, nil
}
