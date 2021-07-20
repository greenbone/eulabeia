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
