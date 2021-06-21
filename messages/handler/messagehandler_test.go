package handler

import (
	"encoding/json"
	"fmt"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"testing"
)

type OnNilResponse struct{}

func (oct OnNilResponse) On(messageType string, message []byte) (interface{}, error) {
	return nil, nil
}

type OnErrorResponse struct{}

func (oct OnErrorResponse) On(messageType string, message []byte) (interface{}, error) {
	return nil, fmt.Errorf("nope")
}

type OnResponse struct{}

func (oct OnResponse) On(messageType string, message []byte) (interface{}, error) {
	return "something", nil
}
func createMessageHandler(changeEvent OnEvent) connection.OnMessage {
	h, _ := New([]OnEvent{changeEvent})
	return h
}

func createMessage(typus string) messages.Create {
	return messages.Create{
		Message: messages.Message{
			MessageType: typus,
			Created:     7774,
			MessageID:   "1",
			GroupID:     "12",
		},
	}
}

func TestNotContainMessageType(t *testing.T) {
	h, _ := New([]OnEvent{OnErrorResponse{}})
	_, err := h.On([]byte{0})
	if err == nil {
		t.Error("Expected error but got none")
	}

}
func TestReturnError(t *testing.T) {
	h, _ := New([]OnEvent{OnErrorResponse{}})
	b, err := json.Marshal(createMessage("errpr"))
	if err != nil {
		t.Error("marshalling to json failed")
	}
	_, err = h.On(b)
	if err == nil || err.Error() != "nope" {
		t.Errorf("Expected nope error but got %s", err)
	}

}
func TestSuccessResponse(t *testing.T) {
	var tests = []struct {
		on      interface{}
		then    string
		handler connection.OnMessage
	}{
		{createMessage("create.target"), "string", createMessageHandler(OnResponse{})},
		{createMessage("nil.response"), "<nil>", createMessageHandler(OnNilResponse{})},
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
