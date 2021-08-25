<!-- DON'T EDIT THIS FILE; INSTEAD RUN: generate_md -->
# Table of content
- [scan](#scan)
  - [create](#createscan)
  - [start](#startscan)
  - [stop](#stopscan)
  - [get](#getscan)
  - [modify](#modifyscan)
  - [created](#createdscan)
  - [modified](#modifiedscan)
  - [stopped](#stoppedscan)
  - [status](#statusscan)
  - [got](#gotscan)
  - [result](#resultscan)
  - [failure.start](#failurestartscan)
  - [failure.stop](#failurestopscan)
  - [failure.create](#failurecreatescan)
  - [failure.modify](#failuremodifyscan)
  - [failure.get](#failuregetscan)
  - [failure](#failurescan)
- [target](#target)
  - [create](#createtarget)
  - [get](#gettarget)
  - [modify](#modifytarget)
  - [created](#createdtarget)
  - [modified](#modifiedtarget)
  - [got](#gottarget)
  - [failure.create](#failurecreatetarget)
  - [failure.modify](#failuremodifytarget)
  - [failure.get](#failuregettarget)
  - [failure](#failuretarget)


# scan

To get type information for e.g. `modify.scan` or `got.scan` please consolidate [ scan model](../models/scan.go)

As a rule of thumb: each type is as shown in the example.

## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "9d9c0a03-2a7d-499a-8adf-d3e0843425b2",
  "message_type" : "create.scan",
  "group_id" : "9d9c0a03-2a7d-499a-8adf-d3e0843425b2",
  "created" : 1629896000451425361
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "9d1c5eb3-28ed-4afc-acde-b51a1c3fb72f",
  "message_type" : "start.scan",
  "group_id" : "9d1c5eb3-28ed-4afc-acde-b51a1c3fb72f",
  "created" : 1629896000451522911,
  "id" : "example.id.scan"
}
```
Responses:

- [status](#statusscan)
- [failure.start](#failurestartscan)
## stop/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "ef37f9f2-0bc5-45fd-b743-54a9ac9c9fdd",
  "message_type" : "stop.scan",
  "group_id" : "ef37f9f2-0bc5-45fd-b743-54a9ac9c9fdd",
  "created" : 1629896000451549540,
  "id" : "example.id.scan"
}
```
Responses:

- [stopped](#stoppedscan)
- [failure.stop](#failurestopscan)
## get/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "613592e9-74d6-4d6d-af3a-692e58f729a4",
  "message_type" : "get.scan",
  "group_id" : "613592e9-74d6-4d6d-af3a-692e58f729a4",
  "created" : 1629896000451572271,
  "id" : "example.id.scan"
}
```
Responses:

- [got](#gotscan)
- [failure.get](#failuregetscan)
## modify/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "ebb40611-c62a-4bd1-a371-d5efce57756e",
  "message_type" : "modify.scan",
  "group_id" : "ebb40611-c62a-4bd1-a371-d5efce57756e",
  "created" : 1629896000451593275,
  "id" : "example.scan.id",
  "values" : {
    "temporary" : false,
    "target_id" : "example.target.id"
  }
}
```
To get type information please consolidate [ scan model](../models/scan.go)


Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8666f202-abde-4ba3-950b-6a18907602ba",
  "message_type" : "created.scan",
  "group_id" : "8666f202-abde-4ba3-950b-6a18907602ba",
  "created" : 1629896000451619396,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d28115b4-9348-437f-881f-f8893cf823ff",
  "message_type" : "modified.scan",
  "group_id" : "d28115b4-9348-437f-881f-f8893cf823ff",
  "created" : 1629896000451636621,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b9d963fa-8ffc-4dc2-b2ee-1af5a4b9867b",
  "message_type" : "stopped.scan",
  "group_id" : "b9d963fa-8ffc-4dc2-b2ee-1af5a4b9867b",
  "created" : 1629896000451648980,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8466eb24-f666-49b7-93af-6fdf021749ae",
  "message_type" : "status.scan",
  "group_id" : "8466eb24-f666-49b7-93af-6fdf021749ae",
  "created" : 1629896000451660386,
  "id" : "example.id.scan",
  "status" : "requested"
}
```
Valid `status` are:
- `requested`
- `queued`
- `init`
- `running`
- `stopping`
- `stopped`
- `interrupted`
- `failed`
- `finished`

## got/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a3f74910-a3c3-4d97-a56c-2ba6a48fe702",
  "message_type" : "got.scan",
  "group_id" : "a3f74910-a3c3-4d97-a56c-2ba6a48fe702",
  "created" : 1629896000451692548,
  "id" : "example.id.scan",
  "temporary" : false,
  "sensor" : "example.sensor.1",
  "alive" : true,
  "hosts" : [
    "example.host.to.scan.com"
  ],
  "plugins" : {
    "single_vts" : [
      {
        "oid" : "example.oid.1"
      }
    ]
  },
  "ports" : [
    "1337"
  ]
}
```
To get type information please consolidate [ scan model](../models/scan.go)


## result/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d9c531d7-ae49-4438-9c30-fdde713d0106",
  "message_type" : "result.scan",
  "group_id" : "d9c531d7-ae49-4438-9c30-fdde713d0106",
  "created" : 1629896000451722556,
  "result_type" : "LOG",
  "host_ip" : "192.168.1.1",
  "host_name" : "example.host.domain",
  "port" : "1337",
  "id" : "example.id.scan",
  "oid" : "example.oid.1",
  "value" : "This an example log message",
  "uri" : "uri.to.oid.description"
}
```
Valid `result_type` are:
- `UNKNOWN`
- `HOST_COUNT`
- `DEADHOST`
- `HOST_START`
- `HOST_END`
- `ERRMSG`
- `LOG`
- `HOST_DETAIL`
- `ALARM`


For more specific information please consolidate [result model](../models/result.go)
## failure.start/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "00f6dd8a-9c6f-4d65-ab53-f563af6914f0",
  "message_type" : "failure.start.scan",
  "group_id" : "00f6dd8a-9c6f-4d65-ab53-f563af6914f0",
  "created" : 1629896000451742896,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "30cd5c6c-2470-4151-81c8-4d598c73d503",
  "message_type" : "failure.stop.scan",
  "group_id" : "30cd5c6c-2470-4151-81c8-4d598c73d503",
  "created" : 1629896000451756745,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8cdecd5f-c017-47e8-9373-e327458f1608",
  "message_type" : "failure.create.scan",
  "group_id" : "8cdecd5f-c017-47e8-9373-e327458f1608",
  "created" : 1629896000451769987,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "448c4cea-48fa-4538-aae8-1d92b279305b",
  "message_type" : "failure.modify.scan",
  "group_id" : "448c4cea-48fa-4538-aae8-1d92b279305b",
  "created" : 1629896000451782349,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0e1128c6-1f4d-447a-827e-0a718404fc39",
  "message_type" : "failure.get.scan",
  "group_id" : "0e1128c6-1f4d-447a-827e-0a718404fc39",
  "created" : 1629896000451794551,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7517bd59-9eaf-4249-aead-93f2717f44e8",
  "message_type" : "failure.scan",
  "group_id" : "7517bd59-9eaf-4249-aead-93f2717f44e8",
  "created" : 1629896000451808686,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
# target

To get type information for e.g. `modify.target` or `got.target` please consolidate [ target model](../models/target.go)

As a rule of thumb: each type is as shown in the example.

## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "bde26616-4502-4dcc-86ae-2e3fab4f25cc",
  "message_type" : "create.target",
  "group_id" : "bde26616-4502-4dcc-86ae-2e3fab4f25cc",
  "created" : 1629896000451826473
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "658f3136-5238-4e1c-a434-3b4f5660238f",
  "message_type" : "get.target",
  "group_id" : "658f3136-5238-4e1c-a434-3b4f5660238f",
  "created" : 1629896000451845722,
  "id" : "example.id.target"
}
```
Responses:

- [got](#gottarget)
- [failure.get](#failuregettarget)
## modify/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "974e72b9-4783-4e8c-806b-32b37547e67f",
  "message_type" : "modify.target",
  "group_id" : "974e72b9-4783-4e8c-806b-32b37547e67f",
  "created" : 1629896000451864631,
  "id" : "example.id.target",
  "sensor" : "example.sensor.1",
  "alive" : true,
  "hosts" : [
    "example.host.to.scan.com"
  ],
  "plugins" : {
    "single_vts" : [
      {
        "oid" : "example.oid.1"
      }
    ]
  },
  "ports" : [
    "1337"
  ]
}
```
To get type information please consolidate [ target model](../models/target.go)


Responses:

- [modified](#modifiedtarget)
- [failure.modify](#failuremodifytarget)
## created/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1f479776-af51-4248-9d6c-6992df5f6ed7",
  "message_type" : "created.target",
  "group_id" : "1f479776-af51-4248-9d6c-6992df5f6ed7",
  "created" : 1629896000451903369,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "f35f60c4-51d3-4442-bfcd-c1f22adc2f75",
  "message_type" : "modified.target",
  "group_id" : "f35f60c4-51d3-4442-bfcd-c1f22adc2f75",
  "created" : 1629896000451920901,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1b86d15b-0049-4c7d-90d0-29a322794c1d",
  "message_type" : "got.target",
  "group_id" : "1b86d15b-0049-4c7d-90d0-29a322794c1d",
  "created" : 1629896000451934689,
  "id" : "example.id.target",
  "sensor" : "example.sensor.1",
  "alive" : true,
  "hosts" : [
    "example.host.to.scan.com"
  ],
  "plugins" : {
    "single_vts" : [
      {
        "oid" : "example.oid.1"
      }
    ]
  },
  "ports" : [
    "1337"
  ]
}
```
To get type information please consolidate [ target model](../models/target.go)


## failure.create/target

Topic: eulabeia/target/info
```
{
  "message_id" : "a824e4cb-3d0b-4a9e-bf76-e3c3db75e802",
  "message_type" : "failure.create.target",
  "group_id" : "a824e4cb-3d0b-4a9e-bf76-e3c3db75e802",
  "created" : 1629896000451957906,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "894240e2-6fd2-41b4-b0a3-6de9090b77be",
  "message_type" : "failure.modify.target",
  "group_id" : "894240e2-6fd2-41b4-b0a3-6de9090b77be",
  "created" : 1629896000451970858,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1575e4c3-fa83-4146-9307-491488c3cec9",
  "message_type" : "failure.get.target",
  "group_id" : "1575e4c3-fa83-4146-9307-491488c3cec9",
  "created" : 1629896000451983267,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "752c9df4-0e2f-4e70-b7f4-ff3a0c4f887d",
  "message_type" : "failure.target",
  "group_id" : "752c9df4-0e2f-4e70-b7f4-ff3a0c4f887d",
  "created" : 1629896000451995707,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
