// Package handler contains various message handler for sensors and initializes MQTT connection
package handler

import (
	"encoding/json"
	"log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
)

var MQTT connection.PubSub

// Handler for the scheduler
type SchedulerHandler struct {
	Channel chan string
}

// Implementation for the On method for handling incoming messages via MQTT
func (s SchedulerHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	var data map[string]string
	if err := json.Unmarshal(message, data); err != nil {
		log.Printf("Sensor cannot read data on Topic %s: %s\n", topic, err)
		return nil, nil
	}

	scan, ok := data["scan_id"]
	if !ok {
		log.Printf("Unable to get Scan ID from message on topic %s.\n", topic)
		return nil, nil
	}

	s.Channel <- scan
	return nil, nil
}

// Setup MQTT for message handling
func init() {
	var err error
	MQTT, err = mqtt.New("localhost:1883", "sensor", "", "", &mqtt.LastWillMessage{
		Topic: "sensor.status",
		MSG: map[string]string{
			"error": "sensor got disconnected from MQTT",
		},
	})
	if err != nil {
		log.Panicf("Unable to connect to MQTT Broker: %s", err)
	}
}
