// Package scan implements handler for scans
package scan

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

// Storage is for poutting and getting a models.Scan
type Storage interface {
	Put(models.Scan) error            // Overrides existing or creates a models.Scan
	Get(string) (*models.Scan, error) // Gets a models.Scan via ID
	Delete(string) error
}

// Depositary stores models.Scan as json within a given StorageDir
// The filename is a uuid without suffix.
type Depositary struct {
	Device storage.Json
}

func (ts Depositary) Put(scan models.Scan) error {
	return ts.Device.Put(scan.ID, scan)
}

func (ts Depositary) Delete(id string) error {
	return ts.Device.Delete(id)
}

func (ts Depositary) Get(id string) (*models.Scan, error) {
	var scan models.Scan
	err := ts.Device.Get(id, &scan)
	return &scan, err
}

type scanAggregate struct {
	storage     Storage
	target      target.Storage
	sensorTopic string
}

func (t scanAggregate) Start(s messages.Start) (interface{}, *messages.Failure, error) {
	scan, err := t.storage.Get(s.ID)
	if err != nil {
		return nil, nil, err
	}
	if scan == nil {
		return nil, messages.GetFailureResponse(s.Message, "scan", s.ID), nil
	}

	return &connection.SendResponse{
		MSG: &messages.Start{
			Message: messages.NewMessage(fmt.Sprintf("start.scan.%s", scan.Sensor), s.MessageID, s.GroupID),
			ID:      s.ID,
		},
		Topic: t.sensorTopic,
	}, nil, nil
}

func (t scanAggregate) Create(c messages.Create) (*messages.Created, error) {
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

func (t scanAggregate) Modify(m messages.Modify) (*messages.Modified, *messages.Failure, error) {
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

func (t scanAggregate) Delete(d messages.Delete) (*messages.Deleted, *messages.Failure, error) {
	if err := t.storage.Delete(d.ID); err != nil {
		return nil, messages.DeleteFailureResponse(d.Message, "target", d.ID), nil
	}
	return &messages.Deleted{
		Message: messages.NewMessage("deleted.target", d.MessageID, d.GroupID),
		ID:      d.ID,
	}, nil, nil
}

func (t scanAggregate) Get(g messages.Get) (interface{}, *messages.Failure, error) {
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
		storage:     Depositary{Device: storage},
		target:      target.Depositary{Device: storage}}
	h := handler.FromAggregate("scan", s)
	h.Starter = s
	return h
}