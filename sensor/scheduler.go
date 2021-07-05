package sensor

import (
	"fmt"
	"log"
	"time"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/sensor/handler"

	"github.com/greenbone/eulabeia/util"
)

const (
	MAX_SCANS       = 4
	MEMORY_FOR_SCAN = 0
	NICENESS        = 10
)

var addChan = make(chan string)   // Channel to insert scan into queue
var delChan = make(chan string)   // Channel to delete scan from queue
var runChan = make(chan string)   // Channel to delete scan from init and insert it into running
var finChan = make(chan string)   // Channel to delete scan from running
var verChan = make(chan struct{}) // Channel to get OpenVAS Version
var vtsChan = make(chan struct{}) // Channel to load VTs into Redis (via OpenVAS)

func schedule() {
	queue := make([]string, 0)
	init := make([]string, 0)
	running := make([]string, 0)
	for {
		time.Sleep(50 * time.Millisecond)
	checkQueues:
		for { // Check for new stuff in queus
			select {
			case scan := <-addChan: // start scan
				queue = append(queue, scan)
				handler.MQTT.Publish("scans.status", map[string]string{
					"scan_id": scan,
					"status":  "Queued",
				})

			case scan := <-delChan: // stop scan
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
				handler.MQTT.Publish("scans.status", map[string]string{
					"scan_id": scan,
					"status":  "Stopped",
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
				handler.MQTT.Publish("scanner.version", map[string]string{
					"version": ret,
				})

			case <-vtsChan:
				go LoadVTsIntoRedis()

			default:
				break checkQueues
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
		handler.MQTT.Publish("scans.status", map[string]string{
			"scan_id": queue[0],
			"status":  "Init",
		})
		init = append(init, queue[0])
		queue = queue[1:]

	}
}

// Init MQTT Message handling
func init() {
	// MQTT OnMessage Types
	var mqttStartScan = handler.SchedulerHandler{
		Channel: addChan,
	}
	var mqttStopScan = handler.SchedulerHandler{
		Channel: delChan,
	}
	var mqttScanStarted = handler.SchedulerHandler{
		Channel: runChan,
	}
	var mqttScanFinished = handler.SchedulerHandler{
		Channel: finChan,
	}

	// MQTT Subscription Map
	var subMap = map[string]connection.OnMessage{
		"sensor.startScan":     mqttStartScan,
		"sensor.stopScan":      mqttStopScan,
		"scanner.scanStarted":  mqttScanStarted,
		"scanner.scanFinished": mqttScanFinished,
	}

	err := handler.MQTT.Subscribe(subMap)
	log.Panicf("Sensor cannot subscribe to topics: %s", err)

	go schedule()
}
