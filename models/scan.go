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

// Scan contains Target as well as volatile information for a specific scan
type Scan struct {
	Target
	ID        string   `json:"id"`        // ID of a Scan
	Finished  []string `json:"exclude"`   // Finished hosts from previous scan progress
	Temporary bool     `json:"temporary"` // Temporary defines if the Target as well as the Scan should be deleted when a scan finished
}

// GotScan is a response for get.scan
type GotScan struct {
	messages.Message
	info.EventType
	Scan
}

// GotMemory is the response on get.memory and contains memory information
//
// GotMemory is needed to actually start a scan since only sensor which
// sufficient
// memory should be started
type GotMemory struct {
	messages.Message
	info.EventType
	ID     string `json:"id"`     // Contains the ID from get event, usually sensor use the scanid
	Total  string `json:"total"`  // Total memory in bytes available
	Used   string `json:"used"`   // Used memory in bytes
	Cached string `json:"cached"` //Cached memory in bytes
	Free   string `json:"free"`   // Free memory in bytes
}
