package messagehandler

import (
	"encoding/json"
	"fmt"
	"github.com/greenbone/eulabeia/messages"
	"io"
	"testing"
)

type FakeWriter struct {
	data []byte
}

func (fw *FakeWriter) Write(data []byte) (n int, err error) {
	fw.data = data
	return 42, nil
}

func createMessageHandler(writer io.Writer, changeEvent OnChangeEvent) OnMessage {
	h, _ := NewMessageHandler(writer, []OnChangeEvent{OnCreateTarget{}})
	return h
}

func TestSuccessResponse(t *testing.T) {
	writer := FakeWriter{}
	var tests = []struct {
		on      interface{}
		then    string
		handler OnMessage
	}{
		{messages.Create{
			Message: messages.Message{
				MessageType: "create.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
		},
			"messages.Created",
			createMessageHandler(&writer, OnCreateTarget{})},
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
