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

package models

import (
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/info"
)

// Target contains all information needed to start a scan
type Target struct {
	ID          string                       `json:"id"`            // ID of a Target
	Hosts       []string                     `json:"hosts"`         // Hosts to scan
	Ports       []string                     `json:"ports"`         // Ports to scan
	Plugins     []string                     `json:"plugins"`       // OID of plugins
	Sensor      string                       `json:"sensor"`        // Sensor to use
	Alive       bool                         `json:"alive"`         // Alive when true only alive hosts get scanned
	Parallel    bool                         `json:"parallel"`      // Parallel when true mulitple scans run in parallel
	Exclude     []string                     `json:"exclude_hosts"` // Exclude hosts from a scan
	Credentials map[string]map[string]string `json:"credentials"`   // Credentials to login into a target
}

// GotTarget is response for get.target
type GotTarget struct {
	messages.Message
	info.EventType
	Target
}
