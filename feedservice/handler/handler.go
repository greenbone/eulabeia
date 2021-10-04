package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
)

// FeedHandler handles incoming request regarding the feed
type FeedHandler struct {
	GetVT         func(msg cmds.Get) (models.VT, *info.Failure, error) // Get VT metadata from feedservice
	ResolveFilter func(filter []models.VTFilter) ([]string, error)     //
	Context       string
	ID            string
}

func (handler FeedHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
	log.Printf("Got VT request: %s\n", message)
	// determine message type
	var msg messages.Message
	if err := json.Unmarshal(message, &msg); err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	if mt.Aggregate == "vt" {
		switch mt.Function {
		case "get": // Get single VT metadata
			var msg cmds.Get
			var response messages.Event
			if err := json.Unmarshal(message, &msg); err != nil {
				return nil, err
			}
			vt, f, err := handler.GetVT(msg)
			if err != nil {
				return nil, err
			}
			if f != nil {
				response = f
			} else {
				response = models.GotVT{
					Message: messages.NewMessage("got.vt", "", msg.GroupID),
					VT:      vt,
				}
			}
			return messages.EventToResponse(handler.Context, response), nil
		case "resolve": // Get OIDs that matches given filter
			var msg models.ResolveFilter
			if err := json.Unmarshal(message, &msg); err != nil {
				return nil, err
			}

			oids, err := handler.ResolveFilter(msg.Filter)
			if err != nil {
				return nil, err
			}

			return &connection.SendResponse{
				MSG: models.ResolvedFilter{
					Identifier: messages.Identifier{
						Message: messages.NewMessage("resolved.vt", "", msg.GroupID),
						ID:      msg.ID,
					},
					OIDs: oids,
				},
				Topic: fmt.Sprintf("%s/vt/info", handler.Context),
			}, nil

		}
	}
	return nil, nil
}
