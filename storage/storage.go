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
	"sync"
)

// Json is the base wrapper for json based storage devices
type Json interface {
	Put(string, interface{}) error
	Get(string, interface{}) error
	Delete(string) error
}

// InMemory is a Json implementation to hold structs in memory
type InMemory struct {
	Pretend bool // Set Pretend to true to simulate a found even when it's not previously stored
	sync.RWMutex
	lookup map[string][]byte
}

func (n *InMemory) Put(id string, data interface{}) error {
	n.Lock()
	defer n.Unlock()
	if n.lookup == nil {
		n.lookup = make(map[string][]byte)
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	n.lookup[id] = j
	return nil
}

func (n *InMemory) Delete(id string) error {
	n.Lock()
	defer n.Unlock()
	if n.lookup == nil {
		n.lookup = make(map[string][]byte)
	}
	delete(n.lookup, id)
	return nil
}

func (n *InMemory) Get(id string, v interface{}) error {
	n.RLock()
	defer n.RUnlock()
	if n.Pretend {
		return nil
	}
	if j, ok := n.lookup[id]; ok {
		return json.Unmarshal(j, v)
	}
	return &os.PathError{}

}

// File is a file system based Json implementation
type File struct {
	Dir   string
	Crypt Crypt
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
	if fs.Crypt != nil {
		b, err = fs.Crypt.Encrypt(b)
		if err != nil {
			return err
		}
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

func (fs File) Get(id string, v interface{}) error {
	if p, err := fs.path(id); err == nil {
		b, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		if fs.Crypt != nil {
			b, err = fs.Crypt.Decrypt(b)
			if err != nil {
				return err
			}
		}
		return json.Unmarshal(b, v)
	} else {
		return err
	}
}

// Returns new file system based JSON implementation. The dir is created if it
// does not exist.
func New(dir string, crypt Crypt) (*File, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &File{
		Dir:   dir,
		Crypt: crypt,
	}, nil
}
