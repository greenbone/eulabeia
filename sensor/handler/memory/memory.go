// Package memory contains message handler for get.memory
package memory

import (
	"fmt"

	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	mem "github.com/mackerelio/go-osstat/memory"
)

type GotMemory struct {
	messages.Message
	Total  string `json:"total"`
	Used   string `json:"used"`
	Cached string `json:"cached"`
	Free   string `json:"free"`
}

type getMemory struct {
	stats func() (*mem.Stats, error)
}

func (gm getMemory) Get(g messages.Get) (interface{}, *messages.Failure, error) {
	if g.MessageType != "get.memory" {
		return nil, nil, nil
	}
	s, err := gm.stats()
	if err != nil {
		return nil, nil, err
	}
	if s == nil {
		return nil, nil, fmt.Errorf("unable to get memory on message %s", g.MessageID)
	}
	response := GotMemory{
		Message: messages.NewMessage("got.memory", g.MessageID, g.GroupID),
		Total:   fmt.Sprintf("%dB", s.Total),
		Used:    fmt.Sprintf("%dB", s.Used),
		Cached:  fmt.Sprintf("%dB", s.Cached),
		Free:    fmt.Sprintf("%dB", s.Free),
	}
	return &response, nil, nil
}

func New() handler.Getter {
	return getMemory{
		stats: mem.Get,
	}
}
