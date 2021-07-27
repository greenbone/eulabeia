# Topics

In Eulabeia we follow the topic structure:

`<group>/<aggregate>/<event>/<destination>`

- group is the actual project (`eulabeia`)
- aggregate is the aggregate this message belongs to (e.g. `scan`)
- event indicats if the message is an info or a cmd (either `cmd` or `info`)
- destination if the message is for a specific `sensor` it contains the id of that (e.d. `openvas_sensor1`)

## Meta Information within each `eulabeia` message

To be able to trace a message although the topic information are lost and for easier transition to another principle of message handling the information of a message type, aggregate and destination is also stored within the message.

Each message must contain those fields:

```
{
  "created": 443947894,
  "message_type": "modify.target",
  "message_id": "3b1489a9-849a-42d4-9506-5a13b5912eb9",
  "group_id": "12"
}
```

- created is a timestamp when this message got created
- message_type identifies the type of message and is following the pattern `<cmd/info_indidactor>.<aggregate>.<destination>`
- message_id the id of the message
- group_id of the message, indicating that multiple events belong to the same group


## `cmd` message structure

`cmds` are limited to

- create, created a new entity. The ID of the newly created entity will be send with an created info message.
- modify, modifies a existing or creates a new entityt when not found. The values to set muste be send within a `values` object.

- delete, deletes an existing entity.
- get, retrieves an existing entityt.
- start, starts an event chain based on the information of aggregate and entity ID
- stop, stops an event chain.

All but the modify and create message just contain meta information about the message as well as the ID to identify an aggregate entity.

Additionally to the ID the modify message contains map of values to change. The key of the values will be mapped to the key of the actual aggregate and the value of that map will then override the value of the entity found via ID.

The create cmd will create a new entity and the id of that entity will be returned by the created information message. 

### Create Example

```
{
  "id": "f704d1e0-768d-4a86-ab6a-dd28e3f45776",
  "created": 443397956,
  "message_type": "created.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
}
```

### Modify Example

```
{
  "created": 443947894,
  "message_type": "modify.target",
  "message_id": "3b1489a9-849a-42d4-9506-5a13b5912eb9",
  "group_id": "12",
  "id": "f704d1e0-768d-4a86-ab6a-dd28e3f45776",
  "values": {
    "credentials": {
        "ssh": {
          "password": "admin",
          "username": "admin"
        }
    },
    "hosts": [
      "localhorst"
    ],
    "plugins": [
      "someoids"
    ]
  }
}
```

### Get Example

```
{
  "id": "f704d1e0-768d-4a86-ab6a-dd28e3f45776",
  "created": 445205931,
  "message_type": "get.target",
  "message_id": "dbbf93b6-6a46-4a7b-8d9c-39ef041aa0a5",
  "group_id": "13b8b480-6121-4b48-9d05-17dab6b76359"
}
```

