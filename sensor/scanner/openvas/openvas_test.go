package openvas

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
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
// simulate a succesfully terminating programm. E.g. IsSudo will return true
func TestCommandSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, "test")
	os.Exit(0)
}

// helperLongCommander creates a Command to execute a programm with a endless
// runtime
type helperLongCommander struct {
}

func (exe helperLongCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandEndless", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestCommandEndless is not a real test. It is only used as a helper process to
// simulate long running programm such as a scan in openvas
func TestCommandEndless(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	for {
		time.Sleep(time.Second)
	}
}

// helperFailCommander creates a Command to execute a programm with a exit code 1
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
// simulate a failing programm. E.g. IsSudo will return false
func TestCommandFail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}

// helperVersionCommander creates a Command to execute a programm which prints an OpenVAS sample version
type helperVersionCommander struct {
}

func (exe helperVersionCommander) Command(name string, arg ...string) *exec.Cmd {
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

// helperSettingsCommander creates an Command to execute a programm which prints sample settings
type helperSettingsCommander struct {
}

func (exe helperSettingsCommander) Command(name string, arg ...string) *exec.Cmd {
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

// TestStartStopScanSudo tests the procedure of creating an openvas process and stopping it with sudo privileges
func TestStartStopScanSudo(t *testing.T) {
	// Create Process List for test
	var processes = ProcessList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}

	// Get sudo rights
	sudo := IsSudo(helperShortCommander{})
	if !sudo {
		t.Fatalf("Error: Sudo is %t, but should be %t", sudo, true)
	}

	// Start scan
	if err := StartScan("foo", 10, sudo, helperLongCommander{}, processes); err != nil {
		t.Fatalf("Error: Cannot run StartScan: %s", err)
	}

	// Check if process is added to the list
	if _, ok := processes.procs["foo"]; !ok {
		t.Fatal("Error: Unable to locate Process")
	}

	// Stop Scan
	if err := StopScan("foo", sudo, helperShortCommander{}, processes); err != nil {
		t.Fatalf("Error: Unable to stop process: %s", err)
	}

	// Check if process is removed from the list
	if _, ok := processes.procs["foo"]; ok {
		t.Fatalf("Error: Process still in process list")
	}
}

// TestNonSudoStopScanFail tests if it fails to stop a scan when there is no scan to stop
func TestNonSudoStopScanFail(t *testing.T) {
	// Create Process List for test
	var processes = ProcessList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}

	// Test if sudo is unavailable
	sudo := IsSudo(helperFailCommander{})
	if sudo {
		t.Fatalf("Error: Sudo is %t, but should be %t", sudo, false)
	}

	// Stop Scan should fail because ther is no process
	if err := StopScan("foo", sudo, helperShortCommander{}, processes); err == nil {
		t.Fatalf("Error: Should be unable to successfully stop scan")
	}
}

// TestScanFinishedSuccess tests if it can mark a scan as finished
func TestScanFinishedSuccess(t *testing.T) {
	// Create Process List for test
	var processes = ProcessList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}

	processes.addProcess("foo", nil)

	if err := ScanFinished("foo", processes); err != nil {
		t.Fatalf("Error: Unable to finish scan: %s", err)
	}
}

// TestScanFinishedFail tests if marking a scan as finished fails if there is no scan to finish
func TestScanFinishedFail(t *testing.T) {
	// Create Process List for test
	var processes = ProcessList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}

	if err := ScanFinished("foo", processes); err == nil {
		t.Fatalf("Error: ScanFinished should return an error")
	}
}

// TestGetVersion tests if the information getting from the openvas version is extracted correctly
func TestGetVersion(t *testing.T) {
	ver, err := GetVersion(helperVersionCommander{})
	if err != nil {
		t.Fatalf("Error: Unable to get Version")
	}
	if ver != "OpenVAS Test" {
		t.Fatalf("Error: Expected `OpenVAS Test`, GOT `%s`", ver)
	}
}

// TestGetSettings tests if the information getting from the openvas settings is extracted correctly
func TestGetSettings(t *testing.T) {
	set, err := GetSettings(helperSettingsCommander{})
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
