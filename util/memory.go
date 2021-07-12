package util

import (
	"fmt"

	mem "github.com/mackerelio/go-osstat/memory"
)

type Memory uint64

func (m Memory) String() string {
	calc := float64(m)
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

func GetAvailableMemory() (Memory, error) {
	s, err := mem.Get()
	if err != nil {
		return 0, err
	}
	return Memory(s.Available), nil
}
