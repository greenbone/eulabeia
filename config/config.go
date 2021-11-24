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
	Server        string // The server to connect to
	QOS           byte   // Setting the default QOS for most cases it should be 1
	CleanStart    bool   // When set to true the broker will not store session information
	SessionExpiry uint64 // How long a session will be stored; when 0 and CleanStart false it will be one day
	Timeout       int64
	Username      string // Username for authentication
	Password      string // Password used with Username for authentication
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

type Feedservice struct {
	RedisDbAddress string
}

type Director struct {
	Id          string // The Id (a uuid) of this director
	StoragePath string // The path to store the json into
	KeyFile     string // The path to the private RSA key used to crypt json
	VTSensor    string // Sensor used to send get vt messages
}

type Configuration struct {
	Context            string
	Certificate        Certificate
	Connection         Connection
	ScannerPreferences ScannerPreferences
	Preferences        Preferences
	Feedservice        Feedservice
	Director           Director
	path               string
}
