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

// Package storage contains abstract storage implementations
package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Json is the base wrapper for json based storage devices
type Json interface {
	Put(string, interface{}) error
	Get(string, interface{}) error
	Delete(string) error
}

// Noop No operation is a Json implementations without changing an device
type Noop struct {
}

func (n Noop) Put(id string, data interface{}) error {
	_, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return nil
}

func (n Noop) Delete(id string) error {
	return nil
}

func (n Noop) Get(id string, target interface{}) error {
	return nil
}

// File is a file system based Json implementation
type File struct {
	Dir string
}

func (fs File) path(id string) (string, error) {
	if id == "" {
		return "", errors.New("missing ID")
	}
	return fmt.Sprintf("%s/%s", fs.Dir, id), nil
}

func (fs File) Put(id string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if p, err := fs.path(id); err == nil {
		return ioutil.WriteFile(p, b, 0640)
	} else {
		return err
	}
}

func (fs File) Delete(id string) error {
	if p, err := fs.path(id); err == nil {
		return os.Remove(p)
	} else {
		return err
	}
}

func (fs File) Get(id string, target interface{}) error {
	if p, err := fs.path(id); err == nil {
		b, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, target)
	} else {
		return err
	}
}
