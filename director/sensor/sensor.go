// Package sensor implements handler for sensors
package sensor

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

type sensorAggregate struct {
	storage Storage
}

func (t sensorAggregate) Create(c cmds.Create) (*info.Created, error) {
	sensor := models.Sensor{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(sensor); err != nil {
		return nil, err
	}
	return &info.Created{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("created.sensor", c.MessageID, c.GroupID),
			ID:      sensor.ID,
		},
	}, nil
}

func (t sensorAggregate) Modify(m cmds.Modify) (*info.Modified, *info.Failure, error) {
	sensor, err := t.storage.Get(m.ID)
	if err != nil {
		return nil, nil, err
	} else if sensor == nil {
		log.Printf("Scan %s not found, creating a new one.\n", m.ID)
		sensor = &models.Sensor{
			ID: m.ID,
		}
	}
	if f := handler.ModifySetValueOf(sensor, m, nil); f != nil {
		return nil, f, nil
	}
	if err := t.storage.Put(*sensor); err != nil {
		return nil, nil, err
	}

	return &info.Modified{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modified.sensor", m.MessageID, m.GroupID),
			ID:      m.ID,
		},
	}, nil, nil

}
func (t sensorAggregate) Get(g cmds.Get) (interface{}, *info.Failure, error) {
	if sensor, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if sensor == nil {
		return nil, &info.Failure{
			Message: messages.NewMessage("failure.get.sensor", g.MessageID, g.GroupID),
			Error:   fmt.Sprintf("%s not found.", g.ID),
		}, nil
	} else {
		return &models.GotSensor{
			Message: messages.NewMessage("got.sensor", g.MessageID, g.GroupID),
			Sensor:  *sensor,
		}, nil, nil
	}
}

func (t sensorAggregate) Delete(d cmds.Delete) (*info.Deleted, *info.Failure, error) {
	if err := t.storage.Delete(d.ID); err != nil {
		return nil, info.DeleteFailureResponse(d.Message, "sensor", d.ID), nil
	}
	return &info.Deleted{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("deleted.sensor", d.MessageID, d.GroupID),
			ID:      d.ID,
		},
	}, nil, nil
}

// New returns the type of aggregate as string and Aggregate
func New(store storage.Json) handler.Holder {
	return handler.FromAggregate("sensor", sensorAggregate{storage: NewStorage(store)})
}
