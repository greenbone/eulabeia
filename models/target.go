package models

import "github.com/greenbone/eulabeia/messages"

// Target contains all information needed for a scanner
type Target struct {
	ID      string   `json:"id"`
	Hosts   []string `json:"hosts"`
	Ports   []string `json:"ports"`
	Plugins []string `json:"plugins"`
	Scanner string   `json:"scanner"`
	//Exclude  []string `json:"exclude"` rather needed on scan for start/stop
	Alive    bool `json:"alive"`
	Parallel bool `json:"parallel"`
}

// GotTarget is response for get.target
type GotTarget struct {
	messages.Message
	Target
}
