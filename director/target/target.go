// Package target implements handler for targets
package target

import (
	"fmt"
	"log"

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
	if f := handler.ModifySetValueOf(target, m, nil); f != nil {
		return nil, f, nil
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
		return nil, &info.Failure{
			Message: messages.NewMessage("failure.get.target", g.MessageID, g.GroupID),
			Error:   fmt.Sprintf("%s not found.", g.ID),
		}, nil
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
