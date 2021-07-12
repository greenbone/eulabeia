package cmds

import (
	"github.com/greenbone/eulabeia/messages"
)

type eventType struct{}

func (eventType) Event() messages.EventType {
	return messages.CMD
}

// IDCMD is a command with just an ID to identify a specific entity
type IDCMD struct {
	eventType
	messages.Identifier
}

// Create indicates that a new entity should be created.
// The type of of entity is indicated by `message_type`
// e.g. "message_type": "create.target" creates a target.
type Create struct {
	eventType
	messages.Message
}

// NewCreate creates a new Create message for given aggregate in given groupID
//
// When the groupID is empty then a random uuid will be chosen.
func NewCreate(aggregate string, groupID string) Create {
	return Create{
		Message: messages.NewMessage("create." + aggregate, "", groupID),
	}
}

// Get is used by a client to get the latest snapshot of an aggregate.
//
// The response for Get is usually the aggragte with Message information and can be found within a model.
type Get IDCMD

// Delete is used by a client to delete the latest snapshot of an aggregate.
//
// The response of Delete is Deleted.
type Delete IDCMD

// Start indicates that something with the ID should be started.
//
// As an example an event with the stype start.scan with the id 1 would start scan id 1
type Start IDCMD

// Modify indicates that a entity should be modified.
//
// The values to modify are within Values, they typically match the names of the aggregate
// but with lower case starting.
type Modify struct {
	eventType
	messages.Identifier
	Values map[string]interface{} `json:"values"`
}
