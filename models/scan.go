package models

import "github.com/greenbone/eulabeia/messages"

// Scan contains Target as well as volatile information for a specific scan
type Scan struct {
	Target
	ID      string   `json:"id"`      // ID of a Scan
	Exclude []string `json:"exclude"` // Exclude hosts from scan
}

// GotScan is a response for get.scan
type GotScan struct {
	messages.Message
	Scan
}
