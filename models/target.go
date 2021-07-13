package models

import (
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/info"
)

// Target contains all information needed to start a scan
type Target struct {
	ID          string            `json:"id"`          // ID of a Target
	Hosts       []string          `json:"hosts"`       // Hosts to scan
	Ports       []string          `json:"ports"`       // Ports to scan
	Plugins     []string          `json:"plugins"`     // OID of plugins
	Sensor      string            `json:"sensor"`      // Sensor to use
	Alive       bool              `json:"alive"`       // Alive when true only alive hosts get scanned
	Parallel    bool              `json:"parallel"`    // Parallel when true mulitple scans run in parallel
	Exclude     []string          `json:"exclude"`     // Exclude hosts from a scan
	Credentials map[string]string `json:"credentials"` // Credentials to login into a target
}

// GotTarget is response for get.target
type GotTarget struct {
	messages.Message
	info.EventType
	Target
}
