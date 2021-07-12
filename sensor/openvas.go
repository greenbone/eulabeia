package sensor

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var processes = make(map[string]*os.Process)
var mutex = &sync.Mutex{}

// TODO: Remove global
var sudo bool

// StartScan starts scan with given scan-ID and process priority (-20 - 19,
// lower is more prioritized)
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
		return fmt.Errorf("Unable to start openvas process: %s", err)
	}
	go waitForProcessToEnd(cmd.Process, scan)
	return nil
}

// StopScan stops a scan with given scan-ID
func StopScan(scan string) error {
	err := removeProcess(scan)
	if err != nil {
		return err
	}
	log.Printf("%s: Stopping scan.\n", scan)

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

// EndScan must be called when a Openvas Process succesfully finished
func EndScan(scan string) {
	removeProcess(scan)
	log.Printf("%s: Scan successfully finished.\n", scan)
}

// GetVersion returns the Version of OpenVAS
func GetVersion() (string, error) {
	out, err := exec.Command("openvas", "-V").CombinedOutput()
	if err != nil {
		return "", err
	}
	split := strings.Split(string(out), "\n")
	return split[0], nil
}

// GetSettings returns the Settings of OpenVAS as a map
func GetSettings() (map[string]string, error) {
	out, err := exec.Command("openvas", "-s").CombinedOutput()
	if err != nil {
		return nil, err
	}
	settingsList := strings.Split(string(out), "\n")
	settingsMap := make(map[string]string)
	for _, setting := range settingsList {
		settingSplit := strings.Split(setting, "=")
		if len(settingSplit) != 2 {
			continue
		}
		settingsMap[settingSplit[0]] = settingSplit[1]
	}
	return settingsMap, nil
}

// LoadVTsIntoRedis starts openvas which then loads new VTs into Redis
func LoadVTsIntoRedis(loadedChan chan struct{}) {
	log.Printf("Loading VTs into Redis DB...\n")

	err := exec.Command("openvas", "--update-vt-info").Run()
	if err != nil {
		log.Printf("OpenVAS Scanner failed to load VTs: %s", err)
		return
	}
	log.Printf("Finished loading VTs into Redis DB.\n")
	loadedChan <- struct{}{}
}

// IsSudo checks for sudo permissions
func IsSudo() bool {
	cmd := exec.Command("sudo", []string{"-n", "openvas", "-s"}...)
	err := cmd.Run()
	if err != nil {
		log.Printf("Cannot start openvas as sudo: %s", err)
		return false
	}
	return true
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
		return errors.New("process already exist")
	}
	processes[scan] = p
	return nil
}

// removeProcess removes a Process from the Process list
func removeProcess(scan string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := processes[scan]; !ok {
		return errors.New("process does not exist")
	}
	delete(processes, scan)
	return nil
}

func init() {
	// Check for sudo rights
	sudo = IsSudo()
}
