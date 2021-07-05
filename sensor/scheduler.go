package sensor

import (
	"log"
	"time"

	"github.com/greenbone/eulabeia/util"
)

const (
	MAX_SCANS       = 4
	MEMORY_FOR_SCAN = 0
	SUDO            = true
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
			case scan := <-addChan:
				queue = append(queue, scan)

			case scan := <-delChan:
				util.RemoveListItem(queue, scan)

			case scan := <-runChan:
				running = append(running, scan)
				util.RemoveListItem(init, scan)

			case scan := <-finChan:
				util.RemoveListItem(running, scan)

			default:
				break checkQueues
			}
		}

		// Check for free scanner slot
		if len(init)+len(running) == MAX_SCANS {
			log.Printf("Unable to start scan, no free slots")
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
				log.Printf("Unable to start scan, not enough memory\n")
				continue
			}
		}

		// try to initalize scan
		StartScan(queue[0], SUDO, NICENESS)
		init = append(init, queue[0])
		queue = queue[1:]
	}
}

func init() {
	go schedule()
	// TODO: Setup MQTT
}
