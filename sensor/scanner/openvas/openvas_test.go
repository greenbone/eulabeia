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

package openvas

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

// This section only contains helpers to test functionality of some methods

// helperShortCommander creates a Command to execute a programm with a short
// runtime
type helperShortCommander struct {
}

func (exe helperShortCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandSuccess", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestCommandSuccess is not a real test. It is only used as a helper process to
// simulate a succesful terminating programm. E.g. IsSudo will return true
func TestCommandSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, "test")
	os.Exit(0)
}

// helperLongCommander creates a Command to execute a programm with a long
// runtime
type helperLongCommander struct {
}

func (exe helperLongCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandLong", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestCommandEndless is not a real test. It is only used as a helper process to
// simulate long running programm such as a scan in openvas
func TestCommandLong(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	time.Sleep(time.Second)
}

// helperFailCommander creates a Command to execute a programm with a exit code
// 1
type helperFailCommander struct {
}

func (exe helperFailCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandFail", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestCommandSuccess is not a real test. It is only used as a helper process to
// simulate a failing programm.
func TestCommandFail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}

// helperVersionCommander creates a Command to execute a programm which prints
// an OpenVAS sample version
type helperVersionCommander struct {
}

func (exe helperVersionCommander) Command(
	name string,
	arg ...string,
) *exec.Cmd {
	cs := []string{"-test.run=TestCommandVersion", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestCommandVersion is not a real test. It is only used as a helper process to
// simulate openvas printing its version
func TestCommandVersion(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Printf("OpenVAS Test\nFooBar\n")
	os.Exit(0)
}

// helperSettingsCommander creates an Command to execute a programm which prints
// sample settings
type helperSettingsCommander struct {
}

func (exe helperSettingsCommander) Command(
	name string,
	arg ...string,
) *exec.Cmd {
	cs := []string{"-test.run=TestCommandSettings", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestCommandVersion is not a real test. It is only used as a helper process to
// simulate openvas printing its settings
func TestCommandSettings(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Printf("Foo = Bar\nnumber = 1\nactivate item = yes\n")
	os.Exit(0)
}

// At this point the real tests begin.

// TestStartStopScanSudo tests the procedure of creating an openvas process and
// stopping it with sudo privileges
func TestStartStopScanSudo(t *testing.T) {
	// OpenVASScanner instance
	ovas := NewOpenVASScanner(make(chan string))
	sudo := IsSudo(helperShortCommander{})

	// Test for sudo rights
	if !sudo {
		t.Fatalf("Error: Sudo is %t, but should be %t", sudo, true)
	}

	// Start scan
	if err := ovas.StartScan("foo", 10, sudo, helperLongCommander{}); err != nil {
		t.Fatalf("Error: Cannot run StartScan: %s", err)
	}

	// Check if process is added to the list
	if _, ok := ovas.procs["foo"]; !ok {
		t.Fatal("Error: Unable to locate Process")
	}

	// Stop Scan
	if err := ovas.StopScan("foo", sudo, helperShortCommander{}); err != nil {
		t.Fatalf("Error: Unable to stop process: %s", err)
	}

	// Check if process is removed from the list
	if _, ok := ovas.procs["foo"]; ok {
		t.Fatalf("Error: Process still in process list")
	}
}

// TestNonSudoStopScanFail tests if it fails to stop a scan when there is no
// scan to stop
func TestNonSudoStopScanFail(t *testing.T) {
	// OpenVASScanner instance
	ovas := NewOpenVASScanner(make(chan string))
	sudo := IsSudo(helperFailCommander{})

	// Test if sudo is unavailable
	if sudo {
		t.Fatalf("Error: Sudo is %t, but should be %t", sudo, false)
	}

	// Stop Scan should fail because ther is no process
	if err := ovas.StopScan("foo", sudo, helperShortCommander{}); err == nil {
		t.Fatalf("Error: Should be unable to successfully stop scan")
	}
}

// TestScanFinishedSuccess tests if it can mark a scan as finished
func TestScanFinishedSuccess(t *testing.T) {
	// OpenVASScanner instance
	ovas := NewOpenVASScanner(make(chan string))

	ovas.addProcess("foo", nil)

	if err := ovas.ScanFinished("foo"); err != nil {
		t.Fatalf("Error: Unable to finish scan: %s", err)
	}
}

// TestScanFinishedFail tests if marking a scan as finished fails if there is no
// scan to finish
func TestScanFinishedFail(t *testing.T) {
	// OpenVASScanner instance
	ovas := NewOpenVASScanner(make(chan string))

	if err := ovas.ScanFinished("foo"); err == nil {
		t.Fatalf("Error: ScanFinished should return an error")
	}
}

// TestGetVersion tests if the information getting from the openvas version is
// extracted correctly
func TestGetVersion(t *testing.T) {
	// OpenVASScanner instance
	ovas := NewOpenVASScanner(make(chan string))

	ver, err := ovas.GetVersion(helperVersionCommander{})
	if err != nil {
		t.Fatalf("Error: Unable to get Version")
	}
	if ver != "OpenVAS Test" {
		t.Fatalf("Error: Expected `OpenVAS Test`, GOT `%s`", ver)
	}
}

// TestGetSettings tests if the information getting from the openvas settings is
// extracted correctly
func TestGetSettings(t *testing.T) {
	// OpenVASScanner instance
	ovas := NewOpenVASScanner(make(chan string))

	set, err := ovas.GetSettings(helperSettingsCommander{})
	fmt.Printf("%v\n", set)

	if err != nil {
		t.Fatalf("Error, Unable to get Settings")
	}
	if set["Foo"] != "Bar" {
		t.Fatalf("Error: Expected `Bar`, GOT `%s`", set["Foo"])
	}
	if set["number"] != "1" {
		t.Fatalf("Error: Expected `1`, GOT `%s`", set["number"])
	}
	if set["activate item"] != "yes" {
		t.Fatalf("Error: Expected `yes`, GOT `%s`", set["activated"])
	}
}
