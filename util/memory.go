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
	"fmt"

	mem "github.com/mackerelio/go-osstat/memory"
)

type Memory struct {
	Bytes uint64
}

type MemoryManager interface {
	Get() (*mem.Stats, error)
}

type StdMemoryManager struct {
}

func (mm StdMemoryManager) Get() (*mem.Stats, error) {
	return mem.Get()
}

func (m Memory) String() string {
	calc := float64(m.Bytes)
	i := 0
	for ; calc > 1024; i++ {
		calc /= 1024
	}
	size := "B"
	switch i {
	case 1:
		size = "KiB"
	case 2:
		size = "MiB"
	case 3:
		size = "GiB"
	case 4:
		size = "TiB"
	case 5:
		size = "PiB"
	case 6:
		size = "EiB"
	}

	return fmt.Sprintf("%.1f %s", calc, size)
}

func GetAvailableMemory(mm MemoryManager) (Memory, error) {
	s, err := mm.Get()
	if err != nil {
		return Memory{0}, err
	}
	return Memory{s.Available}, nil
}
