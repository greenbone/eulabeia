package sensor

import (
	"encoding/json"
	"testing"

	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

type MockSensor struct {
	scanStarted  bool
	scanStopped  bool
	scanRunning  bool
	scanFinished bool
	vtsLoaded    bool
}

func NewMockSensor() *MockSensor {
	return &MockSensor{
		scanStarted:  false,
		scanStopped:  false,
		scanRunning:  false,
		scanFinished: false,
		vtsLoaded:    false,
	}
}

func (ms *MockSensor) StartScan(scanID string) error {
	ms.scanStarted = true
	return nil
}

func (ms *MockSensor) StopScan(scanID string) error {
	ms.scanStopped = true
	return nil
}

func (ms *MockSensor) ScanRunning(scanID string) error {
	ms.scanRunning = true
	return nil
}

func (ms *MockSensor) ScanFinished(scanID string) error {
	ms.scanFinished = true
	return nil
}

func (ms *MockSensor) LoadVts() {
	ms.vtsLoaded = true
}

// TestStartStop tests the functionality of the StartStop handler
func TestStartStop(t *testing.T) {
	ms := NewMockSensor()
	startStopHandler := StartStop{
		Start: ms.StartScan,
		Stop:  ms.StopScan,
	}

	// Creating Start and Stop Scan messages
	startMsg := cmds.Start{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("start.scan", "", ""),
			ID:      "foo",
		},
	}
	stopMsg := cmds.Start{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("stop.scan", "", ""),
			ID:      "foo",
		},
	}

	// Transforming Messages into Bytes slice
	startMsgJson, err := json.Marshal(startMsg)
	if err != nil {
		t.Fatal("Transform Start Msg into JSON failed\n")
	}
	stopMsgJson, err := json.Marshal(stopMsg)
	if err != nil {
		t.Fatal("Transform Stop Msg into JSON failed\n")
	}

	// Simulate an Event trigger
	startStopHandler.On("", startMsgJson)
	startStopHandler.On("", stopMsgJson)

	// Check if fields are set
	if !ms.scanStarted {
		t.Fatal("Scan should be started\n")
	}
	if !ms.scanStopped {
		t.Fatal("Scan should be stopped\n")
	}
}

func TestStatus(t *testing.T) {
	ms := NewMockSensor()
	statusHandler := Status{
		Run: ms.ScanRunning,
		Fin: ms.ScanFinished,
	}

	// Creating running and finished messages
	runningMsg := info.Status{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("start.scan", "", ""),
			ID:      "foo",
		},
		Status: "running",
	}
	finishMsg := info.Status{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("start.scan", "", ""),
			ID:      "foo",
		},
		Status: "finished",
	}

	// Transforming Messages into Bytes slice
	runningMsgJson, err := json.Marshal(runningMsg)
	if err != nil {
		t.Fatal("Transform Start Msg into JSON failed\n")
	}
	finishMsgJson, err := json.Marshal(finishMsg)
	if err != nil {
		t.Fatal("Transform Stop Msg into JSON failed\n")
	}

	// Simulate an Event trigger
	statusHandler.On("", runningMsgJson)
	statusHandler.On("", finishMsgJson)

	// Check if fields are set
	if !ms.scanRunning {
		t.Fatal("Scan should be started\n")
	}
	if !ms.scanFinished {
		t.Fatal("Scan should be stopped\n")
	}
}

func TestLoadVts(t *testing.T) {
	ms := NewMockSensor()
	vtsHandler := LoadVTs{
		VtsLoad: ms.LoadVts,
	}

	// Creating running and finished messages
	vtsMsg := cmds.LoadVTs{
		Message: messages.NewMessage("load.vts", "", ""),
	}

	// Transforming Messages into Bytes slice
	vtsMsgJson, err := json.Marshal(vtsMsg)
	if err != nil {
		t.Fatal("Transform Start Msg into JSON failed\n")
	}

	// Simulate an Event trigger
	vtsHandler.On("", vtsMsgJson)

	// Check if fields are set
	if !ms.vtsLoaded {
		t.Fatal("Scan should be started\n")
	}
}
