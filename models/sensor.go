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
// Sensor contains registered sensors

package models

import (
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/info"
)

// A sensor is starting and stopping the actual scan process
type Sensor struct {
	ID   string `json:"id"`   // ID of the sensor; can be freely chosen as long as there is no collision
	Type string `json:"type"` // Type of sensor; e.g. openvas.
}

// GotSensor is a response for get.sensor
type GotSensor struct {
	messages.Message
	info.EventType
	Sensor
}
