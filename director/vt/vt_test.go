package vt

import (
	"encoding/json"
	"testing"

	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

const success = "{\"id\":\"testScan\",\"created\":0,\"message_type\":\"get.vt\",\"message_id\":\"0\",\"group_id\":\"0\",\"oid\":\"testOID\"}"

func TestResponse(t *testing.T) {
	s := storage.InMemory{Pretend: false}
	h := New(&s, "test")

	m := models.Scan{
		ID: "testScan",
		Target: models.Target{
			Sensor: "testSensor",
		},
	}

	s.Put("testScan", m)

	resp, err := h.On("", []byte(success))
	if err != nil {
		t.Fatal(err)
	}
	if resp.Topic != "test/vt/cmd/testSensor" {
		t.Fatalf("Wrong topic.\nShould be:\n%s\nIs:\n%s\n", "test/vt/cmd/testSensor", resp.Topic)
	}
	respString, err := json.Marshal(resp.MSG)
	if err != nil {
		t.Fatal(err)
	}
	if string(respString) != success {
		t.Fatalf("Error in response.\nShould be:\n%s\n\nIs:\n%s\n", success, resp.MSG)
	}

}
