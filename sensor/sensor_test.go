package sensor

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
)

// TODO write clean service
var out = make(chan *connection.SendResponse, 100)

// helperShortCommander creates a Command to execute a programm with a short
// runtime
type MockCommander struct {
}

func (exe MockCommander) Command(name string, arg ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandSuccess", "--", name}
	cs = append(cs, arg...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestCommandSuccess is not a real test. It is only used as a helper process to
// simulate a succesful terminating programm.
func TestCommandSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, "test")
	os.Exit(0)
}

// TestQueueInitRunFinishScan is testing the whole process of a start scan event
// inside the sensor.
func TestQueueInitRunFinishScan(t *testing.T) {
	var conf = config.ScannerPreferences{
		ScanInfoStoreTime:   0,
		MaxScan:             0,
		MaxQueuedScans:      0,
		Niceness:            10,
		MinFreeMemScanQueue: 0,
	}
	scheduler := NewScheduler(out, "testID", conf, "")
	scheduler.commander = MockCommander{}

	scanID := "foo"

	// Insert scan into queue
	if err := scheduler.QueueScan(scanID); err != nil {
		t.Fatalf("There should be no error but got: %s\n", err)
	}

	// Check if scan is in queue
	if !scheduler.queue.Contains(scanID) {
		t.Fatal("Scan foo was not added to queue\n")
	}

	// Check if I can add another scan with same ID (should fail)
	if err := scheduler.QueueScan(scanID); err == nil {
		t.Fatalf("There should be an error but got none\n")
	}

	// Check if init list is empty
	if !scheduler.init.IsEmpty() {
		t.Fatalf("Init list should be empty but is not\n")
	}

	// Check if scan can be started
	if err := scheduler.StartScan(scanID); err != nil {
		t.Fatalf("Cannot start scan: %s\n", err)
	}

	// Check if scan is in init list
	if !scheduler.init.Contains(scanID) {
		t.Fatalf("Scan was not added to init\n")
	}

	// Check if queue list is empty again
	if !scheduler.queue.IsEmpty() {
		t.Fatalf("Queue list should be empty but is not\n")
	}
	// Check if the scan can be started again (should fail)
	if err := scheduler.StartScan(scanID); err == nil {
		t.Fatalf("Should be unable to start scan again\n")
	}

	// Check if running list is empty
	if !scheduler.running.IsEmpty() {
		t.Fatalf("Running list should be empty but is not\n")
	}

	// Check if the status of the scan can be switched to running
	if err := scheduler.ScanRunning(scanID); err != nil {
		t.Fatalf("Cannot stwitch state to running: %s\n", err)
	}

	// Check if scan is in running list
	if !scheduler.running.Contains(scanID) {
		t.Fatalf("Scan is not in running list but should be\n")
	}

	// Check if init list is empty again
	if !scheduler.init.IsEmpty() {
		t.Fatalf("Init list should be empty but is not\n")
	}

	// Check if it works another time (should fail)
	if err := scheduler.ScanRunning(scanID); err == nil {
		t.Fatalf("Should be unable to set state to running again\n")
	}

	// Finish the scan
	if err := scheduler.ScanFinished(scanID); err != nil {
		t.Fatalf("Unable to set scan as finished: %s\n", err)
	}

	// Check if running list is empty
	if !scheduler.running.IsEmpty() {
		t.Fatalf("Running list should be empty but is not\n")
	}

	// Check if scan can be finished again (should fail)
	if err := scheduler.ScanFinished(scanID); err == nil {
		t.Fatalf("Should be unable to set state to finished again\n")
	}

}

func TestStopScan(t *testing.T) {
	var conf = config.ScannerPreferences{
		ScanInfoStoreTime:   0,
		MaxScan:             0,
		MaxQueuedScans:      0,
		Niceness:            10,
		MinFreeMemScanQueue: 0,
	}
	scheduler := NewScheduler(out, "testID", conf, "")
	scheduler.commander = MockCommander{}

	scanID := "foo"

	// Test failcase of StopScan
	if err := scheduler.StopScan(scanID); err == nil {
		t.Fatalf("Should not be able to stop scan\n")
	}

	// Test stop scan if scan is in queue
	scheduler.queue.Enqueue(scanID)
	if err := scheduler.StopScan(scanID); err != nil {
		t.Fatalf("Unable to stop scan: %s\n", err)
	}
	if !scheduler.queue.IsEmpty() {
		t.Fatalf("Queue should be empty but ist not\n")
	}

	// Test stop scan if scan is in init
	scheduler.init.Enqueue(scanID)
	if err := scheduler.StopScan(scanID); err.Error() != fmt.Sprintf(
		"process for scan id %s does not exist",
		scanID,
	) {
		t.Fatalf("Unable to stop scan: %s\n", err)
	}
	if !scheduler.init.IsEmpty() {
		t.Fatalf("Init list should be empty but ist not\n")
	}

	// Test stop scan if scan is in running
	scheduler.running.Enqueue(scanID)
	if err := scheduler.StopScan(scanID); err.Error() != fmt.Sprintf(
		"process for scan id %s does not exist",
		scanID,
	) {
		t.Fatalf("Unable to stop scan: %s\n", err)
	}
	if !scheduler.running.IsEmpty() {
		t.Fatalf("Running list should be empty but ist not\n")
	}

}

func TestClose(t *testing.T) {
	var conf = config.ScannerPreferences{
		ScanInfoStoreTime:   0,
		MaxScan:             0,
		MaxQueuedScans:      0,
		Niceness:            10,
		MinFreeMemScanQueue: 0,
	}
	scheduler := NewScheduler(out, "testID", conf, "")
	scheduler.commander = MockCommander{}

	for i := 0; i < 30; i++ {
		switch {
		case i < 10:
			scheduler.queue.Enqueue(fmt.Sprint(i))
		case i < 20:
			scheduler.init.Enqueue(fmt.Sprint(i))
		default:
			scheduler.running.Enqueue(fmt.Sprint(i))
		}
	}

	if scheduler.queue.IsEmpty() || scheduler.init.IsEmpty() ||
		scheduler.running.IsEmpty() {
		t.Fatalf("None of the queueLists should be empty\n")
	}

	scheduler.Close()

	if !scheduler.queue.IsEmpty() || !scheduler.init.IsEmpty() ||
		!scheduler.running.IsEmpty() {
		t.Fatalf("All of the queueLists should be empty\n")
	}
}

func TestInterruptedScan(t *testing.T) {
	var conf = config.ScannerPreferences{
		ScanInfoStoreTime:   0,
		MaxScan:             0,
		MaxQueuedScans:      0,
		Niceness:            10,
		MinFreeMemScanQueue: 0,
	}
	scheduler := NewScheduler(out, "testID", conf, "")

	scan1 := "foo"
	scan2 := "bar"

	if err := scheduler.interruptScan(scan1); err == nil {
		t.Fatalf("Should be unable to interrupt unknsown scan %s", scan1)
	}
	if err := scheduler.interruptScan(scan2); err == nil {
		t.Fatalf("Should be unable to interrupt unknsown scan %s", scan2)
	}

	scheduler.init.Enqueue(scan1)
	scheduler.running.Enqueue(scan2)

	if err := scheduler.interruptScan(scan1); err != nil {
		t.Fatalf("Should be able to interrupt unknsown scan %s", scan1)
	}
	if err := scheduler.interruptScan(scan2); err != nil {
		t.Fatalf("Should be able to interrupt unknsown scan %s", scan2)
	}
}
