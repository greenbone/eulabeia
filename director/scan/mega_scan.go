package scan

import (
	"encoding/json"

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

func toTopicData(v messages.Event) connection.TopicData {
	return connection.TopicData{
		Topic:   messages.EventToResponse("eulabeia", v).Topic,
		Message: toJson(v),
	}
}

func (s StartMegaScan) Preprocess(topic string, payload []byte) ([]connection.TopicData, bool) {
	if topic != "eulabeia/scan/cmd/director" {
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
		"alive":       sms.Alive,
		"parallel":    sms.Parallel,
		"exclude":     sms.Exclude,
		"credentials": sms.Credentials,
	}, "director", sms.GroupID)
	scan := cmds.NewModify("scan", sms.ID, map[string]interface{}{
		"finished":  sms.Finished,
		"temporary": true,
	}, "director", sms.GroupID)
	start := cmds.NewStart("scan", sms.ID, "director", sms.GroupID)
	return []connection.TopicData{
		toTopicData(target),
		toTopicData(scan),
		toTopicData(start),
	}, true

}
