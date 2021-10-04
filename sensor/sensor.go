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
	"github.com/greenbone/eulabeia/models"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/sensor/scanner/openvas"

	"github.com/greenbone/eulabeia/util"
)

// Scheduler is a struct containing functionality to control a sensor
type Scheduler struct {
	queue         *util.QueueList         // queued scans
	init          *util.QueueList         // scans that currently initializes
	running       *util.QueueList         // scans that are currently running
	loadingVTs    bool                    // marks that VTs are currently loading
	ovas          *openvas.OpenVASScanner // openvas
	sudo          bool                    // sudo rights
	sync.Mutex                            // thread safe hadnling when moving scan IDs between lists
	stopped       bool                    // marks that the sensor is stopped
	regChan       chan struct{}           // channel for succesful registation
	commander     openvas.Commander       // commander used for openvas
	context       string                  // context used for mqtt
	interruptChan chan string             // channel for signaling interrupted scans

	mqtt connection.PubSub         // mqtt connection
	id   string                    // ID of the sensor
	conf config.ScannerPreferences // config file

	resolveFilter func([]models.VTFilter) ([]string, error)
}

type scan struct {
	models.ScanPrefs
	Ready bool
}

func (s *scan) Compare(sc util.Comparable) bool {
	if v, ok := sc.(*scan); ok {
		return s.ID == v.ID
	}
	return false
}

func scanModelToScan(sm models.Scan) scan {
	return scan{
		ScanPrefs: models.ScanPrefs{
			ID:          sm.ID,
			Hosts:       sm.Hosts,
			Ports:       sm.Ports,
			Plugins:     sm.Plugins.Single,
			AliveTest:   sm.AliveTest,
			Parallel:    sm.Parallel,
			Exclude:     sm.Exclude,
			Finished:    sm.Finished,
			Credentials: sm.Credentials,
		},
		Ready: false,
	}
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
func (sensor *Scheduler) QueueScan(s models.Scan) error {
	sensor.Lock()
	defer sensor.Unlock()

	scan := scanModelToScan(s)

	if sensor.queue.Contains(&scan) || sensor.init.Contains(&scan) || sensor.running.Contains(&scan) {
		return fmt.Errorf("there is already a running scan with the ID %s", scan.ID)
	}
	sensor.queue.Enqueue(&scan)
	sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
		Identifier: messages.Identifier{
			ID:      scan.ID,
			Message: messages.NewMessage("status.scan", "", ""),
		},
		Status: "queued",
	})

	vtChan := make(chan []string)

	go func() {
		vts, err := sensor.resolveFilter(s.Plugins.Group)
		if err != nil {
			log.Printf("Error while resolving VTs for scan %s: %s", s.ID, err)
		}
		vtChan <- vts
	}()
	go sensor.addVTInfo(s.ID, vtChan)
	return nil
}

// addVTInfo is adding VTs to a scan which it gets through a channel
func (sensor *Scheduler) addVTInfo(scanID string, vtChan chan []string) error {
	sensor.Lock()
	defer sensor.Unlock()

	s, ok := sensor.getScanByID(sensor.queue, scanID)
	if !ok {
		s, ok = sensor.getScanByID(sensor.init, scanID)
	}

	if ok {
		sensor.Unlock()
		log.Println("Waiting for resolved VTs")
		vts := <-vtChan
		log.Println("VTs resolved")
		vtsSingle := make([]models.SingleVT, len(vts))
		for i, v := range vts {
			vtsSingle[i] = models.SingleVT{
				OID: v,
			}
		}
		sensor.Lock()
		s.Plugins = append(s.Plugins, vtsSingle...)
		s.Plugins = removeDuplicateVT(s.Plugins)
		s.Ready = true
		return nil
	}

	return fmt.Errorf("scan with id %s not found", scanID)
}

func removeDuplicateVT(vts []models.SingleVT) []models.SingleVT {
	vtTest := make(map[string]struct{})
	ret := make([]models.SingleVT, 0)
	for _, v := range vts {
		if _, v2 := vtTest[v.OID]; !v2 {
			vtTest[v.OID] = struct{}{}
			ret = append(ret, v)
		}
	}
	return ret
}

func (sensor *Scheduler) getScanByID(ql *util.QueueList, scanID string) (*scan, bool) {
	// Create dummy to get real scan
	sd := scan{
		ScanPrefs: models.ScanPrefs{
			ID: scanID,
		},
	}

	// Check if scan is in queue
	if item, ok := ql.Get(&sd); ok {
		return item.(*scan), true
	}
	return nil, false
}

// StartScan starts a scan process
func (sensor *Scheduler) StartScan(s *scan) error {
	sensor.Lock()
	defer sensor.Unlock()

	if s == nil {
		return fmt.Errorf("scan is nil")
	}
	if !sensor.queue.Contains(s) {
		return fmt.Errorf("scan with id %s not in queue", s.ID)
	}

	if err := sensor.ovas.StartScan(s.ID, int(sensor.conf.Niceness), sensor.sudo, sensor.commander); err != nil {
		return err
	}
	// Update scan status
	sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
		Identifier: messages.Identifier{
			ID:      s.ID,
			Message: messages.NewMessage("status.scan", "", ""),
		},
		Status: "init",
	})
	// Put scan from queue to init
	sensor.queue.RemoveListItem(s)
	sensor.init.Enqueue(s)
	return nil
}

// GetScan searches the init list for the given scan id and returns it as soon
// as the vt information is given
func (sensor *Scheduler) GetScan(scanID string) (models.ScanPrefs, error) {
	sensor.Lock()
	defer sensor.Unlock()
	if s, ok := sensor.getScanByID(sensor.init, scanID); ok {
		for i := 0; !s.Ready; i++ {
			sensor.Unlock()
			time.Sleep(time.Second)
			sensor.Lock()
			if i >= 300 || sensor.stopped {
				return models.ScanPrefs{}, fmt.Errorf("timeout for scan with id %s while waiting for vt information", scanID)
			}
		}
		return models.ScanPrefs(s.ScanPrefs), nil

	}
	return models.ScanPrefs{}, fmt.Errorf("scan with id %s not in init", scanID)

}

// ScanRunning moves a scan from the init to the running state
func (sensor *Scheduler) ScanRunning(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()
	s, _ := sensor.getScanByID(sensor.init, scanID)
	if !sensor.init.RemoveListItem(s) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}
	sensor.running.Enqueue(s)
	return nil
}

func (sensor *Scheduler) ScanFinished(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()
	s, _ := sensor.getScanByID(sensor.running, scanID)
	if !sensor.running.RemoveListItem(s) {
		return fmt.Errorf("scan ID %s unknown", scanID)
	}
	sensor.ovas.ScanFinished(scanID)
	return nil
}

// StopScan will remove the scan from the queue or invoke a stop scan command to scanner
func (sensor *Scheduler) StopScan(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()
	s := scan{
		ScanPrefs: models.ScanPrefs{
			ID: scanID,
		},
	}
	if sensor.queue.RemoveListItem(&s) {
		return nil
	}
	if sensor.init.RemoveListItem(&s) || sensor.running.RemoveListItem(&s) {
		err := sensor.ovas.StopScan(scanID, sensor.sudo, sensor.commander)
		if err == nil {
			sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
				Identifier: messages.Identifier{
					ID:      scanID,
					Message: messages.NewMessage("status.scan", "", ""),
				},
				Status: "stopped",
			})
			return nil
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

// interruptScan removes scan from list and publishes a status MSG
func (sensor *Scheduler) interruptScan(scanID string) error {
	sensor.Lock()
	defer sensor.Unlock()

	s := scan{
		ScanPrefs: models.ScanPrefs{
			ID: scanID,
		},
	}
	if sensor.init.RemoveListItem(&s) || sensor.running.RemoveListItem(&s) {
		sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
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
		if err := sensor.StartScan(sensor.queue.Front().(*scan)); err != nil {
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
// stopped and stops the scheduler
func (sensor *Scheduler) Close() error {
	sensor.Lock()
	defer sensor.Unlock()
	sensor.stopped = true
	log.Print("Cleaning all OpenVAS Processes...\n")
	// Remove all queued scans
	for item := sensor.queue.Dequeue(); item != nil; item = sensor.queue.Dequeue() {
		sensor.mqtt.Publish(fmt.Sprintf("%s/scan/info", sensor.context), info.Status{
			Identifier: messages.Identifier{
				ID:      item.(*scan).ID,
				Message: messages.NewMessage("status.scan", "", ""),
			},
			Status: "stopped",
		})
	}

	var wg sync.WaitGroup
	// Stopping all init processes
	for item := sensor.init.Dequeue(); item != nil; item = sensor.init.Dequeue() {
		log.Printf("Stopping %s\n", item.(*scan).ID)
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			sensor.ovas.StopScan(item, sensor.sudo, sensor.commander)
		}(item.(*scan).ID)
	}
	// Stopping all running processes
	for item := sensor.running.Dequeue(); item != nil; item = sensor.running.Dequeue() {
		log.Printf("Stopping %s\n", item.(*scan).ID)
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			sensor.ovas.StopScan(item, sensor.sudo, sensor.commander)
		}(item.(*scan).ID)
	}
	wg.Wait()
	return nil
}

// Start initializes MQTT handling and starts the scheduler
func (sensor *Scheduler) Start() {
	// Subscribe on Topic to get confirmation about registration
	sensor.mqtt.Subscribe(map[string]connection.OnMessage{
		fmt.Sprintf("%s/sensor/info", sensor.context): Registered{
			Register: sensor.regChan,
			ID:       sensor.id,
		},
	})
	// Register Sensor
	sensor.register()
	// MQTT OnMessage Types
	var startStopHandler = ScanCmd{
		Context: sensor.context,
		Stop:    sensor.StopScan,
		Get:     sensor.GetScan,
	}

	var scanInfoHandler = ScanInfo{
		Context: sensor.context,
		Sensor:  sensor.id,
		Run:     sensor.ScanRunning,
		Fin:     sensor.ScanFinished,
		Start:   sensor.QueueScan,
	}

	var vtsHandler = LoadVTs{
		VtsLoad: sensor.loadVTs,
	}

	// MQTT Subscription Map
	var subMap = map[string]connection.OnMessage{
		fmt.Sprintf("%s/scan/cmd/%s", sensor.context, sensor.id): startStopHandler,
		fmt.Sprintf("%s/scan/info", sensor.context):              scanInfoHandler,
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
func NewScheduler(mqtt connection.PubSub, id string, conf config.ScannerPreferences, context string, rf func([]models.VTFilter) ([]string, error)) *Scheduler {
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

		mqtt: mqtt,
		id:   id,
		conf: conf,

		resolveFilter: rf,
	}
}
