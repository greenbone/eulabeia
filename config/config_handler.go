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

package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

// fileExists checks if file exists
func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// file exists
		return true
	}
	return false
}

// findConfigFile looks  up the config file
//
//The look up order is
// 1. given by parameter --config
// 2. custom user config file in home
// 3. in /usr/etc or /etc/ ...
// Returns the path to the first found file
func findConfigFile(path string, module string) (string, error) {
	if path != "" && fileExists((path)) {
		return path, nil
	}
	// look in the default paths
	var defaultPaths = [...]string{
		os.Getenv("HOME") + "/.config",
		"/usr/etc",
		"/etc",
	}
	for _, path := range defaultPaths {
		path += "/" + module + "/config.toml"
		if fileExists(path) {
			// file exists
			return path, nil
		}
	}
	return "", errors.New("no config file found")
}

// Save the Configuration in its current state into the stored file path
func Save(c *Configuration) error {
	bytes, err := toml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.path, bytes, 0644)
}

// Looks for the configuration file and returns a filled Configuration struct
func New(path string, module string) (*Configuration, error) {
	if p, err := findConfigFile(path, module); err != nil {
		return nil, err
	} else {
		c := Configuration{}
		c.path = p

		bytes, err := ioutil.ReadFile(c.path)
		if err != nil {
			return nil, err
		}

		toml.Unmarshal(bytes, &c)
		if c.Context == "" {
			c.Context = module
		}

		return &c, nil
	}
}
