package messages

// Message contains the meta data for each sent message.
// It should be embedded into all messages send to or received by eulabia.
type Message struct {
	Created     int    `json:"created"`      // Timestamp when this message was created
	MessageType string `json:"message_type"` // Identifier what this message actually contains
	MessageID   string `json:"message_id"`   // The ID of a message, responses will have the same ID
	GroupID     string `json:"group_id"`     // The ID of a group of messages, responses will have the same ID
}

// Create indicates that a new entity should be created.
// The type of of entity is indicated by `message_type`
// e.g. "message_type": "create.target" creates a target.
type Create struct {
	Message
}

// Created is returned by a create event and contains the `created_id` as an identifier for the created entity.
// The type of entity is indicated by `message_type`.
// e.g. on "message_type": "created.target" the `created_id` is a identifier for a target.
type Created struct {
	CreatedID string `json:"created_id"`
	Message
}
