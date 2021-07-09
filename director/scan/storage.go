package scan

import (
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
	"os"
)

// Storage is for putting and getting a models.Scan
type Storage interface {
	Put(models.Scan) error            // Overrides existing or creates a models.Scan
	Get(string) (*models.Scan, error) // Gets a models.Scan via ID
	Delete(string) error
}

// depositary stores models.Scan as json within a given device.
type depositary struct {
	device storage.Json
}

func (d depositary) Put(scan models.Scan) error {
	return d.device.Put(scan.ID, scan)
}

func (d depositary) Delete(id string) error {
	return d.device.Delete(id)
}

func (d depositary) Get(id string) (*models.Scan, error) {
	var scan models.Scan
	err := d.device.Get(id, &scan)
	if _, ok := err.(*os.PathError); ok {
		return nil, nil
	}
	return &scan, err
}

func NewStorage(device storage.Json) depositary {
	return depositary{
		device: device,
	}
}
