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
  "message_id" : "0f0bdd45-b85e-4872-a38d-0e9b95255bdc",
  "message_type" : "create.scan",
  "group_id" : "0f0bdd45-b85e-4872-a38d-0e9b95255bdc",
  "created" : 1629808666889967031
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "d4471faa-2f89-4c9a-9885-502e648f0496",
  "message_type" : "start.scan",
  "group_id" : "d4471faa-2f89-4c9a-9885-502e648f0496",
  "created" : 1629808666890083862,
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
  "message_id" : "6bf12c18-9680-412e-9637-63360b4a62e1",
  "message_type" : "stop.scan",
  "group_id" : "6bf12c18-9680-412e-9637-63360b4a62e1",
  "created" : 1629808666890118656,
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
  "message_id" : "94567d16-835a-4c8b-a3fd-44e5f2bc4361",
  "message_type" : "get.scan",
  "group_id" : "94567d16-835a-4c8b-a3fd-44e5f2bc4361",
  "created" : 1629808666890148173,
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
  "message_id" : "77e3d83e-a41e-47ef-a069-21afa9e9be8a",
  "message_type" : "modify.scan",
  "group_id" : "77e3d83e-a41e-47ef-a069-21afa9e9be8a",
  "created" : 1629808666890175693,
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
  "message_id" : "77dc4ec0-2811-4c56-b850-e078e6fdda66",
  "message_type" : "created.scan",
  "group_id" : "77dc4ec0-2811-4c56-b850-e078e6fdda66",
  "created" : 1629808666890209660,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "044e6939-dd54-46b1-be4c-3e8a8d649e51",
  "message_type" : "modified.scan",
  "group_id" : "044e6939-dd54-46b1-be4c-3e8a8d649e51",
  "created" : 1629808666890230982,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c56355c5-6019-4c98-a1e6-5e0fe308a13d",
  "message_type" : "stopped.scan",
  "group_id" : "c56355c5-6019-4c98-a1e6-5e0fe308a13d",
  "created" : 1629808666890251033,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7b1d2db9-693d-4204-b367-830057e2a7ca",
  "message_type" : "status.scan",
  "group_id" : "7b1d2db9-693d-4204-b367-830057e2a7ca",
  "created" : 1629808666890269642,
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
  "message_id" : "747171e3-5164-41cc-9f76-4dc5d76f6386",
  "message_type" : "got.scan",
  "group_id" : "747171e3-5164-41cc-9f76-4dc5d76f6386",
  "created" : 1629808666890296121,
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
  "message_id" : "6437fe74-c820-4946-a6fa-b275089162b3",
  "message_type" : "result.scan",
  "group_id" : "6437fe74-c820-4946-a6fa-b275089162b3",
  "created" : 1629808666890360206,
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
  "message_id" : "f6212040-7a95-4c48-ae0f-80915bbb990a",
  "message_type" : "failure.start.scan",
  "group_id" : "f6212040-7a95-4c48-ae0f-80915bbb990a",
  "created" : 1629808666890395098,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "527f43bb-593f-4b20-8955-c45802d78188",
  "message_type" : "failure.stop.scan",
  "group_id" : "527f43bb-593f-4b20-8955-c45802d78188",
  "created" : 1629808666890420981,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6961a70c-b393-40ac-965f-e1bd44a3d24d",
  "message_type" : "failure.create.scan",
  "group_id" : "6961a70c-b393-40ac-965f-e1bd44a3d24d",
  "created" : 1629808666890442588,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "aba4dbb2-00b0-40bf-ad18-33fd2ba9e56b",
  "message_type" : "failure.modify.scan",
  "group_id" : "aba4dbb2-00b0-40bf-ad18-33fd2ba9e56b",
  "created" : 1629808666890463115,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "08618959-0437-4943-bc65-931ae35d4cb6",
  "message_type" : "failure.get.scan",
  "group_id" : "08618959-0437-4943-bc65-931ae35d4cb6",
  "created" : 1629808666890483477,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c992fe7e-350c-44ea-9ed8-549669614541",
  "message_type" : "failure.scan",
  "group_id" : "c992fe7e-350c-44ea-9ed8-549669614541",
  "created" : 1629808666890503994,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "a5abb43f-7b7c-49ab-93bc-eea3269fd2ea",
  "message_type" : "create.scan",
  "group_id" : "a5abb43f-7b7c-49ab-93bc-eea3269fd2ea",
  "created" : 1629808666890523931
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "a797f69c-edc2-4714-a72a-54a3a4b3edeb",
  "message_type" : "start.scan",
  "group_id" : "a797f69c-edc2-4714-a72a-54a3a4b3edeb",
  "created" : 1629808666890546125,
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
  "message_id" : "2c2b75f1-0e22-463b-9b9f-78210852adc2",
  "message_type" : "stop.scan",
  "group_id" : "2c2b75f1-0e22-463b-9b9f-78210852adc2",
  "created" : 1629808666890568392,
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
  "message_id" : "8622f4e9-2545-442d-82df-9ac2522262dd",
  "message_type" : "get.scan",
  "group_id" : "8622f4e9-2545-442d-82df-9ac2522262dd",
  "created" : 1629808666890600039,
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
  "message_id" : "d291e2c9-646f-480a-8140-5e2f4ca2872a",
  "message_type" : "modify.scan",
  "group_id" : "d291e2c9-646f-480a-8140-5e2f4ca2872a",
  "created" : 1629808666890642649,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "1b2de578-88ee-4839-89b1-f917cd9e68cb",
  "message_type" : "created.scan",
  "group_id" : "1b2de578-88ee-4839-89b1-f917cd9e68cb",
  "created" : 1629808666890676319,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "bdb8a1dd-3aa7-4bf2-9b44-258c240db12a",
  "message_type" : "modified.scan",
  "group_id" : "bdb8a1dd-3aa7-4bf2-9b44-258c240db12a",
  "created" : 1629808666890697216,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0260943e-99c9-449a-a5e1-c6ebdb245637",
  "message_type" : "stopped.scan",
  "group_id" : "0260943e-99c9-449a-a5e1-c6ebdb245637",
  "created" : 1629808666890716282,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fe136132-ab8b-4473-bf8b-79943a4e6d1c",
  "message_type" : "status.scan",
  "group_id" : "fe136132-ab8b-4473-bf8b-79943a4e6d1c",
  "created" : 1629808666890735697,
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
  "message_id" : "f88dbcda-3d1e-489e-8cbf-aee993c0be73",
  "message_type" : "got.scan",
  "group_id" : "f88dbcda-3d1e-489e-8cbf-aee993c0be73",
  "created" : 1629808666890758711,
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
  "message_id" : "e65a7701-ed2d-4969-952f-23e24159e4a7",
  "message_type" : "result.scan",
  "group_id" : "e65a7701-ed2d-4969-952f-23e24159e4a7",
  "created" : 1629808666890795262,
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
  "message_id" : "dbc6a274-5c0b-43f0-afc4-2ab6b6af9ae0",
  "message_type" : "failure.start.scan",
  "group_id" : "dbc6a274-5c0b-43f0-afc4-2ab6b6af9ae0",
  "created" : 1629808666890831613,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b0d3f795-42a7-4386-b0cd-76113d813111",
  "message_type" : "failure.stop.scan",
  "group_id" : "b0d3f795-42a7-4386-b0cd-76113d813111",
  "created" : 1629808666890854666,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f4c70abb-a437-440b-aa7d-8860794cd9b9",
  "message_type" : "failure.create.scan",
  "group_id" : "f4c70abb-a437-440b-aa7d-8860794cd9b9",
  "created" : 1629808666890875586,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f547d8a8-895a-48b8-b49c-d354bafe8206",
  "message_type" : "failure.modify.scan",
  "group_id" : "f547d8a8-895a-48b8-b49c-d354bafe8206",
  "created" : 1629808666890896634,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "00a0b186-e963-44b9-9553-3b5bebcca9f8",
  "message_type" : "failure.get.scan",
  "group_id" : "00a0b186-e963-44b9-9553-3b5bebcca9f8",
  "created" : 1629808666890917330,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b5ea5f96-d90a-48f8-b9ed-8c2c7b360cf4",
  "message_type" : "failure.scan",
  "group_id" : "b5ea5f96-d90a-48f8-b9ed-8c2c7b360cf4",
  "created" : 1629808666890947091,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "ef9bf125-ce0e-4e58-8507-4074f104e511",
  "message_type" : "create.scan",
  "group_id" : "ef9bf125-ce0e-4e58-8507-4074f104e511",
  "created" : 1629808666890969444
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "1d466354-bd59-47ca-911e-9643f1314829",
  "message_type" : "start.scan",
  "group_id" : "1d466354-bd59-47ca-911e-9643f1314829",
  "created" : 1629808666891000763,
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
  "message_id" : "e235ffd0-4b64-42ba-ba7d-bf066ce6b65e",
  "message_type" : "stop.scan",
  "group_id" : "e235ffd0-4b64-42ba-ba7d-bf066ce6b65e",
  "created" : 1629808666891025029,
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
  "message_id" : "2628bd68-c30d-4798-aa7a-3791f5bec01f",
  "message_type" : "get.scan",
  "group_id" : "2628bd68-c30d-4798-aa7a-3791f5bec01f",
  "created" : 1629808666891051121,
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
  "message_id" : "28e10989-9a89-4d4f-95a5-57343c725a23",
  "message_type" : "modify.scan",
  "group_id" : "28e10989-9a89-4d4f-95a5-57343c725a23",
  "created" : 1629808666891082145,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3ec1ef8d-2a8c-4cfc-b559-d865785f2123",
  "message_type" : "created.scan",
  "group_id" : "3ec1ef8d-2a8c-4cfc-b559-d865785f2123",
  "created" : 1629808666891116326,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2e664c86-4229-4973-b91c-41f34dd07c5b",
  "message_type" : "modified.scan",
  "group_id" : "2e664c86-4229-4973-b91c-41f34dd07c5b",
  "created" : 1629808666891136887,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "e8f94041-5402-4531-a6cb-20632553dacd",
  "message_type" : "stopped.scan",
  "group_id" : "e8f94041-5402-4531-a6cb-20632553dacd",
  "created" : 1629808666891156562,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "279f3874-224e-4556-a617-6ce5d8048e3d",
  "message_type" : "status.scan",
  "group_id" : "279f3874-224e-4556-a617-6ce5d8048e3d",
  "created" : 1629808666891176045,
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
  "message_id" : "542913ea-1b3d-4397-b1bf-ce692ee098db",
  "message_type" : "got.scan",
  "group_id" : "542913ea-1b3d-4397-b1bf-ce692ee098db",
  "created" : 1629808666891201257,
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
  "message_id" : "e3ab36f6-c7ef-43ea-812f-6324793ec999",
  "message_type" : "result.scan",
  "group_id" : "e3ab36f6-c7ef-43ea-812f-6324793ec999",
  "created" : 1629808666891238235,
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
  "message_id" : "c5ba54a9-5fb1-4dbb-913e-23e614b6fccf",
  "message_type" : "failure.start.scan",
  "group_id" : "c5ba54a9-5fb1-4dbb-913e-23e614b6fccf",
  "created" : 1629808666891277633,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "eef93ec3-b5d7-47a4-b040-e741786c0db4",
  "message_type" : "failure.stop.scan",
  "group_id" : "eef93ec3-b5d7-47a4-b040-e741786c0db4",
  "created" : 1629808666891303210,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d5064304-852e-471f-aaf8-2d9822724292",
  "message_type" : "failure.create.scan",
  "group_id" : "d5064304-852e-471f-aaf8-2d9822724292",
  "created" : 1629808666891324185,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "652973f4-506f-4d1b-a005-8dbdfbcbcf72",
  "message_type" : "failure.modify.scan",
  "group_id" : "652973f4-506f-4d1b-a005-8dbdfbcbcf72",
  "created" : 1629808666891344839,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fdf6aa51-df09-46b3-bb20-1cf678ff5641",
  "message_type" : "failure.get.scan",
  "group_id" : "fdf6aa51-df09-46b3-bb20-1cf678ff5641",
  "created" : 1629808666891379154,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a25972e9-fe60-4602-8f0b-b01008018b83",
  "message_type" : "failure.scan",
  "group_id" : "a25972e9-fe60-4602-8f0b-b01008018b83",
  "created" : 1629808666891400047,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "5d597ea2-a887-4ec4-a37f-dd5a1b9cd28b",
  "message_type" : "create.scan",
  "group_id" : "5d597ea2-a887-4ec4-a37f-dd5a1b9cd28b",
  "created" : 1629808666891420313
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "fd696a77-d58b-4c11-bfec-be762eebdc25",
  "message_type" : "start.scan",
  "group_id" : "fd696a77-d58b-4c11-bfec-be762eebdc25",
  "created" : 1629808666891448555,
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
  "message_id" : "ebd02a62-f035-4f21-98c8-b5def1b40c21",
  "message_type" : "stop.scan",
  "group_id" : "ebd02a62-f035-4f21-98c8-b5def1b40c21",
  "created" : 1629808666891475075,
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
  "message_id" : "5d48537d-9eb8-46d2-93cd-23042e6098f9",
  "message_type" : "get.scan",
  "group_id" : "5d48537d-9eb8-46d2-93cd-23042e6098f9",
  "created" : 1629808666891503049,
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
  "message_id" : "8dc6730b-5b86-4abe-8c1e-29949bc5c4f2",
  "message_type" : "modify.scan",
  "group_id" : "8dc6730b-5b86-4abe-8c1e-29949bc5c4f2",
  "created" : 1629808666891530616,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4179b59e-4ad0-4e40-b806-73eba4251911",
  "message_type" : "created.scan",
  "group_id" : "4179b59e-4ad0-4e40-b806-73eba4251911",
  "created" : 1629808666891571016,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "553b5fa8-9ee9-49b6-afa1-b6886dab5226",
  "message_type" : "modified.scan",
  "group_id" : "553b5fa8-9ee9-49b6-afa1-b6886dab5226",
  "created" : 1629808666891592637,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d68c1cfd-fa18-461a-87c2-6b6b162ea8e0",
  "message_type" : "stopped.scan",
  "group_id" : "d68c1cfd-fa18-461a-87c2-6b6b162ea8e0",
  "created" : 1629808666891611717,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "741548d7-5b22-471e-aa54-c8c075ebcac2",
  "message_type" : "status.scan",
  "group_id" : "741548d7-5b22-471e-aa54-c8c075ebcac2",
  "created" : 1629808666891633595,
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
  "message_id" : "5e3727aa-714d-499f-a591-5e7e10308635",
  "message_type" : "got.scan",
  "group_id" : "5e3727aa-714d-499f-a591-5e7e10308635",
  "created" : 1629808666891655853,
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
  "message_id" : "6fc1e2a2-199d-4d00-a670-b3702f6cc400",
  "message_type" : "result.scan",
  "group_id" : "6fc1e2a2-199d-4d00-a670-b3702f6cc400",
  "created" : 1629808666891691493,
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
  "message_id" : "afd3001d-873a-4302-9d9d-9863abb48e4b",
  "message_type" : "failure.start.scan",
  "group_id" : "afd3001d-873a-4302-9d9d-9863abb48e4b",
  "created" : 1629808666891729595,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9264ed94-78a1-4ffb-afed-f5a1a029c660",
  "message_type" : "failure.stop.scan",
  "group_id" : "9264ed94-78a1-4ffb-afed-f5a1a029c660",
  "created" : 1629808666891752760,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "50674d8c-1a98-43d0-9139-b5a2e3203f09",
  "message_type" : "failure.create.scan",
  "group_id" : "50674d8c-1a98-43d0-9139-b5a2e3203f09",
  "created" : 1629808666891773236,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5dcdf1d6-d9f4-4932-adc5-e0a4f10fa1ec",
  "message_type" : "failure.modify.scan",
  "group_id" : "5dcdf1d6-d9f4-4932-adc5-e0a4f10fa1ec",
  "created" : 1629808666891793889,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "891163b2-8c1b-46d3-a409-3f5a37de9eb9",
  "message_type" : "failure.get.scan",
  "group_id" : "891163b2-8c1b-46d3-a409-3f5a37de9eb9",
  "created" : 1629808666891814627,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c242b68b-199c-4ff3-9300-1ecdc6aa2473",
  "message_type" : "failure.scan",
  "group_id" : "c242b68b-199c-4ff3-9300-1ecdc6aa2473",
  "created" : 1629808666891834681,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "87f119f7-628a-455a-90bb-6d51015aaa9e",
  "message_type" : "create.scan",
  "group_id" : "87f119f7-628a-455a-90bb-6d51015aaa9e",
  "created" : 1629808666891863733
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "b6343516-24a9-4ee4-ab48-93a67cc686d9",
  "message_type" : "start.scan",
  "group_id" : "b6343516-24a9-4ee4-ab48-93a67cc686d9",
  "created" : 1629808666891893079,
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
  "message_id" : "e40c4234-b2a5-43f9-98ba-e7663ac27315",
  "message_type" : "stop.scan",
  "group_id" : "e40c4234-b2a5-43f9-98ba-e7663ac27315",
  "created" : 1629808666891924749,
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
  "message_id" : "e2a29256-3479-42d9-b65d-b46acb345368",
  "message_type" : "get.scan",
  "group_id" : "e2a29256-3479-42d9-b65d-b46acb345368",
  "created" : 1629808666891954087,
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
  "message_id" : "4287ec75-6cfb-4453-9b6f-01f1a46c2de6",
  "message_type" : "modify.scan",
  "group_id" : "4287ec75-6cfb-4453-9b6f-01f1a46c2de6",
  "created" : 1629808666891982770,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "974d399c-839a-4636-9e30-94a7121aa26e",
  "message_type" : "created.scan",
  "group_id" : "974d399c-839a-4636-9e30-94a7121aa26e",
  "created" : 1629808666892013749,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3d738ed6-3122-4505-939c-a3c8bb098c14",
  "message_type" : "modified.scan",
  "group_id" : "3d738ed6-3122-4505-939c-a3c8bb098c14",
  "created" : 1629808666892033934,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a151ff3e-cbb4-446b-ad81-069563b51311",
  "message_type" : "stopped.scan",
  "group_id" : "a151ff3e-cbb4-446b-ad81-069563b51311",
  "created" : 1629808666892056317,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "268b862c-5d46-4b75-9eb6-79087405a582",
  "message_type" : "status.scan",
  "group_id" : "268b862c-5d46-4b75-9eb6-79087405a582",
  "created" : 1629808666892075339,
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
  "message_id" : "5c3222d4-ac99-4ac0-ba61-cfd608af6b62",
  "message_type" : "got.scan",
  "group_id" : "5c3222d4-ac99-4ac0-ba61-cfd608af6b62",
  "created" : 1629808666892099879,
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
  "message_id" : "18e8a0b8-d9b1-4adf-a8ba-f87f23a86904",
  "message_type" : "result.scan",
  "group_id" : "18e8a0b8-d9b1-4adf-a8ba-f87f23a86904",
  "created" : 1629808666892136509,
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
  "message_id" : "7355c911-da4f-4d8e-b4ed-e2f7567f41a7",
  "message_type" : "failure.start.scan",
  "group_id" : "7355c911-da4f-4d8e-b4ed-e2f7567f41a7",
  "created" : 1629808666892175490,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6a7338c8-c00e-4d58-89cf-2d496de9c6c1",
  "message_type" : "failure.stop.scan",
  "group_id" : "6a7338c8-c00e-4d58-89cf-2d496de9c6c1",
  "created" : 1629808666892198454,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c4b1a383-5275-448a-af41-58064855103e",
  "message_type" : "failure.create.scan",
  "group_id" : "c4b1a383-5275-448a-af41-58064855103e",
  "created" : 1629808666892218921,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a6b39033-562e-487f-b8c8-17d52c9a7aa3",
  "message_type" : "failure.modify.scan",
  "group_id" : "a6b39033-562e-487f-b8c8-17d52c9a7aa3",
  "created" : 1629808666892239873,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4a6defd9-0c6d-4834-b508-3332ac4c8c01",
  "message_type" : "failure.get.scan",
  "group_id" : "4a6defd9-0c6d-4834-b508-3332ac4c8c01",
  "created" : 1629808666892263521,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5bcd0fab-e181-48ef-9fd7-6b59be41a101",
  "message_type" : "failure.scan",
  "group_id" : "5bcd0fab-e181-48ef-9fd7-6b59be41a101",
  "created" : 1629808666892283508,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "9aa6f6ff-bd03-4169-9ba2-f9e855dcfc2e",
  "message_type" : "create.scan",
  "group_id" : "9aa6f6ff-bd03-4169-9ba2-f9e855dcfc2e",
  "created" : 1629808666892300091
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "6153ee32-60d6-4611-97a0-57663b923655",
  "message_type" : "start.scan",
  "group_id" : "6153ee32-60d6-4611-97a0-57663b923655",
  "created" : 1629808666892322555,
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
  "message_id" : "ea2905e9-27bd-43d8-872d-6aea52ce26ad",
  "message_type" : "stop.scan",
  "group_id" : "ea2905e9-27bd-43d8-872d-6aea52ce26ad",
  "created" : 1629808666892353156,
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
  "message_id" : "96725b5e-b9b9-4ec8-a34f-04bcb90a64bc",
  "message_type" : "get.scan",
  "group_id" : "96725b5e-b9b9-4ec8-a34f-04bcb90a64bc",
  "created" : 1629808666892382692,
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
  "message_id" : "e7ce2f62-c5da-4785-a463-a4880e5f66b1",
  "message_type" : "modify.scan",
  "group_id" : "e7ce2f62-c5da-4785-a463-a4880e5f66b1",
  "created" : 1629808666892410961,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f3bdbdfd-3083-4e7f-98ba-24cb83e6d57c",
  "message_type" : "created.scan",
  "group_id" : "f3bdbdfd-3083-4e7f-98ba-24cb83e6d57c",
  "created" : 1629808666892442102,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a1acc97e-bc24-41e4-80b0-c3dbef32327b",
  "message_type" : "modified.scan",
  "group_id" : "a1acc97e-bc24-41e4-80b0-c3dbef32327b",
  "created" : 1629808666892462318,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "99407098-92e0-4cc2-97d6-63e4ac772536",
  "message_type" : "stopped.scan",
  "group_id" : "99407098-92e0-4cc2-97d6-63e4ac772536",
  "created" : 1629808666892496097,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0d707a3c-f0b3-44ee-9814-b697f8c1eeab",
  "message_type" : "status.scan",
  "group_id" : "0d707a3c-f0b3-44ee-9814-b697f8c1eeab",
  "created" : 1629808666892517254,
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
  "message_id" : "fb56ac11-a990-4844-8aa3-1510c14f17fa",
  "message_type" : "got.scan",
  "group_id" : "fb56ac11-a990-4844-8aa3-1510c14f17fa",
  "created" : 1629808666892544287,
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
  "message_id" : "8532078f-3f22-42da-8ba2-bb1fbc94fa9b",
  "message_type" : "result.scan",
  "group_id" : "8532078f-3f22-42da-8ba2-bb1fbc94fa9b",
  "created" : 1629808666892582138,
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
  "message_id" : "ca0cdfdf-bb85-41df-866d-87ea1a7a21c5",
  "message_type" : "failure.start.scan",
  "group_id" : "ca0cdfdf-bb85-41df-866d-87ea1a7a21c5",
  "created" : 1629808666892612927,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "151b107c-2839-4c8d-a3c5-5a5d9b88c7cd",
  "message_type" : "failure.stop.scan",
  "group_id" : "151b107c-2839-4c8d-a3c5-5a5d9b88c7cd",
  "created" : 1629808666892634434,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6fc33f24-4dc4-4f38-9098-95c168a8773a",
  "message_type" : "failure.create.scan",
  "group_id" : "6fc33f24-4dc4-4f38-9098-95c168a8773a",
  "created" : 1629808666892655273,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "66f6fc0f-9163-4ec9-9c72-c594d61baebd",
  "message_type" : "failure.modify.scan",
  "group_id" : "66f6fc0f-9163-4ec9-9c72-c594d61baebd",
  "created" : 1629808666892675981,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "af678104-b2b4-4b3e-8b49-82ebf8eab2aa",
  "message_type" : "failure.get.scan",
  "group_id" : "af678104-b2b4-4b3e-8b49-82ebf8eab2aa",
  "created" : 1629808666892696505,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "1718c9f9-d3eb-43e5-b9ec-d4e7bf934f84",
  "message_type" : "failure.scan",
  "group_id" : "1718c9f9-d3eb-43e5-b9ec-d4e7bf934f84",
  "created" : 1629808666892716872,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "7d744b7c-a9dd-44fb-b14b-063bf8bb6af9",
  "message_type" : "create.scan",
  "group_id" : "7d744b7c-a9dd-44fb-b14b-063bf8bb6af9",
  "created" : 1629808666892739188
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "e3425eb1-24ca-4fc8-a2b1-38c543608f49",
  "message_type" : "start.scan",
  "group_id" : "e3425eb1-24ca-4fc8-a2b1-38c543608f49",
  "created" : 1629808666892768954,
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
  "message_id" : "932f0504-054a-4615-819a-98aeb2f6758a",
  "message_type" : "stop.scan",
  "group_id" : "932f0504-054a-4615-819a-98aeb2f6758a",
  "created" : 1629808666892800784,
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
  "message_id" : "5efded4e-a722-4e27-ad6f-27cddff27dc9",
  "message_type" : "get.scan",
  "group_id" : "5efded4e-a722-4e27-ad6f-27cddff27dc9",
  "created" : 1629808666892829982,
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
  "message_id" : "72828134-ff72-47d2-aa63-5fcba9557c3b",
  "message_type" : "modify.scan",
  "group_id" : "72828134-ff72-47d2-aa63-5fcba9557c3b",
  "created" : 1629808666892858132,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7f471b93-fbd6-41ac-923c-75426b4cab41",
  "message_type" : "created.scan",
  "group_id" : "7f471b93-fbd6-41ac-923c-75426b4cab41",
  "created" : 1629808666892889520,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6bcc9a68-888b-4121-8d0b-e598fbd047c3",
  "message_type" : "modified.scan",
  "group_id" : "6bcc9a68-888b-4121-8d0b-e598fbd047c3",
  "created" : 1629808666892909765,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "025b2cbe-29ae-4f0d-b4f5-a3e09158238d",
  "message_type" : "stopped.scan",
  "group_id" : "025b2cbe-29ae-4f0d-b4f5-a3e09158238d",
  "created" : 1629808666892928528,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2dc33c7b-b33c-41ac-9bd2-1233f3571032",
  "message_type" : "status.scan",
  "group_id" : "2dc33c7b-b33c-41ac-9bd2-1233f3571032",
  "created" : 1629808666892947480,
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
  "message_id" : "7809dd13-847b-4ed4-ac87-59c8b8073703",
  "message_type" : "got.scan",
  "group_id" : "7809dd13-847b-4ed4-ac87-59c8b8073703",
  "created" : 1629808666892972473,
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
  "message_id" : "a28f5974-5743-49c7-b99c-8392a9f8abee",
  "message_type" : "result.scan",
  "group_id" : "a28f5974-5743-49c7-b99c-8392a9f8abee",
  "created" : 1629808666893009747,
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
  "message_id" : "7fad16c5-e4a6-4a2e-ad45-2867a2b322f8",
  "message_type" : "failure.start.scan",
  "group_id" : "7fad16c5-e4a6-4a2e-ad45-2867a2b322f8",
  "created" : 1629808666893040578,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c537828b-036d-48ee-aea9-fd40171866eb",
  "message_type" : "failure.stop.scan",
  "group_id" : "c537828b-036d-48ee-aea9-fd40171866eb",
  "created" : 1629808666893070639,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "021c79a7-cab3-46ea-a055-4681bef585d3",
  "message_type" : "failure.create.scan",
  "group_id" : "021c79a7-cab3-46ea-a055-4681bef585d3",
  "created" : 1629808666893092445,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f1e7c330-db7b-4bb1-aca5-a34ea1c29d6a",
  "message_type" : "failure.modify.scan",
  "group_id" : "f1e7c330-db7b-4bb1-aca5-a34ea1c29d6a",
  "created" : 1629808666893112779,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "dec26823-cff4-4be7-960a-75035f14909a",
  "message_type" : "failure.get.scan",
  "group_id" : "dec26823-cff4-4be7-960a-75035f14909a",
  "created" : 1629808666893133046,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "49be9613-ba52-44a8-8203-725893820f1c",
  "message_type" : "failure.scan",
  "group_id" : "49be9613-ba52-44a8-8203-725893820f1c",
  "created" : 1629808666893152733,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "f160cb95-3240-4115-baf7-65ff8835d84b",
  "message_type" : "create.scan",
  "group_id" : "f160cb95-3240-4115-baf7-65ff8835d84b",
  "created" : 1629808666893172647
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "4ac217e8-17d6-4bd4-b82a-415faee0c0ba",
  "message_type" : "start.scan",
  "group_id" : "4ac217e8-17d6-4bd4-b82a-415faee0c0ba",
  "created" : 1629808666893197288,
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
  "message_id" : "ad0906c0-9d51-4a50-8031-eeb3b8f33c59",
  "message_type" : "stop.scan",
  "group_id" : "ad0906c0-9d51-4a50-8031-eeb3b8f33c59",
  "created" : 1629808666893231895,
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
  "message_id" : "48345936-b516-4165-989d-f84bc6ff1acc",
  "message_type" : "get.scan",
  "group_id" : "48345936-b516-4165-989d-f84bc6ff1acc",
  "created" : 1629808666893260768,
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
  "message_id" : "d65ca44c-46c0-4659-9feb-2f9802b482a0",
  "message_type" : "modify.scan",
  "group_id" : "d65ca44c-46c0-4659-9feb-2f9802b482a0",
  "created" : 1629808666893289290,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "090f9272-500b-4cef-9fd8-9b6f261751a0",
  "message_type" : "created.scan",
  "group_id" : "090f9272-500b-4cef-9fd8-9b6f261751a0",
  "created" : 1629808666893320554,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "557f6d17-5289-429f-b2b5-66769eaef209",
  "message_type" : "modified.scan",
  "group_id" : "557f6d17-5289-429f-b2b5-66769eaef209",
  "created" : 1629808666893340372,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d9b2e517-de3b-4826-bd0e-d5a065554911",
  "message_type" : "stopped.scan",
  "group_id" : "d9b2e517-de3b-4826-bd0e-d5a065554911",
  "created" : 1629808666893359786,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fc804287-73ec-4bed-867f-69612068d9ce",
  "message_type" : "status.scan",
  "group_id" : "fc804287-73ec-4bed-867f-69612068d9ce",
  "created" : 1629808666893387582,
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
  "message_id" : "c1d36071-9a35-4f3b-aaf3-bfc00c6dc088",
  "message_type" : "got.scan",
  "group_id" : "c1d36071-9a35-4f3b-aaf3-bfc00c6dc088",
  "created" : 1629808666893413149,
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
  "message_id" : "12529733-8954-4eca-9ea5-1f52cb7cc2b0",
  "message_type" : "result.scan",
  "group_id" : "12529733-8954-4eca-9ea5-1f52cb7cc2b0",
  "created" : 1629808666893450066,
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
  "message_id" : "38483b87-2997-4cde-925a-8fb9f0c9b048",
  "message_type" : "failure.start.scan",
  "group_id" : "38483b87-2997-4cde-925a-8fb9f0c9b048",
  "created" : 1629808666893483351,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "006a4fb9-3ce4-4e44-9c33-9e1613a8c569",
  "message_type" : "failure.stop.scan",
  "group_id" : "006a4fb9-3ce4-4e44-9c33-9e1613a8c569",
  "created" : 1629808666893504830,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f2f278f6-cf40-4d04-bf49-190320c2c5af",
  "message_type" : "failure.create.scan",
  "group_id" : "f2f278f6-cf40-4d04-bf49-190320c2c5af",
  "created" : 1629808666893525103,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5fd93460-2d4a-423e-939a-bcb55a0777eb",
  "message_type" : "failure.modify.scan",
  "group_id" : "5fd93460-2d4a-423e-939a-bcb55a0777eb",
  "created" : 1629808666893545530,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9ca4f63f-377a-4117-90dc-abeb5852da10",
  "message_type" : "failure.get.scan",
  "group_id" : "9ca4f63f-377a-4117-90dc-abeb5852da10",
  "created" : 1629808666893565834,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "878da502-7e9e-44da-8d73-a800009e327f",
  "message_type" : "failure.scan",
  "group_id" : "878da502-7e9e-44da-8d73-a800009e327f",
  "created" : 1629808666893586131,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "80520ddf-4694-4048-a1a1-ed9dfba42008",
  "message_type" : "create.scan",
  "group_id" : "80520ddf-4694-4048-a1a1-ed9dfba42008",
  "created" : 1629808666893606113
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "022ed544-9688-422c-99a1-d379a8e3ec9f",
  "message_type" : "start.scan",
  "group_id" : "022ed544-9688-422c-99a1-d379a8e3ec9f",
  "created" : 1629808666893628915,
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
  "message_id" : "7ca9f30e-3f76-44ba-abc4-4b09ee2001c7",
  "message_type" : "stop.scan",
  "group_id" : "7ca9f30e-3f76-44ba-abc4-4b09ee2001c7",
  "created" : 1629808666893673651,
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
  "message_id" : "048531ec-779e-4c24-a36e-a52280d7efb1",
  "message_type" : "get.scan",
  "group_id" : "048531ec-779e-4c24-a36e-a52280d7efb1",
  "created" : 1629808666893704561,
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
  "message_id" : "946c30e3-bb43-4e8c-90f3-a7139c94faf6",
  "message_type" : "modify.scan",
  "group_id" : "946c30e3-bb43-4e8c-90f3-a7139c94faf6",
  "created" : 1629808666893738637,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b30fa9d8-a39e-4561-b579-ed01e7eaa2b2",
  "message_type" : "created.scan",
  "group_id" : "b30fa9d8-a39e-4561-b579-ed01e7eaa2b2",
  "created" : 1629808666893769990,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fee2ceea-6cfc-4db6-9281-b948d51ebfbe",
  "message_type" : "modified.scan",
  "group_id" : "fee2ceea-6cfc-4db6-9281-b948d51ebfbe",
  "created" : 1629808666893789931,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "46263c0a-c923-4925-bd0f-77e5cf08431b",
  "message_type" : "stopped.scan",
  "group_id" : "46263c0a-c923-4925-bd0f-77e5cf08431b",
  "created" : 1629808666893808740,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7682c8e1-8929-4008-ab85-d9cb3e07ffc0",
  "message_type" : "status.scan",
  "group_id" : "7682c8e1-8929-4008-ab85-d9cb3e07ffc0",
  "created" : 1629808666893828950,
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
  "message_id" : "59f6392d-96f9-4c8e-bf06-4e4688372fdb",
  "message_type" : "got.scan",
  "group_id" : "59f6392d-96f9-4c8e-bf06-4e4688372fdb",
  "created" : 1629808666893851520,
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
  "message_id" : "f7c321c5-563e-4144-95e2-d1698ca5343e",
  "message_type" : "result.scan",
  "group_id" : "f7c321c5-563e-4144-95e2-d1698ca5343e",
  "created" : 1629808666893887963,
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
  "message_id" : "5dadba50-280e-4b6f-9125-e4e3f83d0493",
  "message_type" : "failure.start.scan",
  "group_id" : "5dadba50-280e-4b6f-9125-e4e3f83d0493",
  "created" : 1629808666893920374,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "faea0fe6-20cc-4a4e-a7f3-c6bf60229e2e",
  "message_type" : "failure.stop.scan",
  "group_id" : "faea0fe6-20cc-4a4e-a7f3-c6bf60229e2e",
  "created" : 1629808666893942219,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2faf3e45-b5df-440b-86ae-bbc2e63bbebe",
  "message_type" : "failure.create.scan",
  "group_id" : "2faf3e45-b5df-440b-86ae-bbc2e63bbebe",
  "created" : 1629808666893962941,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d6d9c008-e1bf-4aab-8d8f-9acfca54b644",
  "message_type" : "failure.modify.scan",
  "group_id" : "d6d9c008-e1bf-4aab-8d8f-9acfca54b644",
  "created" : 1629808666893991701,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2070a694-2f4c-4c01-9145-e7f4b77015ce",
  "message_type" : "failure.get.scan",
  "group_id" : "2070a694-2f4c-4c01-9145-e7f4b77015ce",
  "created" : 1629808666894013412,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "213c3f19-cd3f-4461-85a0-853e418acbee",
  "message_type" : "failure.scan",
  "group_id" : "213c3f19-cd3f-4461-85a0-853e418acbee",
  "created" : 1629808666894033214,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "456c4c51-5a85-43a8-8e9f-a36217951378",
  "message_type" : "create.scan",
  "group_id" : "456c4c51-5a85-43a8-8e9f-a36217951378",
  "created" : 1629808666894049231
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "4d68a485-7a9a-4f12-bae7-635d2139e19e",
  "message_type" : "start.scan",
  "group_id" : "4d68a485-7a9a-4f12-bae7-635d2139e19e",
  "created" : 1629808666894072099,
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
  "message_id" : "be13a64e-4a34-4465-a32a-d4122ece417c",
  "message_type" : "stop.scan",
  "group_id" : "be13a64e-4a34-4465-a32a-d4122ece417c",
  "created" : 1629808666894103209,
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
  "message_id" : "42c60e96-8557-448e-ad32-2dc3674983cc",
  "message_type" : "get.scan",
  "group_id" : "42c60e96-8557-448e-ad32-2dc3674983cc",
  "created" : 1629808666894131986,
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
  "message_id" : "22f71986-b3b3-40ce-875f-7169c62710b4",
  "message_type" : "modify.scan",
  "group_id" : "22f71986-b3b3-40ce-875f-7169c62710b4",
  "created" : 1629808666894160116,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c63fb884-2e83-4640-a16c-d886081af11a",
  "message_type" : "created.scan",
  "group_id" : "c63fb884-2e83-4640-a16c-d886081af11a",
  "created" : 1629808666894194112,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d9283d4c-1eb7-4680-9be2-26399d7708f1",
  "message_type" : "modified.scan",
  "group_id" : "d9283d4c-1eb7-4680-9be2-26399d7708f1",
  "created" : 1629808666894214705,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "ba72e4fd-6251-45e4-840f-de9139b23078",
  "message_type" : "stopped.scan",
  "group_id" : "ba72e4fd-6251-45e4-840f-de9139b23078",
  "created" : 1629808666894233646,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "36efcef7-40df-4c93-b23b-3e49c835959a",
  "message_type" : "status.scan",
  "group_id" : "36efcef7-40df-4c93-b23b-3e49c835959a",
  "created" : 1629808666894252480,
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
  "message_id" : "9819772f-39cf-43be-a3ce-16250d604078",
  "message_type" : "got.scan",
  "group_id" : "9819772f-39cf-43be-a3ce-16250d604078",
  "created" : 1629808666894283801,
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
  "message_id" : "d2d7869f-43fc-4b40-bbc2-0afb757d4e18",
  "message_type" : "result.scan",
  "group_id" : "d2d7869f-43fc-4b40-bbc2-0afb757d4e18",
  "created" : 1629808666894321775,
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
  "message_id" : "a4ade97c-7fc2-4e29-a63d-ea9f5ba91de8",
  "message_type" : "failure.start.scan",
  "group_id" : "a4ade97c-7fc2-4e29-a63d-ea9f5ba91de8",
  "created" : 1629808666894354203,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "bb539f70-59ae-41d3-9f3d-66dc785a1848",
  "message_type" : "failure.stop.scan",
  "group_id" : "bb539f70-59ae-41d3-9f3d-66dc785a1848",
  "created" : 1629808666894376375,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d34f3752-e602-4ab6-a181-13f1cc0416a0",
  "message_type" : "failure.create.scan",
  "group_id" : "d34f3752-e602-4ab6-a181-13f1cc0416a0",
  "created" : 1629808666894396558,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a26c1bbe-77dc-4a78-b8a9-7814ac9f2588",
  "message_type" : "failure.modify.scan",
  "group_id" : "a26c1bbe-77dc-4a78-b8a9-7814ac9f2588",
  "created" : 1629808666894417141,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6639eb7b-fdd4-4b87-a6ff-aab93858d0d0",
  "message_type" : "failure.get.scan",
  "group_id" : "6639eb7b-fdd4-4b87-a6ff-aab93858d0d0",
  "created" : 1629808666894438003,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "aa5ee6c7-f556-41d5-a7a0-64ace742ba44",
  "message_type" : "failure.scan",
  "group_id" : "aa5ee6c7-f556-41d5-a7a0-64ace742ba44",
  "created" : 1629808666894457833,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "e4f82eed-572c-4941-9c96-08393ac82f87",
  "message_type" : "create.scan",
  "group_id" : "e4f82eed-572c-4941-9c96-08393ac82f87",
  "created" : 1629808666894478258
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "ca320057-9b22-4b58-9f04-e63dae115f7e",
  "message_type" : "start.scan",
  "group_id" : "ca320057-9b22-4b58-9f04-e63dae115f7e",
  "created" : 1629808666894501576,
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
  "message_id" : "fd3cda95-7c7a-44a1-99e1-bce6c5879533",
  "message_type" : "stop.scan",
  "group_id" : "fd3cda95-7c7a-44a1-99e1-bce6c5879533",
  "created" : 1629808666894532883,
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
  "message_id" : "473f7038-80b3-4b35-99ec-5043169e14e2",
  "message_type" : "get.scan",
  "group_id" : "473f7038-80b3-4b35-99ec-5043169e14e2",
  "created" : 1629808666894562168,
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
  "message_id" : "67aee1c6-c68d-4f7f-9213-73a301be31e5",
  "message_type" : "modify.scan",
  "group_id" : "67aee1c6-c68d-4f7f-9213-73a301be31e5",
  "created" : 1629808666894599835,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8fb184ae-b887-4f03-8b81-c5f0e1a484b3",
  "message_type" : "created.scan",
  "group_id" : "8fb184ae-b887-4f03-8b81-c5f0e1a484b3",
  "created" : 1629808666894632240,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "039957d7-d9bf-4d78-94f6-b554a0dfd6bb",
  "message_type" : "modified.scan",
  "group_id" : "039957d7-d9bf-4d78-94f6-b554a0dfd6bb",
  "created" : 1629808666894652138,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f689216a-8462-4c11-b63c-e70c76c0fe55",
  "message_type" : "stopped.scan",
  "group_id" : "f689216a-8462-4c11-b63c-e70c76c0fe55",
  "created" : 1629808666894671522,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c65b2edb-20fc-4dd9-833e-80ad44e165fa",
  "message_type" : "status.scan",
  "group_id" : "c65b2edb-20fc-4dd9-833e-80ad44e165fa",
  "created" : 1629808666894694213,
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
  "message_id" : "92fbdea0-1ebd-4816-959a-a66944142a0a",
  "message_type" : "got.scan",
  "group_id" : "92fbdea0-1ebd-4816-959a-a66944142a0a",
  "created" : 1629808666894717803,
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
  "message_id" : "a8a5f1e2-1627-4ef7-8f39-06dcbdcce940",
  "message_type" : "result.scan",
  "group_id" : "a8a5f1e2-1627-4ef7-8f39-06dcbdcce940",
  "created" : 1629808666894754653,
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
  "message_id" : "7187e916-0327-4650-9b4f-226b43d7fd0b",
  "message_type" : "failure.start.scan",
  "group_id" : "7187e916-0327-4650-9b4f-226b43d7fd0b",
  "created" : 1629808666894784600,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4e6cffbb-7a16-4896-804d-d769678a48e6",
  "message_type" : "failure.stop.scan",
  "group_id" : "4e6cffbb-7a16-4896-804d-d769678a48e6",
  "created" : 1629808666894806278,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f1558f13-63fc-48b2-8827-54a9a636a919",
  "message_type" : "failure.create.scan",
  "group_id" : "f1558f13-63fc-48b2-8827-54a9a636a919",
  "created" : 1629808666894827255,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b9b09993-8459-4674-acfa-8659be21796e",
  "message_type" : "failure.modify.scan",
  "group_id" : "b9b09993-8459-4674-acfa-8659be21796e",
  "created" : 1629808666894847688,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c0884da8-1731-4460-8d3d-136a7da76372",
  "message_type" : "failure.get.scan",
  "group_id" : "c0884da8-1731-4460-8d3d-136a7da76372",
  "created" : 1629808666894879015,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3fe33480-6c20-498c-b715-9a47c6f2280f",
  "message_type" : "failure.scan",
  "group_id" : "3fe33480-6c20-498c-b715-9a47c6f2280f",
  "created" : 1629808666894900320,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "62681ab6-9bd7-4302-82fa-9a9b474fa1ac",
  "message_type" : "create.scan",
  "group_id" : "62681ab6-9bd7-4302-82fa-9a9b474fa1ac",
  "created" : 1629808666894915002
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "7057f2d9-b85f-449c-be2e-2eb0619e7a15",
  "message_type" : "start.scan",
  "group_id" : "7057f2d9-b85f-449c-be2e-2eb0619e7a15",
  "created" : 1629808666894937510,
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
  "message_id" : "c38379b2-b549-498f-a833-ded4475ebebc",
  "message_type" : "stop.scan",
  "group_id" : "c38379b2-b549-498f-a833-ded4475ebebc",
  "created" : 1629808666894968162,
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
  "message_id" : "f2aa1f89-6533-429e-94fe-f71ff5b05de9",
  "message_type" : "get.scan",
  "group_id" : "f2aa1f89-6533-429e-94fe-f71ff5b05de9",
  "created" : 1629808666894997022,
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
  "message_id" : "01b1092c-34fe-4466-9413-8bc38aa67f86",
  "message_type" : "modify.scan",
  "group_id" : "01b1092c-34fe-4466-9413-8bc38aa67f86",
  "created" : 1629808666895025534,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0eb99be1-1355-4933-ab57-11bc3789bc78",
  "message_type" : "created.scan",
  "group_id" : "0eb99be1-1355-4933-ab57-11bc3789bc78",
  "created" : 1629808666895056709,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "11859783-b843-4ce0-90b8-bbfb20289349",
  "message_type" : "modified.scan",
  "group_id" : "11859783-b843-4ce0-90b8-bbfb20289349",
  "created" : 1629808666895077045,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6590dc4f-ea32-4eb9-a807-673c86e3b78f",
  "message_type" : "stopped.scan",
  "group_id" : "6590dc4f-ea32-4eb9-a807-673c86e3b78f",
  "created" : 1629808666895096412,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a1d06389-76c6-43cc-8a13-cb3b51b8980b",
  "message_type" : "status.scan",
  "group_id" : "a1d06389-76c6-43cc-8a13-cb3b51b8980b",
  "created" : 1629808666895119132,
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
  "message_id" : "e43e5225-4be1-4bf4-a409-ecdc98a35d6f",
  "message_type" : "got.scan",
  "group_id" : "e43e5225-4be1-4bf4-a409-ecdc98a35d6f",
  "created" : 1629808666895143616,
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
  "message_id" : "d5c673f3-2762-4025-86e0-f52d2a5f1ca2",
  "message_type" : "result.scan",
  "group_id" : "d5c673f3-2762-4025-86e0-f52d2a5f1ca2",
  "created" : 1629808666895189724,
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
  "message_id" : "948cdecb-2b71-4ca7-bc33-4cfbad081357",
  "message_type" : "failure.start.scan",
  "group_id" : "948cdecb-2b71-4ca7-bc33-4cfbad081357",
  "created" : 1629808666895221421,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2cc77561-9c01-4a1d-9e58-18b5957d3d94",
  "message_type" : "failure.stop.scan",
  "group_id" : "2cc77561-9c01-4a1d-9e58-18b5957d3d94",
  "created" : 1629808666895243810,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "db913eb9-7a84-48c9-aff9-4cf2aa4a8f65",
  "message_type" : "failure.create.scan",
  "group_id" : "db913eb9-7a84-48c9-aff9-4cf2aa4a8f65",
  "created" : 1629808666895264630,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3d6899e2-1ebf-463f-9b41-d5057851d678",
  "message_type" : "failure.modify.scan",
  "group_id" : "3d6899e2-1ebf-463f-9b41-d5057851d678",
  "created" : 1629808666895285238,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2f995c87-3b4f-4785-aa16-6b297c64dfa1",
  "message_type" : "failure.get.scan",
  "group_id" : "2f995c87-3b4f-4785-aa16-6b297c64dfa1",
  "created" : 1629808666895308535,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "bc1698c0-3fdf-4f20-b731-6747f1164f84",
  "message_type" : "failure.scan",
  "group_id" : "bc1698c0-3fdf-4f20-b731-6747f1164f84",
  "created" : 1629808666895328089,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "0713c510-0fe2-439c-b158-3baf1fafa604",
  "message_type" : "create.scan",
  "group_id" : "0713c510-0fe2-439c-b158-3baf1fafa604",
  "created" : 1629808666895344803
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "e11e9e23-d776-4f03-8183-17691357048d",
  "message_type" : "start.scan",
  "group_id" : "e11e9e23-d776-4f03-8183-17691357048d",
  "created" : 1629808666895376343,
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
  "message_id" : "8f731a20-7416-451c-ac2f-f3a1af2d4c71",
  "message_type" : "stop.scan",
  "group_id" : "8f731a20-7416-451c-ac2f-f3a1af2d4c71",
  "created" : 1629808666895407297,
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
  "message_id" : "d6770046-c5f7-41d8-9762-a51d979973e3",
  "message_type" : "get.scan",
  "group_id" : "d6770046-c5f7-41d8-9762-a51d979973e3",
  "created" : 1629808666895436305,
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
  "message_id" : "e29d978b-c7c5-419e-8407-c1f4719f0f44",
  "message_type" : "modify.scan",
  "group_id" : "e29d978b-c7c5-419e-8407-c1f4719f0f44",
  "created" : 1629808666895464544,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a2a55698-421e-4b91-83f4-9e0aad60cc57",
  "message_type" : "created.scan",
  "group_id" : "a2a55698-421e-4b91-83f4-9e0aad60cc57",
  "created" : 1629808666895505428,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "ff8e57fd-fd1e-4383-90a6-216de6c38d7d",
  "message_type" : "modified.scan",
  "group_id" : "ff8e57fd-fd1e-4383-90a6-216de6c38d7d",
  "created" : 1629808666895526706,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4bbb8edf-e5be-47d3-8a75-3b5e6a48fee4",
  "message_type" : "stopped.scan",
  "group_id" : "4bbb8edf-e5be-47d3-8a75-3b5e6a48fee4",
  "created" : 1629808666895548783,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "dc15f8ba-0f63-426b-9f4c-1777f45c5f8d",
  "message_type" : "status.scan",
  "group_id" : "dc15f8ba-0f63-426b-9f4c-1777f45c5f8d",
  "created" : 1629808666895568424,
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
  "message_id" : "92e469a6-4520-4690-b712-abed1ffcc2b6",
  "message_type" : "got.scan",
  "group_id" : "92e469a6-4520-4690-b712-abed1ffcc2b6",
  "created" : 1629808666895592279,
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
  "message_id" : "7399bce0-88fe-4140-a45e-7cef630e12df",
  "message_type" : "result.scan",
  "group_id" : "7399bce0-88fe-4140-a45e-7cef630e12df",
  "created" : 1629808666895629679,
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
  "message_id" : "d9243c9e-8a60-4c9e-b898-e6df79ace493",
  "message_type" : "failure.start.scan",
  "group_id" : "d9243c9e-8a60-4c9e-b898-e6df79ace493",
  "created" : 1629808666895665324,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2f5df939-7083-446f-81ba-045515bb6ea2",
  "message_type" : "failure.stop.scan",
  "group_id" : "2f5df939-7083-446f-81ba-045515bb6ea2",
  "created" : 1629808666895687841,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "34649d76-8e49-43ad-ad0b-89521fdb5416",
  "message_type" : "failure.create.scan",
  "group_id" : "34649d76-8e49-43ad-ad0b-89521fdb5416",
  "created" : 1629808666895708716,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7a980b92-ee64-4826-8028-336bf4e71ea8",
  "message_type" : "failure.modify.scan",
  "group_id" : "7a980b92-ee64-4826-8028-336bf4e71ea8",
  "created" : 1629808666895729262,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fcedddd0-5c17-446e-898d-3e17981ec4d6",
  "message_type" : "failure.get.scan",
  "group_id" : "fcedddd0-5c17-446e-898d-3e17981ec4d6",
  "created" : 1629808666895749951,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "719a1827-271d-4329-98d2-e41075373d5a",
  "message_type" : "failure.scan",
  "group_id" : "719a1827-271d-4329-98d2-e41075373d5a",
  "created" : 1629808666895778299,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "9ca5beff-a9ca-4c54-bf73-476c5b9eb5bf",
  "message_type" : "create.scan",
  "group_id" : "9ca5beff-a9ca-4c54-bf73-476c5b9eb5bf",
  "created" : 1629808666895794374
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "8d444ab9-974d-4da9-abcb-4e42221de109",
  "message_type" : "start.scan",
  "group_id" : "8d444ab9-974d-4da9-abcb-4e42221de109",
  "created" : 1629808666895816466,
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
  "message_id" : "87bbfcf1-792c-41ad-b48e-dc463cd8842c",
  "message_type" : "stop.scan",
  "group_id" : "87bbfcf1-792c-41ad-b48e-dc463cd8842c",
  "created" : 1629808666895847973,
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
  "message_id" : "d8772de2-ab21-4c13-b9a0-1eec632d33b9",
  "message_type" : "get.scan",
  "group_id" : "d8772de2-ab21-4c13-b9a0-1eec632d33b9",
  "created" : 1629808666895877213,
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
  "message_id" : "2eb7a9d2-bcd9-4790-93b6-2f5a92a7b0a5",
  "message_type" : "modify.scan",
  "group_id" : "2eb7a9d2-bcd9-4790-93b6-2f5a92a7b0a5",
  "created" : 1629808666895905389,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4cac6807-a137-4bdd-8436-1ebc1276ef4f",
  "message_type" : "created.scan",
  "group_id" : "4cac6807-a137-4bdd-8436-1ebc1276ef4f",
  "created" : 1629808666895936743,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "222dbb39-dde3-4ab3-9a96-a5ee125c5dad",
  "message_type" : "modified.scan",
  "group_id" : "222dbb39-dde3-4ab3-9a96-a5ee125c5dad",
  "created" : 1629808666895959920,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "50589a4d-5f8b-4839-a979-1fe2b836ad43",
  "message_type" : "stopped.scan",
  "group_id" : "50589a4d-5f8b-4839-a979-1fe2b836ad43",
  "created" : 1629808666895979447,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "94b98279-1e75-4c4f-aed7-090ce43a461f",
  "message_type" : "status.scan",
  "group_id" : "94b98279-1e75-4c4f-aed7-090ce43a461f",
  "created" : 1629808666895998550,
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
  "message_id" : "a3f751d2-a600-4aaf-8a07-198ce9df1994",
  "message_type" : "got.scan",
  "group_id" : "a3f751d2-a600-4aaf-8a07-198ce9df1994",
  "created" : 1629808666896021592,
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
  "message_id" : "6e1ce7de-98c5-45ba-ad6d-72afb0ab7bd8",
  "message_type" : "result.scan",
  "group_id" : "6e1ce7de-98c5-45ba-ad6d-72afb0ab7bd8",
  "created" : 1629808666896057885,
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
  "message_id" : "a5ea4797-586f-49d6-8840-f546a0f5cbc8",
  "message_type" : "failure.start.scan",
  "group_id" : "a5ea4797-586f-49d6-8840-f546a0f5cbc8",
  "created" : 1629808666896099785,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c77bed96-cb0f-437a-998a-9d7d2a4a5c20",
  "message_type" : "failure.stop.scan",
  "group_id" : "c77bed96-cb0f-437a-998a-9d7d2a4a5c20",
  "created" : 1629808666896123029,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b8df207a-8601-45b4-ba16-af3bd13c33fb",
  "message_type" : "failure.create.scan",
  "group_id" : "b8df207a-8601-45b4-ba16-af3bd13c33fb",
  "created" : 1629808666896143909,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "96fdd597-7a1e-44b3-ae81-b767ef108877",
  "message_type" : "failure.modify.scan",
  "group_id" : "96fdd597-7a1e-44b3-ae81-b767ef108877",
  "created" : 1629808666896164824,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "37c31ef7-ef11-41b9-abfb-2d5a5930d54a",
  "message_type" : "failure.get.scan",
  "group_id" : "37c31ef7-ef11-41b9-abfb-2d5a5930d54a",
  "created" : 1629808666896185162,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "e995db49-f11f-4c18-ad12-6c6507632dcc",
  "message_type" : "failure.scan",
  "group_id" : "e995db49-f11f-4c18-ad12-6c6507632dcc",
  "created" : 1629808666896204750,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "c5025add-cfc6-483c-b7f1-30296052d126",
  "message_type" : "create.scan",
  "group_id" : "c5025add-cfc6-483c-b7f1-30296052d126",
  "created" : 1629808666896223376
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "730dc2b9-075d-4119-8e0c-8e61d7bf4278",
  "message_type" : "start.scan",
  "group_id" : "730dc2b9-075d-4119-8e0c-8e61d7bf4278",
  "created" : 1629808666896246223,
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
  "message_id" : "9a03f139-d2f7-44b6-9fe2-3400b1fb3cb6",
  "message_type" : "stop.scan",
  "group_id" : "9a03f139-d2f7-44b6-9fe2-3400b1fb3cb6",
  "created" : 1629808666896277125,
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
  "message_id" : "44ae75e7-340a-4854-bca1-354a2c91c6e3",
  "message_type" : "get.scan",
  "group_id" : "44ae75e7-340a-4854-bca1-354a2c91c6e3",
  "created" : 1629808666896305930,
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
  "message_id" : "38f9ed0e-54da-447c-a963-ee1c95feaf50",
  "message_type" : "modify.scan",
  "group_id" : "38f9ed0e-54da-447c-a963-ee1c95feaf50",
  "created" : 1629808666896334206,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "19aebd11-7802-4449-9044-7483e35ed758",
  "message_type" : "created.scan",
  "group_id" : "19aebd11-7802-4449-9044-7483e35ed758",
  "created" : 1629808666896368215,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "88fa8380-4e94-42ad-91c9-2e079815520b",
  "message_type" : "modified.scan",
  "group_id" : "88fa8380-4e94-42ad-91c9-2e079815520b",
  "created" : 1629808666896396403,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "63a68742-b9fd-444a-8d30-2f40b2d25b47",
  "message_type" : "stopped.scan",
  "group_id" : "63a68742-b9fd-444a-8d30-2f40b2d25b47",
  "created" : 1629808666896416659,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4413e4b8-a11d-4b41-8e73-e75badd4984e",
  "message_type" : "status.scan",
  "group_id" : "4413e4b8-a11d-4b41-8e73-e75badd4984e",
  "created" : 1629808666896435754,
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
  "message_id" : "4b513cab-246f-4db5-aa1a-ce143517857e",
  "message_type" : "got.scan",
  "group_id" : "4b513cab-246f-4db5-aa1a-ce143517857e",
  "created" : 1629808666896459396,
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
  "message_id" : "d358ffab-b251-4800-81a6-aa1dbbba643b",
  "message_type" : "result.scan",
  "group_id" : "d358ffab-b251-4800-81a6-aa1dbbba643b",
  "created" : 1629808666896496268,
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
  "message_id" : "d7be1ec0-a513-4de9-893b-c135c07257d8",
  "message_type" : "failure.start.scan",
  "group_id" : "d7be1ec0-a513-4de9-893b-c135c07257d8",
  "created" : 1629808666896528958,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "902f063b-896b-4ec1-a383-454ea633d520",
  "message_type" : "failure.stop.scan",
  "group_id" : "902f063b-896b-4ec1-a383-454ea633d520",
  "created" : 1629808666896551260,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3a17c034-6382-4632-87bd-efcc5c1379ae",
  "message_type" : "failure.create.scan",
  "group_id" : "3a17c034-6382-4632-87bd-efcc5c1379ae",
  "created" : 1629808666896571751,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "92e00884-ec4f-4ca9-8f8b-e5c0afb66768",
  "message_type" : "failure.modify.scan",
  "group_id" : "92e00884-ec4f-4ca9-8f8b-e5c0afb66768",
  "created" : 1629808666896592475,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c8232fde-89cb-41a0-903b-2ae149a41ca3",
  "message_type" : "failure.get.scan",
  "group_id" : "c8232fde-89cb-41a0-903b-2ae149a41ca3",
  "created" : 1629808666896612755,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c70dcbc5-b17b-4f10-aa0a-8d3d91779b2b",
  "message_type" : "failure.scan",
  "group_id" : "c70dcbc5-b17b-4f10-aa0a-8d3d91779b2b",
  "created" : 1629808666896632501,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "93bba3c2-7f8e-4fd5-95e9-55187ac02c7a",
  "message_type" : "create.scan",
  "group_id" : "93bba3c2-7f8e-4fd5-95e9-55187ac02c7a",
  "created" : 1629808666896658856
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "7e1b6568-8d23-4676-8121-b72aad79044a",
  "message_type" : "start.scan",
  "group_id" : "7e1b6568-8d23-4676-8121-b72aad79044a",
  "created" : 1629808666896682476,
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
  "message_id" : "5f30dc7d-726c-44fe-b1cf-598ba3776166",
  "message_type" : "stop.scan",
  "group_id" : "5f30dc7d-726c-44fe-b1cf-598ba3776166",
  "created" : 1629808666896713853,
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
  "message_id" : "da849829-6165-47f8-b955-5d4f2daaac87",
  "message_type" : "get.scan",
  "group_id" : "da849829-6165-47f8-b955-5d4f2daaac87",
  "created" : 1629808666896743244,
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
  "message_id" : "bf2983b7-0feb-4cb6-a8cd-96a446a7e951",
  "message_type" : "modify.scan",
  "group_id" : "bf2983b7-0feb-4cb6-a8cd-96a446a7e951",
  "created" : 1629808666896771594,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "ec395309-5bbe-4277-b5f6-511f4e628a62",
  "message_type" : "created.scan",
  "group_id" : "ec395309-5bbe-4277-b5f6-511f4e628a62",
  "created" : 1629808666896806090,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "1ffdb053-e14c-4524-8824-021e194c9635",
  "message_type" : "modified.scan",
  "group_id" : "1ffdb053-e14c-4524-8824-021e194c9635",
  "created" : 1629808666896825882,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3252df68-ee7e-4a83-a10e-4d5b9b53e6a4",
  "message_type" : "stopped.scan",
  "group_id" : "3252df68-ee7e-4a83-a10e-4d5b9b53e6a4",
  "created" : 1629808666896844977,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a92e478a-6722-41f5-8c21-d72e2ee11f5a",
  "message_type" : "status.scan",
  "group_id" : "a92e478a-6722-41f5-8c21-d72e2ee11f5a",
  "created" : 1629808666896863930,
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
  "message_id" : "af819657-d013-4c4e-8695-c6d11a97be94",
  "message_type" : "got.scan",
  "group_id" : "af819657-d013-4c4e-8695-c6d11a97be94",
  "created" : 1629808666896887283,
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
  "message_id" : "9f5688b2-c783-4c8e-aec9-5370a1c9350e",
  "message_type" : "result.scan",
  "group_id" : "9f5688b2-c783-4c8e-aec9-5370a1c9350e",
  "created" : 1629808666896923934,
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
  "message_id" : "f04ee3fd-2fc3-4c1f-8b96-73966cb4b97f",
  "message_type" : "failure.start.scan",
  "group_id" : "f04ee3fd-2fc3-4c1f-8b96-73966cb4b97f",
  "created" : 1629808666896956712,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "17c2e155-f267-4d9f-84bd-ae70f5a1451b",
  "message_type" : "failure.stop.scan",
  "group_id" : "17c2e155-f267-4d9f-84bd-ae70f5a1451b",
  "created" : 1629808666896987130,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "56580797-8d16-44e3-b511-ee770e7d892c",
  "message_type" : "failure.create.scan",
  "group_id" : "56580797-8d16-44e3-b511-ee770e7d892c",
  "created" : 1629808666897008975,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5bb50b2f-ee31-4389-b45a-953023b2e8df",
  "message_type" : "failure.modify.scan",
  "group_id" : "5bb50b2f-ee31-4389-b45a-953023b2e8df",
  "created" : 1629808666897030113,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "72ae1b24-0311-4728-8a35-1e2c9b04155b",
  "message_type" : "failure.get.scan",
  "group_id" : "72ae1b24-0311-4728-8a35-1e2c9b04155b",
  "created" : 1629808666897050066,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5b35e2c2-f9fc-4560-acb0-8b980dc7dce9",
  "message_type" : "failure.scan",
  "group_id" : "5b35e2c2-f9fc-4560-acb0-8b980dc7dce9",
  "created" : 1629808666897069649,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "f1486a1b-044e-42e5-92ff-3286e6f127b1",
  "message_type" : "create.scan",
  "group_id" : "f1486a1b-044e-42e5-92ff-3286e6f127b1",
  "created" : 1629808666897088596
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "653a7cdb-3754-4dbd-8db1-c108f9f1d120",
  "message_type" : "start.scan",
  "group_id" : "653a7cdb-3754-4dbd-8db1-c108f9f1d120",
  "created" : 1629808666897111110,
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
  "message_id" : "61b2cf01-7dc3-4a0c-8798-a56f34dbb92d",
  "message_type" : "stop.scan",
  "group_id" : "61b2cf01-7dc3-4a0c-8798-a56f34dbb92d",
  "created" : 1629808666897142061,
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
  "message_id" : "d16bce7f-28aa-4727-ae71-6e49eeac9570",
  "message_type" : "get.scan",
  "group_id" : "d16bce7f-28aa-4727-ae71-6e49eeac9570",
  "created" : 1629808666897171235,
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
  "message_id" : "40b19687-dc4c-4d78-a032-5381da652621",
  "message_type" : "modify.scan",
  "group_id" : "40b19687-dc4c-4d78-a032-5381da652621",
  "created" : 1629808666897206833,
  "id" : "example.scan.id",
  "temporary" : false,
  "target_id" : "example.target.id"
}
```
Responses:

- [modified](#modifiedscan)
- [failure.modify](#failuremodifyscan)
## created/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "960853de-99b9-45c0-bc72-7a66a084aed8",
  "message_type" : "created.scan",
  "group_id" : "960853de-99b9-45c0-bc72-7a66a084aed8",
  "created" : 1629808666897238496,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3944ef4d-5ed4-4dec-835a-1fb1de692204",
  "message_type" : "modified.scan",
  "group_id" : "3944ef4d-5ed4-4dec-835a-1fb1de692204",
  "created" : 1629808666897258746,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f21e1f9e-1d8a-44ce-b1e6-d2222870efed",
  "message_type" : "stopped.scan",
  "group_id" : "f21e1f9e-1d8a-44ce-b1e6-d2222870efed",
  "created" : 1629808666897286157,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "62328650-84bc-4a59-91c3-01a61915f6de",
  "message_type" : "status.scan",
  "group_id" : "62328650-84bc-4a59-91c3-01a61915f6de",
  "created" : 1629808666897306795,
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
  "message_id" : "119266c2-51bd-4096-8f4f-8c73ffca266e",
  "message_type" : "got.scan",
  "group_id" : "119266c2-51bd-4096-8f4f-8c73ffca266e",
  "created" : 1629808666897330112,
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
  "message_id" : "ffac140b-988c-464f-acaa-797caa6ee0b2",
  "message_type" : "result.scan",
  "group_id" : "ffac140b-988c-464f-acaa-797caa6ee0b2",
  "created" : 1629808666897370305,
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
  "message_id" : "6a4e5aee-3d7d-4b8b-921a-518234193fd8",
  "message_type" : "failure.start.scan",
  "group_id" : "6a4e5aee-3d7d-4b8b-921a-518234193fd8",
  "created" : 1629808666897400872,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3df5894d-9aeb-47a9-a101-1b9fad1a95ff",
  "message_type" : "failure.stop.scan",
  "group_id" : "3df5894d-9aeb-47a9-a101-1b9fad1a95ff",
  "created" : 1629808666897422762,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3d31faf7-a317-4990-81a0-25e1e531e9f8",
  "message_type" : "failure.create.scan",
  "group_id" : "3d31faf7-a317-4990-81a0-25e1e531e9f8",
  "created" : 1629808666897443350,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "30984e3d-0cec-46a0-8b26-c0e43c6942d6",
  "message_type" : "failure.modify.scan",
  "group_id" : "30984e3d-0cec-46a0-8b26-c0e43c6942d6",
  "created" : 1629808666897463377,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7ce35fba-17a9-470a-bdbb-b040f6aabcfc",
  "message_type" : "failure.get.scan",
  "group_id" : "7ce35fba-17a9-470a-bdbb-b040f6aabcfc",
  "created" : 1629808666897483212,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f7b3488a-9048-4274-91c7-b2aaa387c75a",
  "message_type" : "failure.scan",
  "group_id" : "f7b3488a-9048-4274-91c7-b2aaa387c75a",
  "created" : 1629808666897502590,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
# target

## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "d33f88ba-f014-448c-96b2-1cf99d2e461b",
  "message_type" : "create.target",
  "group_id" : "d33f88ba-f014-448c-96b2-1cf99d2e461b",
  "created" : 1629808666897522014
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "c5928e02-326c-428f-8e81-ff85edcf4437",
  "message_type" : "get.target",
  "group_id" : "c5928e02-326c-428f-8e81-ff85edcf4437",
  "created" : 1629808666897545318,
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
  "message_id" : "972a5a60-816c-4c65-9a8a-fa1bb5b7a261",
  "message_type" : "modify.target",
  "group_id" : "972a5a60-816c-4c65-9a8a-fa1bb5b7a261",
  "created" : 1629808666897587591,
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
  "message_id" : "05baf4ed-4673-4f19-83e4-48f894d6cc1d",
  "message_type" : "created.target",
  "group_id" : "05baf4ed-4673-4f19-83e4-48f894d6cc1d",
  "created" : 1629808666897636247,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "db8f2b58-8470-4319-851b-91b80cabec15",
  "message_type" : "modified.target",
  "group_id" : "db8f2b58-8470-4319-851b-91b80cabec15",
  "created" : 1629808666897657534,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "29683699-00f3-4787-b47a-32ff553e2a68",
  "message_type" : "got.target",
  "group_id" : "29683699-00f3-4787-b47a-32ff553e2a68",
  "created" : 1629808666897677054,
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
  "message_id" : "c1ef1d79-c8b6-41f4-ab71-542c71415be1",
  "message_type" : "failure.create.target",
  "group_id" : "c1ef1d79-c8b6-41f4-ab71-542c71415be1",
  "created" : 1629808666897709025,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "9991a29d-b3b8-43d8-b3a2-4546db2ddf6e",
  "message_type" : "failure.modify.target",
  "group_id" : "9991a29d-b3b8-43d8-b3a2-4546db2ddf6e",
  "created" : 1629808666897730435,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "6ab8afb9-3558-481b-9537-e4354042f8f2",
  "message_type" : "failure.get.target",
  "group_id" : "6ab8afb9-3558-481b-9537-e4354042f8f2",
  "created" : 1629808666897756042,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "f6b687b1-0711-4cff-b55e-bea4f4f05cc6",
  "message_type" : "failure.target",
  "group_id" : "f6b687b1-0711-4cff-b55e-bea4f4f05cc6",
  "created" : 1629808666897778478,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "832323c7-ec26-4991-a508-a248f4a8129d",
  "message_type" : "create.target",
  "group_id" : "832323c7-ec26-4991-a508-a248f4a8129d",
  "created" : 1629808666897799937
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "776a9fb8-4eb9-4633-ad19-381ae268e84d",
  "message_type" : "get.target",
  "group_id" : "776a9fb8-4eb9-4633-ad19-381ae268e84d",
  "created" : 1629808666897829370,
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
  "message_id" : "1e5dce14-e23e-4d2c-9f34-1caf8c9e5501",
  "message_type" : "modify.target",
  "group_id" : "1e5dce14-e23e-4d2c-9f34-1caf8c9e5501",
  "created" : 1629808666897858918,
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
  "message_id" : "ce3ffdba-aff7-49c5-a4c0-0fb4527a2757",
  "message_type" : "created.target",
  "group_id" : "ce3ffdba-aff7-49c5-a4c0-0fb4527a2757",
  "created" : 1629808666897908612,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "622ae53c-c5f4-4fa8-ad4c-3ac0534e8ad8",
  "message_type" : "modified.target",
  "group_id" : "622ae53c-c5f4-4fa8-ad4c-3ac0534e8ad8",
  "created" : 1629808666897930114,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "bae00fb0-ebff-43bc-806e-7a2dc2883d58",
  "message_type" : "got.target",
  "group_id" : "bae00fb0-ebff-43bc-806e-7a2dc2883d58",
  "created" : 1629808666897950247,
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
  "message_id" : "6b9c4972-5186-4f31-acf8-7c0e9d6ef22f",
  "message_type" : "failure.create.target",
  "group_id" : "6b9c4972-5186-4f31-acf8-7c0e9d6ef22f",
  "created" : 1629808666897987492,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "6be9668e-fe50-428e-9e0c-55c53ebe284e",
  "message_type" : "failure.modify.target",
  "group_id" : "6be9668e-fe50-428e-9e0c-55c53ebe284e",
  "created" : 1629808666898008781,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "c27fc6d1-d454-49ed-b4dc-20083a0309d3",
  "message_type" : "failure.get.target",
  "group_id" : "c27fc6d1-d454-49ed-b4dc-20083a0309d3",
  "created" : 1629808666898029193,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "9bea09ef-a099-4b78-831d-c8db19d65c5b",
  "message_type" : "failure.target",
  "group_id" : "9bea09ef-a099-4b78-831d-c8db19d65c5b",
  "created" : 1629808666898049911,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "6e5b4b69-61a5-4992-8b8d-3f43d3a4da8f",
  "message_type" : "create.target",
  "group_id" : "6e5b4b69-61a5-4992-8b8d-3f43d3a4da8f",
  "created" : 1629808666898070697
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "ea0a5d59-3e47-4b9c-a70e-063c0083ba4a",
  "message_type" : "get.target",
  "group_id" : "ea0a5d59-3e47-4b9c-a70e-063c0083ba4a",
  "created" : 1629808666898099023,
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
  "message_id" : "dfde8132-4032-4643-9b9b-e14bbad3a7b2",
  "message_type" : "modify.target",
  "group_id" : "dfde8132-4032-4643-9b9b-e14bbad3a7b2",
  "created" : 1629808666898128090,
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
  "message_id" : "5fe1099f-8a71-419a-8391-c44e053e5468",
  "message_type" : "created.target",
  "group_id" : "5fe1099f-8a71-419a-8391-c44e053e5468",
  "created" : 1629808666898176054,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "cbc72b21-037c-42f5-986c-234d82deacd1",
  "message_type" : "modified.target",
  "group_id" : "cbc72b21-037c-42f5-986c-234d82deacd1",
  "created" : 1629808666898205008,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "5d46c1a0-e7bb-4446-949c-408e37346fe1",
  "message_type" : "got.target",
  "group_id" : "5d46c1a0-e7bb-4446-949c-408e37346fe1",
  "created" : 1629808666898253293,
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
  "message_id" : "7a2e0c8e-d5e9-46b2-9702-beae45777d41",
  "message_type" : "failure.create.target",
  "group_id" : "7a2e0c8e-d5e9-46b2-9702-beae45777d41",
  "created" : 1629808666898285380,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "80e403aa-b11f-4ec6-9552-198b54f0ed66",
  "message_type" : "failure.modify.target",
  "group_id" : "80e403aa-b11f-4ec6-9552-198b54f0ed66",
  "created" : 1629808666898306769,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "504c70bc-dff4-4afd-ab14-eed2cb2e2faf",
  "message_type" : "failure.get.target",
  "group_id" : "504c70bc-dff4-4afd-ab14-eed2cb2e2faf",
  "created" : 1629808666898327364,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "22b69f6b-5726-498e-9fa2-8d291563f20a",
  "message_type" : "failure.target",
  "group_id" : "22b69f6b-5726-498e-9fa2-8d291563f20a",
  "created" : 1629808666898347575,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "c6afaa25-cefb-4a26-9eba-87dc962aa2ab",
  "message_type" : "create.target",
  "group_id" : "c6afaa25-cefb-4a26-9eba-87dc962aa2ab",
  "created" : 1629808666898369755
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "a891a2fa-0048-4bb2-bc67-53400437d0dc",
  "message_type" : "get.target",
  "group_id" : "a891a2fa-0048-4bb2-bc67-53400437d0dc",
  "created" : 1629808666898398842,
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
  "message_id" : "44e27eee-4a17-49aa-b9f6-4097057c6a3b",
  "message_type" : "modify.target",
  "group_id" : "44e27eee-4a17-49aa-b9f6-4097057c6a3b",
  "created" : 1629808666898427517,
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
  "message_id" : "ca28dd8b-b699-4407-a4db-c3dbe28ce4fc",
  "message_type" : "created.target",
  "group_id" : "ca28dd8b-b699-4407-a4db-c3dbe28ce4fc",
  "created" : 1629808666898472770,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "fefbb91e-2f24-42ea-8cc7-d45fc5f32918",
  "message_type" : "modified.target",
  "group_id" : "fefbb91e-2f24-42ea-8cc7-d45fc5f32918",
  "created" : 1629808666898493076,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "cd85a363-af66-4724-9e3d-2b0598e235b1",
  "message_type" : "got.target",
  "group_id" : "cd85a363-af66-4724-9e3d-2b0598e235b1",
  "created" : 1629808666898512611,
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
  "message_id" : "e72abef4-3e06-409e-927e-77870d5a36ce",
  "message_type" : "failure.create.target",
  "group_id" : "e72abef4-3e06-409e-927e-77870d5a36ce",
  "created" : 1629808666898552786,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "a68285f2-872a-4913-9133-56b3dc7dabfb",
  "message_type" : "failure.modify.target",
  "group_id" : "a68285f2-872a-4913-9133-56b3dc7dabfb",
  "created" : 1629808666898577100,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "07147c16-7b5b-4ef8-bd63-ac21a79cf6c5",
  "message_type" : "failure.get.target",
  "group_id" : "07147c16-7b5b-4ef8-bd63-ac21a79cf6c5",
  "created" : 1629808666898598139,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "2edd65f6-107f-40cb-b672-215e74b02710",
  "message_type" : "failure.target",
  "group_id" : "2edd65f6-107f-40cb-b672-215e74b02710",
  "created" : 1629808666898621339,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "a7e73944-ff3f-40c3-a33d-f8687740fd96",
  "message_type" : "create.target",
  "group_id" : "a7e73944-ff3f-40c3-a33d-f8687740fd96",
  "created" : 1629808666898642134
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "594e0fef-f510-4bf4-9e35-afc4ab7694d3",
  "message_type" : "get.target",
  "group_id" : "594e0fef-f510-4bf4-9e35-afc4ab7694d3",
  "created" : 1629808666898671193,
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
  "message_id" : "301fe2a6-4024-4138-8428-aac9450ffe38",
  "message_type" : "modify.target",
  "group_id" : "301fe2a6-4024-4138-8428-aac9450ffe38",
  "created" : 1629808666898700501,
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
  "message_id" : "aaf75b52-c8d4-41ef-9597-e8401b2f77ff",
  "message_type" : "created.target",
  "group_id" : "aaf75b52-c8d4-41ef-9597-e8401b2f77ff",
  "created" : 1629808666898741734,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "3f506b72-c00d-40ae-a3c9-18ab9378f705",
  "message_type" : "modified.target",
  "group_id" : "3f506b72-c00d-40ae-a3c9-18ab9378f705",
  "created" : 1629808666898764283,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "6f7a7db8-4bca-4121-8e43-6b7b2c9c4360",
  "message_type" : "got.target",
  "group_id" : "6f7a7db8-4bca-4121-8e43-6b7b2c9c4360",
  "created" : 1629808666898783663,
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
  "message_id" : "db497d4b-55bf-41f9-b4a8-352c9a3d0114",
  "message_type" : "failure.create.target",
  "group_id" : "db497d4b-55bf-41f9-b4a8-352c9a3d0114",
  "created" : 1629808666898818482,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "3f97f4a4-4340-4ebe-92d4-cef04751ffb5",
  "message_type" : "failure.modify.target",
  "group_id" : "3f97f4a4-4340-4ebe-92d4-cef04751ffb5",
  "created" : 1629808666898847550,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "424222d6-33a8-4c18-b9c6-d51f9e7b30af",
  "message_type" : "failure.get.target",
  "group_id" : "424222d6-33a8-4c18-b9c6-d51f9e7b30af",
  "created" : 1629808666898869165,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "d3d0d1ba-2f43-401f-8547-9b976dfb71f2",
  "message_type" : "failure.target",
  "group_id" : "d3d0d1ba-2f43-401f-8547-9b976dfb71f2",
  "created" : 1629808666898890035,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "e64db3b7-c005-48bc-b4a8-5136ca4a3aa5",
  "message_type" : "create.target",
  "group_id" : "e64db3b7-c005-48bc-b4a8-5136ca4a3aa5",
  "created" : 1629808666898910312
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "8a8d28dd-d89b-445d-bd55-67243a699a22",
  "message_type" : "get.target",
  "group_id" : "8a8d28dd-d89b-445d-bd55-67243a699a22",
  "created" : 1629808666898939715,
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
  "message_id" : "3c89c38d-bd16-4ef2-8a88-248454d85d08",
  "message_type" : "modify.target",
  "group_id" : "3c89c38d-bd16-4ef2-8a88-248454d85d08",
  "created" : 1629808666898969922,
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
  "message_id" : "da04d124-5379-49f2-85fc-1e381bf1a048",
  "message_type" : "created.target",
  "group_id" : "da04d124-5379-49f2-85fc-1e381bf1a048",
  "created" : 1629808666899014870,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "bcb69d44-b60c-4e6f-9e17-b9636372a5ec",
  "message_type" : "modified.target",
  "group_id" : "bcb69d44-b60c-4e6f-9e17-b9636372a5ec",
  "created" : 1629808666899034986,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1ec543cf-c11c-469b-afe9-61c744d764bb",
  "message_type" : "got.target",
  "group_id" : "1ec543cf-c11c-469b-afe9-61c744d764bb",
  "created" : 1629808666899054330,
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
  "message_id" : "ab15f251-e8a0-42bf-bd2e-9b8c744ea228",
  "message_type" : "failure.create.target",
  "group_id" : "ab15f251-e8a0-42bf-bd2e-9b8c744ea228",
  "created" : 1629808666899085413,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "37a48f87-beec-479a-8232-a91ea1cce3ef",
  "message_type" : "failure.modify.target",
  "group_id" : "37a48f87-beec-479a-8232-a91ea1cce3ef",
  "created" : 1629808666899106116,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "b0c4e30f-283f-42e6-993e-1e9c5e83ae24",
  "message_type" : "failure.get.target",
  "group_id" : "b0c4e30f-283f-42e6-993e-1e9c5e83ae24",
  "created" : 1629808666899127113,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "2dc4eec9-e06b-4ff3-9d7b-55a2ca59e07c",
  "message_type" : "failure.target",
  "group_id" : "2dc4eec9-e06b-4ff3-9d7b-55a2ca59e07c",
  "created" : 1629808666899157735,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "8d295068-b64e-49bc-93a8-86683331d839",
  "message_type" : "create.target",
  "group_id" : "8d295068-b64e-49bc-93a8-86683331d839",
  "created" : 1629808666899179983
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "d2f9637e-d7a7-4313-80b6-b05bb8e8c662",
  "message_type" : "get.target",
  "group_id" : "d2f9637e-d7a7-4313-80b6-b05bb8e8c662",
  "created" : 1629808666899208362,
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
  "message_id" : "4782a2ef-fbc7-4336-90c0-3ab761b86c66",
  "message_type" : "modify.target",
  "group_id" : "4782a2ef-fbc7-4336-90c0-3ab761b86c66",
  "created" : 1629808666899237459,
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
  "message_id" : "e0a845ee-a44e-46f5-aab7-692e44a3f8c1",
  "message_type" : "created.target",
  "group_id" : "e0a845ee-a44e-46f5-aab7-692e44a3f8c1",
  "created" : 1629808666899282130,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "ccf01bb1-49c0-4ea4-85cf-7fd126ece1e7",
  "message_type" : "modified.target",
  "group_id" : "ccf01bb1-49c0-4ea4-85cf-7fd126ece1e7",
  "created" : 1629808666899303531,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "98f56597-5e14-4f63-a88e-e1d509913c08",
  "message_type" : "got.target",
  "group_id" : "98f56597-5e14-4f63-a88e-e1d509913c08",
  "created" : 1629808666899323741,
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
  "message_id" : "833a296d-fd4e-44d6-8a7f-a76d0b6ef002",
  "message_type" : "failure.create.target",
  "group_id" : "833a296d-fd4e-44d6-8a7f-a76d0b6ef002",
  "created" : 1629808666899363469,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "a4b109e7-d540-40f0-9193-15a26718ab5b",
  "message_type" : "failure.modify.target",
  "group_id" : "a4b109e7-d540-40f0-9193-15a26718ab5b",
  "created" : 1629808666899385700,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "7bb1ab36-d73b-46c4-a4e5-60afac762bd4",
  "message_type" : "failure.get.target",
  "group_id" : "7bb1ab36-d73b-46c4-a4e5-60afac762bd4",
  "created" : 1629808666899406210,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "c43c4f64-d266-442d-8c56-843bfb255277",
  "message_type" : "failure.target",
  "group_id" : "c43c4f64-d266-442d-8c56-843bfb255277",
  "created" : 1629808666899432487,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "ceccd68e-cf84-4abe-a878-f4c6174c59f0",
  "message_type" : "create.target",
  "group_id" : "ceccd68e-cf84-4abe-a878-f4c6174c59f0",
  "created" : 1629808666899463359
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "8b7aa6d5-cd82-4329-857e-cb4713569142",
  "message_type" : "get.target",
  "group_id" : "8b7aa6d5-cd82-4329-857e-cb4713569142",
  "created" : 1629808666899493172,
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
  "message_id" : "1d4a2cbb-aad6-481e-8ec8-f3150d187bb1",
  "message_type" : "modify.target",
  "group_id" : "1d4a2cbb-aad6-481e-8ec8-f3150d187bb1",
  "created" : 1629808666899525485,
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
  "message_id" : "ccff4ab9-8cbf-488e-b397-640d2c5932aa",
  "message_type" : "created.target",
  "group_id" : "ccff4ab9-8cbf-488e-b397-640d2c5932aa",
  "created" : 1629808666899569663,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "eb740e4f-bce0-4863-a079-f97027dca747",
  "message_type" : "modified.target",
  "group_id" : "eb740e4f-bce0-4863-a079-f97027dca747",
  "created" : 1629808666899590613,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "c84b1280-2829-4a26-af42-059f63327863",
  "message_type" : "got.target",
  "group_id" : "c84b1280-2829-4a26-af42-059f63327863",
  "created" : 1629808666899609732,
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
  "message_id" : "e74bace7-428a-4c64-8c49-5c55aa6762f7",
  "message_type" : "failure.create.target",
  "group_id" : "e74bace7-428a-4c64-8c49-5c55aa6762f7",
  "created" : 1629808666899644135,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "66f63052-2ce1-4a80-bb04-0e9053871d40",
  "message_type" : "failure.modify.target",
  "group_id" : "66f63052-2ce1-4a80-bb04-0e9053871d40",
  "created" : 1629808666899665530,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "a267465a-d903-4bb9-9fed-af0328d9e238",
  "message_type" : "failure.get.target",
  "group_id" : "a267465a-d903-4bb9-9fed-af0328d9e238",
  "created" : 1629808666899685936,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "83c9c8d0-1668-4ea5-9f26-ee039d3cbda7",
  "message_type" : "failure.target",
  "group_id" : "83c9c8d0-1668-4ea5-9f26-ee039d3cbda7",
  "created" : 1629808666899708492,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "03433a9f-3e79-4aab-8e23-2496e4986d95",
  "message_type" : "create.target",
  "group_id" : "03433a9f-3e79-4aab-8e23-2496e4986d95",
  "created" : 1629808666899729992
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "e212d87a-8220-4f51-9cd6-d3a458c27b51",
  "message_type" : "get.target",
  "group_id" : "e212d87a-8220-4f51-9cd6-d3a458c27b51",
  "created" : 1629808666899758719,
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
  "message_id" : "d1b53bf0-582f-4b6c-807a-2115c0f50dd8",
  "message_type" : "modify.target",
  "group_id" : "d1b53bf0-582f-4b6c-807a-2115c0f50dd8",
  "created" : 1629808666899796897,
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
  "message_id" : "e19d1d88-9393-4e0b-bd2b-ccf340f96576",
  "message_type" : "created.target",
  "group_id" : "e19d1d88-9393-4e0b-bd2b-ccf340f96576",
  "created" : 1629808666899839301,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "970e9951-e432-439c-964d-31f555c7c849",
  "message_type" : "modified.target",
  "group_id" : "970e9951-e432-439c-964d-31f555c7c849",
  "created" : 1629808666899859612,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1a1d7d40-52e1-4df5-ba92-49d6dbc24cfa",
  "message_type" : "got.target",
  "group_id" : "1a1d7d40-52e1-4df5-ba92-49d6dbc24cfa",
  "created" : 1629808666899882192,
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
  "message_id" : "d1c3b54a-38fe-4f0a-8304-99c593bd0afc",
  "message_type" : "failure.create.target",
  "group_id" : "d1c3b54a-38fe-4f0a-8304-99c593bd0afc",
  "created" : 1629808666899914817,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "c84cd9c7-1a0a-4e05-ae4b-59829bd9dcd2",
  "message_type" : "failure.modify.target",
  "group_id" : "c84cd9c7-1a0a-4e05-ae4b-59829bd9dcd2",
  "created" : 1629808666899936552,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "dd86567f-0f53-4f94-86ef-b202f4c46223",
  "message_type" : "failure.get.target",
  "group_id" : "dd86567f-0f53-4f94-86ef-b202f4c46223",
  "created" : 1629808666899956888,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "7feda9e1-f419-4fed-8d0b-9e148a264dda",
  "message_type" : "failure.target",
  "group_id" : "7feda9e1-f419-4fed-8d0b-9e148a264dda",
  "created" : 1629808666899977846,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "4a4dc207-d54c-45fc-b688-eecb51e61358",
  "message_type" : "create.target",
  "group_id" : "4a4dc207-d54c-45fc-b688-eecb51e61358",
  "created" : 1629808666899998521
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "0610fa74-778e-49f7-a8f8-9be2b28cd66f",
  "message_type" : "get.target",
  "group_id" : "0610fa74-778e-49f7-a8f8-9be2b28cd66f",
  "created" : 1629808666900026570,
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
  "message_id" : "e9649527-1bfd-4e2a-becb-63efbac4c374",
  "message_type" : "modify.target",
  "group_id" : "e9649527-1bfd-4e2a-becb-63efbac4c374",
  "created" : 1629808666900055647,
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
  "message_id" : "5b054a3a-833d-4b97-904f-b2b637ce4c4c",
  "message_type" : "created.target",
  "group_id" : "5b054a3a-833d-4b97-904f-b2b637ce4c4c",
  "created" : 1629808666900110586,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "00550ce8-69fd-429b-a5b1-dd60270c4157",
  "message_type" : "modified.target",
  "group_id" : "00550ce8-69fd-429b-a5b1-dd60270c4157",
  "created" : 1629808666900133604,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "c6536059-4914-4b8c-b446-5684cfee5dce",
  "message_type" : "got.target",
  "group_id" : "c6536059-4914-4b8c-b446-5684cfee5dce",
  "created" : 1629808666900153072,
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
  "message_id" : "221d2ce2-81a7-400a-b0d8-68f02b682874",
  "message_type" : "failure.create.target",
  "group_id" : "221d2ce2-81a7-400a-b0d8-68f02b682874",
  "created" : 1629808666900184190,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1ed4196f-b70a-4e8e-ab35-4a6b907b0c14",
  "message_type" : "failure.modify.target",
  "group_id" : "1ed4196f-b70a-4e8e-ab35-4a6b907b0c14",
  "created" : 1629808666900204917,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "49fac8bc-c6f2-4e98-838c-93a2ed86056a",
  "message_type" : "failure.get.target",
  "group_id" : "49fac8bc-c6f2-4e98-838c-93a2ed86056a",
  "created" : 1629808666900225176,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "9b8a96e3-2987-40c3-bc2c-6ff5b0642b9c",
  "message_type" : "failure.target",
  "group_id" : "9b8a96e3-2987-40c3-bc2c-6ff5b0642b9c",
  "created" : 1629808666900248887,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
