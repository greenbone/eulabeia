package handler

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
	startStopHandler := StartStop{
		StartFunc: ms.StartScan,
		StopFunc:  ms.StopScan,
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
		t.Fatalf("Transform Start Msg into JSON failed")
	}
	stopMsgJson, err := json.Marshal(stopMsg)
	if err != nil {
		t.Fatalf("Transform Stop Msg into JSON failed")
	}

	// Simulate an Event trigger
	startStopHandler.On("", startMsgJson)
	startStopHandler.On("", stopMsgJson)

	// Check if fields are set
	if !ms.scanStarted {
		t.Fatalf("Scan should be started")
	}
	if !ms.scanStopped {
		t.Fatalf("Scan should be stopped")
	}
}

// TestRegistered tests the functionality of the Registered handler
func TestRegistered(t *testing.T) {
	regChan := make(chan struct{}, 1)
	sensorID := "foo"
	registered := &Registered{
		RegChan: regChan,
		ID:      sensorID,
	}

	registeredMsg := info.Modified{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modified.sensor", "", ""),
			ID:      sensorID,
		},
	}

	registeredMsgJSON, err := json.Marshal(registeredMsg)

	if err != nil {
		t.Fatalf("Transform registered Msg into JSON failed")
	}

	registered.On("", registeredMsgJSON)

	select {
	case <-regChan:
	case <-time.After(time.Millisecond * 50):
		t.Fatalf("Unable to register")
	}
}

func TestStatus(t *testing.T) {
	ms := NewMockSensor()
	statusHandler := Status{
		RunFunc: ms.ScanRunning,
		FinFunc: ms.ScanFinished,
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
		t.Fatalf("Transform Start Msg into JSON failed")
	}
	finishMsgJson, err := json.Marshal(finishMsg)
	if err != nil {
		t.Fatalf("Transform Stop Msg into JSON failed")
	}

	// Simulate an Event trigger
	statusHandler.On("", runningMsgJson)
	statusHandler.On("", finishMsgJson)

	// Check if fields are set
	if !ms.scanRunning {
		t.Fatalf("Scan should be started")
	}
	if !ms.scanFinished {
		t.Fatalf("Scan should be stopped")
	}
}

func TestLoadVts(t *testing.T) {
	ms := NewMockSensor()
	vtsHandler := LoadVTs{
		VtsFunc: ms.LoadVts,
	}

	// Creating running and finished messages
	vtsMsg := cmds.LoadVTs{
		Message: messages.NewMessage("load.vts", "", ""),
	}

	// Transforming Messages into Bytes slice
	vtsMsgJson, err := json.Marshal(vtsMsg)
	if err != nil {
		t.Fatalf("Transform Start Msg into JSON failed")
	}

	// Simulate an Event trigger
	vtsHandler.On("", vtsMsgJson)

	// Check if fields are set
	if !ms.vtsLoaded {
		t.Fatalf("Scan should be started")
	}
}
