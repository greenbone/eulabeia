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

package sensor

import (
	"testing"

	"github.com/greenbone/eulabeia/internal/test"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/storage"
)

func TestSensor(t *testing.T) {
	h := handler.New("eulabeia", New(storage.Noop{}))
	tests := []test.HandleTests{
		{
			Input: cmds.Create{
				Message: messages.NewMessage("create.sensor", "1", "1"),
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("created.sensor", "1", "1"),
		},
		{
			Input: cmds.Modify{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("modify.sensor", "1", "2"),
					ID:      "123",
				},
				Values: map[string]interface{}{
					"type": "openvas",
				},
			},
			ExpectedMessage: messages.NewMessage("modified.sensor", "1", "2"),
			Handler:         h,
		},
		{
			Input: cmds.Get{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("get.sensor", "1", "2"),
					ID:      "123",
				},
			},
			ExpectedMessage: messages.NewMessage("got.sensor", "1", "2"),
			Handler:         h,
		},
	}
	for _, test := range tests {
		test.Verify(t)
	}

}
