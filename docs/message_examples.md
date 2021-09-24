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

Topic: scanner/scan/cmd/director
```
{
  "message_id" : "312e2177-2a15-4b4c-ba92-4d217ebb3144",
  "message_type" : "create.scan",
  "group_id" : "312e2177-2a15-4b4c-ba92-4d217ebb3144",
  "message_created" : 1632752557579025545
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: scanner/scan/cmd/director
```
{
  "message_id" : "393667bc-3bfb-4c1e-888e-c77631296431",
  "message_type" : "start.scan",
  "group_id" : "393667bc-3bfb-4c1e-888e-c77631296431",
  "message_created" : 1632752557579208756,
  "id" : "example.id.scan"
}
```
Responses:

- [status](#statusscan)
- [failure.start](#failurestartscan)
## stop/scan

Topic: scanner/scan/cmd/director
```
{
  "message_id" : "3b35ed9d-12fd-4bbe-be6e-b9ee77a4bc32",
  "message_type" : "stop.scan",
  "group_id" : "3b35ed9d-12fd-4bbe-be6e-b9ee77a4bc32",
  "message_created" : 1632752557579252445,
  "id" : "example.id.scan"
}
```
Responses:

- [stopped](#stoppedscan)
- [failure.stop](#failurestopscan)
## get/scan

Topic: scanner/scan/cmd/director
```
{
  "message_id" : "616d7b08-4028-4a24-acd2-4304d422bf5f",
  "message_type" : "get.scan",
  "group_id" : "616d7b08-4028-4a24-acd2-4304d422bf5f",
  "message_created" : 1632752557579291643,
  "id" : "example.id.scan"
}
```
Responses:

- [got](#gotscan)
- [failure.get](#failuregetscan)
## modify/scan

Topic: scanner/scan/cmd/director
```
{
  "message_id" : "61cde83b-c544-429c-9e1d-f3010f788b98",
  "message_type" : "modify.scan",
  "group_id" : "61cde83b-c544-429c-9e1d-f3010f788b98",
  "message_created" : 1632752557579344720,
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

Topic: scanner/scan/info
```
{
  "message_id" : "b83edc47-cb4d-4d56-84a3-c8175ab41fe8",
  "message_type" : "created.scan",
  "group_id" : "b83edc47-cb4d-4d56-84a3-c8175ab41fe8",
  "message_created" : 1632752557579404136,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: scanner/scan/info
```
{
  "message_id" : "fef10c1c-1e19-49de-bfa4-d1912fba2ca2",
  "message_type" : "modified.scan",
  "group_id" : "fef10c1c-1e19-49de-bfa4-d1912fba2ca2",
  "message_created" : 1632752557579434808,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: scanner/scan/info
```
{
  "message_id" : "9ded600d-b225-42d1-b3df-2bb246fb3164",
  "message_type" : "stopped.scan",
  "group_id" : "9ded600d-b225-42d1-b3df-2bb246fb3164",
  "message_created" : 1632752557579457401,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: scanner/scan/info
```
{
  "message_id" : "ac5d6860-9010-48b4-bd2c-8bccc0e2c286",
  "message_type" : "status.scan",
  "group_id" : "ac5d6860-9010-48b4-bd2c-8bccc0e2c286",
  "message_created" : 1632752557579478475,
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

Topic: scanner/scan/info
```
{
  "message_id" : "9623ec9b-ff31-499b-9981-03d5c86a8ed0",
  "message_type" : "got.scan",
  "group_id" : "9623ec9b-ff31-499b-9981-03d5c86a8ed0",
  "message_created" : 1632752557579535485,
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

Topic: scanner/scan/info
```
{
  "message_id" : "0ef38736-ecc9-4814-b436-9c0806fa28a9",
  "message_type" : "result.scan",
  "group_id" : "0ef38736-ecc9-4814-b436-9c0806fa28a9",
  "message_created" : 1632752557579589301,
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
- `HOST_START`
- `HOST_END`
- `ERRMSG`
- `LOG`
- `HOST_DETAIL`
- `ALARM`


For more specific information please consolidate [result model](../models/result.go)
## failure.start/scan

Topic: scanner/scan/info
```
{
  "message_id" : "db6607db-6d2e-4885-a10a-62c86e0f5e24",
  "message_type" : "failure.start.scan",
  "group_id" : "db6607db-6d2e-4885-a10a-62c86e0f5e24",
  "message_created" : 1632752557579624911,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: scanner/scan/info
```
{
  "message_id" : "65ecbcca-9bd0-4150-80ac-3ee19347ccd3",
  "message_type" : "failure.stop.scan",
  "group_id" : "65ecbcca-9bd0-4150-80ac-3ee19347ccd3",
  "message_created" : 1632752557579649788,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: scanner/scan/info
```
{
  "message_id" : "1249a1db-1eda-4818-bd9e-d77be9205f0b",
  "message_type" : "failure.create.scan",
  "group_id" : "1249a1db-1eda-4818-bd9e-d77be9205f0b",
  "message_created" : 1632752557579673270,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: scanner/scan/info
```
{
  "message_id" : "9e8f382d-b363-4cca-8a9e-1b07cea1dc57",
  "message_type" : "failure.modify.scan",
  "group_id" : "9e8f382d-b363-4cca-8a9e-1b07cea1dc57",
  "message_created" : 1632752557579695181,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: scanner/scan/info
```
{
  "message_id" : "7898db21-657f-4ad5-a728-ecd905a5bdf7",
  "message_type" : "failure.get.scan",
  "group_id" : "7898db21-657f-4ad5-a728-ecd905a5bdf7",
  "message_created" : 1632752557579720463,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: scanner/scan/info
```
{
  "message_id" : "4a6f931e-e201-4e50-a333-2cee63115676",
  "message_type" : "failure.scan",
  "group_id" : "4a6f931e-e201-4e50-a333-2cee63115676",
  "message_created" : 1632752557579742638,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
# target

To get type information for e.g. `modify.target` or `got.target` please consolidate [ target model](../models/target.go)

As a rule of thumb: each type is as shown in the example.

## create/target

Topic: scanner/target/cmd/director
```
{
  "message_id" : "1fb1a03c-c1ae-4d67-bb61-5f80eb3d211b",
  "message_type" : "create.target",
  "group_id" : "1fb1a03c-c1ae-4d67-bb61-5f80eb3d211b",
  "message_created" : 1632752557579773365
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: scanner/target/cmd/director
```
{
  "message_id" : "23c6bba7-bb92-48ad-aa8b-346196c9da6d",
  "message_type" : "get.target",
  "group_id" : "23c6bba7-bb92-48ad-aa8b-346196c9da6d",
  "message_created" : 1632752557579807912,
  "id" : "example.id.target"
}
```
Responses:

- [got](#gottarget)
- [failure.get](#failuregettarget)
## modify/target

Topic: scanner/target/cmd/director
```
{
  "message_id" : "604c6f86-e114-41f1-bb4d-85823c84b9b9",
  "message_type" : "modify.target",
  "group_id" : "604c6f86-e114-41f1-bb4d-85823c84b9b9",
  "message_created" : 1632752557579855794,
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

Topic: scanner/target/info
```
{
  "message_id" : "0b6111bf-f738-4da8-b16d-254ba6b5d932",
  "message_type" : "created.target",
  "group_id" : "0b6111bf-f738-4da8-b16d-254ba6b5d932",
  "message_created" : 1632752557579914340,
  "id" : "example.id.target"
}
```
## modified/target

Topic: scanner/target/info
```
{
  "message_id" : "f0fcd300-3e04-4d88-9437-9042e1634859",
  "message_type" : "modified.target",
  "group_id" : "f0fcd300-3e04-4d88-9437-9042e1634859",
  "message_created" : 1632752557579939237,
  "id" : "example.id.target"
}
```
## got/target

Topic: scanner/target/info
```
{
  "message_id" : "2bca2b44-cb31-470c-afa7-05cb571109be",
  "message_type" : "got.target",
  "group_id" : "2bca2b44-cb31-470c-afa7-05cb571109be",
  "message_created" : 1632752557579960951,
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

Topic: scanner/target/info
```
{
  "message_id" : "e590b272-d905-4082-bd68-f83afcec2737",
  "message_type" : "failure.create.target",
  "group_id" : "e590b272-d905-4082-bd68-f83afcec2737",
  "message_created" : 1632752557579997367,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: scanner/target/info
```
{
  "message_id" : "f3780712-1687-4929-b1e6-2dbcf0b1b16a",
  "message_type" : "failure.modify.target",
  "group_id" : "f3780712-1687-4929-b1e6-2dbcf0b1b16a",
  "message_created" : 1632752557580020924,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: scanner/target/info
```
{
  "message_id" : "b51abf24-508e-4e9c-b0cd-62957c1ba50c",
  "message_type" : "failure.get.target",
  "group_id" : "b51abf24-508e-4e9c-b0cd-62957c1ba50c",
  "message_created" : 1632752557580043434,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: scanner/target/info
```
{
  "message_id" : "4108e236-f21c-48a9-9c12-f728d0af140f",
  "message_type" : "failure.target",
  "group_id" : "4108e236-f21c-48a9-9c12-f728d0af140f",
  "message_created" : 1632752557580069703,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
