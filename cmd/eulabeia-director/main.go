package main

import (
	"flag"
	"log"

	"github.com/greenbone/eulabeia/config"
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
	clientid := flag.String("clientid", "", "A clientid for the connection")
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	configuration := config.New(*configPath, "eulabeia")
	server := configuration.Connection.Server

	log.Println("Starting director")
	client, err := mqtt.New(server, *clientid, "", "", nil)
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = client.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	device := storage.File{Dir: "/tmp/"}
	context := configuration.Context
	err = client.Subscribe(map[string]connection.OnMessage{
		"eulabeia/sensor/cmd/director": handler.New(context, sensor.New(device)),
		"eulabeia/target/cmd/director": handler.New(context, target.New(device)),
		"eulabeia/scan/cmd/director":   handler.New(context, scan.New(device)),
	})
	if err != nil {
		panic(err)
	}

	process.BlockAndClose(client)
}
