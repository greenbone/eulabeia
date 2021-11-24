package vt

import (
	"testing"

	"github.com/greenbone/eulabeia/messages/cmds"
)

func TestResponse(t *testing.T) {
	h := New("testSensor")

	resp, _, _ := h.Getter.Get(cmds.NewGet("vt", "testOID", "", ""))
	if resp.MessageType().Destination != "testSensor" {
		t.Fatalf(
			"Destination %s is wrong it should be testSensor",
			resp.MessageType().Destination,
		)
	}

}
