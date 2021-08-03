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

// Package target implements handler for targets
package target

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

type targetAggregate struct {
	storage Storage
}

func (t targetAggregate) Create(c cmds.Create) (*info.Created, error) {
	target := models.Target{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(target); err != nil {
		return nil, err
	}
	return &info.Created{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("created.target", c.MessageID, c.GroupID),
			ID:      target.ID,
		},
	}, nil
}

func (t targetAggregate) Modify(m cmds.Modify) (*info.Modified, *info.Failure, error) {
	var target *models.Target
	target, err := t.storage.Get(m.ID)
	if err != nil {
		return nil, nil, err
	} else if target == nil {
		log.Printf("Target %s not found, creating a new one.", m.ID)
		target = &models.Target{
			ID: m.ID,
		}
	}

	for k, v := range m.Values {
		// normalize field name
		nk := strings.Title(k)
		switch nk {
		case "Hosts":
			target.Hosts = handler.InterfaceArrayToStringArray(v)
		case "Ports":
			target.Ports = handler.InterfaceArrayToStringArray(v)
		case "Plugins":
			jsonbody, err := json.Marshal(v)
			if err != nil {
				log.Printf("%s: Given Plugins for target not valid\n", m.ID)
				continue
			}
			plugins := models.VTsList{}
			if err := json.Unmarshal(jsonbody, &plugins); err != nil {
				log.Printf("%s: Given Plugins for target not valid\n", m.ID)
				continue
			}
			target.Plugins = plugins
		case "Sensor":
			if cv, ok := v.(string); ok {
				target.Sensor = cv
			}
		case "Alive":
			if cv, ok := v.(bool); ok {
				target.Alive = cv
			}
		case "Parallel":
			if cv, ok := v.(bool); ok {
				target.Parallel = cv
			}
		case "Exclude":
			target.Exclude = handler.InterfaceArrayToStringArray(v)
		case "Credentials":
			if cv, ok := v.(map[string]interface{}); ok {
				c := make(map[string]map[string]string)
				for k, v := range cv {
					if vs, ok := v.(map[string]interface{}); ok {
						buf := make(map[string]string)
						for k2, v2 := range vs {
							if vs2, ok := v2.(string); ok {
								buf[k2] = vs2
							}
						}
						c[k] = buf
					}
				}
				target.Credentials = c
			}
		}
	}

	if err := t.storage.Put(*target); err != nil {
		return nil, nil, err
	}

	return &info.Modified{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modified.target", m.MessageID, m.GroupID),
			ID:      m.ID,
		},
	}, nil, nil

}
func (t targetAggregate) Get(g cmds.Get) (messages.Event, *info.Failure, error) {
	if target, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if target == nil {
		return nil, info.GetFailureResponse(g.Message, "target", g.ID), nil
	} else {
		return &models.GotTarget{
			Message: messages.NewMessage("got.target", g.MessageID, g.GroupID),
			Target:  *target,
		}, nil, nil
	}
}

func (t targetAggregate) Delete(d cmds.Delete) (*info.Deleted, *info.Failure, error) {
	if err := t.storage.Delete(d.ID); err != nil {
		return nil, info.DeleteFailureResponse(d.Message, "target", d.ID), nil
	}
	return &info.Deleted{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("deleted.target", d.MessageID, d.GroupID),
			ID:      d.ID,
		},
	}, nil, nil
}

// New creates a target aggregate as a handler.Container
func New(storage storage.Json) handler.Container {
	return handler.FromAggregate("target", targetAggregate{storage: NewStorage(storage)})
}
