package handler

import (
	"encoding/json"
	"fmt"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
)

type FeedHandler struct {
	GetVt   func(oids string) (models.VT, error) // Function to get VTs from feedservice
	Context string
	ID      string
}

func (handler FeedHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	var msg models.GetVT
	if err := json.Unmarshal(message, &msg); err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	if mt.Aggregate == "vt" {
		switch mt.Function {
		case "get":
			vt, err := handler.GetVt(msg.ID)
			if err != nil {
				return nil, err
			}

			return &connection.SendResponse{
				MSG: models.GotVT{
					Identifier: messages.Identifier{
						Message: messages.NewMessage("got.vt", "", msg.GroupID),
						ID:      handler.ID,
					},
					VT: vt,
				},
				Topic: fmt.Sprintf("%s/feed/info", handler.Context),
			}, nil

		}
	}
	return nil, nil
}
