// Package scan implements handler for scans
package scan

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

type scanAggregate struct {
	storage     Storage
	target      target.Storage
	sensorTopic string
}

func (t scanAggregate) Start(s cmds.Start) (interface{}, *messages.Failure, error) {
	scan, err := t.storage.Get(s.ID)
	if err != nil {
		return nil, nil, err
	}
	if scan == nil {
		return nil, messages.GetFailureResponse(s.Message, "scan", s.ID), nil
	}

	return &connection.SendResponse{

		MSG: &cmds.Start{
			Identifier: messages.Identifier{
				Message: messages.NewMessage(fmt.Sprintf("start.scan.%s", scan.Sensor), s.MessageID, s.GroupID),
				ID:      s.ID,
			},
		},
		Topic: t.sensorTopic,
	}, nil, nil
}

func (t scanAggregate) Create(c cmds.Create) (*messages.Created, error) {
	scan := models.Scan{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(scan); err != nil {
		return nil, err
	}
	return &messages.Created{
		ID:      scan.ID,
		Message: messages.NewMessage("created.scan", c.MessageID, c.GroupID),
	}, nil
}

func (t scanAggregate) Modify(m cmds.Modify) (*messages.Modified, *messages.Failure, error) {
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
	applyTargetID := func(k string, v interface{}) error {
		switch k {
		case "target_id":
			if str, ok := v.(string); ok {
				target, err := t.target.Get(str)
				if err != nil {
					return err
				}
				if target == nil {
					return fmt.Errorf("target %s not found", str)
				}
				scan.Target = *target
				return nil
			} else {
				return fmt.Errorf("[%T] %v is not a target ID", v, v)
			}
		default:
			return fmt.Errorf("%s is unknown", k)
		}
	}
	if f := handler.ModifySetValueOf(scan, m, applyTargetID); f != nil {
		return nil, f, nil
	}
	if err := t.storage.Put(*scan); err != nil {
		return nil, nil, err
	}

	return &messages.Modified{
		ID:      m.ID,
		Message: messages.NewMessage("modified.scan", m.MessageID, m.GroupID),
	}, nil, nil

}

func (t scanAggregate) Delete(d cmds.Delete) (*messages.Deleted, *messages.Failure, error) {
	if err := t.storage.Delete(d.ID); err != nil {
		return nil, messages.DeleteFailureResponse(d.Message, "target", d.ID), nil
	}
	return &messages.Deleted{
		Message: messages.NewMessage("deleted.target", d.MessageID, d.GroupID),
		ID:      d.ID,
	}, nil, nil
}

func (t scanAggregate) Get(g cmds.Get) (interface{}, *messages.Failure, error) {
	if scan, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if scan == nil {
		return nil, messages.GetFailureResponse(g.Message, "scan", g.ID), nil
	} else {
		return &models.GotScan{
			Message: messages.NewMessage("got.scan", g.MessageID, g.GroupID),
			Scan:    *scan,
		}, nil, nil
	}
}

// New returns the type of aggregate as string and Aggregate
func New(sensorTopic string, storage storage.Json) handler.Holder {
	s := scanAggregate{
		sensorTopic: sensorTopic,
		storage:     NewStorage(storage),
		target:      target.NewStorage(storage)}
	h := handler.FromAggregate("scan", s)
	h.Starter = s
	return h
}
