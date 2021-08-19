# Topics

In Eulabeia we follow the topic structure:

`<context>/<aggregate>/<event>/<destination>`

- context is the context where the message is exchanged (e.g. `scanner`)
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

## Aggregates

1. [target](./../models/target.go)
1. [scan](./../models/scan.go)
1. [sensor](./../models/scan.go)

## Message structure

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


Besides commands there also info messages:
- created, success message for creation of an entity 
- modified, success message for modification of an entity 
- deleted, success message for deletion of an entity
- got, contains the entity of an aggregate
- failure, indicates a failure and contains an error field
- status, contains status updates
- result, contains result information

### Create Example

Create is used to create an entity it is a event that does not contain an ID because the ID is generated.

```
{
  "created": 443397956,
  "message_type": "create.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
}
```

#### Response

The response of a create event is either `created.aggregate` or `failure.create.aggregate`.

Created:

```
{
  "created": 443397956,
  "message_type": "create.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "id": "example_scan_1234",
}
```

Failure:

```
{
  "created": 443397956,
  "message_type": "failure.create.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "error": "insufficient space on device."
  "id": "example_scan_1234",
}
```

### Delete Example

Delete is used to delete an entity it is a event that does not contain an ID because the ID is generated.

```
{
  "created": 443397956,
  "message_type": "delete.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12",
  "id": "example_scan1"
}
```

#### Response

The response of a delete event is either `deleted.aggregate` or `failure.delete.aggregate`.

Created:

```
{
  "created": 443397956,
  "message_type": "delete.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "id": "example_scan_1234",
}
```

Failure:

```
{
  "created": 443397956,
  "message_type": "failure.delete.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "error": "Not found"
  "id": "example_scan_1234",
}
```
### Start Example

Start is used to start an entity it is a event that does not contain an ID because the ID is generated.

```
{
  "created": 443397956,
  "message_type": "start.scan",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12",
  "id": "example_scan1"
}
```

#### Response

The response of a start event are:
- status.aggregate
- result.aggregate
- failure

Status:

```
{
  "id": "classic_scan_1",
  "created": 846648159,
  "message_type": "status.scan",
  "message_id": "a738a1d6-aa8b-4fc5-95eb-b633b14aa437",
  "group_id": "6e48650d-2f97-43b8-9469-8264eecfaaef",
  "status": "queued"
}
```

Result:

```
{
  "message_id": "3bdf1db0-dddc-4979-a0c4-f897a18c7422",
  "message_type": "result.scan",
  "group_id": "3bdf1db0-dddc-4979-a0c4-f897a18c7422",
  "created": 1629411808329199900,
  "result_type": "ERRMSG",
  "host_ip": "127.0.0.1",
  "host_name": "localhost",
  "port": "general/tcp",
  "id": "classic_scan_1",
  "oid": "1.3.6.1.4.1.25623.1.0.90022",
  "value": "this is a error message\n",
  "uri": ""
}
```

Failure:

```
{
  "created": 443397956,
  "message_type": "failure.start.scan",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "error": "Not found"
  "id": "example_scan_1234",
}
```

### Modify Example

Modify is used to either modify or create a new entity of an aggregate when it is not existing.

The values within modify must match fields within a aggregate.

```
{
  "message_id": "b5e1015b-6883-4c1e-8e86-47c5be8f5d14",
  "message_type": "modify.target",
  "group_id": "b5e1015b-6883-4c1e-8e86-47c5be8f5d14",
  "created": 1629411006415021800,
  "id": "example_target_1234",
  "values": {
    "sensor": "localhorst",
    "alive": true,
    "hosts": [
      "localhost"
    ],
    "plugins": {
      "single_vts": [
        {
          "oid": "1.3.6.1.4.1.25623.1.0.90022"
        }
      ]
    },
    "ports": [
      "22"
    ]
  }
}
```

#### Response

The response of a modify event is either `modified.aggregate` or `failure.modify.aggregate`.

Created:

```
{
  "created": 443397956,
  "message_type": "modify.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "id": "example_target_1234",
}
```

Failure:

```
{
  "created": 443397956,
  "message_type": "failure.modify.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "error": "insufficient space on device."
  "id": "example_target_1234",
}
```


### Get Example

Get is used to get an entity of an aggregate.

```
{
  "id": "f704d1e0-768d-4a86-ab6a-dd28e3f45776",
  "created": 445205931,
  "message_type": "get.target",
  "message_id": "dbbf93b6-6a46-4a7b-8d9c-39ef041aa0a5",
  "group_id": "13b8b480-6121-4b48-9d05-17dab6b76359"
}
```

#### Response

The response of a get event is either `got.aggregate` or `failure.get.aggregate`.

The got message contains the aggregate directly; e.g. for a scan:

```
{
  "created": 668372989,
  "message_type": "got.scan",
  "message_id": "d5d6ee1f-aa60-4b82-90e2-4a11e17cd3e0",
  "group_id": "42e7fdaf-429a-49e2-bf5c-0e76929901e8",
  "hosts": [
    "localhost"
  ],
  "ports": [
    "22"
  ],
  "plugins": {
    "single_vts": [
      {
        "oid": "1.3.6.1.4.1.25623.1.0.90022",
        "prefs_by_id": null,
        "prefs_by_name": null
      }
    ],
    "vt_groups": null
  },
  "sensor": "localhorst",
  "alive": true,
  "parallel": false,
  "exclude_hosts": null,
  "credentials": null,
  "id": "classic_scan_1",
  "exclude": null,
  "temporary": false
}
```

Failure:

```
{
  "created": 443397956,
  "message_type": "failure.get.scan",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12"
  "error": "Not found."
  "id": "classic_scan_1",
}
```
