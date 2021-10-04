package sensor

import (
	"encoding/json"
	"testing"
	"time"

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
	startStopHandler := ScanCmd{
		Stop: ms.StopScan,
	}

	stopMsg := cmds.Stop{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("stop.scan", "", ""),
			ID:      "foo",
		},
	}

	stopMsgJson, err := json.Marshal(stopMsg)
	if err != nil {
		t.Fatal("Transform Stop Msg into JSON failed\n")
	}

	// Simulate an Event trigger
	startStopHandler.On("", stopMsgJson)

	// Check if fields are set
	if !ms.scanStopped {
		t.Fatal("Scan should be stopped\n")
	}
}

// TestRegistered tests the functionality of the Registered handler
func TestRegistered(t *testing.T) {
	regChan := make(chan struct{}, 1)
	sensorID := "foo"
	registered := &Registered{
		Register: regChan,
		ID:       sensorID,
	}

	registeredMsg := info.Modified{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modified.sensor", "", ""),
			ID:      sensorID,
		},
	}

	registeredMsgJSON, err := json.Marshal(registeredMsg)

	if err != nil {
		t.Fatal("Transform registered Msg into JSON failed")
	}

	registered.On("", registeredMsgJSON)

	select {
	case <-regChan:
	case <-time.After(time.Millisecond * 50):
		t.Fatal("Unable to register\n")
	}
}

func TestStatus(t *testing.T) {
	ms := NewMockSensor()
	statusHandler := ScanInfo{
		Run: ms.ScanRunning,
		Fin: ms.ScanFinished,
	}

	// Creating running and finished messages
	runningMsg := info.Status{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("status.scan", "", ""),
			ID:      "foo",
		},
		Status: "running",
	}
	finishMsg := info.Status{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("status.scan", "", ""),
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
