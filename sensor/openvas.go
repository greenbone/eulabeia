package sensor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/greenbone/eulabeia/connection/mqtt"
)

var processes = make(map[string]*os.Process)
var mutex = &sync.Mutex{}

var endProcessChan = make(chan string)

var sudo bool

var _mqtt mqtt.MQTT

type Error struct {
	What string
}

func (e *Error) Error() string {
	return e.What
}

func StartScan(scan string, niceness int) error {
	cmdString := make([]string, 0)

	if niceness != 0 {
		cmdString = append(cmdString, "nice", "-n", fmt.Sprintf("%v", niceness))
	}

	if sudo {
		cmdString = append(cmdString, "sudo", "-n")
	}

	cmdString = append(cmdString, "openvas", "--scan-start", scan)

	head := cmdString[0]
	args := cmdString[1:]

	cmd := exec.Command(head, args...)

	err := cmd.Start()
	if err != nil {
		return &Error{
			fmt.Sprintf("Unable to start openvas process: %s", err),
		}
	}
	go waitForProcessToEnd(cmd.Process, scan)
	return nil
}

// Stops a scan
func StopScan(scan string) error {

	err := removeProcess(scan)
	if err != nil {
		return err
	}
	log.Printf("%s: Stopping scan.\n", scan)
	// TODO: End openvas process
	cmdString := make([]string, 0)

	if sudo {
		cmdString = append(cmdString, "sudo", "-n")
	}

	cmdString = append(cmdString, "openvas", "--scan-stop", scan)

	head := cmdString[0]
	args := cmdString[1:]

	cmd := exec.Command(head, args...)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// EndScan must be called when a Scan Process succesfully finished
func EndScan(scan string) {
	removeProcess(scan)
	log.Printf("%s: Scan successfully finished.\n", scan)
}

// waitForProcessToEnd gets Called as go-routine after OpenVAS Scan Process was
// started
func waitForProcessToEnd(p *os.Process, scan string) {
	addProcess(scan, p)
	p.Wait()
	err := removeProcess(scan)
	if err == nil {
		log.Printf("%s: Scan process got unexpectedly stopped or killed.\n", scan)
		// TODO: Interrupt scan
		return
	}
	log.Printf("%s: Scan process with PID %v terminated correctly.\n", scan, p.Pid)
}

// addProcess adds a Process to the Process list
func addProcess(scan string, p *os.Process) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := processes[scan]; ok {
		return &Error{
			"process already exist",
		}
	}
	processes[scan] = p
	return nil
}

// removeProcess removes a Process from the Process list
func removeProcess(scan string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := processes[scan]; !ok {
		return &Error{
			"process does not exist",
		}
	}
	delete(processes, scan)
	return nil
}

func init() {
	// Check for sudo rights
	cmd := exec.Command("sudo", []string{"-n", "openvas", "-s"}...)
	err := cmd.Run()
	if err != nil {
		log.Printf("Cannot start openvas as sudo: %s", err)
		sudo = false
	} else {
		sudo = true
	}

	// Setup MQTT
}
