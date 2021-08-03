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

// Scheduler is a struct containing functionality to control a sensor
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
	commander  openvas.Commander
	context    string

	mqtt connection.PubSub
	id   string
	conf config.ScannerPreferences
}

// loadVTs commands openvas to load VTs into redis
func (sensor *Scheduler) loadVTs() {
	sensor.loadingVTs = true
	defer func() { sensor.loadingVTs = false }()
	log.Printf("Loading VTs into Redis DB...\n")
	err := sensor.ovas.LoadVTsIntoRedis(sensor.commander)
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
	sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
		Identifier: messages.Identifier{
			ID:      scanID,
			Message: messages.NewMessage("scan.status", "", ""),
		},
		Status: "queued",
	})
	return nil
}

// StartScan starts a scan process
func (sensor *Scheduler) StartScan(scanID string) error {
	sensor.mutex.Lock()
	defer sensor.mutex.Unlock()

	if !sensor.queue.Contains(scanID) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}

	if err := sensor.ovas.StartScan(sensor.queue.Front(), int(sensor.conf.Niceness), sensor.sudo, sensor.commander); err != nil {
		return err
	}
	sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
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
		err := sensor.ovas.StopScan(scanID, sensor.sudo, sensor.commander)
		if err == nil {
			sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
				Identifier: messages.Identifier{
					ID:      scanID,
					Message: messages.NewMessage("scan.status", "", ""),
				},
				Status: "stopped",
			})
		} else {
			return err
		}
	}
	return fmt.Errorf("scan ID %s unknown", scanID)
}

// GetVersion publishes the Version of the scanner
func (sensor *Scheduler) GetVersion() error {
	ver, err := sensor.ovas.GetVersion(sensor.commander)
	if err != nil {
		return err
	}
	err = sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Version{
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

	for !sensor.stopped { // Infinite scheduler Loop
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
		sensor.mqtt.Publish(fmt.Sprintf("%s/sensor/cmd/director", sensor.context), cmds.Modify{
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

// Close cleans all queues and OpenVAS processes, sets all scan stats to
// interrupted and stops the scheduler
func (sensor *Scheduler) Close() error {
	sensor.mutex.Lock()
	defer sensor.mutex.Unlock()
	sensor.stopped = true
	log.Print("Cleaning all OpenVAS Processes...\n")
	// Remove all queued scans
	for item, ok := sensor.queue.Dequeue(); ok; item, ok = sensor.queue.Dequeue() {
		sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
			Identifier: messages.Identifier{
				ID:      item,
				Message: messages.NewMessage("scan.status", "", ""),
			},
			Status: "interrupted",
		})
	}
	// Stopping all init processes
	for item, ok := sensor.init.Dequeue(); ok; item, ok = sensor.init.Dequeue() {
		log.Printf("Stopping %s\n", item)
		sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
			Identifier: messages.Identifier{
				ID:      item,
				Message: messages.NewMessage("scan.status", "", ""),
			},
			Status: "interrupted",
		})
		sensor.ovas.StopScan(item, sensor.sudo, sensor.commander)
	}
	// Stopping all running processes
	for item, ok := sensor.running.Dequeue(); ok; item, ok = sensor.running.Dequeue() {
		log.Printf("Stopping %s\n", item)
		sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
			Identifier: messages.Identifier{
				ID:      item,
				Message: messages.NewMessage("scan.status", "", ""),
			},
			Status: "interrupted",
		})
		sensor.ovas.StopScan(item, sensor.sudo, sensor.commander)
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
		fmt.Sprintf("%s/scan/cmd/%s", sensor.context, sensor.id): startStopHandler,
		fmt.Sprintf("%s/scan/info", sensor.context):              statusHandler,
		fmt.Sprintf("%s/sensor/cmd", sensor.context):             vtsHandler,
	}

	err := sensor.mqtt.Subscribe(subMap)
	if err != nil {
		log.Panicf("Sensor cannot subscribe to topics: %s", err)
	}

	sensor.stopped = false
	go sensor.schedule()
}

// NewScheduler creates a new scheduler
func NewScheduler(mqtt connection.PubSub, id string, conf config.ScannerPreferences, context string) *Scheduler {
	return &Scheduler{
		queue:     util.NewQueueList(),
		init:      util.NewQueueList(),
		running:   util.NewQueueList(),
		ovas:      openvas.NewOpenVASScanner(),
		sudo:      openvas.IsSudo(openvas.StdCommander{}),
		mutex:     &sync.Mutex{},
		stopped:   true,
		regChan:   make(chan struct{}),
		commander: openvas.StdCommander{},
		context:   context,

		mqtt: mqtt,
		id:   id,
		conf: conf,
	}
}
