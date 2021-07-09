package test

import (
	"encoding/json"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
	"reflect"
	"testing"
)

var VerifyNilError = func(e error, _ HandleTests, t *testing.T) {
	if e != nil {
		t.Errorf("returned err (%v) on", e)
	}
}

var VerifyMessageOfResult = func(d *connection.SendResponse, h HandleTests, t *testing.T) {

	tv := reflect.ValueOf(d.MSG)
	if tv.Kind() != reflect.Ptr || tv.IsNil() {
		t.Fatal(&models.InvalidTargetError{Type: reflect.TypeOf(t)})
	}
	tve := tv.Elem()
	if tve.Kind() != reflect.Struct {
		t.Fatal(&models.InvalidTargetError{Type: tve.Type()})
	}
	f := tve.FieldByName("Message")
	if !f.IsValid() || !f.CanSet() {
		t.Fatal(&models.InvalidFieldError{Type: tve.Type(), Field: "Message"})
	}
	rm, ok := f.Interface().(messages.Message)
	if !ok {
		t.Fatalf("Unable to get message from %v: %T", d, d)
	}
	if rm.GroupID != h.ExpectedMessage.GroupID {
		t.Errorf("Expected GroupID to be: %s but was %s", h.ExpectedMessage.GroupID, rm.GroupID)
	}
	if rm.MessageID != h.ExpectedMessage.MessageID {
		t.Errorf("Expected MessageID to be: %s but was %s", h.ExpectedMessage.MessageID, rm.MessageID)
	}
	if rm.Type != h.ExpectedMessage.Type {
		t.Errorf("Expected MessageType to be: %s but was %s", h.ExpectedMessage.Type, rm.Type)
	}
}

type HandleTests struct {
	Input           interface{}
	Handler         connection.OnMessage
	ExpectedMessage messages.Message
	VerifyError     func(error, HandleTests, *testing.T)
	VerifyResult    func(*connection.SendResponse, HandleTests, *testing.T)
}

func (h *HandleTests) Verify(t *testing.T) {
	b, err := json.Marshal(h.Input)
	if err != nil {
		t.Errorf("marshalling [%v] to json failed", h.Input)
	}
	if h.Handler == nil {
		t.Fatalf("Handler is not set")
	}
	r, err := h.Handler.On("some", b)
	if h.VerifyError != nil {
		h.VerifyError(err, *h, t)
	} else {
		VerifyNilError(err, *h, t)
		if h.VerifyResult != nil {
			h.VerifyResult(r, *h, t)
		} else {
			VerifyMessageOfResult(r, *h, t)
		}
	}

}
