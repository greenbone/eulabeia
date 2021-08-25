# Topics

In Eulabeia we follow the topic structure:

`<contex<aggregate>/<event>/<destination>`

- context is the context where the message is exchanged (e.g. `scanner`)
- aggregate is the aggregate this message belongs to (e.g. `scan`)
- event indicats if the message is an info or a cmd (either `cmd` or `info`)
- destination if the message is for a specific `sensor` it contains the id of that (e.d. `openvas_sensor1`)

## Message-Information

To be able to trace a message although the topic information are lost and for easier transition to another principle of message handling the information of a message type, aggregate and destination is also stored within the message.

Each message must contain the fields:

```
{
  "created": 443947894,
  "message_type": "modify.target",
  "message_id": "3b1489a9-849a-42d4-9506-5a13b5912eb9",
  "group_id": "12"
}
```

to be recognized.

- created is a timestamp when this message got created
- message_type identifies the type of message and is following the pattern `<cinfo_indidactor>.<aggregate>.<destination>`
- message_id the id of the message
- group_id of the message, indicating that multiple events belong to the same group

## Aggregates

1. [target](../models/target.go)
1. [scan](../models/scan.go)
1. [sensor](../models/scan.go)

## Message structure

`cmds` are limited to

| CMD | Description | Example |
| --- | --- | --- |
| create | creates a new entity. The ID of the newly created entity will be send with an created info message. | [create scan](message_examples.md#createscan)
| modify | modifies a existing or creates a new entityt when not found. The values to set muste be send within a `values` object. | [modify scan](message_examples.md#modifyscan)
| delete | deletes an existing entity. | [delete scan](message_examples.md#deletescan)
| get | retrieves an existing entityt. | [get scan](message_examples.md#getscan)
| start | starts an event chain based on the information of aggregate and entity ID | [start scan](message_examples.md#startscan)
| stop | stops an event chain. | [stop scan](message_examples.md#stopscan)

All but the modify and create contain an `id` to identify an aggregate entity.

Additionally to the ID the modify message contains map of values to change.
The key of the values will be mapped to the key of the actual aggregate and the value of that map will then override the value of the entity found via ID.

The create cmd will create a new entity and the id of that entity will be returned by the created information message. 


Besides commands there also info messages:

| INFO | Description | Example |
| --- | --- | --- |
| created | success message for creation of an entity  | [created scan](message_examples.md#createdscan)
| modified | success message for modification of an entity  | [modified scan](message_examples.md#modifiedscan)
| deleted | success message for deletion of an entity | [deleted scan](message_examples.md#deletedscan)
| got | contains the entity of an aggregate | [got scan](message_examples.md#gotscan)
| failure | indicates a failure and contains an error field | [failure scan](message_examples.md#failurescan)
| status | contains status updates | [status scan](message_examples.md#statusscan)
| result | contains result information | [result scan](message_examples.md#resultscan)


