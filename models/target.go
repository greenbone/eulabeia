// Package models contains various aggregate structs
package models

import "github.com/greenbone/eulabeia/messages"

// Target contains all information needed to start a scan
type Target struct {
	ID       string   `json:"id"`       // ID of a Target
	Hosts    []string `json:"hosts"`    // Hosts to scan
	Ports    []string `json:"ports"`    // Ports to scan
	Plugins  []string `json:"plugins"`  // OID of plugins
	Scanner  string   `json:"scanner"`  // Scanner to use (to identify sensor)
	Alive    bool     `json:"alive"`    // Alive when true only alive hosts get scanned
	Parallel bool     `json:"parallel"` // Parallel when true mulitple scans run in parallel
}

// GotTarget is response for get.target
type GotTarget struct {
	messages.Message
	Target
}
