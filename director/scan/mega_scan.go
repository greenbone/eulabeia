package scan

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
)

// MeagScan is a start.scan event containing the scan aggregate directly.
//
// It will be used to identify if it should be split up to:
// - modify.target
// - modify.scan
// - start.scan
type StartMegaScan struct {
	cmds.EventType
	messages.Message
	models.Scan
}

func toJson(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func toTopicData(context string, v messages.Event) connection.TopicData {
	return connection.TopicData{
		Topic:   messages.EventToResponse("scanner", v).Topic,
		Message: toJson(v),
	}
}

type ScanPreprocessor struct {
	Context string
}

func (s ScanPreprocessor) Preprocess(topic string, payload []byte) ([]connection.TopicData, bool) {
	if !strings.HasSuffix(topic, "/scan/cmd/director") {
		return nil, false
	}
	mt, err := handler.ParseMessageType(payload)
	if err != nil || mt.Function != "start" && mt.Aggregate != "scan" {
		return nil, false
	}
	var sms StartMegaScan
	if err = json.Unmarshal(payload, &sms); err != nil {
		return nil, false
	}
	if len(sms.Hosts) == 0 {
		return nil, false
	}
	// we use the principle that on modify it will be created when not found
	target := cmds.NewModify("target", sms.ID, map[string]interface{}{
		"hosts":       sms.Hosts,
		"ports":       sms.Ports,
		"plugins":     sms.Plugins,
		"sensor":      sms.Sensor,
		"aliveTest":   sms.AliveTest,
		"parallel":    sms.Parallel,
		"exclude":     sms.Exclude,
		"credentials": sms.Credentials,
	}, "director", sms.GroupID)
	scan := cmds.NewModify("scan", sms.ID, map[string]interface{}{
		"finished":  sms.Finished,
		"temporary": true,
	}, "director", sms.GroupID)
	start := cmds.NewStart("scan", sms.ID, "director", sms.GroupID)
	fmt.Printf("Using context: %s\n", s.Context)
	return []connection.TopicData{
		toTopicData(s.Context, target),
		toTopicData(s.Context, scan),
		toTopicData(s.Context, start),
	}, true

}
