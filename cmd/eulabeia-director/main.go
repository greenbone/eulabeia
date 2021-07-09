package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/director/scan"
	"github.com/greenbone/eulabeia/director/sensor"
	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/storage"
)

func main() {
	confHandler := config.ConfigurationHandler{}
	clientid := flag.String("clientid", "", "A clientid for the connection")
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	confHandler.Load(*configPath, "eulabeia")
	server := confHandler.Configuration.Connection.Server

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
		"greenbone.sensor": handler.New(sensor.New(device)),
		"greenbone.director": handler.New(
			target.New(device),
			scan.New("greenbone.sensor", device)),
	})
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
