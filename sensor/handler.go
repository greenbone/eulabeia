package sensor

import (
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

type StartStop struct {
	Start func(scanID string) error // Function to Start a scan
	Stop  func(scanID string) error // Function to Stop a scan
}

func (handler StartStop) On(
	topic string,
	message []byte,
) (*connection.SendResponse, error) {
	var msg cmds.IDCMD
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	if mt.Aggregate == "scan" {
		switch mt.Function {
		case "start":
			if err := handler.Start(msg.ID); err != nil {
				log.Printf("Unable to start scan: %s", err)
			}
		case "stop":
			if err := handler.Stop(msg.ID); err != nil {
				log.Printf("Unable to stop scan: %s", err)
			}
		}
	}
	return nil, nil
}

type Registered struct {
	Register chan struct{} // Channel to signal succesful registration
	ID       string        // SensorID to compare registered ID with own
}

func (handler Registered) On(
	topic string,
	message []byte,
) (*connection.SendResponse, error) {
	var msg info.Created
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	mt, err := messages.ParseMessageType(msg.Type)
	if err != nil {
		return nil, err
	}
	if msg.ID == handler.ID && mt.Function == "modified" &&
		mt.Aggregate == "sensor" {
		select {
		case handler.Register <- struct{}{}:
		default:
			log.Trace().Msgf("Ignoring modified sensor (%s); it is already registered", msg.ID)
		}

	}
	return nil, nil
}

type Status struct {
	Run func(string) error // Function to mark a scan as running
	Fin func(string) error // Function to mark a scan as finished
}

func (handler Status) On(
	topic string,
	message []byte,
) (*connection.SendResponse, error) {
	var msg info.Status
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	switch msg.Status {
	case "running":
		if err := handler.Run(msg.ID); err != nil {
			log.Printf("Unable to set status to running: %s", err)
		}
	case "finished":
		if err := handler.Fin(msg.ID); err != nil {
			log.Printf("Unable to set status to finished: %s", err)
		}
	}
	return nil, nil
}

type LoadVTs struct {
	VtsLoad func() // Function to start LoadingVTs (into redis by openvas)
}

func (handler LoadVTs) On(
	topic string,
	message []byte,
) (*connection.SendResponse, error) {
	handler.VtsLoad()
	return nil, nil
}
