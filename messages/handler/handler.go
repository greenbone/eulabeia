// Package handler implements various handler for events and messages
package handler

import (
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
)

// Starter is the interface that wraps the basic Start method.
//
// Start is used to start an new event chain (e.g. start.scan)
type Starter interface {
	Start(cmds.Start) (messages.Event, *info.Failure, error)
}

// Creater is the interface that wraps the basic Create method.
//
// Create is used on aggregate handler to handle messages.Create.
// Creates a new entity of a given type via messages.Message.MessageType.
// It responds with info.Created which contains the id of the entity.
type Creater interface {
	Create(cmds.Create) (*info.Created, error)
}

// Modifier is the interface that wraps the basic Modify method.
//
// Modifies a entity of a given type via messages.Message.MessageType and ID.
// It responds with info.Modified on successful alteration
// info.Failure on incorrect Values
type Modifier interface {
	Modify(cmds.Modify) (*info.Modified, *info.Failure, error)
}

// Getter is the interface that wraps the basic Get method.
//
// Gets a entity of a given type via messages.Message.MessageType and ID.
// It responds with interface{} on success and info.Failure when not found.
type Getter interface {
	Get(cmds.Get) (messages.Event, *info.Failure, error)
}

// Deleter is the interface that wraps the basic Delete method.
//
// Delets a entity of a given type found via messages.Message,MessageType and ID.
// It responds with info.Deleted on successful removal
type Deleter interface {
	Delete(cmds.Delete) (*info.Deleted, *info.Failure, error)
}

// Aggregate is the interface to handle Aggregate messages
type Aggregate interface {
	Creater
	Modifier
	Deleter
	Getter
}
