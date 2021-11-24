package vt

import (
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/messages/info"
)

type vtHandler struct {
	sensor string
}

func (t vtHandler) Get(g cmds.Get) (messages.Event, *info.Failure, error) {
	return cmds.NewGet("vt", "", t.sensor, g.GroupID), nil, nil
}

func New(sensor string) handler.Container {
	return handler.FromGetter("vt", vtHandler{sensor})
}
