package main

import (
	"flag"
	"log"
	"os"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/process"
	"github.com/greenbone/eulabeia/sensor/memory"
)

func main() {
	topic := "eulabeia/+/+/sensor"
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	configuration, err := config.New(*configPath, "eulabeia")
	if err != nil {
		panic(err)
	}

	server := configuration.Connection.Server
	if configuration.Sensor.Id == "" {
		sensor_id, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		configuration.Sensor.Id = sensor_id
	}

	log.Println("Starting sensor")
	c, err := mqtt.New(server, configuration.Sensor.Id, "", "",
		&mqtt.LastWillMessage{
			Topic: "eulabeia/sensor/cmd/director",
			MSG: cmds.Delete{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("delete.sensor", "", ""),
					ID:      configuration.Sensor.Id,
				},
			}})
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	c.Publish("eulabeia/sensor/cmd/director", cmds.Modify{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modify.sensor", "", ""),
			ID:      configuration.Sensor.Id,
		},
		Values: map[string]interface{}{
			"type": "undefined",
		},
	})
	err = c.Subscribe(map[string]connection.OnMessage{
		topic: handler.New(configuration.Context, memory.New()),
	})
	if err != nil {
		panic(err)
	}
	process.Block(c)
}
