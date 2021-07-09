// Package handler contains various message handler for sensors and initializes MQTT connection
package handler

import (
	"github.com/greenbone/eulabeia/messages"
	"encoding/json"
	"log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
)

var MQTT connection.PubSub

type InvalidCommandError struct {
	cmd string
}

func (err InvalidCommandError) Error() string {
	if err.cmd == ""
		return "missing command"
	return fmt.Sprintf("invalid command %s used", err.cmd)
}

// Handler for Messages regardings commands running scanner
type CommandHandler struct {
	startChan chan string
	stopChan  chan string
	verChan   chan struct{}
	vtsChan   chan struct{}
}


// Implementation for the On method for handling incoming messages via MQTT
func (handler CommandHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	var data messages.Command
	if err := json.Unmarshal(message, &data); err != nil {
		log.Printf("Sensor cannot read command on Topic %s\n", topic)
		return nil, err
	}

	switch data.Cmd {
	case: "start"
		handler.startChan <- data.ID
	case: "stop"
		handler.stopChan <- data.ID
	case: "version"
		handler.verChan <- struct{}{}
	case: "loadvts"
		handler.vtsChan <- struct{}{}
	default:
		return nil, &InvalidCommandError {
			cmd: data.Cmd
		}
	}
	return nil, nil
}

// Handler for Messages which do not regard scans (e.g. get version)
type InfoHandler struct {
	runChan   chan string
	finChan   chan string
}

func (handler OpenVASHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	var data messages.ScanInfo
	if err := json.Unmarshal(message, &data); err != nil {
		log.Printf("Sensor cannot read info on topic %s\n", topic)
		return nil, err
	}

	if data.InfoType == "status" {
		switch data.Info{
		case "running":
			InfoHandler.runChan <- data.ID
		case "finished":
			InfoHandler.finChan <- data.ID
		}
	}
	return nil, nil
}
