package main

import (
	"flag"
	"log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/director/scan"
	"github.com/greenbone/eulabeia/director/sensor"
	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/process"
	"github.com/greenbone/eulabeia/storage"
)

func main() {
	server := flag.String("server", "localhost:1883", "A clientid for the connection")
	clientid := flag.String("clientid", "", "A clientid for the connection")
	flag.Parse()

	log.Println("Starting director")
	client, err := mqtt.New(*server, *clientid, "", "", nil)
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = client.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	device := storage.File{Dir: "/tmp/"}
	err = client.Subscribe(map[string]connection.OnMessage{
		"eulabeia/sensor/cmd/director": handler.New(sensor.New(device)),
		"eulabeia/target/cmd/director": handler.New(target.New(device)),
		"eulabeia/scan/cmd/director":   handler.New(scan.New(device)),
	})
	if err != nil {
		panic(err)
	}

	process.Block(client)
}
