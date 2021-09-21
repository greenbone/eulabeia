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

package config

type Certificate struct {
	DefaultKeyFile  string // Path to the default location of the private Key
	DefaultCertFile string // Path to the default location of the private Cert
}

type Connection struct {
	Server  string // Bind address of server in format 133.713.371.337:1337
	Timeout int64  // TODO
}

type ScannerPreferences struct {
	ScanInfoStoreTime   int64  // Time (h) before a scan is considere forgotten
	MaxScan             int64  // Maxi number of parallel scans
	MaxQueuedScans      int64  // Maxi number of scans that can be queued
	Niceness            int64  // Niceness of the openvas Process
	MinFreeMemScanQueue uint64 // Min Memory necessary for a Scan to start

}

type Preferences struct {
	LogLevel string // Loglevel (Debug, Info ...)
	LogFile  string // Path to logfile
}

type Sensor struct {
	Id string // The Id (a uuid) of this sensor
}

type Feedservice struct {
	RedisDbAddress string
}

type Director struct {
	Id          string // The Id (a uuid) of this director
	StoragePath string // The path to store the json into
	KeyFile     string // The path to the private RSA key used to crypt json
}

type Configuration struct {
	Context            string
	Certificate        Certificate
	Connection         Connection
	ScannerPreferences ScannerPreferences
	Preferences        Preferences
	Sensor             Sensor
	Feedservice        Feedservice
	Director           Director
	path               string
}
