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
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

// References are the bid, cve or xrefs entries in a nasl script.
//eg. an xref from nmap.nasl is type 'URL' with id: 'https://nmap.org/book/man-performance.html'
type RefType struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// VT's parameters have an ID. The ID 0 (zero) is reserved for the script timeout which is not add to Redis cache as a preference but as script parameter.
type VTParamType struct {
	ID           int    `json:"id"`          // Parameter ID. ID:0 is always to specify a timeout.
	Name         string `json:"name"`        // Parameter Name
	Value        string `json:"value"`       // Parameter Value
	Type         string `json:"type"`        // Parameter Type
	Description  string `json:"description"` // Parameter description
	DefaultValue string `json:"default"`     // Parameter default value
}

// Severities are stored as Tag. Because tag names can not be repeated, only one severity is supported.
// The old tag name for severities are cvss_base and cvss_base_vector. The new extended severities tags has priority over the old format.
type SeverityType struct {
	Vector  string `json:"severity_vector"` // CVSS vector. Supported are CVSSv2 and CVSSv3.x
	Version string `json:"severity_type"`   // CVSS version
	Date    string `json:"severity_date"`   // CVE creation date. Default to VT creation date
	Origin  string `json:"severity_origin"` // Serverity Origin
}

// Strucure to store the information of a single VT.
type VT struct {
	OID                string        `json:"oid"`                // Script OID
	Name               string        `json:"name"`               // Script name
	FileName           string        `json:"filename"`           // Script filename
	RequiredKeys       string        `json:"required_keys"`      // Required keys
	MandatoryKeys      string        `json:"mandatory_keys"`     // Mandatory keys required to run the script
	ExcludedKeys       string        `json:"excluded_keys"`      // Excluded keys
	RequiredPorts      string        `json:"required_ports"`     // Required open ports to run the script
	RequiredUDPPorts   string        `json:"required_udp_ports"` // Required open UDP ports to run the script
	Category           string        `json:"category"`           // Script category (e.g. ACT_ATTACK, ACT_END)
	Family             string        `json:"family"`             // Script family (Debian LSC, Port scanners)
	VTCreationTime     string        `json:"created"`            // Script creation date
	VTModificationTime string        `json:"modified"`           // Last time the script was modified
	Summary            string        `json:"summary"`            // Description of the vulnerability test
	Solution           string        `json:"solution"`           // Script
	SolutionType       string        `json:"solution_type"`      // This information shows possible solutions for the remediation of the vulnerability
	SolutionMethod     string        `json:"solution_method"`    // Script
	Impact             string        `json:"impact"`             // Details about the impact of the vulnerability
	Insight            string        `json:"insight"`            // Some more details about the vulnerability
	Affected           string        `json:"affected"`           // Script
	Vuldetect          string        `json:"vuldetect"`          // Description on the method used to detect the vulnerability
	QoDType            string        `json:"qod_type"`           // Quality of detection
	QoDValue           string        `json:"qod"`                // Quality of detection as percentage
	References         []RefType     `json:"references"`         // See above RefType
	VTParameters       []VTParamType `json:"vt_parameters"`      // See VTParamType
	VTDependencies     []string      `json:"vt_dependencies"`    // List of plugin's filenames which a VT depends on.-
	Severity           SeverityType  `json:"severety"`           // Script severity. See SeverityType.
}

type GetVT struct {
	cmds.EventType
	messages.Identifier
}

type GotVT struct {
	info.EventType
	messages.Identifier
	VT VT `json:"vt"`
}

// SingleVT contains a single VT and its preferences
type SingleVT struct {
	OID         string                 `json:"oid"`
	PrefsByID   map[int]interface{}    `json:"prefs_by_id"`
	PrefsByName map[string]interface{} `json:"prefs_by_name"`
}

// VTsList list to support multiple VTs with own preferences
type VTsList struct {
	Single []SingleVT        `json:"single_vts"`
	Group  map[string]string `json:"vt_groups"`
}

// Target contains all information needed to start a scan
type Target struct {
	ID          string                       `json:"id"`            // ID of a Target
	Hosts       []string                     `json:"hosts"`         // Hosts to scan
	Ports       []string                     `json:"ports"`         // Ports to scan
	Plugins     VTsList                      `json:"plugins"`       // OID of plugins
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
