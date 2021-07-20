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
	"os"
	"strings"
)

func server(server string, c *Configuration) {
	c.Connection.Server = server
}

// lookup table that binds an environment variable to a function that overrides the configuration variable in the config file struct
var lookup = map[string]func(string, *Configuration){
	"MQTT_SERVER": server,
}

// OverrideViaENV overrides configuration settings with environment variables, if they are set.
//
// Uses a defined lookup table to identify and get environment variables to override.
func OverrideViaENV(c *Configuration) {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if f, ok := lookup[pair[0]]; ok {
			f(pair[1], c)
		}
	}
}
