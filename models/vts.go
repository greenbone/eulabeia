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

// This file define the structures to represent a VT element.

package models

// References are the bid, cve or xrefs entries in a nasl script.
//eg. an xref from nmap.nasl is type 'URL' with id: 'https://nmap.org/book/man-performance.html'
type RefType struct {
	Type string
	ID   string
}

// VT's parameters have an ID. The ID 0 (zero) is reserved for the script timeout which is not add to Redis cache as a preference but as script parameter.
type VTParamType struct {
	ParameterID           int    // Parameter ID. ID:0 is always to specify a timeout.
	ParameterName         string // Parameter Name
	ParameterValue        string // Parameter Value
	ParameterType         string // Parameter Type
	ParameterDescription  string // Parameter description
	ParameterDefaultValue string // Parameter default value
}

// Severities are stored as Tag. Because tag names can not be repeated, only one severity is supported.
// The old tag name for severities are cvss_base and cvss_base_vector. The new extended severities tags has priority over the old format.
type SeverityType struct {
	SeverityVector  string // CVSS vector. Supported are CVSSv2 and CVSSv3.x
	SeverityVersion string // CVSS version
	SeverityDate    string // CVE creation date. Default to VT creation date
	SeverityOrigin  string // Serverity Origin
}

// Strucure to store the information of a single VT.
type VT struct {
	OID                string        // Script OID
	Name               string        // Script name
	FileName           string        // Script filename
	MandatoryKeys      string        // Mandatory keys required to run the script
	ExcludedKeys       string        // Excluded keys
	RequiredPorts      string        // Required open ports to run the script
	RequiredUDPPorts   string        // Required open UDP ports to run the script
	Category           string        // Script category (e.g. ACT_ATTACK, ACT_END)
	Family             string        // Script family (Debian LSC, Port scanners)
	VTCreationTime     string        // Script creation date
	VTModificationTime string        // Last time the script was modified
	Summary            string        // Description of the vulnerability test
	Solution           string        // Script
	SolutionType       string        // This information shows possible solutions for the remediation of the vulnerability
	SolutionMethod     string        // Script
	Impact             string        // Details about the impact of the vulnerability
	Insight            string        // Some more details about the vulnerability
	Affected           string        // Script
	Vuldectect         string        // Description on the method used to detect the vulnerability
	QoDType            string        // Quality of detection
	QoDValue           string        // Quality of detection as percentage
	References         []RefType     // See above RefType
	VTParameters       []VTParamType // See VTParamType
	VTDependencies     []string      // List of plugin's filenames which a VT depends on.-
	Severity           SeverityType  // Script severity. See SeverityType.
}
