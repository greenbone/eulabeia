package sensor

import (
	"os"
	"os/exec"
	"time"
	"fmt"

	"github.com/greenbone/eulabeia/models"
)

var processes = make(map[string]os.Process)
var endProcessChan = make(chan string)
var removeProcessChan = make(chan string)

type OpenvasError struct {
	What string
}

func (e *OpenvasError) Error() string {
	return e.What
}

// InitScan tries to initialize scan. Fails if max_scans is reached.
func InitScan(scan_id string) error {
	// TODO: Request Target and Scan Prefs via MQTT
}

func StartScan(t models.Target, scan string) error {
	// TODO: Start an openvas process and put process ID into processes
	cmd := exec.Command("openvas")

	err := cmd.Start()
	if err != nil {
		return &OpenvasError {
			fmt.Sprintf("Unable to start openvas process: %s", err)
		}
	}
	go waitForProcessToEnd(cmd.Process, scan)
}

func StopScan(scan string) {
	p := processes[scan]
	removeProcessChan <- scan
	// TODO: End openvas process
}

func EndScan(scan string) {
	removeProcessChan <- scan
}

func waitForProcessToEnd(p os.Process, scan string) {
	p.Wait()
	endProcessChan <- scan
}

func checkScanProcesses() {
	for {
		select {
		case scan <- endProcessChan:
			if val, ok := processes[scan]; ok {
				// openvas process terminated unexpectedly
				// TODO: Interrupt scan
				removeProcessChan <- scan
			}
		}
	}
}

func removeProcess() {
	for {
		select {
		case scan <- removeProcessChan:
			delete(processes, scan)
		}
	}
}

func init() {
	go checkScanProcesses()
	go removeProcess()
}