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
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/sensor/handler/memory"
)

func main() {
	server := flag.String("server", "localhost:1883", "A clientid for the connection")
	clientid := flag.String("clientid", "", "A clientid for the connection")
	flag.Parse()

	log.Println("Starting sensor")
	c, err := mqtt.New(*server, *clientid, "", "")
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	mh := handler.New(handler.FromGetter(memory.New()))
	err = c.Subscribe(map[string]connection.OnMessage{"greenbone.sensor": mh})
	if err != nil {
		panic(err)
	}
	ic := make(chan os.Signal, 1)
	signal.Notify(ic, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP)
	<-ic
	fmt.Println("signal received, exiting")
	if c != nil {
		err = c.Close()
		if err != nil {
			log.Fatalf("failed to send Disconnect: %s", err)
		}
	}
}
