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

package models

import (
	"testing"
)

type TestStruct struct {
	A string
	B int
}

func TestSetValueOf(t *testing.T) {
	var tests = []struct {
		name     string
		val      interface{}
		err      string
		expected interface{}
	}{
		{
			"A",
			"b",
			"",
			TestStruct{"b", 0},
		},
		{
			"B",
			"b",
			"field type (int) does not match value type (string)",
			TestStruct{"b", 0},
		},
		{
			"C",
			"b",
			"field (C) not found on target (models.TestStruct)",
			TestStruct{"b", 0},
		},
	}
	for i, test := range tests {
		target := &TestStruct{"a", 0}
		err := SetValueOf(target, test.name, test.val)

		if err != nil {
			if test.err == "" || test.err != err.Error() {
				t.Errorf("[%d] returned '%s' but expected '%v'", i, err, test.err)
			}
		}
		if test.err == "" {
			if test.expected != *target {
				t.Errorf("[%d] returned %v but expected %v", i, target, test.expected)

			}
		}
	}
}
