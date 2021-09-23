package vt

import (
	"encoding/json"
	"fmt"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/director/scan"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

type vtHandler struct {
	storage scan.Storage
	context string
	sensor  string
}

func (vt vtHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	// Determine message type
	var msg messages.Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	if mt.Aggregate == "vt" {
		switch mt.Function {
		case "get":
			// send get vt request to corresponding sensor
			var getVT models.GetVT
			err := json.Unmarshal(message, &getVT)
			if err != nil {
				return nil, err
			}
			if err != nil {
				return nil, err
			}

			return &connection.SendResponse{
				Topic: fmt.Sprintf("%s/%s/%s/%s", vt.context, "vt", "cmd", vt.sensor),
				MSG:   getVT,
			}, nil
		}
	}
	return nil, nil
}

func New(storage storage.Json, context string, sensor string) vtHandler {
	return vtHandler{
		storage: scan.NewStorage(storage),
		context: context,
		sensor:  sensor,
	}
}
