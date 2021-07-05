package target

import (
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

// Storage is for poutting and getting a models.Target
type Storage interface {
	Put(models.Target) error            // Overrides existing or creates a models.Target
	Get(string) (*models.Target, error) // Gets a models.Target via ID
	Delete(string) error
}

// depositary stores models.Target as json
type depositary struct {
	device storage.Json
}

func (ts depositary) Put(target models.Target) error {
	return ts.device.Put(target.ID, target)
}

func (ts depositary) Delete(id string) error {
	return ts.device.Delete(id)
}

func (ts depositary) Get(id string) (*models.Target, error) {
	var target models.Target
	err := ts.device.Get(id, &target)
	return &target, err
}

func NewStorage(device storage.Json) Storage {
	return depositary{
		device: device,
	}
}
