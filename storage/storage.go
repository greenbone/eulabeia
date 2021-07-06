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
