package sensor

import (
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
	"os"
)

// Storage is for putting and getting a models.Sensor
type Storage interface {
	Put(models.Sensor) error            // Overrides existing or creates a models.Sensor
	Get(string) (*models.Sensor, error) // Gets a models.Sensor via ID
	Delete(string) error
}

// depositary stores models.Sensor as json within a given device.
type depositary struct {
	device storage.Json
}

func (d depositary) Put(sensor models.Sensor) error {
	return d.device.Put(sensor.ID, sensor)
}

func (d depositary) Delete(id string) error {
	return d.device.Delete(id)
}

func (d depositary) Get(id string) (*models.Sensor, error) {
	var sensor models.Sensor
	err := d.device.Get(id, &sensor)
	if _, ok := err.(*os.PathError); ok {
		return nil, nil
	}
	return &sensor, err
}

func NewStorage(device storage.Json) depositary {
	return depositary{
		device: device,
	}
}
