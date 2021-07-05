package sensor

import (
	"log"
	"time"

	"github.com/greenbone/eulabeia/util"
)

const (
	MAX_SCANS       = 4
	MEMORY_FOR_SCAN = 0
	NICENESS        = 10
)

var addChan = make(chan string) // Channel to insert scan into queue
var delChan = make(chan string) // Channel to delete scan from queue
var runChan = make(chan string) // Channel to delete scan from init and insert it into running
var finChan = make(chan string) // Channel to delete scan from running

func Add(scan string) {
	addChan <- scan
}

func Del(scan string) {
	delChan <- scan
}

func Run(scan string) {
	runChan <- scan
}

func Fin(scan string) {
	finChan <- scan
}

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

			case scan := <-delChan: // stop scan
				var ok bool
				queue, ok = util.RemoveListItem(queue, scan)
				if !ok { // scan was not queued
					init, ok = util.RemoveListItem(init, scan)
					if !ok { // scan was not in init, scan should be in running
						running, ok = util.RemoveListItem(running, scan)
						if !ok {
							log.Printf("%s: Scan cannot be stopped: Scan ID unknown.\n", scan)
						}
					}
					err := StopScan(scan)
					if err != nil {
						log.Printf("%s: Scan cannot be stopped: %s.\n", scan, err)
					}
				}

			case scan := <-runChan: // scan runs
				running = append(running, scan)
				util.RemoveListItem(init, scan)

			case scan := <-finChan: // scan finished
				util.RemoveListItem(running, scan)

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

		// try to initalize scan
		StartScan(queue[0], NICENESS)
		init = append(init, queue[0])
		queue = queue[1:]
	}
}

func init() {
	go schedule()
	// TODO: Setup MQTT
}
