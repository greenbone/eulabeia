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

// Got* structures in this file do not respond to any event. Results and host/scan status messages are published in the way they are generated.

// ResultType is used to identify a result
type ResultType string

const (
	HOST_COUNT  ResultType = "HOST_COUNT"  // Result is a host count
	DEADHOST    ResultType = "DEADHOST"    // Result is a dead host
	HOST_START  ResultType = "HOST_START"  // Result is a host start
	HOST_END    ResultType = "HOST_END"    // Result is a host end
	ERRMSG      ResultType = "ERRMSG"      // Result is a error message
	LOG         ResultType = "LOG"         // Result is a log message
	HOST_DETAIL ResultType = "HOST_DETAIL" // Result is a host detail result
	ALARM       ResultType = "ALARM"       // Result is an alarm result
)

// Result
type Result struct {
	ScanId   string     `json:"scan_id"`     // Scan id
	Type     ResultType `json:"result_type"` // Result type
	Host     string     `json:"host"`        // Host's IP the result belongs to
	Hostname string     `json:"hostname"`    // Hostname
	Port     string     `json:"ports"`       // Port scanned
	OID      string     `json:"oid"`         // VT's OID
	Name     string     `json:"name"`        // VT's name
	QOD      string     `json:"qod"`         // Quality of detection
	Score    string     `json:"score"`       // Severity score
	Value    string     `json:"value"`       // The result value
	URI      string     `json:"uri"`         // Location of the vulnerability. Commonly a path to a installed package
}

type GotResult struct {
	messages.Message
	info.EventType
	Result
}

// Host information usefull for the host progress calculation
type HostInfoType string   // host progress message
type HostStatusType string // the different possible status for a host

const (
	HOST_PROGRESS HostInfoType   = "HOST_PROGRESS" // Message contains information about host scan progress
	HOST_DEAD     HostStatusType = "HOST_DEAD"     // Status for a dead host
	HOST_STOPPED  HostStatusType = "HOST_STOPPED"  // Status for a stopped host
	HOST_FINISHED HostStatusType = "HOST_FINISHED" // Status for a finished host
)

type GotHostProgress struct {
	messages.Message
	info.EventType
	ScanId string         `json:"scan_id"`      // Scan id
	Type   HostInfoType   `json:"message_type"` // HostProgress message type
	Host   string         `json:"host"`         // Host's IP the progress belongs to
	Count  string         `json:"current"`      // Current amount of plugins that have been launched against the host
	Max    string         `json:"max"`          // Total plugins to be launched against the host
	Status HostStatusType `json:"status"`       // Host status. E.g. dead, finished.
}
