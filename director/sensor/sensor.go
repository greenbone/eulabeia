// Package sensor implements handler for sensors
package sensor

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

type sensorAggregate struct {
	storage Storage
}

func (t sensorAggregate) Create(c messages.Create) (*messages.Created, error) {
	sensor := models.Sensor{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(sensor); err != nil {
		return nil, err
	}
	return &messages.Created{
		ID:      sensor.ID,
		Message: messages.NewMessage("created.sensor", c.MessageID, c.GroupID),
	}, nil
}

func (t sensorAggregate) Modify(m messages.Modify) (*messages.Modified, *messages.Failure, error) {
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

	return &messages.Modified{
		ID:      m.ID,
		Message: messages.NewMessage("modified.sensor", m.MessageID, m.GroupID),
	}, nil, nil

}
func (t sensorAggregate) Get(g messages.Get) (interface{}, *messages.Failure, error) {
	if sensor, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if sensor == nil {
		return nil, &messages.Failure{
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

func (t sensorAggregate) Delete(d messages.Delete) (*messages.Deleted, *messages.Failure, error) {
	if err := t.storage.Delete(d.ID); err != nil {
		return nil, messages.DeleteFailureResponse(d.Message, "sensor", d.ID), nil
	}
	return &messages.Deleted{
		Message: messages.NewMessage("deleted.sensor", d.MessageID, d.GroupID),
		ID:      d.ID,
	}, nil, nil
}

// New returns the type of aggregate as string and Aggregate
func New(store storage.Json) handler.Holder {
	return handler.FromAggregate("sensor", sensorAggregate{storage: NewStorage(store)})
}
