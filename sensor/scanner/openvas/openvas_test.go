package openvas

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"
)

type TestShortCommander struct {
}

func (exe TestShortCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandSuccess", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestCommandSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, "test")
	os.Exit(0)
}

type TestLongCommander struct {
}

func (exe TestLongCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandEndless", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestCommandEndless(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	for {
		time.Sleep(time.Second)
	}
}

type TestFailCommander struct {
}

func (exe TestFailCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandFail", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestCommandFail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}

type TestVersionCommander struct {
}

func (exe TestVersionCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandVersion", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestCommandVersion(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Printf("OpenVAS Test\nFooBar\n")
	os.Exit(0)
}

type TestSettingsCommander struct {
}

func (exe TestSettingsCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandSettings", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestCommandSettings(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Printf("Foo = Bar\nnumber = 1\nactivate item = yes\n")
	os.Exit(0)
}

func TestStartStopScanSudo(t *testing.T) {
	// Create Process List for test
	var processes = ProcessList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}

	// Get sudo rights
	sudo := IsSudo(TestShortCommander{})
	if !sudo {
		t.Fatalf("Error: Sudo is %t, but should be %t", sudo, true)
	}

	// Start scan
	if err := StartScan("foo", 10, sudo, TestLongCommander{}, processes); err != nil {
		t.Fatalf("Error: Cannot run StartScan: %s", err)
	}

	// Check if process is added to the list
	if _, ok := processes.procs["foo"]; !ok {
		t.Fatal("Error: Unable to locate Process")
	}

	// Stop Scan
	if err := StopScan("foo", sudo, TestShortCommander{}, processes); err != nil {
		t.Fatalf("Error: Unable to stop process: %s", err)
	}

	// Check if process is removed from the list
	if _, ok := processes.procs["foo"]; ok {
		t.Fatalf("Error: Process still in process list")
	}
}

func TestNonSudoStopScanFail(t *testing.T) {
	// Create Process List for test
	var processes = ProcessList{
		procs: make(map[string]*os.Process),
		mutex: &sync.Mutex{},
	}

	// Test if sudo is unavailable
	sudo := IsSudo(TestFailCommander{})
	if sudo {
		t.Fatalf("Error: Sudo is %t, but should be %t", sudo, false)
	}

	// Stop Scan should fail because ther is no process
	if err := StopScan("foo", sudo, TestShortCommander{}, processes); err == nil {
		t.Fatalf("Error: Should be unable to successfully stop scan")
	}
}

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

func TestGetVersion(t *testing.T) {
	ver, err := GetVersion(TestVersionCommander{})
	if err != nil {
		t.Fatalf("Error: Unable to get Version")
	}
	if ver != "OpenVAS Test" {
		t.Fatalf("Error: Expected `OpenVAS Test`, GOT `%s`", ver)
	}
}

func TestGetSettings(t *testing.T) {
	set, err := GetSettings(TestSettingsCommander{})
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
