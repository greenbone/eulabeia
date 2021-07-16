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

// Scheduler component of the sensor. This module is responsible for handling request from the director.
package sensor

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/sensor/handler"
	"github.com/greenbone/eulabeia/sensor/scanner/openvas"

	"github.com/greenbone/eulabeia/util"
)

type Scheduler struct {
	queue      *util.QueueList
	init       *util.QueueList
	running    *util.QueueList
	loadingVTs bool
	ovas       *openvas.OpenVASScanner
	sudo       bool
	mutex      *sync.Mutex
	stopped    bool
	regChan    chan struct{}

	mqtt connection.PubSub
	id   string
	conf config.ScannerPreferences
}

// loadVTs commands openvas to load VTs into redis
func (sensor *Scheduler) loadVTs() {
	sensor.loadingVTs = true
	defer func() { sensor.loadingVTs = false }()
	log.Printf("Loading VTs into Redis DB...\n")
	err := sensor.ovas.LoadVTsIntoRedis(openvas.StdCommander{})
	if err != nil {
		log.Panicf("Unable to load VTs into redis: %s", err)
	}
	log.Printf("Loading VTs into Redis DB finished\n")
}

// QueueScan queues a scan
func (sensor *Scheduler) QueueScan(scanID string) error {
	sensor.mutex.Lock()
	defer sensor.mutex.Unlock()
	if sensor.queue.Contains(scanID) || sensor.init.Contains(scanID) || sensor.running.Contains(scanID) {
		return fmt.Errorf("there is already a running scan with the ID %s", scanID)
	}
	sensor.queue.Enqueue(scanID)
	sensor.mqtt.Publish("eulabeia/scan/info", info.Status{
		Identifier: messages.Identifier{
			ID:      scanID,
			Message: messages.NewMessage("scan.status", "", ""),
		},
		Status: "queued",
	})
	return nil
}

func (sensor *Scheduler) StartScan(scanID string) error {
	sensor.mutex.Lock()
	defer sensor.mutex.Unlock()

	if !sensor.queue.Contains(scanID) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}

	if err := sensor.ovas.StartScan(sensor.queue.Front(), int(sensor.conf.Niceness), sensor.sudo, openvas.StdCommander{}); err != nil {
		return err
	}
	sensor.mqtt.Publish("eulabeia/scan/info", info.Status{
		Identifier: messages.Identifier{
			ID:      sensor.queue.Front(),
			Message: messages.NewMessage("scan.status", "", ""),
		},
		Status: "init",
	})

	sensor.queue.RemoveListItem(scanID)
	sensor.init.Enqueue(scanID)
	return nil
}

// ScanRunning moves a scan from the init to the running state
func (sensor *Scheduler) ScanRunning(scanID string) error {
	sensor.mutex.Lock()
	defer sensor.mutex.Unlock()
	if !sensor.init.RemoveListItem(scanID) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}
	sensor.running.Enqueue(scanID)
	return nil
}

func (sensor *Scheduler) ScanFinished(scanID string) error {
	sensor.mutex.Lock()
	defer sensor.mutex.Unlock()
	if !sensor.running.RemoveListItem(scanID) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}
	return nil
}

// StopScan will remove the scan from the queue or invoke a stop scan command to scanner
func (sensor *Scheduler) StopScan(scanID string) error {
	sensor.mutex.Lock()
	defer sensor.mutex.Unlock()
	if sensor.queue.RemoveListItem(scanID) {
		return nil
	}
	if sensor.init.RemoveListItem(scanID) || sensor.running.RemoveListItem(scanID) {
		err := sensor.ovas.StopScan(scanID, sensor.sudo, openvas.StdCommander{})
		if err == nil {
			sensor.mqtt.Publish("eulabeia/scan/info", info.Status{
				Identifier: messages.Identifier{
					ID:      scanID,
					Message: messages.NewMessage("scan.status", "", ""),
				},
				Status: "stopped",
			})
		}
	}
	return fmt.Errorf("scan ID %s unknown", scanID)
}

func (sensor *Scheduler) GetVersion() error {
	ver, err := sensor.ovas.GetVersion(openvas.StdCommander{})
	if err != nil {
		return err
	}
	err = sensor.mqtt.Publish("eulabeia/scan/info", info.Version{
		Identifier: messages.Identifier{
			ID:      "",
			Message: messages.NewMessage("sensor.version", "", ""),
		},
		Version: ver,
	})
	return err
}

// schedule is checking the queue and starts new scans
func (sensor *Scheduler) schedule() {

	sensor.loadVTs()

	for { // Infinite scheduler Loop
		time.Sleep(time.Second)

		if sensor.queue.IsEmpty() {
			continue
		}

		// Check for free scanner slot
		if sensor.conf.MaxScan > 0 && sensor.init.Size()+sensor.running.Size() == int(sensor.conf.MaxScan) {
			log.Printf("Unable to start a scan from queue, Max number of scans reached.\n")
			continue
		}

		// get memory stats and check for memory
		if sensor.conf.MinFreeMemScanQueue > 0 {
			m, err := util.GetAvailableMemory(util.StdMemoryManager{})
			memoryNeeded := m.Bytes + uint64(sensor.init.Size())*sensor.conf.MinFreeMemScanQueue
			if err != nil {
				log.Panicf("Unable to get memory stats: %s\n", err)
			}
			if m.Bytes < memoryNeeded {
				log.Printf("Unable to start scan, not enough memory.\n")
				continue
			}
		}

		// try to run scan process
		if err := sensor.StartScan(sensor.queue.Front()); err != nil {
			log.Printf("%s: unable to start scan: %s", sensor.queue.Front(), err)
		}

	}
}

// register loops until its ID is registrated
func (sensor *Scheduler) register() {
	for { // loop until sensor is registered
		sensor.mqtt.Publish("eulabeia/sensor/cmd/director", cmds.Modify{
			Identifier: messages.Identifier{
				ID:      sensor.id,
				Message: messages.NewMessage("modify.sensor", "", ""),
			},
		})
		select {
		case <-sensor.regChan:
			return
		// Send new registration mqtt message each second
		case <-time.After(time.Second):
		}
	}
}

// Close cleans all OpenVAS processes
func (sensor *Scheduler) Close() error {
	log.Print("Cleaning all OpenVAS Processes...\n")
	// Stopping all init processes
	for item, ok := sensor.init.Dequeue(); ok; {
		sensor.StopScan(item)
	}
	// Stopping all running processes
	for item, ok := sensor.running.Dequeue(); ok; {
		sensor.StopScan(item)
	}
	return nil
}

// Start initializes MQTT handling and starts the scheduler
func (sensor *Scheduler) Start() {
	// Subscribe on Topic to get confirmation about registration
	sensor.mqtt.Subscribe(map[string]connection.OnMessage{
		"eulabeia/sensor/info": handler.Registered{
			RegChan: sensor.regChan,
			ID:      sensor.id,
		},
	})
	// Register Sensor
	sensor.register()
	// MQTT OnMessage Types
	var startStopHandler = handler.StartStop{
		StartFunc: sensor.QueueScan,
		StopFunc:  sensor.StopScan,
	}

	var statusHandler = handler.Status{
		RunFunc: sensor.ScanRunning,
		FinFunc: sensor.ScanFinished,
	}

	var vtsHandler = handler.LoadVTs{
		VtsFunc: sensor.loadVTs,
	}

	// MQTT Subscription Map
	var subMap = map[string]connection.OnMessage{
		fmt.Sprintf("eulabeia/scan/cmd/%s", sensor.id): startStopHandler,
		"eulabeia/scan/info":                           statusHandler,
		"eulabeia/sensor/cmd":                          vtsHandler,
	}

	err := sensor.mqtt.Subscribe(subMap)
	if err != nil {
		log.Panicf("Sensor cannot subscribe to topics: %s", err)
	}

	go sensor.schedule()
	sensor.stopped = false
}

func NewScheduler(mqtt connection.PubSub, id string, conf config.ScannerPreferences) *Scheduler {
	return &Scheduler{
		queue:   util.NewQueueList(),
		init:    util.NewQueueList(),
		running: util.NewQueueList(),
		ovas:    openvas.NewOpenVASScanner(),
		sudo:    openvas.IsSudo(openvas.StdCommander{}),
		mutex:   &sync.Mutex{},
		stopped: true,
		regChan: make(chan struct{}),

		mqtt: mqtt,
		id:   id,
		conf: conf,
	}
}
