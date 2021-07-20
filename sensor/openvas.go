// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// OpenVAS component of the sensor. This module is responsible fot everything regarding OpenVAS
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

// rocessList represents a list of processes. It is used to manage processes
// within different go routines.
type processList struct {
	procs map[string]*os.Process
	mutex *sync.Mutex
}

// addProcess adds a Process to the Process list
func (pl processList) addProcess(scan string, p *os.Process) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	if _, ok := pl.procs[scan]; ok {
		return errors.New("process already exist")
	}
	pl.procs[scan] = p
	return nil
}

// removeProcess removes a Process from the Process list
func (pl processList) removeProcess(scan string) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	if _, ok := pl.procs[scan]; !ok {
		return errors.New("process does not exist")
	}
	delete(pl.procs, scan)
	return nil
}

var processes processList

// StartScan starts scan with given scan-ID and process priority (-20 to 19,
// lower is more prioritized)
func StartScan(scan string, niceness int, sudo bool) error {
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
		return fmt.Errorf("unable to start openvas process: %s", err)
	}
	go waitForProcessToEnd(cmd.Process, scan)
	return nil
}

// StopScan stops a scan with given scan-ID
func StopScan(scan string, sudo bool) error {
	err := processes.removeProcess(scan)
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
func EndScan(scan string) error {
	return processes.removeProcess(scan)
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
func LoadVTsIntoRedis() error {
	return exec.Command("openvas", "--update-vt-info").Run()
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
	processes.addProcess(scan, p)
	p.Wait()
	err := processes.removeProcess(scan)
	if err == nil {
		log.Printf("%s: Scan process got unexpectedly stopped or killed.\n", scan)
		// TODO: Interrupt scan
		return
	}
	log.Printf("%s: Scan process with PID %v terminated correctly.\n", scan, p.Pid)
}

func init() {
	processes = processList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}
}
