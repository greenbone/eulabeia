package test

import (
	"encoding/json"
	"testing"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
)

var VerifyNilError = func(e error, _ HandleTests, t *testing.T) {
	if e != nil {
		t.Errorf("returned err (%v) on", e)
	}
}

var VerifyMessageOfResult = func(d interface{}, h HandleTests, t *testing.T) {
	var rm messages.Message
	switch cv := d.(type) {
	case *messages.Created:
		rm = cv.Message
	case *messages.Modified:
		rm = cv.Message
	case *models.GotScan:
		rm = cv.Message
	case *models.GotTarget:
		rm = cv.Message
	case *messages.Failure:
		rm = cv.Message
	default:
		t.Fatalf("Unable to get message from %v", d)
	}
	if rm.GroupID != h.ExpectedMessage.GroupID {
		t.Errorf("Expected GroupID to be: %s but was %s", h.ExpectedMessage.GroupID, rm.GroupID)
	}
	if rm.MessageID != h.ExpectedMessage.MessageID {
		t.Errorf("Expected MessageID to be: %s but was %s", h.ExpectedMessage.MessageID, rm.MessageID)
	}
	if rm.MessageType != h.ExpectedMessage.MessageType {
		t.Errorf("Expected MessageType to be: %s but was %s", h.ExpectedMessage.MessageType, rm.MessageType)
	}
}

type HandleTests struct {
	Input           interface{}
	Handler         connection.OnMessage
	ExpectedMessage messages.Message
	VerifyError     func(error, HandleTests, *testing.T)
	VerifyResult    func(interface{}, HandleTests, *testing.T)
}

func (h *HandleTests) Verify(t *testing.T) {
	b, err := json.Marshal(h.Input)
	if err != nil {
		t.Errorf("marshalling [%v] to json failed", h.Input)
	}
	if h.Handler == nil {
		t.Fatalf("Handler is not set")
	}
	r, err := h.Handler.On(b)
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
