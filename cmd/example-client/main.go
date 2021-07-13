package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/process"
	"github.com/tidwall/gjson"
)

type OnEvent interface {
	On(string, []byte) (interface{}, error)
}

type ExampleHandler struct {
	handler []OnEvent
}

func (e ExampleHandler) On(topic string, msg []byte) (*connection.SendResponse, error) {
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
	modifyMSGChan chan cmds.Modify
}

const topic = "eulabeia/+/#"

func (oct OnCreatedTarget) On(messageType string, message []byte) (interface{}, error) {
	if messageType != "created.target" {
		return nil, nil
	}
	var created info.Created
	if err := json.Unmarshal(message, &created); err != nil {
		return nil, err
	}
	modify := cmds.Modify{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modify.target", "", created.GroupID),
			ID:      created.ID,
		},
		Values: map[string]interface{}{
			"hosts":   []string{"localhorst", "nebenan"},
			"plugins": []string{"someoids"},
			"credentials": map[string]string{
				"username": "admin",
				"password": "admin",
			},
		},
	}
	if err := oct.publisher.Publish("eulabeia/target/cmd/director", modify); err != nil {
		return nil, err
	}
	oct.modifyMSGChan <- modify
	return nil, nil
}

type OnModifiedTarget struct {
	publisher     connection.Publisher
	modifyMSGChan chan cmds.Modify
}

func (omt OnModifiedTarget) On(messageType string, message []byte) (interface{}, error) {
	if messageType != "modified.target" {
		return nil, nil
	}
	original, ok := <-omt.modifyMSGChan
	if !ok {
		return nil, errors.New("closed modify channel")
	}
	var modified info.Modified
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
	omt.publisher.Publish("eulabeia/target/cmd/director", cmds.Get{
		Identifier: messages.Identifier{

			Message: messages.NewMessage("get.target", "", ""),
			ID:      original.ID,
		},
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
	topic := "greenbone.sensor"
	clientid := flag.String("clientid", "", "A clientid for the connection")
	configPath := flag.String("config", "", "Path to config file, default: search for config file in TODO")
	flag.Parse()
	configuration := config.New(*configPath, "eulabeia")
	server := configuration.Connection.Server

	log.Println("Starting example client")
	c, err := mqtt.New(server, *clientid+uuid.NewString(), "", "", nil)
	if err != nil {
		log.Panicf("Failed to create MQTT: %s", err)
	}
	err = c.Connect()
	if err != nil {
		log.Panicf("Failed to connect: %s", err)
	}
	err = c.Publish("eulabeia/target/cmd/director", cmds.Create{
		Message: messages.Message{
			Type:      "create.target",
			Created:   7774,
			MessageID: "1",
			GroupID:   "12",
		},
	})
	if err != nil {
		log.Panicf("Failed to publish: %s", err)
	}
	modifyChan := make(chan cmds.Modify, 1)
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
	process.Block(c)
}
