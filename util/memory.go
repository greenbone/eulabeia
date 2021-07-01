package util

import (
	"fmt"

	mem "github.com/mackerelio/go-osstat/memory"
)

type Memory struct {
	Bytes uint64
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

func GetAvailableMemory() (m Memory, err error) {
	var s *mem.Stats
	s, err = mem.Get()
	if err == nil {
		m = Memory{
			s.Available,
		}
	}
	return
}
