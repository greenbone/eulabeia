// Package memory contains message handler for get.memory
package memory

import (
	"fmt"

	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	mem "github.com/mackerelio/go-osstat/memory"
)

type getMemory struct {
	stats func() (*mem.Stats, error)
}

func (gm getMemory) Get(g messages.Get) (interface{}, *messages.Failure, error) {
	s, err := gm.stats()
	if err != nil {
		return nil, nil, err
	}
	if s == nil {
		return nil, nil, fmt.Errorf("unable to get memory on message %s", g.MessageID)
	}
	response := models.GotMemory{
		Message: messages.NewMessage("got.memory", g.MessageID, g.GroupID),
		ID:      g.ID,
		Total:   fmt.Sprintf("%dB", s.Total),
		Used:    fmt.Sprintf("%dB", s.Used),
		Cached:  fmt.Sprintf("%dB", s.Cached),
		Free:    fmt.Sprintf("%dB", s.Free),
	}
	return &response, nil, nil
}

func New() handler.Holder {
	return handler.FromGetter("memory", getMemory{
		stats: mem.Get,
	})
}
