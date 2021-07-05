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
	"github.com/greenbone/eulabeia/director/handler/target"
	"github.com/greenbone/eulabeia/messages/handler"
)

func main() {
	clientid := flag.String("clientid", "", "A clientid for the connection")
	configfile := flag.String("config", "", "Use this config file")
	flag.Parse()
	conf_map := config.Load(*configfile)
	server := conf_map.Get("connection.server").(string)

	log.Println("Starting sensor")
	c, err := mqtt.New(server, *clientid, "", "")
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	entityName, aggregateHandler := target.New(target.FileStorage{StorageDir: "/tmp"})
	mh := handler.New(map[string]handler.Aggregate{
		entityName: aggregateHandler,
	})
	if err != nil {
		log.Panicf("Failed to create handler: %s", err)
	}
	err = c.Subscribe(map[string]connection.OnMessage{"greenbone.target": mh})
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
