package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/tidwall/gjson"
)

type OnEvent interface {
	On(string, []byte) (interface{}, error)
}

type ExampleHandler struct {
	handler []OnEvent
}

func (e ExampleHandler) On(msg []byte) (interface{}, error) {
	messageType := gjson.GetBytes(msg, "message_type")
	for _, h := range e.handler {
		if _, err := h.On(messageType.String(), msg); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

type OnCreatedTarget struct {
	publisher     connection.Publisher
	modifyMSGChan chan messages.Modify
}

const topic = "greenbone.director"

func (oct OnCreatedTarget) On(messageType string, message []byte) (interface{}, error) {
	if messageType != "created.target" {
		return nil, nil
	}
	var created messages.Created
	if err := json.Unmarshal(message, &created); err != nil {
		return nil, err
	}
	modify := messages.Modify{
		Message: messages.NewMessage("modify.target", "", created.GroupID),
		ID:      created.ID,
		Values: map[string]interface{}{
			"hosts":   []string{"localhorst", "nebenan"},
			"plugins": []string{"someoids"},
		},
	}
	if err := oct.publisher.Publish(topic, modify); err != nil {
		return nil, err
	}
	oct.modifyMSGChan <- modify
	return nil, nil
}

type OnModifiedTarget struct {
	publisher     connection.Publisher
	modifyMSGChan chan messages.Modify
}

func (omt OnModifiedTarget) On(messageType string, message []byte) (interface{}, error) {
	if messageType != "modified.target" {
		return nil, nil
	}
	original, ok := <-omt.modifyMSGChan
	if !ok {
		return nil, errors.New("closed modify channel")
	}
	var modified messages.Modified
	if err := json.Unmarshal(message, &modified); err != nil {
		return nil, err
	}
	log.Printf("original message id %v", original.MessageID)
	log.Printf("modified message id %v", modified.MessageID)
	if original.MessageID != modified.MessageID {
		omt.modifyMSGChan <- original
		return nil, nil
	}
	log.Printf("target: %s modified", original.ID)
	omt.publisher.Publish(topic, messages.Get{
		Message: messages.NewMessage("get.target", "", ""),
		ID:      original.ID,
	})
	return nil, nil
}

type OnGotTarget struct{}

func (ogt OnGotTarget) On(messageType string, message []byte) (interface{}, error) {
	if messageType != "got.target" {
		return nil, nil
	}
	log.Printf("Got target:\n%s\n", message)
	return nil, nil
}

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
	err = c.Publish(topic, messages.Create{
		Message: messages.Message{
			MessageType: "create.target",
			Created:     7774,
			MessageID:   "1",
			GroupID:     "12",
		},
	})
	if err != nil {
		log.Panicf("Failed to publish: %s", err)
	}
	modifyChan := make(chan messages.Modify, 1)
	defer close(modifyChan)
	mh := ExampleHandler{
		handler: []OnEvent{
			OnCreatedTarget{publisher: c, modifyMSGChan: modifyChan},
			OnModifiedTarget{publisher: c, modifyMSGChan: modifyChan},
			OnGotTarget{},
		},
	}
	if err != nil {
		log.Panicf("Failed to create handler: %s", err)
	}
	err = c.Subscribe(map[string]connection.OnMessage{topic: mh})
	if err != nil {
		panic(err)
	}
	ic := make(chan os.Signal, 1)
	defer close(ic)
	signal.Notify(ic, os.Interrupt, syscall.SIGTERM)
	<-ic
	log.Println("signal received, exiting")
	if c != nil {
		err = c.Close()
		if err != nil {
			log.Fatalf("failed to send Disconnect: %s", err)
		}
	}
	<-ic
	log.Println("Received message, exiting")
	err = c.Close()
	if err != nil {
		log.Panicf("Error while closing connection")
	}
}
