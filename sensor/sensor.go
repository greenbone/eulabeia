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

// Scheduler component of the sensor. This module is responsible for handling
// request from the director.
package sensor

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/messages/info"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/sensor/scanner/openvas"

	"github.com/greenbone/eulabeia/util"
)

// Scheduler is a struct containing functionality to control a sensor
type Scheduler struct {
	queue         *util.QueueList                 // queued scans
	init          *util.QueueList                 // scans that currently initializes
	running       *util.QueueList                 // scans that are currently running
	loadingVTs    bool                            // marks that VTs are currently loading
	ovas          *openvas.OpenVASScanner         // openvas
	sudo          bool                            // sudo rights
	sync.Mutex                                    // thread safe hadnling when moving scan IDs between lists
	stopped       bool                            // marks that the sensor is stopped
	regChan       chan struct{}                   // channel for succesful registation
	commander     openvas.Commander               // commander used for openvas
	context       string                          // context used for mqtt
	interruptChan chan string                     // channel for signaling interrupted scans
	out           chan<- *connection.SendResponse // Channel to send messages
	id            string                          // ID of the sensor
	conf          config.ScannerPreferences       // config file
}

// loadVTs commands openvas to load VTs into redis
func (sensor *Scheduler) loadVTs() {
	sensor.loadingVTs = true
	defer func() { sensor.loadingVTs = false }()
	log.Printf("Loading VTs into Redis DB...\n")
	err := sensor.ovas.LoadVTsIntoRedis(sensor.commander)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to load VTs into redis: %s", err)
	}
	log.Printf("Loading VTs into Redis DB finished\n")
}

// QueueScan queues a scan
func (sensor *Scheduler) QueueScan(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()
	if sensor.queue.Contains(scanID) || sensor.init.Contains(scanID) ||
		sensor.running.Contains(scanID) {
		return fmt.Errorf(
			"there is already a running scan with the ID %s",
			scanID,
		)
	}
	sensor.queue.Enqueue(scanID)
	sensor.out <- messages.EventToResponse(sensor.context, info.Status{
		Identifier: messages.Identifier{
			ID:      scanID,
			Message: messages.NewMessage("status.scan", "", ""),
		},
		Status: "queued",
	})
	return nil
}

// StartScan starts a scan process
func (sensor *Scheduler) StartScan(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()

	if !sensor.queue.Contains(scanID) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}

	if err := sensor.ovas.StartScan(sensor.queue.Front(), int(sensor.conf.Niceness), sensor.sudo, sensor.commander); err != nil {
		return err
	}
	sensor.out <- messages.EventToResponse(sensor.context, info.Status{
		Identifier: messages.Identifier{
			ID:      sensor.queue.Front(),
			Message: messages.NewMessage("status.scan", "", ""),
		},
		Status: "init",
	})

	sensor.queue.RemoveListItem(scanID)
	sensor.init.Enqueue(scanID)
	return nil
}

// ScanRunning moves a scan from the init to the running state
func (sensor *Scheduler) ScanRunning(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()
	if !sensor.init.RemoveListItem(scanID) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}
	sensor.running.Enqueue(scanID)
	return nil
}

func (sensor *Scheduler) ScanFinished(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()
	if !sensor.running.RemoveListItem(scanID) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}
	sensor.ovas.ScanFinished(scanID)
	return nil
}

// StopScan will remove the scan from the queue or invoke a stop scan command to
// scanner
func (sensor *Scheduler) StopScan(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()
	if sensor.queue.RemoveListItem(scanID) {
		return nil
	}
	if sensor.init.RemoveListItem(scanID) ||
		sensor.running.RemoveListItem(scanID) {
		err := sensor.ovas.StopScan(scanID, sensor.sudo, sensor.commander)
		if err == nil {
			sensor.out <- messages.EventToResponse(sensor.context, info.Status{
				Identifier: messages.Identifier{
					ID:      scanID,
					Message: messages.NewMessage("status.scan", "", ""),
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
	sensor.out <- messages.EventToResponse(sensor.context, info.Version{
		Identifier: messages.Identifier{
			ID:      "",
			Message: messages.NewMessage("sensor.version", "", ""),
		},
		Version: ver,
	})
	return nil
}

// interruptScan removes scan from list and publishes a status MSG
func (sensor *Scheduler) interruptScan(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()

	if sensor.init.RemoveListItem(scanID) ||
		sensor.running.RemoveListItem(scanID) {
		sensor.out <- messages.EventToResponse(sensor.context, info.Status{
			Identifier: messages.Identifier{
				ID:      scanID,
				Message: messages.NewMessage("status.scan", "", ""),
			},
			Status: "interrupted",
		})
		return nil
	}
	return fmt.Errorf("scan %s unknown", scanID)
}

// schedule is checking the queue and starts new scans
func (sensor *Scheduler) schedule() {

	sensor.loadVTs()

	for !sensor.stopped { // Infinite scheduler Loop
		time.Sleep(time.Second)

		// Handle interrupted scans
		select {
		case scanID := <-sensor.interruptChan:
			if err := sensor.interruptScan(scanID); err != nil {
				log.Printf("Unable to interrupt scan: %s", err)
			}
		default:
		}

		if sensor.queue.IsEmpty() {
			continue
		}

		// Check for free scanner slot
		if sensor.conf.MaxScan > 0 &&
			sensor.init.Size()+sensor.running.Size() == int(
				sensor.conf.MaxScan,
			) {
			log.Printf(
				"Unable to start a scan from queue, Max number of scans reached.\n",
			)
			continue
		}

		// get memory stats and check for memory
		if sensor.conf.MinFreeMemScanQueue > 0 {
			m, err := util.GetAvailableMemory(util.StdMemoryManager{})
			memoryNeeded := m.Bytes + uint64(
				sensor.init.Size(),
			)*sensor.conf.MinFreeMemScanQueue
			if err != nil {
				log.Fatal().Err(err).Msg("Unable to get memory stats")
			}
			if m.Bytes < memoryNeeded {
				log.Printf("Unable to start scan, not enough memory.\n")
				continue
			}
		}

		// try to run scan process
		if err := sensor.StartScan(sensor.queue.Front()); err != nil {
			log.Printf(
				"%s: unable to start scan: %s",
				sensor.queue.Front(),
				err,
			)
		}

	}
}

// register loops until its ID is registrated
func (sensor *Scheduler) register(m handler.Register) {
	for { // loop until sensor is registered
		sensor.out <- &connection.SendResponse{
			Topic: fmt.Sprintf("%s/sensor/cmd/director", sensor.context),
			MSG:   cmds.NewModify("sensor", sensor.id, nil, "director", ""),
		}
		select {
		case <-sensor.regChan:
			log.Printf("%s registered", sensor.id)
			return
		// Send new registration mqtt message each second
		case <-time.After(time.Second):
		}
		go m.Check()
	}
}

// Close cleans all queues and OpenVAS processes, sets all scan stats to
// stopped and stops the scheduler
func (sensor *Scheduler) Close() error {
	sensor.Lock()
	defer sensor.Unlock()
	sensor.stopped = true
	log.Print("Cleaning all OpenVAS Processes...\n")
	// Remove all queued scans
	for item, ok := sensor.queue.Dequeue(); ok; item, ok = sensor.queue.Dequeue() {
		sensor.out <- messages.EventToResponse(sensor.context, info.Status{
			Identifier: messages.Identifier{
				ID:      item,
				Message: messages.NewMessage("status.scan", "", ""),
			},
			Status: "stopped",
		})
	}

	var wg sync.WaitGroup
	// Stopping all init processes
	for item, ok := sensor.init.Dequeue(); ok; item, ok = sensor.init.Dequeue() {
		log.Printf("Stopping %s\n", item)
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			sensor.ovas.StopScan(item, sensor.sudo, sensor.commander)
		}(item)
	}
	// Stopping all running processes
	for item, ok := sensor.running.Dequeue(); ok; item, ok = sensor.running.Dequeue() {
		log.Printf("Stopping %s\n", item)
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			sensor.ovas.StopScan(item, sensor.sudo, sensor.commander)
		}(item)
	}
	wg.Wait()
	close(sensor.interruptChan)
	close(sensor.regChan)
	return nil
}

func (sensor *Scheduler) Handler() map[string]connection.OnMessage {
	// TODO separate Register Sensor
	// MQTT OnMessage Types
	startStopHandler := StartStop{
		Start: sensor.QueueScan,
		Stop:  sensor.StopScan,
	}

	statusHandler := Status{
		Run: sensor.ScanRunning,
		Fin: sensor.ScanFinished,
	}

	vtsHandler := LoadVTs{
		VtsLoad: sensor.loadVTs,
	}

	registeredHandler := Registered{
		Register: sensor.regChan,
		ID:       sensor.id,
	}

	return map[string]connection.OnMessage{
		fmt.Sprintf("%s/sensor/info", sensor.context):            registeredHandler,
		fmt.Sprintf("%s/sensor/cmd", sensor.context):             vtsHandler,
		fmt.Sprintf("%s/scan/cmd/%s", sensor.context, sensor.id): startStopHandler,
		fmt.Sprintf("%s/scan/info", sensor.context):              statusHandler,
	}
}

// Start registers a sensor and starts scheduler
//
// Uses the out channel to send register messages
func (sensor *Scheduler) Start(m handler.Register) {
	sensor.register(m)
	sensor.stopped = false
	go sensor.schedule()
}

// NewScheduler creates a new scheduler
func NewScheduler(
	out chan<- *connection.SendResponse,
	id string,
	conf config.ScannerPreferences,
	context string,
) *Scheduler {
	interruptChan := make(chan string)
	return &Scheduler{
		queue:         util.NewQueueList(),
		init:          util.NewQueueList(),
		running:       util.NewQueueList(),
		ovas:          openvas.NewOpenVASScanner(interruptChan),
		sudo:          openvas.IsSudo(openvas.StdCommander{}),
		stopped:       true,
		regChan:       make(chan struct{}),
		commander:     openvas.StdCommander{},
		context:       context,
		interruptChan: interruptChan,
		out:           out,
		id:            id,
		conf:          conf,
	}
}
