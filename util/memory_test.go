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

package util

import (
	"testing"

	mem "github.com/mackerelio/go-osstat/memory"
)

type helperMemoryManager struct {
}

func (mm helperMemoryManager) Get() (*mem.Stats, error) {
	return &mem.Stats{
		Available: 4294967296, // = 4GiB
	}, nil
}

func TestAvailableMemory(t *testing.T) {
	m, _ := GetAvailableMemory(helperMemoryManager{})

	if m.Bytes != 4294967296 {
		t.Fatalf("Error: expected %d, got %d", 4294967296, m.Bytes)
	}

	mString := m.String()

	if mString != "4.0 GiB" {
		t.Fatalf("Error: expected %s, got %s", "4.0GiB", mString)
	}
}
