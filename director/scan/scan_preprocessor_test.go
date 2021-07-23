// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package scan

import (
	"encoding/json"
	"testing"

	"github.com/greenbone/eulabeia/director/target"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	"github.com/greenbone/eulabeia/storage"
)

type testData struct {
	topic   string
	payload []byte
	handled bool
}

func asModify(payload []byte) cmds.Modify {
	var modify cmds.Modify
	if err := json.Unmarshal(payload, &modify); err != nil {
		panic(err)
	}
	return modify
}

func TestStartScanPreprocessor(t *testing.T) {
	mega := StartMegaScan{
		Message: messages.NewMessage("start.scan.director", "", ""),
		Scan: models.Scan{
			ID: "f123",
			Target: models.Target{
				ID:       "f123",
				Hosts:    []string{"hosts1"},
				Ports:    []string{"ports1"},
				Plugins:  []string{"plugins1"},
				Exclude:  []string{"exclude1"},
				Sensor:   "sensor",
				Alive:    true,
				Parallel: true,
				Credentials: map[string]map[string]string{
					"ssh": {
						"private_key": "denkste",
					},
				},
			},
			Finished: []string{"hosts2"},
		},
	}

	td := []testData{
		{
			payload: toJson(cmds.NewStart("scan", "f1", "director", "")),
			handled: false,
		},
		{
			payload: toJson(mega),
			handled: true,
		},
	}
	device := &storage.InMemory{}
	targetHandler := target.New(device)
	scanHandler := New(device)
	preprocessor := StartMegaScan{}
	dataVerifier := func(
		m handler.Modifier,
		payload []byte,
		verify func(interface{}) interface{},
		retriever func() (interface{}, error)) {
		m.Modify(asModify(payload))
		ta, err := retriever()
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}
		if expected := verify(ta); expected != nil {
			t.Fatalf("expected (%v) after modifying equal to %v", ta, expected)
		}

	}
	for _, test := range td {
		topic := "eulabeia/scan/cmd/director"
		if test.topic != "" {
			topic = test.topic
		}
		topicData, ok := preprocessor.Preprocess(topic, test.payload)
		if ok != test.handled {
			t.Fatalf("Expected %v to be %v but it is not.", ok, test.handled)
		}
		if test.handled {
			if len(topicData) != 3 {
				t.Fatalf("Expected 3 new events but got %d", len(topicData))
			}
			dataVerifier(targetHandler.Modifier,
				topicData[0].Message,
				func(i interface{}) interface{} {
					actual := i.(*models.Target)
					if actual.ID == mega.Scan.Target.ID &&
						len(actual.Ports) == len(mega.Scan.Target.Ports) &&
						len(actual.Hosts) == len(mega.Scan.Target.Hosts) &&
						len(actual.Plugins) == len(mega.Scan.Target.Plugins) &&
						actual.Sensor == mega.Scan.Target.Sensor &&
						actual.Alive == mega.Scan.Target.Alive &&
						actual.Parallel == mega.Scan.Target.Parallel &&
						len(actual.Exclude) == len(mega.Scan.Target.Exclude) &&
						len(actual.Credentials) == len(mega.Scan.Target.Credentials) {
						return nil
					}
					return mega.Scan.Target
				},
				func() (interface{}, error) { return target.NewStorage(device).Get(mega.ID) })
			dataVerifier(scanHandler.Modifier,
				topicData[1].Message,
				func(i interface{}) interface{} {
					actual := i.(*models.Scan)
					if actual.ID == mega.Scan.ID &&
						len(actual.Finished) == len(mega.Scan.Finished) {
						return nil
					}
					return mega.Scan
				},
				func() (interface{}, error) { return NewStorage(device).Get(mega.ID) })
			var startScan cmds.Start
			if json.Unmarshal(topicData[2].Message, &startScan) != nil {
				t.Fatal("Expected no error while unmarshalling start.scan")
			}
			if startScan.ID != mega.ID {
				t.Fatalf("Expected %s to be %s", startScan.ID, mega.ID)
			}
			if startScan.Type != "start.scan.director" {
				t.Fatalf("Expected %s to be start.scan.director", startScan.MessageID)

			}

		}

	}

}
