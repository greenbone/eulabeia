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
