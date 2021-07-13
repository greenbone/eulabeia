// Scheduler component of the sensor. This module is responsible for handling request from the director.
package sensor

import (
	"fmt"
	"log"
	"time"

	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/sensor/handler"

	"github.com/greenbone/eulabeia/util"
)

// TODO: Replace Consts with Config File
const (
	MAX_SCANS       = 4
	MEMORY_FOR_SCAN = 0
	NICENESS        = 10
)

type schedulerChannels struct {
	startChan chan string   // Channel to insert scan into queue
	stopChan  chan string   // Channel to delete scan from queue
	runChan   chan string   // Channel to delete scan from init and insert it into running
	finChan   chan string   // Channel to delete scan from running
	verChan   chan struct{} // Channel to get OpenVAS Version
	vtsChan   chan struct{} // Channel to load VTs into Redis (via OpenVAS)
	regChan   chan struct{} // Channel to mark Sensor as registered
	discChan  chan struct{} // Channel to mark disconnected director
}

// Checks for new instructions for the sensor and starts queued scans.
func schedule(channels schedulerChannels, mqtt connection.PubSub) {
	queue := make([]string, 0)
	init := make([]string, 0)
	running := make([]string, 0)
	sudo := IsSudo()

	var vtsLoadedChan = make(chan struct{})
	vtsLoading := true
	go LoadVTsIntoRedis(vtsLoadedChan)

	for { // Infinite scheduler Loop
		for vtsLoading || len(queue) == 0 { // Check for new stuff in Channels
			select {
			case scan := <-channels.startChan: // start scan
				queue = append(queue, scan)
				mqtt.Publish("eulabeia/scan/info", info.Status{
					Identifier: messages.Identifier{
						ID:      scan,
						Message: messages.NewMessage("scan.status", "", ""),
					},
					Status: "queued",
				})

			case scan := <-channels.stopChan: // stop scan
				var ok bool
				queue, ok = util.RemoveListItem(queue, scan)
				if !ok { // scan was not queued
					init, ok = util.RemoveListItem(init, scan)
					if !ok { // scan was not in init, scan should be in running
						running, ok = util.RemoveListItem(running, scan)
						if !ok {
							log.Printf("%s: Scan cannot be stopped: Scan ID unknown.\n", scan)
							continue
						}
					}
					err := StopScan(scan, sudo)
					if err != nil {
						log.Printf("%s: Scan cannot be stopped: %s.\n", scan, err)
						continue
					}
				}
				mqtt.Publish("eulabeia/scan/info", info.Status{
					Identifier: messages.Identifier{
						ID:      scan,
						Message: messages.NewMessage("scan.status", "", ""),
					},
					Status: "stopped",
				})

			case scan := <-channels.runChan: // scan runs
				running = append(running, scan)
				util.RemoveListItem(init, scan)

			case scan := <-channels.finChan: // scan finished
				util.RemoveListItem(running, scan)

			case <-channels.verChan:
				ver, err := GetVersion()
				var ret string
				if err != nil {
					ret = fmt.Sprintf("%s", err)
				} else {
					ret = ver
				}
				mqtt.Publish("eulabeia/scan/info", info.Version{
					Identifier: messages.Identifier{
						ID:      "",
						Message: messages.NewMessage("sensor.version", "", ""),
					},
					Version: ret,
				})

			case <-channels.vtsChan:
				go LoadVTsIntoRedis(vtsLoadedChan)
				vtsLoading = true

			case <-vtsLoadedChan:
				vtsLoading = false

			// TODO: When terminating the sensor clear all openvas processes
			// and interrupt all scans
			// TODO: When connection to Director breaks stop all scans, clear
			// all lists and try to register scheduler again

			// Check each second if scans can be started
			case <-time.After(time.Second):
				continue
			}
		}

		// Check for free scanner slot
		if len(init)+len(running) == MAX_SCANS {
			log.Printf("Unable to start a scan from queue, Max number of scans reached.\n")
			continue
		}

		// get memory stats and check for memory
		if MEMORY_FOR_SCAN > 0 {
			m, err := util.GetAvailableMemory()
			memoryNeeded := m.Bytes + uint64(len(init))*MEMORY_FOR_SCAN
			if err != nil {
				log.Panicf("Unable to get memory stats: %s\n", err)
			}
			if m.Bytes < memoryNeeded {
				log.Printf("Unable to start scan, not enough memory.\n")
				continue
			}
		}

		// try to run scan process
		err := StartScan(queue[0], NICENESS, sudo)
		if err != nil {
			log.Printf("%s: Scan could not start: %s", queue[0], err)
			continue
		}
		mqtt.Publish("eulabeia/scan/info", info.Status{
			Identifier: messages.Identifier{
				ID:      queue[0],
				Message: messages.NewMessage("scan.status", "", ""),
			},
			Status: "init",
		})
		init = append(init, queue[0])
		queue = queue[1:]

	}
}

// register loops until its ID is registrated
func register(mqtt connection.PubSub, id string, regChan chan struct{}) {
	for { // loop until sensor is registered
		mqtt.Publish("eulabeia/sensor/cmd/director", cmds.Register{
			Identifier: messages.Identifier{
				ID:      id,
				Message: messages.NewMessage("sensor.register", "", ""),
			},
		})
		select {
		case <-regChan:
			return
		// Send new registration mqtt message each second
		case <-time.After(time.Second):
		}
	}
}

// Start MQTT Message handling
func Start(mqtt connection.PubSub, id string) {
	// Setup Channels
	channels := schedulerChannels{
		startChan: make(chan string),
		stopChan:  make(chan string),
		runChan:   make(chan string),
		finChan:   make(chan string),
		verChan:   make(chan struct{}),
		vtsChan:   make(chan struct{}),
		regChan:   make(chan struct{}),
		discChan:  make(chan struct{}),
	}
	// Subscribe on Topic to get confirmation about registration
	mqtt.Subscribe(map[string]connection.OnMessage{
		fmt.Sprintf("eulabeia/sensor/info/%s", id): handler.Registered{
			RegChan: channels.regChan,
		},
	})
	// Register Sensor
	register(mqtt, id, channels.regChan)
	// MQTT OnMessage Types
	var startStopHandler = handler.StartStop{
		StartChan: channels.startChan,
		StopChan:  channels.stopChan,
	}

	var statusHandler = handler.Status{
		RunChan: channels.runChan,
		FinChan: channels.finChan,
	}

	var vtsHandler = handler.LoadVTs{
		VtsChan: channels.vtsChan,
	}

	// MQTT Subscription Map
	var subMap = map[string]connection.OnMessage{
		fmt.Sprintf("eulabeia/scan/cmd/%s", id): startStopHandler,
		"eulabeia/scan/info":                    statusHandler,
		"eulabeia/sensor/cmd":                   vtsHandler,
	}

	err := mqtt.Subscribe(subMap)
	if err != nil {
		log.Panicf("Sensor cannot subscribe to topics: %s", err)
	}

	// TODO: Maybe without go routine. This will be the demon process.
	go schedule(channels, mqtt)
}
