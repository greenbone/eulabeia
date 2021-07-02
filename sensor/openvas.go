package sensor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var processes = make(map[string]*os.Process)
var mutex = &sync.Mutex{}

var endProcessChan = make(chan string)

type Error struct {
	What string
}

func (e *Error) Error() string {
	return e.What
}

func StartScan(scan string, sudo bool, niceness int) error {
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

func StopScan(scan string, sudo bool) error {

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

func EndScan(scan string) {
	removeProcess(scan)
	log.Printf("%s: Scan successfully finished.\n", scan)
}

func waitForProcessToEnd(p *os.Process, scan string) {
	addProcess(scan, p)
	p.Wait()
	endProcessChan <- scan
	log.Printf("%s: Scan process with PID %v finished.\n", scan, p.Pid)
}

func checkScanProcesses() {
	for {
		select {
		case scan := <-endProcessChan:
			err := removeProcess(scan)
			if err == nil {
				// openvas process terminated unexpectedly
				// TODO: Interrupt scan
				log.Printf("%s: Scan process got unexpectedly stopped or killed.\n", scan)
			}
		}
	}
}

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
	go checkScanProcesses()
}
