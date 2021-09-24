package handler

import (
	"encoding/json"
	"fmt"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/models"
)

// FeedHandler handles incoming request regarding the feed
type FeedHandler struct {
	GetVT         func(oids string) (models.VT, error)             // Get VT metadata from feedservice
	ResolveFilter func(filter []models.VTFilter) ([]string, error) //
	Context       string
	ID            string
}

func (handler FeedHandler) On(topic string, message []byte) (*connection.SendResponse, error) {
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
			if err := json.Unmarshal(message, &msg); err != nil {
				return nil, err
			}
			vt, err := handler.GetVT(msg.ID)
			if err != nil {
				return nil, err
			}

			return &connection.SendResponse{
				MSG: models.GotVT{
					Message: messages.NewMessage("got.vt", "", msg.GroupID),
					VT:      vt,
				},
				Topic: fmt.Sprintf("%s/vt/info", handler.Context),
			}, nil
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
					Message: messages.NewMessage("resolved.vt", "", msg.GroupID),
					OIDs:    oids,
				},
				Topic: fmt.Sprintf("%s/vt/info", handler.Context),
			}, nil

		}
	}
	return nil, nil
}
