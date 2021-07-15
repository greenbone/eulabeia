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

// OpenVASScanner is the eulabeia scanner implementation of openvas. It is
// responsible for handling processes of openvas. It also is able to start and
// stop scans via openvas.
type OpenVASScanner struct {
	procs map[string]*os.Process // Process List for running scans
	mutex *sync.Mutex            // For thread save management of processes
	exe   Commander              // Commander to use to run commands
	sudo  bool                   // Decides if scans should be run as sudo
}

// Commander is an inferace to manage different ways to handle calls to openvas.
// It is mostly used for testing purposes.
type Commander interface {
	Command(name string, arg ...string) *exec.Cmd
}

// stdCommander is the standard commander.
type StdCommander struct {
}

func (exe StdCommander) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

// CreateNewOpenVASScanner creates a new instance of an OpenVASScanner with the
// specified settings.
func CreateNewOpenVASScanner(cmd Commander, sudo bool) *OpenVASScanner {

	return &OpenVASScanner{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
		exe:   cmd,
		sudo:  sudo,
	}
}

// addProcess adds a Process to the Process list
func (ovas OpenVASScanner) addProcess(scan string, p *os.Process) error {
	ovas.mutex.Lock()
	defer ovas.mutex.Unlock()
	if _, ok := ovas.procs[scan]; ok {
		return errors.New("process already exist")
	}
	ovas.procs[scan] = p
	return nil
}

// removeProcess removes a Process from the Process list
func (ovas OpenVASScanner) removeProcess(scan string) error {
	ovas.mutex.Lock()
	defer ovas.mutex.Unlock()
	if _, ok := ovas.procs[scan]; !ok {
		return errors.New("process does not exist")
	}
	delete(ovas.procs, scan)
	return nil
}

// StartScan starts scan with given scan-ID and process priority (-20 to 19,
// lower is more prioritized)
func (ovas OpenVASScanner) StartScan(scan string, niceness int) error {
	cmdString := make([]string, 0)

	cmdString = append(cmdString, "nice", "-n", fmt.Sprintf("%v", niceness))

	if ovas.sudo {
		cmdString = append(cmdString, "sudo", "-n")
	}

	cmdString = append(cmdString, "openvas", "--scan-start", scan)

	head := cmdString[0]
	args := cmdString[1:]

	cmd := ovas.exe.Command(head, args...)

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("unable to start openvas process: %s", err)
	}
	ovas.addProcess(scan, cmd.Process)
	go ovas.waitForProcessToEnd(cmd.Process, scan)
	return nil
}

// StopScan stops a scan with given scan-ID
func (ovas OpenVASScanner) StopScan(scan string) error {
	err := ovas.removeProcess(scan)
	if err != nil {
		return err
	}

	cmdString := make([]string, 0)

	if ovas.sudo {
		cmdString = append(cmdString, "sudo", "-n")
	}

	cmdString = append(cmdString, "openvas", "--scan-stop", scan)

	head := cmdString[0]
	args := cmdString[1:]

	cmd := ovas.exe.Command(head, args...)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// ScanFinished must be called when a Openvas Process succesfully finished
func (ovas OpenVASScanner) ScanFinished(scan string) error {
	return ovas.removeProcess(scan)
}

// GetVersion returns the Version of OpenVAS
func (ovas OpenVASScanner) GetVersion() (string, error) {
	out, err := ovas.exe.Command("openvas", "-V").CombinedOutput()
	if err != nil {
		return "", err
	}
	split := strings.Split(string(out), "\n")
	return split[0], nil
}

// GetSettings returns the Settings of OpenVAS as a map
func (ovas OpenVASScanner) GetSettings() (map[string]string, error) {
	out, err := ovas.exe.Command("openvas", "-s").CombinedOutput()
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
func (ovas OpenVASScanner) LoadVTsIntoRedis() error {
	return ovas.exe.Command("openvas", "--update-vt-info").Run()
}

// IsSudo checks for sudo permissions
func IsSudo(exe Commander) bool {
	cmd := exe.Command("sudo", []string{"-n", "openvas", "-s"}...)
	err := cmd.Run()
	return err == nil
}

// waitForProcessToEnd gets Called as go-routine after OpenVAS Scan Process was
// started
func (ovas OpenVASScanner) waitForProcessToEnd(p *os.Process, scan string) {
	p.Wait()
	err := ovas.removeProcess(scan)
	if err == nil {
		log.Printf("%s: Scan process with PID %d got unexpectedly stopped or killed.\n", scan, p.Pid)
		// TODO: Interrupt scan
		return
	}
	log.Printf("%s: Scan process with PID %d terminated correctly.\n", scan, p.Pid)
}
