// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// OpenVAS component of the sensor. This module is responsible fot everything regarding OpenVAS
package openvas

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Commander interface {
	Command(name string, arg ...string) *exec.Cmd
}

type StdCommander struct {
}

func (exe StdCommander) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

// ProcessList represents a list of processes. It is used to manage processes
// within different go routines.
type ProcessList struct {
	procs map[string]*os.Process
	mutex *sync.Mutex
}

func CreateEmptyProcessList() ProcessList {
	return ProcessList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}
}

// addProcess adds a Process to the Process list
func (pl ProcessList) addProcess(scan string, p *os.Process) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	if _, ok := pl.procs[scan]; ok {
		return errors.New("process already exist")
	}
	pl.procs[scan] = p
	return nil
}

// removeProcess removes a Process from the Process list
func (pl ProcessList) removeProcess(scan string) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	if _, ok := pl.procs[scan]; !ok {
		return errors.New("process does not exist")
	}
	delete(pl.procs, scan)
	return nil
}

// StartScan starts scan with given scan-ID and process priority (-20 to 19,
// lower is more prioritized)
func StartScan(scan string, niceness int, sudo bool, exe Commander, procList ProcessList) error {
	cmdString := make([]string, 0)

	cmdString = append(cmdString, "nice", "-n", fmt.Sprintf("%v", niceness))

	if sudo {
		cmdString = append(cmdString, "sudo", "-n")
	}

	cmdString = append(cmdString, "openvas", "--scan-start", scan)

	head := cmdString[0]
	args := cmdString[1:]

	cmd := exe.Command(head, args...)

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("unable to start openvas process: %s", err)
	}
	procList.addProcess(scan, cmd.Process)
	go waitForProcessToEnd(cmd.Process, scan, procList)
	return nil
}

// StopScan stops a scan with given scan-ID
func StopScan(scan string, sudo bool, exe Commander, procList ProcessList) error {
	err := procList.removeProcess(scan)
	if err != nil {
		return err
	}

	cmdString := make([]string, 0)

	if sudo {
		cmdString = append(cmdString, "sudo", "-n")
	}

	cmdString = append(cmdString, "openvas", "--scan-stop", scan)

	head := cmdString[0]
	args := cmdString[1:]

	cmd := exe.Command(head, args...)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// ScanFinished must be called when a Openvas Process succesfully finished
func ScanFinished(scan string, procList ProcessList) error {
	return procList.removeProcess(scan)
}

// GetVersion returns the Version of OpenVAS
func GetVersion(exe Commander) (string, error) {
	out, err := exe.Command("openvas", "-V").CombinedOutput()
	if err != nil {
		return "", err
	}
	split := strings.Split(string(out), "\n")
	return split[0], nil
}

// GetSettings returns the Settings of OpenVAS as a map
func GetSettings(exe Commander) (map[string]string, error) {
	out, err := exe.Command("openvas", "-s").CombinedOutput()
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
		settingsMap[strings.TrimSpace(settingSplit[0])] = strings.TrimSpace(settingSplit[1])
	}
	return settingsMap, nil
}

// LoadVTsIntoRedis starts openvas which then loads new VTs into Redis
func LoadVTsIntoRedis(exe Commander) error {
	return exe.Command("openvas", "--update-vt-info").Run()
}

// IsSudo checks for sudo permissions
func IsSudo(exe Commander) bool {
	cmd := exe.Command("sudo", []string{"-n", "openvas", "-s"}...)
	err := cmd.Run()
	return err == nil
}

// waitForProcessToEnd gets Called as go-routine after OpenVAS Scan Process was
// started
func waitForProcessToEnd(p *os.Process, scan string, procList ProcessList) {
	p.Wait()
	err := procList.removeProcess(scan)
	if err == nil {
		log.Printf("%s: Scan process got unexpectedly stopped or killed.\n", scan)
		// TODO: Interrupt scan
		return
	}
	log.Printf("%s: Scan process with PID %v terminated correctly.\n", scan, p.Pid)
}
