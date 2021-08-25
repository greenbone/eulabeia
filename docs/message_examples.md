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

## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "dd024f77-6793-40a5-94c6-1ee53ac197f4",
  "message_type" : "create.scan",
  "group_id" : "dd024f77-6793-40a5-94c6-1ee53ac197f4",
  "created" : 1629893733766194904
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "8af76e2d-0be1-4f33-b20e-f7c5417e5d27",
  "message_type" : "start.scan",
  "group_id" : "8af76e2d-0be1-4f33-b20e-f7c5417e5d27",
  "created" : 1629893733766307057,
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
  "message_id" : "c93546dc-5c07-4802-8b9f-2daa424baeaa",
  "message_type" : "stop.scan",
  "group_id" : "c93546dc-5c07-4802-8b9f-2daa424baeaa",
  "created" : 1629893733766338295,
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
  "message_id" : "6f448d60-347d-494c-bde4-2b5a5bb1b27c",
  "message_type" : "get.scan",
  "group_id" : "6f448d60-347d-494c-bde4-2b5a5bb1b27c",
  "created" : 1629893733766365698,
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
  "message_id" : "41ce8ae2-42a0-44e5-94e7-79c5c833dc08",
  "message_type" : "modify.scan",
  "group_id" : "41ce8ae2-42a0-44e5-94e7-79c5c833dc08",
  "created" : 1629893733766391419,
  "id" : "example.scan.id",
  "values" : {
    "temporary" : false,
    "target_id" : "example.target.id"
  }
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0ef9db84-104a-4323-8176-e41872ee3aee",
  "message_type" : "created.scan",
  "group_id" : "0ef9db84-104a-4323-8176-e41872ee3aee",
  "created" : 1629893733766422592,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0c02a87e-b942-4c5f-841e-68bb3b9541d7",
  "message_type" : "modified.scan",
  "group_id" : "0c02a87e-b942-4c5f-841e-68bb3b9541d7",
  "created" : 1629893733766440723,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "bab5f054-ae7d-48db-bd83-46305a401100",
  "message_type" : "stopped.scan",
  "group_id" : "bab5f054-ae7d-48db-bd83-46305a401100",
  "created" : 1629893733766457031,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5cbf1dd6-6e47-4a2c-8e85-0c7fc042ab0b",
  "message_type" : "status.scan",
  "group_id" : "5cbf1dd6-6e47-4a2c-8e85-0c7fc042ab0b",
  "created" : 1629893733766472761,
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
  "message_id" : "f5894600-eb16-4d65-be70-765fc02d8346",
  "message_type" : "got.scan",
  "group_id" : "f5894600-eb16-4d65-be70-765fc02d8346",
  "created" : 1629893733766495360,
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
## result/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3b5e13a4-9f2c-4f42-b013-316c0aed2f3e",
  "message_type" : "result.scan",
  "group_id" : "3b5e13a4-9f2c-4f42-b013-316c0aed2f3e",
  "created" : 1629893733766553023,
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

## failure.start/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b7b08cf5-6256-4034-9f34-8639d2b4af90",
  "message_type" : "failure.start.scan",
  "group_id" : "b7b08cf5-6256-4034-9f34-8639d2b4af90",
  "created" : 1629893733766583551,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9e31363d-feb1-4cbe-bc3e-d7dbdb984b3a",
  "message_type" : "failure.stop.scan",
  "group_id" : "9e31363d-feb1-4cbe-bc3e-d7dbdb984b3a",
  "created" : 1629893733766605156,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "70f6eef7-d56f-40be-a409-6ac5fc91d192",
  "message_type" : "failure.create.scan",
  "group_id" : "70f6eef7-d56f-40be-a409-6ac5fc91d192",
  "created" : 1629893733766623289,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "18643fa9-c8ba-4fc8-868a-c00a3986dfb9",
  "message_type" : "failure.modify.scan",
  "group_id" : "18643fa9-c8ba-4fc8-868a-c00a3986dfb9",
  "created" : 1629893733766639982,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "14cad595-72ef-41d0-b402-2fc1c298c69a",
  "message_type" : "failure.get.scan",
  "group_id" : "14cad595-72ef-41d0-b402-2fc1c298c69a",
  "created" : 1629893733766656698,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "de5d37eb-4083-4b75-b410-3f2ec541e4c0",
  "message_type" : "failure.scan",
  "group_id" : "de5d37eb-4083-4b75-b410-3f2ec541e4c0",
  "created" : 1629893733766672870,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
# target

## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "66bc5ff4-6abb-4aa0-b258-3c568d724824",
  "message_type" : "create.target",
  "group_id" : "66bc5ff4-6abb-4aa0-b258-3c568d724824",
  "created" : 1629893733766694811
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "68c9df24-b09a-4dc5-bf5f-dbaedc1e42c4",
  "message_type" : "get.target",
  "group_id" : "68c9df24-b09a-4dc5-bf5f-dbaedc1e42c4",
  "created" : 1629893733766721104,
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
  "message_id" : "6c36a4a1-424d-4393-ab2e-a9a388dfda7d",
  "message_type" : "modify.target",
  "group_id" : "6c36a4a1-424d-4393-ab2e-a9a388dfda7d",
  "created" : 1629893733766746902,
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
Responses:

- [modified](#modifiedtarget)
- [failure.modify](#failuremodifytarget)
## created/target

Topic: eulabeia/target/info
```
{
  "message_id" : "3d547238-636b-4e23-8608-92649d0f1827",
  "message_type" : "created.target",
  "group_id" : "3d547238-636b-4e23-8608-92649d0f1827",
  "created" : 1629893733766787731,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "097cca9f-03ad-48da-bb5c-4999483c654a",
  "message_type" : "modified.target",
  "group_id" : "097cca9f-03ad-48da-bb5c-4999483c654a",
  "created" : 1629893733766817002,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "875e067a-9f0c-48c5-a024-3372db4530a8",
  "message_type" : "got.target",
  "group_id" : "875e067a-9f0c-48c5-a024-3372db4530a8",
  "created" : 1629893733766835441,
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
## failure.create/target

Topic: eulabeia/target/info
```
{
  "message_id" : "d3e88893-7579-44fb-b12c-c8958c0ecf23",
  "message_type" : "failure.create.target",
  "group_id" : "d3e88893-7579-44fb-b12c-c8958c0ecf23",
  "created" : 1629893733766862467,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "ee9d2244-08c6-45c2-9031-9a0b6b5552ab",
  "message_type" : "failure.modify.target",
  "group_id" : "ee9d2244-08c6-45c2-9031-9a0b6b5552ab",
  "created" : 1629893733766880251,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "d8095970-f458-4bf3-ac6e-8418cc86073b",
  "message_type" : "failure.get.target",
  "group_id" : "d8095970-f458-4bf3-ac6e-8418cc86073b",
  "created" : 1629893733766897153,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "f77b3d23-dd05-47d0-8ffd-7d029d6424c0",
  "message_type" : "failure.target",
  "group_id" : "f77b3d23-dd05-47d0-8ffd-7d029d6424c0",
  "created" : 1629893733766913752,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
