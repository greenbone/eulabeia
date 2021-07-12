// Package handler contains various message handler for sensors and initializes MQTT connection
package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

var MQTT connection.PubSub

type InvalidCommandError struct {
	cmd string
}

func (err InvalidCommandError) Error() string {
	if err.cmd == "" {
		return "missing command"
	}
	return fmt.Sprintf("invalid command %s used", err.cmd)
}

// Handler for Messages regardings commands running scanner
type CommandHandler struct {
	StartChan chan string
	StopChan  chan string
	VerChan   chan struct{}
	VtsChan   chan struct{}
}

// Implementation for the On method for handling incoming messages via MQTT
func (handler CommandHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	var data cmds.Command
	if err := json.Unmarshal(message, &data); err != nil {
		log.Printf("Sensor cannot read command on Topic %s\n", topic)
		return nil, err
	}

	switch data.Cmd {
	case "start":
		handler.StartChan <- data.ID
	case "stop":
		handler.StopChan <- data.ID
	case "version":
		handler.VerChan <- struct{}{}
	case "loadvts":
		handler.VtsChan <- struct{}{}
	default:
		return nil, &InvalidCommandError{
			cmd: data.Cmd,
		}
	}
	return nil, nil
}

// Handler for Messages which do not regard scans (e.g. get version)
type InfoHandler struct {
	RunChan chan string
	FinChan chan string
}

func (handler InfoHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	var data info.ScanInfo
	if err := json.Unmarshal(message, &data); err != nil {
		log.Printf("Sensor cannot read info on topic %s\n", topic)
		return nil, err
	}

	if data.InfoType == "status" {
		switch data.Info {
		case "running":
			handler.RunChan <- data.ID
		case "finished":
			handler.FinChan <- data.ID
		}
	}
	return nil, nil
}

type RegisterHandler struct {
	RegChan chan struct{}
}

func (handler RegisterHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	handler.RegChan <- struct{}{}
	return nil, nil
}
