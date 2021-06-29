package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/director/handler/scan"
	"github.com/greenbone/eulabeia/director/handler/sensor"
	"github.com/greenbone/eulabeia/director/handler/target"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/storage"
)

func main() {
	server := flag.String("server", "localhost:1883", "A clientid for the connection")
	clientid := flag.String("clientid", "", "A clientid for the connection")
	flag.Parse()

	log.Println("Starting director")
	c, err := mqtt.New(*server, *clientid, "", "", nil)
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	device := storage.File{Dir: "/tmp/"}
	err = c.Subscribe(map[string]connection.OnMessage{
		"greenbone.sensor": handler.New(handler.FromAggregate(sensor.New(device))),
		"greenbone.director": handler.New(
			handler.FromAggregate(target.New(device)),
			handler.FromAggregate(scan.New(device))),
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
