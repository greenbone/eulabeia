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
