// Package target implements handler for targets
package target

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

// Storage is for poutting and getting a models.Target
type Storage interface {
	Put(models.Target) error            // Overrides existing or creates a models.Target
	Get(string) (*models.Target, error) // Gets a models.Target via ID
	Delete(string) error
}

// Depositary stores models.Target as json
type Depositary struct {
	Device storage.Json
}

func (ts Depositary) Put(target models.Target) error {
	return ts.Device.Put(target.ID, target)
}

func (ts Depositary) Delete(id string) error {
	return ts.Device.Delete(id)
}

func (ts Depositary) Get(id string) (*models.Target, error) {
	var target models.Target
	err := ts.Device.Get(id, &target)
	return &target, err
}

type targetAggregate struct {
	storage Storage
}

func (t targetAggregate) Create(c messages.Create) (*messages.Created, error) {
	target := models.Target{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(target); err != nil {
		return nil, err
	}
	return &messages.Created{
		ID:      target.ID,
		Message: messages.NewMessage("created.target", c.MessageID, c.GroupID),
	}, nil
}

func (t targetAggregate) Modify(m messages.Modify) (*messages.Modified, *messages.Failure, error) {
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

	return &messages.Modified{
		ID:      m.ID,
		Message: messages.NewMessage("modified.target", m.MessageID, m.GroupID),
	}, nil, nil

}
func (t targetAggregate) Get(g messages.Get) (interface{}, *messages.Failure, error) {
	if target, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if target == nil {
		return nil, &messages.Failure{
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

func (t targetAggregate) Delete(d messages.Delete) (*messages.Deleted, *messages.Failure, error) {
	if err := t.storage.Delete(d.ID); err != nil {
		return nil, messages.DeleteFailureResponse(d.Message, "target", d.ID), nil
	}
	return &messages.Deleted{
		Message: messages.NewMessage("deleted.target", d.MessageID, d.GroupID),
		ID:      d.ID,
	}, nil, nil
}

// New returns the type of aggregate as string and Aggregate
func New(storage storage.Json) (string, handler.Aggregate) {
	return "target", targetAggregate{storage: Depositary{Device: storage}}
}
