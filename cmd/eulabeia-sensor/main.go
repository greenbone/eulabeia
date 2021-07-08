package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/sensor/memory"
)

func main() {
	topic := "greenbone.sensor"
	server := flag.String("server", "localhost:1883", "A clientid for the connection")
	clientid := flag.String("clientid", "", "A clientid for the connection")
	sensorID := flag.String("sensorID", "bla", "A sensorID for the registration")
	flag.Parse()

	log.Println("Starting sensor")
	c, err := mqtt.New(*server, *clientid+uuid.NewString(), "", "",
		&mqtt.LastWillMessage{
			Topic: topic,
			MSG: cmds.Delete{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("delete.sensor", "", ""),
					ID:      *sensorID,
				},
			}})
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	c.Publish(topic, cmds.Modify{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modify.sensor", "", ""),
			ID:      *sensorID,
		},
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
