package cmds

import (
	"github.com/greenbone/eulabeia/messages"
)

// Create indicates that a new entity should be created.
// The type of of entity is indicated by `message_type`
// e.g. "message_type": "create.target" creates a target.
type Create struct {
	messages.Message
}

// Get is used by a client to get the latest snapshot of an aggregate.
//
// The response for Get is usually the aggragte with Message information and can be found within a model.
type Get struct {
	messages.Identifier
}

// Delete is used by a client to delete the latest snapshot of an aggregate.
//
// The response of Delete is Deleted.
type Delete struct {
	messages.Identifier
}

// Start indicates that something with the ID should be started.
//
// As an example an event with the stype start.scan with the id 1 would start scan id 1
type Start struct {
	messages.Identifier
}

// Modify indicates that a entity should be modified.
//
// The values to modify are within Values, they typically match the names of the aggregate
// but with lower case starting.
type Modify struct {
	messages.Identifier
	Values map[string]interface{} `json:"values"`
}