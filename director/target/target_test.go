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

package target

import (
	"testing"

	"github.com/greenbone/eulabeia/internal/test"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/storage"
)

func TestSuccessResponse(t *testing.T) {
	h := New(&storage.InMemory{Pretend: true})
	tests := []test.HandleTests{
		{
			Input: cmds.Create{
				Message: messages.NewMessage("create.target", "1", "1"),
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("created.target", "1", "1"),
		},
		{
			Input: cmds.Get{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("get.target", "1", "1"),
					ID:      "someid",
				},
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("got.target", "1", "1"),
		},
		{
			Input: cmds.Delete{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("delete.target", "1", "1"),
					ID:      "someid",
				},
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("deleted.target", "1", "1"),
		},
		{
			Input: cmds.Modify{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("modify.target", "1", "1"),
					ID:      "1",
				},
				Values: map[string]interface{}{
					"sensor":   "openvas",
					"hosts":    []string{"a", "b"},
					"plugins":  []string{"a", "b"},
					"alive":    true,
					"parallel": false,
					"exclude":  []string{"host1"},
					"credentials": map[string]map[string]string{
						"ssh": {"username": "nobody"},
					},
				},
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("modified.target", "1", "1"),
		},
	}
	for _, test := range tests {
		test.Verify(t)
	}
}
