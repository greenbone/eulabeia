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
	"testing"

	"github.com/greenbone/eulabeia/internal/test"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/storage"
)

func TestCreateScan(t *testing.T) {
	h := handler.New("eulabeia", New(storage.Noop{}))
	tests := []test.HandleTests{
		{
			Input: cmds.Create{
				Message: messages.NewMessage("create.scan", "1", "1"),
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("created.scan", "1", "1"),
		},
		{
			Input: cmds.Start{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("start.scan", "1", "1"),
					ID:      "1234",
				},
			},
			Handler: h,
			// although NoopStorage for target doesn't have sensor it should just
			// empty string and extend it that way
			ExpectedMessage: messages.NewMessage("start.scan.", "1", "1"),
		},
		{
			Input: cmds.Modify{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("modify.scan", "1", "2"),
					ID:      "123",
				},
				Values: map[string]interface{}{
					"finished":  []string{"1", "2"},
					"target_id": "1",
				},
			},
			ExpectedMessage: messages.NewMessage("modified.scan", "1", "2"),
			Handler:         h,
		},
		{
			Input: cmds.Modify{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("modify.scan", "1", "2"),
					ID:      "123",
				},
				Values: map[string]interface{}{
					"exclude":   []string{"1", "2"},
					"target_id": 1,
				},
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("failure.modify.scan", "1", "2"),
		},
		{
			Input: cmds.Get{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("get.scan", "1", "2"),
					ID:      "123",
				},
			},
			ExpectedMessage: messages.NewMessage("got.scan", "1", "2"),
			Handler:         h,
		},
	}
	for _, test := range tests {
		test.Verify(t)
	}

}
