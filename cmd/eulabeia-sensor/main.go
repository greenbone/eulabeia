package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/sensor/memory"
)

func main() {
	topic := "greenbone.sensor"
	clientid := flag.String("clientid", "", "A clientid for the connection")
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	configuration := config.New(*configPath, "eulabeia")
	configuration.Sensor.Id = uuid.NewString()
	server := configuration.Connection.Server
	sensor_id := configuration.Sensor.Id

	log.Println("Starting sensor")
	c, err := mqtt.New(server, *clientid+sensor_id, "", "",
		&mqtt.LastWillMessage{
			Topic: topic,
			MSG: messages.Delete{
				ID:      sensor_id,
				Message: messages.NewMessage("delete.sensor", "", ""),
			}})
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	c.Publish(topic, messages.Modify{
		Message: messages.NewMessage("modify.sensor", "", ""),
		ID:      sensor_id,
		Values: map[string]interface{}{
			"type": "undefined",
		},
	})
	err = c.Subscribe(map[string]connection.OnMessage{
		topic: handler.New(memory.New()),
	})
	if err != nil {
		panic(err)
	}
	ic := make(chan os.Signal, 1)
	signal.Notify(ic, os.Interrupt, syscall.SIGTERM)
	<-ic
	fmt.Println("signal received, exiting")
	if c != nil {
		err = c.Close()
		if err != nil {
			log.Fatalf("failed to send Disconnect: %s", err)
		}
	}
}
