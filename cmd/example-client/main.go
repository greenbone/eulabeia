// This example will create and modify
// 1. a target
// 2. a scan
// when a scan has been modified it starts a scan.
// On closing it will send delete target and scan message.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/process"
)


// State is used to indicate 
type State int

const (
	CREATE State = iota
	CREATED
	MODIFY
	MODIFIED
	DELETE
	DELETED
)

type MTID string

const (
	CREATED_TARGET = "created.target"
	CREATED_SCAN = "created.scan"
	MODIFIED_TARGET = "modified.target"
	MODIFIED_SCAN = "modified.scan"
)

type OnDo struct {
	on MTID
	do func(msg []byte) *connection.SendResponse
}

type ExampleHandler struct {
	do map[MTID]func([]byte) *connection.SendResponse
	handled []MTID
}

func (e *ExampleHandler) On(topic string, msg []byte) (*connection.SendResponse, error) {
	mt, err := handler.ParseMessageType(msg)
	if err != nil {
		// In this example we end the program on a unexpected message so that we can
		// reuse it as a smoke test.
		// However in a production environment you want to either log and ignore or
		// just ignore unparseable messages.
		panic(err)
	}
	log.Printf("Got message: %s", mt)
	f, ok := e.do[MTID(mt.String())];
	if !ok {
		log.Fatalf("No handler for %s found", mt)
	}
	response := f(msg)
	e.handled = append(e.handled, MTID(mt.String()))

	return response, nil
}


const topic = "eulabeia/+/#"


func main() {

	server := flag.String("server", "localhost:1883", "A clientid for the connection")
	clientid := flag.String("clientid", "", "A clientid for the connection")
	flag.Parse()
	log.Println("Starting example client")
	c, err := mqtt.New(*server, *clientid+uuid.NewString(), "", "", nil)
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	err = c.Publish("eulabeia/target/cmd/director", cmds.NewCreate("target", ""))
	if err != nil {
		log.Panicf("Failed to publish: %s", err)
	}
	mh := ExampleHandler{
		
	}
	if err != nil {
		log.Panicf("Failed to create handler: %s", err)
	}
	err = c.Subscribe(map[string]connection.OnMessage{topic: &mh})
	if err != nil {
		panic(err)
	}
	process.Block(c)
}
