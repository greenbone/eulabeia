// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handler

import (
	"encoding/json"
	"fmt"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
	"github.com/rs/zerolog/log"
)

// FeedHandler handles incoming request regarding the feed
type FeedHandler struct {
	GetVT         func(oid string) (models.VT, error) // Get VT metadata from feedservice
	GetVTs        func() ([]models.VT, error)
	ResolveFilter func(filter []models.VTFilter) ([]string, error) //
	Context       string
	ID            string
}

func (handler FeedHandler) On(
	topic string,
	message []byte,
) (*connection.SendResponse, error) {
	// determine message type
	var msg messages.Message
	if err := json.Unmarshal(message, &msg); err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	log.Trace().Msgf("[%s] %s", topic, string(message))
	if mt.Aggregate == "vt" {
		switch mt.Function {
		case "get": // Get VT Metadata
			var msg cmds.Get
			var response messages.Event
			if err := json.Unmarshal(message, &msg); err != nil {
				return nil, err
			}
			if msg.ID != "" { // Get Single VT
				vt, err := handler.GetVT(msg.ID)
				if err != nil {
					return messages.EventToResponse(handler.Context, info.GetFailureResponse(msg.Message, msg.ID)), nil
				}
				response = models.GotVT{
					Message: messages.NewMessage("got.vt", "", msg.GroupID),
					VT:      vt,
				}
			} else { // Get all VTs
				vts, err := handler.GetVTs()
				if err != nil {
					return nil, err
				}
				response = models.GotVTs{
					Message: messages.NewMessage("gotall.vt", "", msg.GroupID),
					VTs:     vts,
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
					Message: messages.NewMessage(
						"resolved.vt",
						"",
						msg.GroupID,
					),
					OIDs: oids,
				},
				Topic: fmt.Sprintf("%s/vt/info", handler.Context),
			}, nil

		}
	}
	return nil, nil
}
