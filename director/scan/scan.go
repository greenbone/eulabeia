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

// Package scan implements handler for scans
package scan

import (
	"fmt"
	"github.com/greenbone/eulabeia/logging"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

var log = logging.Logger()

type scanAggregate struct {
	storage Storage
	target  target.Storage
}

func (t scanAggregate) Start(s cmds.Start) (messages.Event, *info.Failure, error) {
	scan, err := t.storage.Get(s.ID)
	if err != nil {
		return nil, nil, err
	}
	if scan == nil {
		return nil, info.GetFailureResponse(s.Message, "scan", s.ID), nil
	}

	log.Printf("Starting scan (%s) on sensor (%s)", s.ID, scan.Sensor)
	return &cmds.Start{
		Identifier: messages.Identifier{
			Message: messages.NewMessage(fmt.Sprintf("start.scan.%s", scan.Sensor), s.MessageID, s.GroupID),
			ID:      s.ID,
		},
	}, nil, nil
}

func (t scanAggregate) Create(c cmds.Create) (*info.Created, error) {
	scan := models.Scan{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(scan); err != nil {
		return nil, err
	}
	return &info.Created{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("created.scan", c.MessageID, c.GroupID),
			ID:      scan.ID,
		},
	}, nil
}

func (t scanAggregate) Modify(m cmds.Modify) (*info.Modified, *info.Failure, error) {
	var scan *models.Scan
	scan, err := t.storage.Get(m.ID)
	if err != nil {
		return nil, nil, err
	} else if scan == nil {
		log.Printf("Scan %s not found, creating a new one.", m.ID)
		scan = &models.Scan{
			ID: m.ID,
		}
	}
	for k, v := range m.Values {
		switch k {
		case "target_id", "target":
			if str, ok := v.(string); ok {
				target, err := t.target.Get(str)
				if err != nil {
					return nil, info.GetFailureResponse(m.Message, "target", str), nil
				}
				scan.Target = *target
			} else {
				return nil, info.GetFailureResponse(m.Message, "target", "invalid"), nil
			}
		case "temporary":
			if b, ok := v.(bool); ok {
				scan.Temporary = b
			}
		case "finished":
			scan.Finished = handler.InterfaceArrayToStringArray(v)

		}
	}
	if err := t.storage.Put(*scan); err != nil {
		return nil, nil, err
	}

	return &info.Modified{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modified.scan", m.MessageID, m.GroupID),
			ID:      m.ID,
		},
	}, nil, nil

}

func (t scanAggregate) Delete(d cmds.Delete) (*info.Deleted, *info.Failure, error) {
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

func (t scanAggregate) Get(g cmds.Get) (messages.Event, *info.Failure, error) {
	if scan, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if scan == nil {
		return nil, info.GetFailureResponse(g.Message, "scan", g.ID), nil
	} else {
		return &models.GotScan{
			Message: messages.NewMessage("got.scan", g.MessageID, g.GroupID),
			Scan:    *scan,
		}, nil, nil
	}
}

// New returns the type of aggregate as string and Aggregate
func New(storage storage.Json) handler.Container {
	s := scanAggregate{
		storage: NewStorage(storage),
		target:  target.NewStorage(storage)}
	h := handler.FromAggregate("scan", s)
	h.Starter = s
	return h
}
