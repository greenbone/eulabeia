// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
