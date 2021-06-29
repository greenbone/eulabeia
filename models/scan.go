package models

import "github.com/greenbone/eulabeia/messages"

// Scan contains Target as well as volatile information for a specific scan
type Scan struct {
	Target
	ID      string   `json:"id"`      // ID of a Scan
	Exclude []string `json:"exclude"` // Exclude hosts from scan
}

// Sensor contains registered sensors
//
// A sensor is starting and stopping the actual scan process
type Sensor struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// GotSensor is a response for get.sensor
type GotSensor struct {
	messages.Message
	Sensor
}

// GotScan is a response for get.scan
type GotScan struct {
	messages.Message
	Scan
}

// GotMemory is the response on get.memory and contains memory information
//
// GotMemory is needed to actually start a scan since only sensor which sufficient
// memory should be started
type GotMemory struct {
	messages.Message
	ID     string `json:"id"`     // Contains the ID from get event, usually sensor use the scanid
	Total  string `json:"total"`  // Total memory in bytes available
	Used   string `json:"used"`   // Used memory in bytes
	Cached string `json:"cached"` //Cached memory in bytes
	Free   string `json:"free"`   // Free memory in bytes
}
