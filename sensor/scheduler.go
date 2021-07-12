package sensor

import (
	"fmt"
	"log"
	"time"

	"github.com/greenbone/eulabeia/messages"

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

var startChan = make(chan string) // Channel to insert scan into queue
var stopChan = make(chan string)  // Channel to delete scan from queue
var runChan = make(chan string)   // Channel to delete scan from init and insert it into running
var finChan = make(chan string)   // Channel to delete scan from running
var verChan = make(chan struct{}) // Channel to get OpenVAS Version
var vtsChan = make(chan struct{}) // Channel to load VTs into Redis (via OpenVAS)
var regChan = make(chan struct{}) // Channel to mark Sensor as registered

func schedule() {
	queue := make([]string, 0)
	init := make([]string, 0)
	running := make([]string, 0)
	for { // Infinite scheduler Loop
		for len(queue) == 0 { // Check for new stuff in Channels
			select {
			case scan := <-startChan: // start scan
				queue = append(queue, scan)
				handler.MQTT.Publish("eulabeia/scan/info", messages.ScanInfo{
					ID:       scan,
					InfoType: "status",
					Info:     "queued",
					Message:  messages.NewMessage("scan.info", "", ""),
				})

			case scan := <-stopChan: // stop scan
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
					err := StopScan(scan)
					if err != nil {
						log.Printf("%s: Scan cannot be stopped: %s.\n", scan, err)
						continue
					}
				}
				handler.MQTT.Publish("eulabeia/scan/info", messages.ScanInfo{
					ID:       scan,
					InfoType: "status",
					Info:     "stopped",
					Message:  messages.NewMessage("scan.info", "", ""),
				})

			case scan := <-runChan: // scan runs
				running = append(running, scan)
				util.RemoveListItem(init, scan)

			case scan := <-finChan: // scan finished
				util.RemoveListItem(running, scan)

			case <-verChan:
				ver, err := GetVersion()
				var ret string
				if err != nil {
					ret = fmt.Sprintf("%s", err)
				} else {
					ret = ver
				}
				handler.MQTT.Publish("eulabeia/scan/info", messages.ScanInfo{
					ID:       "",
					InfoType: "version",
					Info:     ret,
					Message:  messages.NewMessage("scan.version", "", ""),
				})

			case <-vtsChan:
				go LoadVTsIntoRedis()

			// TODO: Clear all openvas Processes when terminating the sensor
			// and interrupt all scans
			// TODO: When connection to Director breaks stop all scans, clear
			// all lists and try to register scheduler again

			case <-time.After(time.Second):
				continue
			}
		}

		// Check for free scanner slot
		if len(init)+len(running) == MAX_SCANS {
			log.Printf("Unable to start a scan from queue, no free slots.\n")
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
		err := StartScan(queue[0], NICENESS)
		if err != nil {
			log.Printf("%s: Scan could not start: %s", queue[0], err)
			continue
		}
		handler.MQTT.Publish("eulabeia/scan/info", messages.ScanInfo{
			ID:       queue[0],
			InfoType: "status",
			Info:     "init",
			Message:  messages.NewMessage("scan.info", "", ""),
		})
		init = append(init, queue[0])
		queue = queue[1:]

	}
}

// register loops until its ID is registrated
func register(mqtt connection.PubSub, id string) {
	mqtt.Subscribe(map[string]connection.OnMessage{
		fmt.Sprintf("eulabeia/sensor/info/%s", id): handler.RegisterHandler{
			RegChan: regChan,
		},
	})
	for {
		mqtt.Publish("eulabeia/sensor/cmd/director", messages.Command{
			ID:      "myID",
			Cmd:     "register",
			Message: messages.NewMessage("sensor.register", "", ""),
		})
		select {
		case <-regChan:
			return
		case <-time.After(time.Second):
		}
	}
}

// Start MQTT Message handling
func Start(mqtt connection.PubSub, id string) {
	// Register Sensor
	register(mqtt, id)
	// MQTT OnMessage Types
	var cmdHandler = handler.CommandHandler{
		StartChan: startChan,
		StopChan:  stopChan,
		VerChan:   verChan,
		VtsChan:   vtsChan,
	}

	var infoHandler = handler.InfoHandler{
		RunChan: runChan,
		FinChan: finChan,
	}

	// MQTT Subscription Map
	var subMap = map[string]connection.OnMessage{
		fmt.Sprintf("eulabeia/scan/cmd/%s", id): cmdHandler,
		"eulabeia/scan/info":                    infoHandler,
	}

	err := mqtt.Subscribe(subMap)
	if err != nil {
		log.Panicf("Sensor cannot subscribe to topics: %s", err)
	}
	go schedule()
}
