package sensor

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/models"
)

type MockPubSub struct{}

func (mps MockPubSub) Close() error {
	return nil
}

func (mps MockPubSub) Connect() error {
	return nil
}

func (mps MockPubSub) Publish(topic string, message interface{}) error {
	return nil
}

func (mps MockPubSub) Preprocess(topic string, message []byte) ([]connection.TopicData, bool) {
	return nil, false
}

func (mps MockPubSub) Subscribe(handler map[string]connection.OnMessage) error {
	return nil
}

func MockResolveVT([]models.VTFilter) ([]string, error) {
	return nil, nil
}

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
	scheduler := NewScheduler(MockPubSub{}, "testID", conf, "", MockResolveVT)
	scheduler.commander = MockCommander{}

	scanM := models.Scan{
		ID: "foo",
	}

	s := &scan{
		ScanPrefs: models.ScanPrefs{
			ID: "foo",
		},
	}

	// Insert scan into queue
	if err := scheduler.QueueScan(scanM); err != nil {
		t.Fatalf("There should be no error but got: %s\n", err)
	}

	// Check if scan is in queue
	if !scheduler.queue.Contains(s) {
		t.Fatal("Scan foo was not added to queue\n")
	}

	// Check if I can add another scan with same ID (should fail)
	if err := scheduler.QueueScan(scanM); err == nil {
		t.Fatalf("There should be an error but got none\n")
	}

	// Check if init list is empty
	if !scheduler.init.IsEmpty() {
		t.Fatalf("Init list should be empty but is not\n")
	}

	// Check if scan can be started
	if err := scheduler.StartScan(s); err != nil {
		t.Fatalf("Cannot start scan: %s\n", err)
	}

	// Check if scan is in init list
	if !scheduler.init.Contains(s) {
		t.Fatalf("Scan was not added to init\n")
	}

	// Check if queue list is empty again
	if !scheduler.queue.IsEmpty() {
		t.Fatalf("Queue list should be empty but is not\n")
	}
	// Check if the scan can be started again (should fail)
	if err := scheduler.StartScan(s); err == nil {
		t.Fatalf("Should be unable to start scan again\n")
	}

	// Check if running list is empty
	if !scheduler.running.IsEmpty() {
		t.Fatalf("Running list should be empty but is not\n")
	}

	// Check if the status of the scan can be switched to running
	if err := scheduler.ScanRunning(s.ID); err != nil {
		t.Fatalf("Cannot stwitch state to running: %s\n", err)
	}

	// Check if scan is in running list
	if !scheduler.running.Contains(s) {
		t.Fatalf("Scan is not in running list but should be\n")
	}

	// Check if init list is empty again
	if !scheduler.init.IsEmpty() {
		t.Fatalf("Init list should be empty but is not\n")
	}

	// Check if it works another time (should fail)
	if err := scheduler.ScanRunning(s.ID); err == nil {
		t.Fatalf("Should be unable to set state to running again\n")
	}

	// Finish the scan
	if err := scheduler.ScanFinished(s.ID); err != nil {
		t.Fatalf("Unable to set scan as finished: %s\n", err)
	}

	// Check if running list is empty
	if !scheduler.running.IsEmpty() {
		t.Fatalf("Running list should be empty but is not\n")
	}

	// Check if scan can be finished again (should fail)
	if err := scheduler.ScanFinished(s.ID); err == nil {
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
	scheduler := NewScheduler(MockPubSub{}, "testID", conf, "", MockResolveVT)
	scheduler.commander = MockCommander{}

	s := &scan{
		ScanPrefs: models.ScanPrefs{
			ID: "foo",
		},
	}

	// Test failcase of StopScan
	if err := scheduler.StopScan(s.ID); err == nil {
		t.Fatalf("Should not be able to stop scan\n")
	}

	// Test stop scan if scan is in queue
	scheduler.queue.Enqueue(s)
	if err := scheduler.StopScan(s.ID); err != nil {
		t.Fatalf("Unable to stop scan: %s\n", err)
	}
	if !scheduler.queue.IsEmpty() {
		t.Fatalf("Queue should be empty but ist not\n")
	}

	// Test stop scan if scan is in init
	scheduler.init.Enqueue(s)
	if err := scheduler.StopScan(s.ID); err.Error() != fmt.Sprintf("process for scan id %s does not exist", s.ID) {
		t.Fatalf("Unable to stop scan: %s\n", err)
	}
	if !scheduler.init.IsEmpty() {
		t.Fatalf("Init list should be empty but ist not\n")
	}

	// Test stop scan if scan is in running
	scheduler.running.Enqueue(s)
	if err := scheduler.StopScan(s.ID); err.Error() != fmt.Sprintf("process for scan id %s does not exist", s.ID) {
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
	scheduler := NewScheduler(MockPubSub{}, "testID", conf, "", MockResolveVT)
	scheduler.commander = MockCommander{}

	for i := 0; i < 30; i++ {
		switch {
		case i < 10:
			scheduler.queue.Enqueue(&scan{ScanPrefs: models.ScanPrefs{ID: fmt.Sprint(i)}})
		case i < 20:
			scheduler.init.Enqueue(&scan{ScanPrefs: models.ScanPrefs{ID: fmt.Sprint(i)}})
		default:
			scheduler.running.Enqueue(&scan{ScanPrefs: models.ScanPrefs{ID: fmt.Sprint(i)}})
		}
	}

	if scheduler.queue.IsEmpty() || scheduler.init.IsEmpty() || scheduler.running.IsEmpty() {
		t.Fatalf("None of the queueLists should be empty\n")
	}

	scheduler.Close()

	if !scheduler.queue.IsEmpty() || !scheduler.init.IsEmpty() || !scheduler.running.IsEmpty() {
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
	scheduler := NewScheduler(MockPubSub{}, "testID", conf, "", MockResolveVT)

	scan1 := &scan{
		ScanPrefs: models.ScanPrefs{
			ID: "foo",
		},
	}
	scan2 := &scan{
		ScanPrefs: models.ScanPrefs{
			ID: "bar",
		},
	}

	if err := scheduler.interruptScan(scan1.ID); err == nil {
		t.Fatalf("Should be unable to interrupt unknsown scan %s", scan1.ID)
	}
	if err := scheduler.interruptScan(scan2.ID); err == nil {
		t.Fatalf("Should be unable to interrupt unknsown scan %s", scan2.ID)
	}

	scheduler.init.Enqueue(scan1)
	scheduler.running.Enqueue(scan2)

	if err := scheduler.interruptScan(scan1.ID); err != nil {
		t.Fatalf("Should be able to interrupt unknsown scan %s", scan1.ID)
	}
	if err := scheduler.interruptScan(scan2.ID); err != nil {
		t.Fatalf("Should be able to interrupt unknsown scan %s", scan2.ID)
	}
}

func TestAddVTInfo(t *testing.T) {
	var conf = config.ScannerPreferences{
		ScanInfoStoreTime:   0,
		MaxScan:             0,
		MaxQueuedScans:      0,
		Niceness:            10,
		MinFreeMemScanQueue: 0,
	}
	scheduler := NewScheduler(MockPubSub{}, "testID", conf, "", MockResolveVT)
	scheduler.commander = MockCommander{}

	vt := make(chan []string)

	if err := scheduler.addVTInfo("foo", vt); err == nil {
		t.Errorf("There should be an error, but there is none")
	}

	s := &scan{ScanPrefs: models.ScanPrefs{
		ID: "foo",
	}}
	scheduler.init.Enqueue(s)
	errChan := make(chan error)
	go func() {
		errChan <- scheduler.addVTInfo("foo", vt)
	}()

	vt <- []string{"oid1", "oid2", "oid2"}
	if err := <-errChan; err != nil {
		t.Fatalf("Error while Adding VT Info: %s", err)
	}

	if len(s.Plugins) != 2 {
		t.Fatalf("Wrong number of Plugins. Expected %d, got %d", 2, len(s.Plugins))
	}
	if s.Plugins[0].OID != "oid1" {
		t.Errorf("Wrong Plugin OID. Expected %s, got %s", "odi1", s.Plugins[0].OID)
	}
	if s.Plugins[1].OID != "oid2" {
		t.Errorf("Wrong Plugin OID. Expected %s, got %s", "odi2", s.Plugins[1].OID)
	}

}

func TestGetScan(t *testing.T) {
	var conf = config.ScannerPreferences{
		ScanInfoStoreTime:   0,
		MaxScan:             0,
		MaxQueuedScans:      0,
		Niceness:            10,
		MinFreeMemScanQueue: 0,
	}
	scheduler := NewScheduler(MockPubSub{}, "testID", conf, "", MockResolveVT)
	scheduler.commander = MockCommander{}

	if _, err := scheduler.GetScan("foo"); err == nil {
		t.Errorf("Expected error but got none")
	}

	s := &scan{
		ScanPrefs: models.ScanPrefs{
			ID: "foo",
		},
		Ready: true}
	scheduler.init.Enqueue(s)
	scan, err := scheduler.GetScan("foo")
	if err != nil {
		t.Fatalf("Error while getting scan: %s", err)
	}
	if scan.ID != s.ID {
		t.Fatalf("Different IDs. Original %s, altered %s", s.ID, scan.ID)
	}
}
