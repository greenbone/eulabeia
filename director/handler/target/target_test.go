package target

import (
	"encoding/json"
	"fmt"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"testing"
)

func createMessageHandler(changeEvent handler.OnEvent) connection.OnMessage {
	h, _ := handler.New([]handler.OnEvent{changeEvent})
	return h
}

func TestSuccessResponse(t *testing.T) {
	var tests = []struct {
		on      interface{}
		then    string
		handler connection.OnMessage
	}{
		{messages.Get{
			Message: messages.Message{
				MessageType: "get.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
			ID: "someid",
		},
			"models.GotTarget",
			createMessageHandler(OnGet{storage: NoopStorage{}})},
		{messages.Create{
			Message: messages.Message{
				MessageType: "create.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
		},
			"messages.Created",
			createMessageHandler(OnCreate{storage: NoopStorage{}})},
		{messages.Create{
			Message: messages.Message{
				MessageType: "created.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
		},
			"<nil>",
			createMessageHandler(OnCreate{storage: NoopStorage{}})},
		{messages.Modify{
			Message: messages.Message{
				MessageType: "modify.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
			ID: "1",
			Values: map[string]interface{}{
				"scanner":  "openvas",
				"hosts":    []string{"a", "b"},
				"plugins":  []string{"a", "b"},
				"alive":    true,
				"parallel": false,
			},
		},
			"messages.Modified",
			createMessageHandler(OnModify{storage: NoopStorage{}})},
	}
	for i, test := range tests {
		b, err := json.Marshal(test.on)
		if err != nil {
			t.Errorf("[%d] marshalling [%v] to json failed", i, test.on)
		}
		r, err := test.handler.On(b)
		if err != nil {
			t.Errorf("[%d] returned err (%v) on: %v", i, err, test.on)
		}
		ts := fmt.Sprintf("%T", r)
		if ts != test.then {
			t.Errorf("[%d] returned %v while expecting %v on: %v", i, ts, test.then, test.on)
		}
	}
}
