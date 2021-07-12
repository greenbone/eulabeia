package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
)

func main() {
	topic := "greenbone.sensor"
	server := flag.String("server", "localhost:1883", "A clientid for the connection")
	clientid := flag.String("clientid", "", "A clientid for the connection")
	sensorID := flag.String("sensorID", "bla", "A sensorID for the registration")
	flag.Parse()

	log.Println("Starting sensor")
	client, err := mqtt.New(*server, *clientid+uuid.NewString(), "", "",
		&mqtt.LastWillMessage{
			Topic: topic,
			MSG: messages.Delete{
				ID:      *sensorID,
				Message: messages.NewMessage("delete.sensor", "", ""),
			}})
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = client.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	client.Publish(topic, messages.Modify{
		Message: messages.NewMessage("modify.sensor", "", ""),
		ID:      *sensorID,
		Values: map[string]interface{}{
			"type": "undefined",
		},
	})
	// err = client.Subscribe(map[string]connection.OnMessage{
	// 	topic: handler.New(memory.New()),
	// })
	if err != nil {
		panic(err)
	}
	ic := make(chan os.Signal, 1)
	signal.Notify(ic, os.Interrupt, syscall.SIGTERM)
	<-ic
	fmt.Println("signal received, exiting")
	if client != nil {
		err = client.Close()
		if err != nil {
			log.Fatalf("failed to send Disconnect: %s", err)
		}
	}
}
