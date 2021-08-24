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
  "message_id" : "c2598e5e-6882-4aa7-8818-618aa162f24f",
  "message_type" : "create.scan",
  "group_id" : "c2598e5e-6882-4aa7-8818-618aa162f24f",
  "created" : 1629809644498304206
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "df9ece80-1e7e-490f-b418-1ca9a0d637e8",
  "message_type" : "start.scan",
  "group_id" : "df9ece80-1e7e-490f-b418-1ca9a0d637e8",
  "created" : 1629809644498404498,
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
  "message_id" : "71b20d76-2d22-4ee7-81b4-a72f35e37ce2",
  "message_type" : "stop.scan",
  "group_id" : "71b20d76-2d22-4ee7-81b4-a72f35e37ce2",
  "created" : 1629809644498429253,
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
  "message_id" : "03b84263-197f-450b-bafb-c90ea0f8e955",
  "message_type" : "get.scan",
  "group_id" : "03b84263-197f-450b-bafb-c90ea0f8e955",
  "created" : 1629809644498451193,
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
  "message_id" : "be36dabb-e91d-4ece-82c3-808c393ea56a",
  "message_type" : "modify.scan",
  "group_id" : "be36dabb-e91d-4ece-82c3-808c393ea56a",
  "created" : 1629809644498470535,
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
  "message_id" : "ce710e89-756a-405b-a524-629f9e2eb0b8",
  "message_type" : "created.scan",
  "group_id" : "ce710e89-756a-405b-a524-629f9e2eb0b8",
  "created" : 1629809644498497231,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8b67db3f-ca59-4777-ab9b-0bb46403f279",
  "message_type" : "modified.scan",
  "group_id" : "8b67db3f-ca59-4777-ab9b-0bb46403f279",
  "created" : 1629809644498512905,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4254d734-0ad8-47f8-b3f1-516199104342",
  "message_type" : "stopped.scan",
  "group_id" : "4254d734-0ad8-47f8-b3f1-516199104342",
  "created" : 1629809644498527530,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "98f7d13e-8180-49f6-8b2a-681e20adb96f",
  "message_type" : "status.scan",
  "group_id" : "98f7d13e-8180-49f6-8b2a-681e20adb96f",
  "created" : 1629809644498541399,
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
  "message_id" : "83b2942f-8038-4fd0-b526-41a1b05b3a93",
  "message_type" : "got.scan",
  "group_id" : "83b2942f-8038-4fd0-b526-41a1b05b3a93",
  "created" : 1629809644498561489,
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
  "message_id" : "54c0fa42-8158-4df3-b4f2-852ad5c0dd55",
  "message_type" : "result.scan",
  "group_id" : "54c0fa42-8158-4df3-b4f2-852ad5c0dd55",
  "created" : 1629809644498611851,
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
  "message_id" : "ac0a6e59-39a5-4fce-956c-fefcb26acd90",
  "message_type" : "failure.start.scan",
  "group_id" : "ac0a6e59-39a5-4fce-956c-fefcb26acd90",
  "created" : 1629809644498638161,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8ca6a1a7-d84b-455d-a694-b3850adf1788",
  "message_type" : "failure.stop.scan",
  "group_id" : "8ca6a1a7-d84b-455d-a694-b3850adf1788",
  "created" : 1629809644498657563,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4289336f-7897-4749-a38e-dbf2fca2e311",
  "message_type" : "failure.create.scan",
  "group_id" : "4289336f-7897-4749-a38e-dbf2fca2e311",
  "created" : 1629809644498673792,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fa7a8a11-d0f2-4c17-a588-6431c1ca22bb",
  "message_type" : "failure.modify.scan",
  "group_id" : "fa7a8a11-d0f2-4c17-a588-6431c1ca22bb",
  "created" : 1629809644498688834,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2082f5e9-96c1-42f9-aaba-d697e02bda81",
  "message_type" : "failure.get.scan",
  "group_id" : "2082f5e9-96c1-42f9-aaba-d697e02bda81",
  "created" : 1629809644498703472,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "432fd3e9-f64b-4ba7-ac1b-b0ad05a891d7",
  "message_type" : "failure.scan",
  "group_id" : "432fd3e9-f64b-4ba7-ac1b-b0ad05a891d7",
  "created" : 1629809644498717840,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "ff1749bf-2cc8-42de-843d-4c39da981208",
  "message_type" : "create.scan",
  "group_id" : "ff1749bf-2cc8-42de-843d-4c39da981208",
  "created" : 1629809644498732394
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "ed2ece68-733d-4d8a-baa8-e9e3c406e323",
  "message_type" : "start.scan",
  "group_id" : "ed2ece68-733d-4d8a-baa8-e9e3c406e323",
  "created" : 1629809644498753938,
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
  "message_id" : "bf281411-5376-4fdd-9f9b-a6d124649121",
  "message_type" : "stop.scan",
  "group_id" : "bf281411-5376-4fdd-9f9b-a6d124649121",
  "created" : 1629809644498773407,
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
  "message_id" : "fab791ee-c45b-4e91-bb3b-3b779ddc79c8",
  "message_type" : "get.scan",
  "group_id" : "fab791ee-c45b-4e91-bb3b-3b779ddc79c8",
  "created" : 1629809644498794325,
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
  "message_id" : "044284dc-afcc-4d0e-85cd-0c8b8a47e18e",
  "message_type" : "modify.scan",
  "group_id" : "044284dc-afcc-4d0e-85cd-0c8b8a47e18e",
  "created" : 1629809644498826271,
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
  "message_id" : "c67248a1-519d-44b5-a06b-4001448556f8",
  "message_type" : "created.scan",
  "group_id" : "c67248a1-519d-44b5-a06b-4001448556f8",
  "created" : 1629809644498852500,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "e9885547-3ee7-4c4f-9ade-f7092611bd58",
  "message_type" : "modified.scan",
  "group_id" : "e9885547-3ee7-4c4f-9ade-f7092611bd58",
  "created" : 1629809644498868302,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d572e7a8-6322-4d84-8b13-765e12775ec8",
  "message_type" : "stopped.scan",
  "group_id" : "d572e7a8-6322-4d84-8b13-765e12775ec8",
  "created" : 1629809644498882396,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "10f8ce2b-9113-48d4-85ee-a6caf6d98f30",
  "message_type" : "status.scan",
  "group_id" : "10f8ce2b-9113-48d4-85ee-a6caf6d98f30",
  "created" : 1629809644498896285,
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
  "message_id" : "a40d9356-3de0-4b8c-b9f8-b0e1d2b99e37",
  "message_type" : "got.scan",
  "group_id" : "a40d9356-3de0-4b8c-b9f8-b0e1d2b99e37",
  "created" : 1629809644498913081,
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
  "message_id" : "4c61b8c6-7419-4148-afaa-865df186515c",
  "message_type" : "result.scan",
  "group_id" : "4c61b8c6-7419-4148-afaa-865df186515c",
  "created" : 1629809644498939888,
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
  "message_id" : "08cd8ca4-846f-442f-bfbe-f8507580e153",
  "message_type" : "failure.start.scan",
  "group_id" : "08cd8ca4-846f-442f-bfbe-f8507580e153",
  "created" : 1629809644498965424,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "faf199a8-676e-4d92-81d2-8737d09622c2",
  "message_type" : "failure.stop.scan",
  "group_id" : "faf199a8-676e-4d92-81d2-8737d09622c2",
  "created" : 1629809644498981842,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0d4dd93a-02ee-4960-ba08-4bc622034aa3",
  "message_type" : "failure.create.scan",
  "group_id" : "0d4dd93a-02ee-4960-ba08-4bc622034aa3",
  "created" : 1629809644498997294,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "70fcf816-2295-4d9b-a339-09a14123e98f",
  "message_type" : "failure.modify.scan",
  "group_id" : "70fcf816-2295-4d9b-a339-09a14123e98f",
  "created" : 1629809644499012013,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a9be57d2-8ef6-469f-8695-6069d41f91f9",
  "message_type" : "failure.get.scan",
  "group_id" : "a9be57d2-8ef6-469f-8695-6069d41f91f9",
  "created" : 1629809644499026971,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8979fbbb-1099-481c-920c-11a49964f558",
  "message_type" : "failure.scan",
  "group_id" : "8979fbbb-1099-481c-920c-11a49964f558",
  "created" : 1629809644499049336,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "3213eeb3-2c31-4e6d-b503-1d4da7229c56",
  "message_type" : "create.scan",
  "group_id" : "3213eeb3-2c31-4e6d-b503-1d4da7229c56",
  "created" : 1629809644499065419
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "f5cdfaea-fe2f-44b5-a497-d32a92765b6c",
  "message_type" : "start.scan",
  "group_id" : "f5cdfaea-fe2f-44b5-a497-d32a92765b6c",
  "created" : 1629809644499089703,
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
  "message_id" : "2dbcf24c-9554-4119-974a-70b60e65ee88",
  "message_type" : "stop.scan",
  "group_id" : "2dbcf24c-9554-4119-974a-70b60e65ee88",
  "created" : 1629809644499112162,
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
  "message_id" : "4f665457-6bd7-4ff2-911f-9963cf34ac7a",
  "message_type" : "get.scan",
  "group_id" : "4f665457-6bd7-4ff2-911f-9963cf34ac7a",
  "created" : 1629809644499133131,
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
  "message_id" : "7ced8aa4-fbbd-4c34-bbc6-988427878c3a",
  "message_type" : "modify.scan",
  "group_id" : "7ced8aa4-fbbd-4c34-bbc6-988427878c3a",
  "created" : 1629809644499153819,
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
  "message_id" : "48c1739e-22ab-4ebd-ac26-f750465894d4",
  "message_type" : "created.scan",
  "group_id" : "48c1739e-22ab-4ebd-ac26-f750465894d4",
  "created" : 1629809644499177281,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9584979d-082a-4a55-b7b7-e0cf0034304d",
  "message_type" : "modified.scan",
  "group_id" : "9584979d-082a-4a55-b7b7-e0cf0034304d",
  "created" : 1629809644499192073,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0b2c54a1-34cc-49e6-b6ca-1997f92cc902",
  "message_type" : "stopped.scan",
  "group_id" : "0b2c54a1-34cc-49e6-b6ca-1997f92cc902",
  "created" : 1629809644499205788,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "e3e8bfec-ed5d-4a31-a842-63304bf498e1",
  "message_type" : "status.scan",
  "group_id" : "e3e8bfec-ed5d-4a31-a842-63304bf498e1",
  "created" : 1629809644499219305,
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
  "message_id" : "4f602fbd-d344-4037-8958-a55169248a80",
  "message_type" : "got.scan",
  "group_id" : "4f602fbd-d344-4037-8958-a55169248a80",
  "created" : 1629809644499237139,
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
  "message_id" : "a589d472-6fc5-4679-a979-e8b3b00cfdec",
  "message_type" : "result.scan",
  "group_id" : "a589d472-6fc5-4679-a979-e8b3b00cfdec",
  "created" : 1629809644499262102,
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
  "message_id" : "c2cce6c0-514d-44e5-bbf6-5036e6a802aa",
  "message_type" : "failure.start.scan",
  "group_id" : "c2cce6c0-514d-44e5-bbf6-5036e6a802aa",
  "created" : 1629809644499291360,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "99b8a240-57bf-4999-9260-b6392eb6b5e8",
  "message_type" : "failure.stop.scan",
  "group_id" : "99b8a240-57bf-4999-9260-b6392eb6b5e8",
  "created" : 1629809644499308079,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7aee2fad-38ec-43a5-a194-9b2bffb4ac27",
  "message_type" : "failure.create.scan",
  "group_id" : "7aee2fad-38ec-43a5-a194-9b2bffb4ac27",
  "created" : 1629809644499322694,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "bc4c1a5a-f7d6-4772-b627-76c756561fc9",
  "message_type" : "failure.modify.scan",
  "group_id" : "bc4c1a5a-f7d6-4772-b627-76c756561fc9",
  "created" : 1629809644499336959,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "02d77a5d-ebf3-4d4c-8da9-3f12c4b870da",
  "message_type" : "failure.get.scan",
  "group_id" : "02d77a5d-ebf3-4d4c-8da9-3f12c4b870da",
  "created" : 1629809644499369714,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "465de658-1219-44f7-8231-bf9b28584034",
  "message_type" : "failure.scan",
  "group_id" : "465de658-1219-44f7-8231-bf9b28584034",
  "created" : 1629809644499385487,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "dedb22fd-6206-454c-bd75-cdaae6c7863c",
  "message_type" : "create.scan",
  "group_id" : "dedb22fd-6206-454c-bd75-cdaae6c7863c",
  "created" : 1629809644499399921
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "3f0541bd-ab56-4734-a317-05e8f3d4e44c",
  "message_type" : "start.scan",
  "group_id" : "3f0541bd-ab56-4734-a317-05e8f3d4e44c",
  "created" : 1629809644499421384,
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
  "message_id" : "a3721099-0b6e-4796-8a13-3dd04ce87240",
  "message_type" : "stop.scan",
  "group_id" : "a3721099-0b6e-4796-8a13-3dd04ce87240",
  "created" : 1629809644499440585,
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
  "message_id" : "771b7ea5-d722-4a2c-9c4e-0c28381ffd92",
  "message_type" : "get.scan",
  "group_id" : "771b7ea5-d722-4a2c-9c4e-0c28381ffd92",
  "created" : 1629809644499461045,
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
  "message_id" : "7f766085-34a6-457f-a4f9-430f963bbd29",
  "message_type" : "modify.scan",
  "group_id" : "7f766085-34a6-457f-a4f9-430f963bbd29",
  "created" : 1629809644499482594,
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
  "message_id" : "34f0e681-b25c-480c-9881-cc31ef2338b4",
  "message_type" : "created.scan",
  "group_id" : "34f0e681-b25c-480c-9881-cc31ef2338b4",
  "created" : 1629809644499515197,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a8fdd541-ff16-47c5-a471-863724a4a301",
  "message_type" : "modified.scan",
  "group_id" : "a8fdd541-ff16-47c5-a471-863724a4a301",
  "created" : 1629809644499531208,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "579252fc-c52b-4d7d-b51e-16a7259df9a7",
  "message_type" : "stopped.scan",
  "group_id" : "579252fc-c52b-4d7d-b51e-16a7259df9a7",
  "created" : 1629809644499544892,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d0fba391-e159-4c77-ba1b-772f0e0b435d",
  "message_type" : "status.scan",
  "group_id" : "d0fba391-e159-4c77-ba1b-772f0e0b435d",
  "created" : 1629809644499560584,
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
  "message_id" : "bca71927-e77c-4462-9f0f-20bc20a7e3d5",
  "message_type" : "got.scan",
  "group_id" : "bca71927-e77c-4462-9f0f-20bc20a7e3d5",
  "created" : 1629809644499576837,
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
  "message_id" : "a68e24fc-9215-4ce3-844a-15f766fc1459",
  "message_type" : "result.scan",
  "group_id" : "a68e24fc-9215-4ce3-844a-15f766fc1459",
  "created" : 1629809644499601944,
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
  "message_id" : "6f2d166e-3efc-4812-a23d-cf355ea256b1",
  "message_type" : "failure.start.scan",
  "group_id" : "6f2d166e-3efc-4812-a23d-cf355ea256b1",
  "created" : 1629809644499627290,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8f1de0d6-3c4d-455e-acea-baf7d1a1b2ba",
  "message_type" : "failure.stop.scan",
  "group_id" : "8f1de0d6-3c4d-455e-acea-baf7d1a1b2ba",
  "created" : 1629809644499643288,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "ac556ffc-cdc3-4f75-92c4-5d38be563b5b",
  "message_type" : "failure.create.scan",
  "group_id" : "ac556ffc-cdc3-4f75-92c4-5d38be563b5b",
  "created" : 1629809644499657784,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "469d31cf-96f6-4b3d-857b-a99c67a351fc",
  "message_type" : "failure.modify.scan",
  "group_id" : "469d31cf-96f6-4b3d-857b-a99c67a351fc",
  "created" : 1629809644499672211,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c4a79728-c1ea-43cb-9a14-4e8c2f3fa830",
  "message_type" : "failure.get.scan",
  "group_id" : "c4a79728-c1ea-43cb-9a14-4e8c2f3fa830",
  "created" : 1629809644499686371,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d42c8b33-d309-4dfb-a34e-dbb7fe1c997d",
  "message_type" : "failure.scan",
  "group_id" : "d42c8b33-d309-4dfb-a34e-dbb7fe1c997d",
  "created" : 1629809644499700638,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "6153f852-21f3-447b-990b-871eb39435ea",
  "message_type" : "create.scan",
  "group_id" : "6153f852-21f3-447b-990b-871eb39435ea",
  "created" : 1629809644499721794
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "2542bf82-2ec8-4a39-b08a-94d893be3f70",
  "message_type" : "start.scan",
  "group_id" : "2542bf82-2ec8-4a39-b08a-94d893be3f70",
  "created" : 1629809644499744470,
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
  "message_id" : "271d5848-ef4c-4b4d-a664-822f85282570",
  "message_type" : "stop.scan",
  "group_id" : "271d5848-ef4c-4b4d-a664-822f85282570",
  "created" : 1629809644499764961,
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
  "message_id" : "d8ea5c3c-d94f-4ae3-b6d3-f73368fd628f",
  "message_type" : "get.scan",
  "group_id" : "d8ea5c3c-d94f-4ae3-b6d3-f73368fd628f",
  "created" : 1629809644499783688,
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
  "message_id" : "fc187163-5d6f-462d-a4dc-4f6407c4213c",
  "message_type" : "modify.scan",
  "group_id" : "fc187163-5d6f-462d-a4dc-4f6407c4213c",
  "created" : 1629809644499802072,
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
  "message_id" : "55fc7708-366f-4921-9771-2a20fecd8fce",
  "message_type" : "created.scan",
  "group_id" : "55fc7708-366f-4921-9771-2a20fecd8fce",
  "created" : 1629809644499823200,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b683e2e9-12a8-4a88-9463-18c6fd03aaa9",
  "message_type" : "modified.scan",
  "group_id" : "b683e2e9-12a8-4a88-9463-18c6fd03aaa9",
  "created" : 1629809644499837359,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0e4bddd3-4732-4195-b10a-216a415724c7",
  "message_type" : "stopped.scan",
  "group_id" : "0e4bddd3-4732-4195-b10a-216a415724c7",
  "created" : 1629809644499853155,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "498a5b67-7013-4ad4-aaa8-877f594d8ebb",
  "message_type" : "status.scan",
  "group_id" : "498a5b67-7013-4ad4-aaa8-877f594d8ebb",
  "created" : 1629809644499866938,
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
  "message_id" : "04a7d947-d8c3-44dc-b22c-acc7181e3e53",
  "message_type" : "got.scan",
  "group_id" : "04a7d947-d8c3-44dc-b22c-acc7181e3e53",
  "created" : 1629809644499882865,
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
  "message_id" : "c5441a4c-9ccc-49dc-93d2-da5376ba5fe4",
  "message_type" : "result.scan",
  "group_id" : "c5441a4c-9ccc-49dc-93d2-da5376ba5fe4",
  "created" : 1629809644499907420,
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
  "message_id" : "25d000d2-cc2f-42f8-b446-c6f0baf33112",
  "message_type" : "failure.start.scan",
  "group_id" : "25d000d2-cc2f-42f8-b446-c6f0baf33112",
  "created" : 1629809644499943621,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b515777e-c78b-4e73-bd46-3985b9b7d410",
  "message_type" : "failure.stop.scan",
  "group_id" : "b515777e-c78b-4e73-bd46-3985b9b7d410",
  "created" : 1629809644499960961,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7bb9dd56-028f-439f-9b0a-3ef99fafaf59",
  "message_type" : "failure.create.scan",
  "group_id" : "7bb9dd56-028f-439f-9b0a-3ef99fafaf59",
  "created" : 1629809644499975388,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2ab82f6c-c3c1-4626-87c7-01a3d099c847",
  "message_type" : "failure.modify.scan",
  "group_id" : "2ab82f6c-c3c1-4626-87c7-01a3d099c847",
  "created" : 1629809644499990193,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c8116b69-5fb0-41d7-84f5-e26fce174a3a",
  "message_type" : "failure.get.scan",
  "group_id" : "c8116b69-5fb0-41d7-84f5-e26fce174a3a",
  "created" : 1629809644500006847,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a49004fb-cad9-4654-af2b-ba801a229951",
  "message_type" : "failure.scan",
  "group_id" : "a49004fb-cad9-4654-af2b-ba801a229951",
  "created" : 1629809644500021632,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "cbf76ef5-9795-4b12-8af8-8c32b9dc8b58",
  "message_type" : "create.scan",
  "group_id" : "cbf76ef5-9795-4b12-8af8-8c32b9dc8b58",
  "created" : 1629809644500036410
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "1f115864-88a5-4ff2-871c-c68b3862b3c6",
  "message_type" : "start.scan",
  "group_id" : "1f115864-88a5-4ff2-871c-c68b3862b3c6",
  "created" : 1629809644500055359,
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
  "message_id" : "6163ecca-881d-4f65-8cce-59b5dc77434a",
  "message_type" : "stop.scan",
  "group_id" : "6163ecca-881d-4f65-8cce-59b5dc77434a",
  "created" : 1629809644500077551,
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
  "message_id" : "dc535812-248f-4601-950f-d384a1776144",
  "message_type" : "get.scan",
  "group_id" : "dc535812-248f-4601-950f-d384a1776144",
  "created" : 1629809644500099809,
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
  "message_id" : "c026d0b3-a475-4e3c-94b6-e62f465ad63f",
  "message_type" : "modify.scan",
  "group_id" : "c026d0b3-a475-4e3c-94b6-e62f465ad63f",
  "created" : 1629809644500121907,
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
  "message_id" : "66938c07-b49a-4dd3-90ca-34a613265d04",
  "message_type" : "created.scan",
  "group_id" : "66938c07-b49a-4dd3-90ca-34a613265d04",
  "created" : 1629809644500145521,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "558c0995-15c8-4668-9b3c-5bf9b416359e",
  "message_type" : "modified.scan",
  "group_id" : "558c0995-15c8-4668-9b3c-5bf9b416359e",
  "created" : 1629809644500160458,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "96d2102b-6103-4502-a334-6a7c472c13ec",
  "message_type" : "stopped.scan",
  "group_id" : "96d2102b-6103-4502-a334-6a7c472c13ec",
  "created" : 1629809644500181617,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8ad687b6-839a-4b6d-bde4-157c6d6466e2",
  "message_type" : "status.scan",
  "group_id" : "8ad687b6-839a-4b6d-bde4-157c6d6466e2",
  "created" : 1629809644500196505,
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
  "message_id" : "0f7eba76-082c-49ab-802b-eb7a68f1c148",
  "message_type" : "got.scan",
  "group_id" : "0f7eba76-082c-49ab-802b-eb7a68f1c148",
  "created" : 1629809644500215077,
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
  "message_id" : "1fb2d92c-510b-4892-9560-51877963dfee",
  "message_type" : "result.scan",
  "group_id" : "1fb2d92c-510b-4892-9560-51877963dfee",
  "created" : 1629809644500240656,
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
  "message_id" : "bd1837e0-6cc0-4dc9-ae85-3ef3df9b98b0",
  "message_type" : "failure.start.scan",
  "group_id" : "bd1837e0-6cc0-4dc9-ae85-3ef3df9b98b0",
  "created" : 1629809644500262056,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "08d5b1d1-e2d5-486b-874a-23693958bc1f",
  "message_type" : "failure.stop.scan",
  "group_id" : "08d5b1d1-e2d5-486b-874a-23693958bc1f",
  "created" : 1629809644500277829,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "005ff923-a357-4b18-ad85-e67b78d21aab",
  "message_type" : "failure.create.scan",
  "group_id" : "005ff923-a357-4b18-ad85-e67b78d21aab",
  "created" : 1629809644500292449,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "168975d0-0eb6-4e55-ad3f-a4b129b895a0",
  "message_type" : "failure.modify.scan",
  "group_id" : "168975d0-0eb6-4e55-ad3f-a4b129b895a0",
  "created" : 1629809644500307028,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7f925a57-9e1c-4642-bdad-63b38daa3592",
  "message_type" : "failure.get.scan",
  "group_id" : "7f925a57-9e1c-4642-bdad-63b38daa3592",
  "created" : 1629809644500321461,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0db48c52-6bae-4ea3-bed8-f4e9026e38c7",
  "message_type" : "failure.scan",
  "group_id" : "0db48c52-6bae-4ea3-bed8-f4e9026e38c7",
  "created" : 1629809644500335579,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "0d2e24da-22a8-4f4b-a8b1-66864eb6aa59",
  "message_type" : "create.scan",
  "group_id" : "0d2e24da-22a8-4f4b-a8b1-66864eb6aa59",
  "created" : 1629809644500352714
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "1882215f-a824-412b-bfa7-0e0038850dcd",
  "message_type" : "start.scan",
  "group_id" : "1882215f-a824-412b-bfa7-0e0038850dcd",
  "created" : 1629809644500380914,
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
  "message_id" : "d5cda474-102d-41fb-a357-8572e62c6b9a",
  "message_type" : "stop.scan",
  "group_id" : "d5cda474-102d-41fb-a357-8572e62c6b9a",
  "created" : 1629809644500404430,
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
  "message_id" : "2cf4f145-be35-4b02-8866-728e9cbc74e1",
  "message_type" : "get.scan",
  "group_id" : "2cf4f145-be35-4b02-8866-728e9cbc74e1",
  "created" : 1629809644500424934,
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
  "message_id" : "5a8e3fd0-a6bd-45dc-a973-f71c54a9cc4d",
  "message_type" : "modify.scan",
  "group_id" : "5a8e3fd0-a6bd-45dc-a973-f71c54a9cc4d",
  "created" : 1629809644500445369,
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
  "message_id" : "063c38fa-00f1-4653-8843-99d44e9e4652",
  "message_type" : "created.scan",
  "group_id" : "063c38fa-00f1-4653-8843-99d44e9e4652",
  "created" : 1629809644500467418,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "84ff0dd0-100e-49bd-99bd-88e2d75adb8b",
  "message_type" : "modified.scan",
  "group_id" : "84ff0dd0-100e-49bd-99bd-88e2d75adb8b",
  "created" : 1629809644500481682,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8a5898b9-b558-48bf-a477-71d00af527eb",
  "message_type" : "stopped.scan",
  "group_id" : "8a5898b9-b558-48bf-a477-71d00af527eb",
  "created" : 1629809644500495175,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "05aa45d7-a3d5-430e-9568-2c31d0ced36b",
  "message_type" : "status.scan",
  "group_id" : "05aa45d7-a3d5-430e-9568-2c31d0ced36b",
  "created" : 1629809644500508698,
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
  "message_id" : "b56e3f6e-0d28-4a69-9ed3-9f24606d2376",
  "message_type" : "got.scan",
  "group_id" : "b56e3f6e-0d28-4a69-9ed3-9f24606d2376",
  "created" : 1629809644500526458,
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
  "message_id" : "73a8b5da-e889-40c2-8369-4e9c69d8041f",
  "message_type" : "result.scan",
  "group_id" : "73a8b5da-e889-40c2-8369-4e9c69d8041f",
  "created" : 1629809644500550981,
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
  "message_id" : "be14f17d-0c1e-462b-a159-d8b84f7f2aca",
  "message_type" : "failure.start.scan",
  "group_id" : "be14f17d-0c1e-462b-a159-d8b84f7f2aca",
  "created" : 1629809644500571675,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5eed3b0d-269e-4774-9dd9-a16bf464b2c2",
  "message_type" : "failure.stop.scan",
  "group_id" : "5eed3b0d-269e-4774-9dd9-a16bf464b2c2",
  "created" : 1629809644500593751,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9a4a7179-4041-472b-bd42-146154031705",
  "message_type" : "failure.create.scan",
  "group_id" : "9a4a7179-4041-472b-bd42-146154031705",
  "created" : 1629809644500609582,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6ea9ab01-a3df-412f-99f3-985d798dc773",
  "message_type" : "failure.modify.scan",
  "group_id" : "6ea9ab01-a3df-412f-99f3-985d798dc773",
  "created" : 1629809644500623846,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "704c5929-7c5d-404d-a41a-224d56bd8e69",
  "message_type" : "failure.get.scan",
  "group_id" : "704c5929-7c5d-404d-a41a-224d56bd8e69",
  "created" : 1629809644500638277,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "44d78b55-a221-4481-a5e0-bcf5d1d7fa8e",
  "message_type" : "failure.scan",
  "group_id" : "44d78b55-a221-4481-a5e0-bcf5d1d7fa8e",
  "created" : 1629809644500652438,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "81654790-a0b2-4937-9436-1b6e18b9be0d",
  "message_type" : "create.scan",
  "group_id" : "81654790-a0b2-4937-9436-1b6e18b9be0d",
  "created" : 1629809644500666889
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "880b236d-a487-4fac-b5b6-43b80a69c1da",
  "message_type" : "start.scan",
  "group_id" : "880b236d-a487-4fac-b5b6-43b80a69c1da",
  "created" : 1629809644500688188,
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
  "message_id" : "fcc6b28f-790c-49e0-9349-a889132fffb6",
  "message_type" : "stop.scan",
  "group_id" : "fcc6b28f-790c-49e0-9349-a889132fffb6",
  "created" : 1629809644500710691,
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
  "message_id" : "1d68df38-4900-4fc7-8d57-1b15f7db8454",
  "message_type" : "get.scan",
  "group_id" : "1d68df38-4900-4fc7-8d57-1b15f7db8454",
  "created" : 1629809644500729488,
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
  "message_id" : "32fd3ae2-1289-4ade-b821-6595b3243811",
  "message_type" : "modify.scan",
  "group_id" : "32fd3ae2-1289-4ade-b821-6595b3243811",
  "created" : 1629809644500750337,
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
  "message_id" : "c672bf3b-c3ce-47a6-ab55-f135daccf1e5",
  "message_type" : "created.scan",
  "group_id" : "c672bf3b-c3ce-47a6-ab55-f135daccf1e5",
  "created" : 1629809644500772853,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "ac272d19-733b-416a-be52-136252e52b26",
  "message_type" : "modified.scan",
  "group_id" : "ac272d19-733b-416a-be52-136252e52b26",
  "created" : 1629809644500787407,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f6888440-8ee2-4ac9-8c38-348a0e4c84ae",
  "message_type" : "stopped.scan",
  "group_id" : "f6888440-8ee2-4ac9-8c38-348a0e4c84ae",
  "created" : 1629809644500801143,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "86b1e6b2-0960-418f-a519-cef09926ad81",
  "message_type" : "status.scan",
  "group_id" : "86b1e6b2-0960-418f-a519-cef09926ad81",
  "created" : 1629809644500821225,
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
  "message_id" : "7e56929f-9785-4967-bf95-a8e175f65d57",
  "message_type" : "got.scan",
  "group_id" : "7e56929f-9785-4967-bf95-a8e175f65d57",
  "created" : 1629809644500838493,
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
  "message_id" : "5d8294b2-cd9d-479f-97de-a305783a6d8b",
  "message_type" : "result.scan",
  "group_id" : "5d8294b2-cd9d-479f-97de-a305783a6d8b",
  "created" : 1629809644500862828,
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
  "message_id" : "09e2c1fc-1e94-47f4-9023-5c09a55361f8",
  "message_type" : "failure.start.scan",
  "group_id" : "09e2c1fc-1e94-47f4-9023-5c09a55361f8",
  "created" : 1629809644500885867,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b54dddba-95b3-411c-ae40-08b8d9c16886",
  "message_type" : "failure.stop.scan",
  "group_id" : "b54dddba-95b3-411c-ae40-08b8d9c16886",
  "created" : 1629809644500929791,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "df391bc9-fc59-4c23-b39b-9bdf55b5eae8",
  "message_type" : "failure.create.scan",
  "group_id" : "df391bc9-fc59-4c23-b39b-9bdf55b5eae8",
  "created" : 1629809644500944162,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "55eda065-c131-4a0c-8b2f-5f7a6b4adc56",
  "message_type" : "failure.modify.scan",
  "group_id" : "55eda065-c131-4a0c-8b2f-5f7a6b4adc56",
  "created" : 1629809644500958492,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "80a2df21-2965-4641-a07d-780416c564ee",
  "message_type" : "failure.get.scan",
  "group_id" : "80a2df21-2965-4641-a07d-780416c564ee",
  "created" : 1629809644500972842,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "11d0922a-b887-490f-aad5-a195b25ded06",
  "message_type" : "failure.scan",
  "group_id" : "11d0922a-b887-490f-aad5-a195b25ded06",
  "created" : 1629809644500987157,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "d8f41a54-df24-4c5e-a0ef-b406b5d0b4fd",
  "message_type" : "create.scan",
  "group_id" : "d8f41a54-df24-4c5e-a0ef-b406b5d0b4fd",
  "created" : 1629809644501001576
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "a7bc2d01-d914-4876-ab5a-23d24e155ea4",
  "message_type" : "start.scan",
  "group_id" : "a7bc2d01-d914-4876-ab5a-23d24e155ea4",
  "created" : 1629809644501023406,
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
  "message_id" : "afd28b70-a72a-43de-ab8a-c9589426274e",
  "message_type" : "stop.scan",
  "group_id" : "afd28b70-a72a-43de-ab8a-c9589426274e",
  "created" : 1629809644501051000,
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
  "message_id" : "0022f363-df0d-4557-9a33-d40df8174050",
  "message_type" : "get.scan",
  "group_id" : "0022f363-df0d-4557-9a33-d40df8174050",
  "created" : 1629809644501074442,
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
  "message_id" : "c69e7192-20dd-4038-9c82-26afd09a693f",
  "message_type" : "modify.scan",
  "group_id" : "c69e7192-20dd-4038-9c82-26afd09a693f",
  "created" : 1629809644501100102,
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
  "message_id" : "c4029137-d4de-47d6-9b48-549e33639ec4",
  "message_type" : "created.scan",
  "group_id" : "c4029137-d4de-47d6-9b48-549e33639ec4",
  "created" : 1629809644501123027,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fb443098-8c04-469b-b3dd-75562bb1a8c9",
  "message_type" : "modified.scan",
  "group_id" : "fb443098-8c04-469b-b3dd-75562bb1a8c9",
  "created" : 1629809644501137920,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "d7fbaf93-a286-424e-b182-b00b83f9fcc2",
  "message_type" : "stopped.scan",
  "group_id" : "d7fbaf93-a286-424e-b182-b00b83f9fcc2",
  "created" : 1629809644501151758,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "353636f8-7f2b-450f-843d-378f3b9b4a6b",
  "message_type" : "status.scan",
  "group_id" : "353636f8-7f2b-450f-843d-378f3b9b4a6b",
  "created" : 1629809644501165316,
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
  "message_id" : "d1e43a2a-122b-45ff-9150-e0812671ef4a",
  "message_type" : "got.scan",
  "group_id" : "d1e43a2a-122b-45ff-9150-e0812671ef4a",
  "created" : 1629809644501181268,
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
  "message_id" : "cafa9a74-8ac9-4c73-88d2-37e6f031a3c5",
  "message_type" : "result.scan",
  "group_id" : "cafa9a74-8ac9-4c73-88d2-37e6f031a3c5",
  "created" : 1629809644501205690,
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
  "message_id" : "f07c7852-a452-4109-a22b-f2c4939fbd6e",
  "message_type" : "failure.start.scan",
  "group_id" : "f07c7852-a452-4109-a22b-f2c4939fbd6e",
  "created" : 1629809644501228618,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "bb39b61d-8df4-44b5-966a-2f620056809d",
  "message_type" : "failure.stop.scan",
  "group_id" : "bb39b61d-8df4-44b5-966a-2f620056809d",
  "created" : 1629809644501244036,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a408d050-d891-4340-bf77-6a0dcbe72e1c",
  "message_type" : "failure.create.scan",
  "group_id" : "a408d050-d891-4340-bf77-6a0dcbe72e1c",
  "created" : 1629809644501258461,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fef64058-d288-4d23-a245-608f86010917",
  "message_type" : "failure.modify.scan",
  "group_id" : "fef64058-d288-4d23-a245-608f86010917",
  "created" : 1629809644501279871,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "01bd9216-4d0f-4be1-a37c-f36cf295f702",
  "message_type" : "failure.get.scan",
  "group_id" : "01bd9216-4d0f-4be1-a37c-f36cf295f702",
  "created" : 1629809644501295481,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9b7d4ac7-8889-4304-9560-b165cfea8cf4",
  "message_type" : "failure.scan",
  "group_id" : "9b7d4ac7-8889-4304-9560-b165cfea8cf4",
  "created" : 1629809644501309685,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "88888e96-b0f4-45e5-81e2-8b66bbb2db2f",
  "message_type" : "create.scan",
  "group_id" : "88888e96-b0f4-45e5-81e2-8b66bbb2db2f",
  "created" : 1629809644501323718
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "9d1962ac-49cb-43be-9e4f-92549139a5d6",
  "message_type" : "start.scan",
  "group_id" : "9d1962ac-49cb-43be-9e4f-92549139a5d6",
  "created" : 1629809644501344969,
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
  "message_id" : "b094330a-ce91-4fa4-b2ba-505d22b8f0ee",
  "message_type" : "stop.scan",
  "group_id" : "b094330a-ce91-4fa4-b2ba-505d22b8f0ee",
  "created" : 1629809644501366929,
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
  "message_id" : "6a81cdea-186f-4057-85fa-617529c58b2f",
  "message_type" : "get.scan",
  "group_id" : "6a81cdea-186f-4057-85fa-617529c58b2f",
  "created" : 1629809644501388213,
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
  "message_id" : "c9cadea2-04c9-4419-9ef7-8a98e39b1f7f",
  "message_type" : "modify.scan",
  "group_id" : "c9cadea2-04c9-4419-9ef7-8a98e39b1f7f",
  "created" : 1629809644501409397,
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
  "message_id" : "86d5e2df-b6b4-4cfd-b7ff-f4deeec42d7a",
  "message_type" : "created.scan",
  "group_id" : "86d5e2df-b6b4-4cfd-b7ff-f4deeec42d7a",
  "created" : 1629809644501434708,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "33292697-58eb-4939-8894-e29e7f2cc6e9",
  "message_type" : "modified.scan",
  "group_id" : "33292697-58eb-4939-8894-e29e7f2cc6e9",
  "created" : 1629809644501449190,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2a362a47-4063-457b-8e9a-d2f54e25bdd9",
  "message_type" : "stopped.scan",
  "group_id" : "2a362a47-4063-457b-8e9a-d2f54e25bdd9",
  "created" : 1629809644501462701,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "194d855d-60af-46d0-8045-4d036cd6cf09",
  "message_type" : "status.scan",
  "group_id" : "194d855d-60af-46d0-8045-4d036cd6cf09",
  "created" : 1629809644501476088,
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
  "message_id" : "a08309cd-8846-4894-9625-9e2431797dff",
  "message_type" : "got.scan",
  "group_id" : "a08309cd-8846-4894-9625-9e2431797dff",
  "created" : 1629809644501498682,
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
  "message_id" : "72df0bcb-379b-433b-b6bb-1968e12c9541",
  "message_type" : "result.scan",
  "group_id" : "72df0bcb-379b-433b-b6bb-1968e12c9541",
  "created" : 1629809644501525083,
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
  "message_id" : "c83df9ff-eb93-4790-bc07-c07a3d16f2b9",
  "message_type" : "failure.start.scan",
  "group_id" : "c83df9ff-eb93-4790-bc07-c07a3d16f2b9",
  "created" : 1629809644501548002,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "5bf47730-20c5-4879-a9cf-4bc5de4fa7d9",
  "message_type" : "failure.stop.scan",
  "group_id" : "5bf47730-20c5-4879-a9cf-4bc5de4fa7d9",
  "created" : 1629809644501563808,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a0c8467b-332a-464e-8ab3-99e4779f7d03",
  "message_type" : "failure.create.scan",
  "group_id" : "a0c8467b-332a-464e-8ab3-99e4779f7d03",
  "created" : 1629809644501578111,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3a23338a-2d0d-4337-877d-5b6589dd22ed",
  "message_type" : "failure.modify.scan",
  "group_id" : "3a23338a-2d0d-4337-877d-5b6589dd22ed",
  "created" : 1629809644501592726,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "1999957c-7a01-4bb8-8a11-449ac612b322",
  "message_type" : "failure.get.scan",
  "group_id" : "1999957c-7a01-4bb8-8a11-449ac612b322",
  "created" : 1629809644501607014,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "94382f63-cfef-4f46-906f-b4755bb27623",
  "message_type" : "failure.scan",
  "group_id" : "94382f63-cfef-4f46-906f-b4755bb27623",
  "created" : 1629809644501621342,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "2cde376e-747f-4c19-8689-6ff7f6c0e390",
  "message_type" : "create.scan",
  "group_id" : "2cde376e-747f-4c19-8689-6ff7f6c0e390",
  "created" : 1629809644501635610
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "7859cb0d-855c-4881-a201-8f16fce31356",
  "message_type" : "start.scan",
  "group_id" : "7859cb0d-855c-4881-a201-8f16fce31356",
  "created" : 1629809644501656867,
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
  "message_id" : "f83392c2-d873-4933-aff7-ed819692e01e",
  "message_type" : "stop.scan",
  "group_id" : "f83392c2-d873-4933-aff7-ed819692e01e",
  "created" : 1629809644501678607,
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
  "message_id" : "66016202-002a-4533-8ced-5cae93bab911",
  "message_type" : "get.scan",
  "group_id" : "66016202-002a-4533-8ced-5cae93bab911",
  "created" : 1629809644501700260,
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
  "message_id" : "b34b99ba-5a4c-40ad-b473-990b62769710",
  "message_type" : "modify.scan",
  "group_id" : "b34b99ba-5a4c-40ad-b473-990b62769710",
  "created" : 1629809644501728505,
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
  "message_id" : "648097ae-c09b-41e6-83dd-a9b35fff94a3",
  "message_type" : "created.scan",
  "group_id" : "648097ae-c09b-41e6-83dd-a9b35fff94a3",
  "created" : 1629809644501753010,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2ce13fba-605b-472f-8437-e866bc9be6be",
  "message_type" : "modified.scan",
  "group_id" : "2ce13fba-605b-472f-8437-e866bc9be6be",
  "created" : 1629809644501767461,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "e0a04909-a659-495b-a805-2ee215ee04c8",
  "message_type" : "stopped.scan",
  "group_id" : "e0a04909-a659-495b-a805-2ee215ee04c8",
  "created" : 1629809644501781342,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9fe02693-dbad-4e09-bb3f-59ad42fd972e",
  "message_type" : "status.scan",
  "group_id" : "9fe02693-dbad-4e09-bb3f-59ad42fd972e",
  "created" : 1629809644501797148,
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
  "message_id" : "7fc90752-02d9-45a6-8491-1cb1dd3f61de",
  "message_type" : "got.scan",
  "group_id" : "7fc90752-02d9-45a6-8491-1cb1dd3f61de",
  "created" : 1629809644501813375,
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
  "message_id" : "16743517-0e7d-4a60-ab43-1b6a00d3033c",
  "message_type" : "result.scan",
  "group_id" : "16743517-0e7d-4a60-ab43-1b6a00d3033c",
  "created" : 1629809644501837708,
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
  "message_id" : "7cde2a2c-1427-43ff-acbb-edfc3c4ae917",
  "message_type" : "failure.start.scan",
  "group_id" : "7cde2a2c-1427-43ff-acbb-edfc3c4ae917",
  "created" : 1629809644501858645,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "23ae73dc-06ee-4906-acea-0f4646694f47",
  "message_type" : "failure.stop.scan",
  "group_id" : "23ae73dc-06ee-4906-acea-0f4646694f47",
  "created" : 1629809644501874210,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "57c9c292-6e6f-4452-bc7f-25d9f497add3",
  "message_type" : "failure.create.scan",
  "group_id" : "57c9c292-6e6f-4452-bc7f-25d9f497add3",
  "created" : 1629809644501889041,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2a57c06f-8c0c-47bd-808b-4dbcd73c4656",
  "message_type" : "failure.modify.scan",
  "group_id" : "2a57c06f-8c0c-47bd-808b-4dbcd73c4656",
  "created" : 1629809644501903248,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6903fd30-5276-4f8e-aca6-3942e8912ac9",
  "message_type" : "failure.get.scan",
  "group_id" : "6903fd30-5276-4f8e-aca6-3942e8912ac9",
  "created" : 1629809644501926389,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f739125e-e6fc-4151-b109-aeaa7973d790",
  "message_type" : "failure.scan",
  "group_id" : "f739125e-e6fc-4151-b109-aeaa7973d790",
  "created" : 1629809644501942424,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "b44f8b41-5530-4cb1-b2b2-f0ad674f130e",
  "message_type" : "create.scan",
  "group_id" : "b44f8b41-5530-4cb1-b2b2-f0ad674f130e",
  "created" : 1629809644501957063
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "6c1b707c-d022-4402-a774-10e991b81b1a",
  "message_type" : "start.scan",
  "group_id" : "6c1b707c-d022-4402-a774-10e991b81b1a",
  "created" : 1629809644501978417,
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
  "message_id" : "2ff58648-cf9c-4d0c-b1a6-786db5d6769a",
  "message_type" : "stop.scan",
  "group_id" : "2ff58648-cf9c-4d0c-b1a6-786db5d6769a",
  "created" : 1629809644501998839,
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
  "message_id" : "f9ac0196-09e5-4d08-9264-fccec6e6931f",
  "message_type" : "get.scan",
  "group_id" : "f9ac0196-09e5-4d08-9264-fccec6e6931f",
  "created" : 1629809644502020000,
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
  "message_id" : "2c65c2cb-0512-43f7-b986-00e9ccfd4755",
  "message_type" : "modify.scan",
  "group_id" : "2c65c2cb-0512-43f7-b986-00e9ccfd4755",
  "created" : 1629809644502041320,
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
  "message_id" : "206572df-abf2-48b7-a85d-f7f88a042af7",
  "message_type" : "created.scan",
  "group_id" : "206572df-abf2-48b7-a85d-f7f88a042af7",
  "created" : 1629809644502064027,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9a276b16-b746-46ee-a280-b6240a43e38e",
  "message_type" : "modified.scan",
  "group_id" : "9a276b16-b746-46ee-a280-b6240a43e38e",
  "created" : 1629809644502078468,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "de9d4c1c-22f5-425f-95ea-14d52961c694",
  "message_type" : "stopped.scan",
  "group_id" : "de9d4c1c-22f5-425f-95ea-14d52961c694",
  "created" : 1629809644502091850,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "1bf75ef3-0d27-46aa-9377-fb2ee69ba646",
  "message_type" : "status.scan",
  "group_id" : "1bf75ef3-0d27-46aa-9377-fb2ee69ba646",
  "created" : 1629809644502108750,
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
  "message_id" : "1cbabab3-da2e-4397-9cca-070da012a69e",
  "message_type" : "got.scan",
  "group_id" : "1cbabab3-da2e-4397-9cca-070da012a69e",
  "created" : 1629809644502124921,
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
  "message_id" : "1176d1bd-afae-4b4f-94ff-520e135ba8c7",
  "message_type" : "result.scan",
  "group_id" : "1176d1bd-afae-4b4f-94ff-520e135ba8c7",
  "created" : 1629809644502157238,
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
  "message_id" : "cf1221f4-676d-4252-afbc-7fcc301c32d8",
  "message_type" : "failure.start.scan",
  "group_id" : "cf1221f4-676d-4252-afbc-7fcc301c32d8",
  "created" : 1629809644502180065,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "0a1acb6d-91cc-454e-a5bb-2c373ae07bfb",
  "message_type" : "failure.stop.scan",
  "group_id" : "0a1acb6d-91cc-454e-a5bb-2c373ae07bfb",
  "created" : 1629809644502196075,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "ab2c64e5-ea8b-48c2-b91d-1d710787fbb8",
  "message_type" : "failure.create.scan",
  "group_id" : "ab2c64e5-ea8b-48c2-b91d-1d710787fbb8",
  "created" : 1629809644502211009,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f7db9e08-f411-42ac-911b-3e2491daf91c",
  "message_type" : "failure.modify.scan",
  "group_id" : "f7db9e08-f411-42ac-911b-3e2491daf91c",
  "created" : 1629809644502225868,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9d1833ec-1a00-4d91-90a5-bd060bde3507",
  "message_type" : "failure.get.scan",
  "group_id" : "9d1833ec-1a00-4d91-90a5-bd060bde3507",
  "created" : 1629809644502242397,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c93cd0c9-d14f-482b-8930-d07a4ed2db72",
  "message_type" : "failure.scan",
  "group_id" : "c93cd0c9-d14f-482b-8930-d07a4ed2db72",
  "created" : 1629809644502256775,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "668e6140-5786-4c3f-b4b5-991e9bc0fd42",
  "message_type" : "create.scan",
  "group_id" : "668e6140-5786-4c3f-b4b5-991e9bc0fd42",
  "created" : 1629809644502270879
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "dff3f83a-9f4a-4a96-b907-6eff4bb35dae",
  "message_type" : "start.scan",
  "group_id" : "dff3f83a-9f4a-4a96-b907-6eff4bb35dae",
  "created" : 1629809644502292737,
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
  "message_id" : "2a5f75fd-b355-44e6-9058-b362e4bd860f",
  "message_type" : "stop.scan",
  "group_id" : "2a5f75fd-b355-44e6-9058-b362e4bd860f",
  "created" : 1629809644502313975,
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
  "message_id" : "bedab55f-b956-43dd-944a-863f338600f3",
  "message_type" : "get.scan",
  "group_id" : "bedab55f-b956-43dd-944a-863f338600f3",
  "created" : 1629809644502335443,
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
  "message_id" : "07d192c0-4b88-4dd4-b5d6-8fbe09a548cf",
  "message_type" : "modify.scan",
  "group_id" : "07d192c0-4b88-4dd4-b5d6-8fbe09a548cf",
  "created" : 1629809644502356306,
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
  "message_id" : "94708e33-155f-4059-9e66-ab2b0d050e14",
  "message_type" : "created.scan",
  "group_id" : "94708e33-155f-4059-9e66-ab2b0d050e14",
  "created" : 1629809644502386628,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a5696838-564f-438d-ad77-5b7928dbb837",
  "message_type" : "modified.scan",
  "group_id" : "a5696838-564f-438d-ad77-5b7928dbb837",
  "created" : 1629809644502402135,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "04cd032c-ec83-401f-9ea9-61953da41d40",
  "message_type" : "stopped.scan",
  "group_id" : "04cd032c-ec83-401f-9ea9-61953da41d40",
  "created" : 1629809644502418073,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fde2c49d-8627-4542-af67-fce9a87539c0",
  "message_type" : "status.scan",
  "group_id" : "fde2c49d-8627-4542-af67-fce9a87539c0",
  "created" : 1629809644502432319,
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
  "message_id" : "b503bec2-ac60-4798-944d-dffb737b606b",
  "message_type" : "got.scan",
  "group_id" : "b503bec2-ac60-4798-944d-dffb737b606b",
  "created" : 1629809644502447957,
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
  "message_id" : "6ee04b4c-fdcb-446b-9e8e-61a26f0fde59",
  "message_type" : "result.scan",
  "group_id" : "6ee04b4c-fdcb-446b-9e8e-61a26f0fde59",
  "created" : 1629809644502472874,
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
  "message_id" : "d57da8ce-b225-46cd-99da-0bc16f47b5ec",
  "message_type" : "failure.start.scan",
  "group_id" : "d57da8ce-b225-46cd-99da-0bc16f47b5ec",
  "created" : 1629809644502497862,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "f97d928c-26f3-4590-bb34-278e74615137",
  "message_type" : "failure.stop.scan",
  "group_id" : "f97d928c-26f3-4590-bb34-278e74615137",
  "created" : 1629809644502513929,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "e12a6639-ba21-4680-8058-cb6e2a0742a1",
  "message_type" : "failure.create.scan",
  "group_id" : "e12a6639-ba21-4680-8058-cb6e2a0742a1",
  "created" : 1629809644502529014,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "95307d52-7107-4394-9d78-3db04617acb2",
  "message_type" : "failure.modify.scan",
  "group_id" : "95307d52-7107-4394-9d78-3db04617acb2",
  "created" : 1629809644502543137,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "794fa4f8-443d-400c-882a-3ab5daa3ac95",
  "message_type" : "failure.get.scan",
  "group_id" : "794fa4f8-443d-400c-882a-3ab5daa3ac95",
  "created" : 1629809644502557372,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "fc5a9487-756f-476b-a2e5-1ef7ec032f11",
  "message_type" : "failure.scan",
  "group_id" : "fc5a9487-756f-476b-a2e5-1ef7ec032f11",
  "created" : 1629809644502578432,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "01ad3e26-07a6-43a9-8b8d-b7f67d7325d8",
  "message_type" : "create.scan",
  "group_id" : "01ad3e26-07a6-43a9-8b8d-b7f67d7325d8",
  "created" : 1629809644502618145
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "16dc5c34-ad3b-4023-9f28-e68bc329986c",
  "message_type" : "start.scan",
  "group_id" : "16dc5c34-ad3b-4023-9f28-e68bc329986c",
  "created" : 1629809644502646254,
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
  "message_id" : "3670f472-09f1-4756-aeb6-46d4cee5f9cc",
  "message_type" : "stop.scan",
  "group_id" : "3670f472-09f1-4756-aeb6-46d4cee5f9cc",
  "created" : 1629809644502675640,
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
  "message_id" : "225e2600-397c-45a0-8f0a-a2b47ea812e7",
  "message_type" : "get.scan",
  "group_id" : "225e2600-397c-45a0-8f0a-a2b47ea812e7",
  "created" : 1629809644502698715,
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
  "message_id" : "3c7ead45-9c18-4dd6-a17a-be6c078a3d3d",
  "message_type" : "modify.scan",
  "group_id" : "3c7ead45-9c18-4dd6-a17a-be6c078a3d3d",
  "created" : 1629809644502721657,
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
  "message_id" : "f71ece42-d4c1-4a7f-8a22-5f0c93885efa",
  "message_type" : "created.scan",
  "group_id" : "f71ece42-d4c1-4a7f-8a22-5f0c93885efa",
  "created" : 1629809644502749610,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a009e7c2-727c-4377-9671-dca1f3edb0c1",
  "message_type" : "modified.scan",
  "group_id" : "a009e7c2-727c-4377-9671-dca1f3edb0c1",
  "created" : 1629809644502768951,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "243ff464-864b-4c1a-8a72-defddcd399d7",
  "message_type" : "stopped.scan",
  "group_id" : "243ff464-864b-4c1a-8a72-defddcd399d7",
  "created" : 1629809644502782871,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a53d8e28-3cce-49c1-b452-4bfe7ba7440b",
  "message_type" : "status.scan",
  "group_id" : "a53d8e28-3cce-49c1-b452-4bfe7ba7440b",
  "created" : 1629809644502796840,
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
  "message_id" : "0ff99c13-5bf6-441c-9e5b-3fc7656dce2a",
  "message_type" : "got.scan",
  "group_id" : "0ff99c13-5bf6-441c-9e5b-3fc7656dce2a",
  "created" : 1629809644502812880,
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
  "message_id" : "cc575717-7995-4ca5-bda4-dd0ee9723f56",
  "message_type" : "result.scan",
  "group_id" : "cc575717-7995-4ca5-bda4-dd0ee9723f56",
  "created" : 1629809644502838671,
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
  "message_id" : "564014c8-d762-4f78-a3f4-67cc34f99530",
  "message_type" : "failure.start.scan",
  "group_id" : "564014c8-d762-4f78-a3f4-67cc34f99530",
  "created" : 1629809644502870145,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "36703f64-5aac-4442-bd5c-fa45e6dcd021",
  "message_type" : "failure.stop.scan",
  "group_id" : "36703f64-5aac-4442-bd5c-fa45e6dcd021",
  "created" : 1629809644502887480,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "40dd5cda-5824-4fa7-9f86-3cca326622a2",
  "message_type" : "failure.create.scan",
  "group_id" : "40dd5cda-5824-4fa7-9f86-3cca326622a2",
  "created" : 1629809644502902343,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "412f9bbc-69e2-4a99-b433-8c4f1744ef5a",
  "message_type" : "failure.modify.scan",
  "group_id" : "412f9bbc-69e2-4a99-b433-8c4f1744ef5a",
  "created" : 1629809644502916902,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "14a02a39-e440-448a-9efc-4b53eee67c6c",
  "message_type" : "failure.get.scan",
  "group_id" : "14a02a39-e440-448a-9efc-4b53eee67c6c",
  "created" : 1629809644502931362,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "78b271cb-199a-46bb-8235-043720204697",
  "message_type" : "failure.scan",
  "group_id" : "78b271cb-199a-46bb-8235-043720204697",
  "created" : 1629809644502945598,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "d85e1b7e-a1f0-4b3d-8788-3f4b3b0c4b57",
  "message_type" : "create.scan",
  "group_id" : "d85e1b7e-a1f0-4b3d-8788-3f4b3b0c4b57",
  "created" : 1629809644502959735
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "587d61eb-acce-428b-9813-a43b00eb2d10",
  "message_type" : "start.scan",
  "group_id" : "587d61eb-acce-428b-9813-a43b00eb2d10",
  "created" : 1629809644502983446,
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
  "message_id" : "0a28d12a-6070-47a0-acba-1cb9fe4d7c86",
  "message_type" : "stop.scan",
  "group_id" : "0a28d12a-6070-47a0-acba-1cb9fe4d7c86",
  "created" : 1629809644503009520,
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
  "message_id" : "53acdc0a-2b91-45b3-9009-b6939b8504c2",
  "message_type" : "get.scan",
  "group_id" : "53acdc0a-2b91-45b3-9009-b6939b8504c2",
  "created" : 1629809644503035048,
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
  "message_id" : "b0ace258-99ba-47fd-b0a6-989d2d5880e3",
  "message_type" : "modify.scan",
  "group_id" : "b0ace258-99ba-47fd-b0a6-989d2d5880e3",
  "created" : 1629809644503062343,
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
  "message_id" : "07841271-9ec6-4ec9-bbc8-e650c509c03e",
  "message_type" : "created.scan",
  "group_id" : "07841271-9ec6-4ec9-bbc8-e650c509c03e",
  "created" : 1629809644503093501,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "560ce131-eea6-4667-a3d4-98086e46a232",
  "message_type" : "modified.scan",
  "group_id" : "560ce131-eea6-4667-a3d4-98086e46a232",
  "created" : 1629809644503118195,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "63b2bf32-f25a-4c07-a33d-5f4696384d8b",
  "message_type" : "stopped.scan",
  "group_id" : "63b2bf32-f25a-4c07-a33d-5f4696384d8b",
  "created" : 1629809644503137312,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "ff8a3621-55e5-4cbc-9da4-bc0573322761",
  "message_type" : "status.scan",
  "group_id" : "ff8a3621-55e5-4cbc-9da4-bc0573322761",
  "created" : 1629809644503151427,
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
  "message_id" : "f1d0485e-f65e-427e-ad39-2814cf93a7eb",
  "message_type" : "got.scan",
  "group_id" : "f1d0485e-f65e-427e-ad39-2814cf93a7eb",
  "created" : 1629809644503167442,
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
  "message_id" : "8d1efcc5-6c6a-4934-9142-65d545d871e0",
  "message_type" : "result.scan",
  "group_id" : "8d1efcc5-6c6a-4934-9142-65d545d871e0",
  "created" : 1629809644503196627,
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
  "message_id" : "0f79948f-e38e-429b-a7c9-dfc238c0db3b",
  "message_type" : "failure.start.scan",
  "group_id" : "0f79948f-e38e-429b-a7c9-dfc238c0db3b",
  "created" : 1629809644503223234,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3e541bf0-df13-40ba-a34f-640fbada018b",
  "message_type" : "failure.stop.scan",
  "group_id" : "3e541bf0-df13-40ba-a34f-640fbada018b",
  "created" : 1629809644503239430,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "7897c267-8a7e-4094-9fc9-0d5d6b36f1af",
  "message_type" : "failure.create.scan",
  "group_id" : "7897c267-8a7e-4094-9fc9-0d5d6b36f1af",
  "created" : 1629809644503254280,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "bf0e5984-ed22-4fea-aee8-d145f499f828",
  "message_type" : "failure.modify.scan",
  "group_id" : "bf0e5984-ed22-4fea-aee8-d145f499f828",
  "created" : 1629809644503268764,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "3e7d6f9f-cb04-4154-bf84-925dec7ce2b2",
  "message_type" : "failure.get.scan",
  "group_id" : "3e7d6f9f-cb04-4154-bf84-925dec7ce2b2",
  "created" : 1629809644503283267,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "db7d8c7a-7584-4710-a8df-b7fa450f47e4",
  "message_type" : "failure.scan",
  "group_id" : "db7d8c7a-7584-4710-a8df-b7fa450f47e4",
  "created" : 1629809644503297377,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "d605677a-9098-41e9-8d9a-5859e419ad4d",
  "message_type" : "create.scan",
  "group_id" : "d605677a-9098-41e9-8d9a-5859e419ad4d",
  "created" : 1629809644503320040
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "6f9305ab-e3ff-4e2c-9da7-27db3597fea2",
  "message_type" : "start.scan",
  "group_id" : "6f9305ab-e3ff-4e2c-9da7-27db3597fea2",
  "created" : 1629809644503354314,
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
  "message_id" : "a24dfc07-ac78-4045-8171-6bfad20344ac",
  "message_type" : "stop.scan",
  "group_id" : "a24dfc07-ac78-4045-8171-6bfad20344ac",
  "created" : 1629809644503380927,
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
  "message_id" : "efe2eabc-c8c5-49bd-a5a7-70df4c5fde19",
  "message_type" : "get.scan",
  "group_id" : "efe2eabc-c8c5-49bd-a5a7-70df4c5fde19",
  "created" : 1629809644503400155,
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
  "message_id" : "b199e0a6-a413-4c27-8db9-999192ab3482",
  "message_type" : "modify.scan",
  "group_id" : "b199e0a6-a413-4c27-8db9-999192ab3482",
  "created" : 1629809644503420167,
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
  "message_id" : "f2dd7af4-698e-4bbe-8c97-784245b39a06",
  "message_type" : "created.scan",
  "group_id" : "f2dd7af4-698e-4bbe-8c97-784245b39a06",
  "created" : 1629809644503448330,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a394fb35-1dab-4ea6-bb02-3bdd30e81393",
  "message_type" : "modified.scan",
  "group_id" : "a394fb35-1dab-4ea6-bb02-3bdd30e81393",
  "created" : 1629809644503464379,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "a6600f8a-f5ba-42dd-a2f9-75cf146eb00b",
  "message_type" : "stopped.scan",
  "group_id" : "a6600f8a-f5ba-42dd-a2f9-75cf146eb00b",
  "created" : 1629809644503478150,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "e33b1094-2c2a-4357-bb81-b3093a3d68de",
  "message_type" : "status.scan",
  "group_id" : "e33b1094-2c2a-4357-bb81-b3093a3d68de",
  "created" : 1629809644503492548,
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
  "message_id" : "bd056c2b-6701-4ca6-a119-395332a72a39",
  "message_type" : "got.scan",
  "group_id" : "bd056c2b-6701-4ca6-a119-395332a72a39",
  "created" : 1629809644503513493,
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
  "message_id" : "b0017325-3610-4712-bc3a-bb1c1ff6f5f6",
  "message_type" : "result.scan",
  "group_id" : "b0017325-3610-4712-bc3a-bb1c1ff6f5f6",
  "created" : 1629809644503539973,
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
  "message_id" : "e6e8f4df-eee6-41a8-bbc9-3a0574070650",
  "message_type" : "failure.start.scan",
  "group_id" : "e6e8f4df-eee6-41a8-bbc9-3a0574070650",
  "created" : 1629809644503567838,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8c173c9e-25f0-4488-9849-35e87747ffd1",
  "message_type" : "failure.stop.scan",
  "group_id" : "8c173c9e-25f0-4488-9849-35e87747ffd1",
  "created" : 1629809644503594960,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "4757c140-f519-48e6-a1e0-70b443ffffcc",
  "message_type" : "failure.create.scan",
  "group_id" : "4757c140-f519-48e6-a1e0-70b443ffffcc",
  "created" : 1629809644503611339,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "c8c4528c-ffa4-45e2-8f89-437156bd90bf",
  "message_type" : "failure.modify.scan",
  "group_id" : "c8c4528c-ffa4-45e2-8f89-437156bd90bf",
  "created" : 1629809644503625711,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "45790a1a-b717-42fc-a621-e9c7d7b3d8cc",
  "message_type" : "failure.get.scan",
  "group_id" : "45790a1a-b717-42fc-a621-e9c7d7b3d8cc",
  "created" : 1629809644503640070,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "31dae011-1988-4e62-9640-009f8243ba53",
  "message_type" : "failure.scan",
  "group_id" : "31dae011-1988-4e62-9640-009f8243ba53",
  "created" : 1629809644503654066,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## create/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "3a812853-c0ce-4621-b5a9-e43df361c4ec",
  "message_type" : "create.scan",
  "group_id" : "3a812853-c0ce-4621-b5a9-e43df361c4ec",
  "created" : 1629809644503668311
}
```
Responses:

- [created](#createdscan)
- [failure.create](#failurecreatescan)
## start/scan

Topic: eulabeia/scan/cmd/director
```
{
  "message_id" : "83e4139a-bd9e-4ac3-adcc-f4704628625b",
  "message_type" : "start.scan",
  "group_id" : "83e4139a-bd9e-4ac3-adcc-f4704628625b",
  "created" : 1629809644503693243,
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
  "message_id" : "25afc58e-0566-41e1-97ed-6ea6aad7c7e6",
  "message_type" : "stop.scan",
  "group_id" : "25afc58e-0566-41e1-97ed-6ea6aad7c7e6",
  "created" : 1629809644503719956,
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
  "message_id" : "17163fb0-16d3-4483-b479-636ef94cbd31",
  "message_type" : "get.scan",
  "group_id" : "17163fb0-16d3-4483-b479-636ef94cbd31",
  "created" : 1629809644503743554,
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
  "message_id" : "c45a6783-ad2e-4387-95ef-64e71561c08c",
  "message_type" : "modify.scan",
  "group_id" : "c45a6783-ad2e-4387-95ef-64e71561c08c",
  "created" : 1629809644503772256,
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
  "message_id" : "a59603dd-89f0-46a2-b1e1-f5db4226c94e",
  "message_type" : "created.scan",
  "group_id" : "a59603dd-89f0-46a2-b1e1-f5db4226c94e",
  "created" : 1629809644503796757,
  "id" : "example.id.scan"
}
```
## modified/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "9734efaa-0ab5-4357-bb7f-648b91bcbd3e",
  "message_type" : "modified.scan",
  "group_id" : "9734efaa-0ab5-4357-bb7f-648b91bcbd3e",
  "created" : 1629809644503811293,
  "id" : "example.id.scan"
}
```
## stopped/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "52a40973-545c-419d-81a5-c599fc13f054",
  "message_type" : "stopped.scan",
  "group_id" : "52a40973-545c-419d-81a5-c599fc13f054",
  "created" : 1629809644503832625,
  "id" : "example.id.scan"
}
```
## status/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "1ed649f9-1507-4793-a44d-0375316354b5",
  "message_type" : "status.scan",
  "group_id" : "1ed649f9-1507-4793-a44d-0375316354b5",
  "created" : 1629809644503847832,
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
  "message_id" : "fd867047-cdcb-401b-9356-7aac5df3f23f",
  "message_type" : "got.scan",
  "group_id" : "fd867047-cdcb-401b-9356-7aac5df3f23f",
  "created" : 1629809644503864616,
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
  "message_id" : "14ac0b27-990d-4001-8cfe-95c253c4aa0d",
  "message_type" : "result.scan",
  "group_id" : "14ac0b27-990d-4001-8cfe-95c253c4aa0d",
  "created" : 1629809644503893142,
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
  "message_id" : "2aa3b4c9-87de-4ad3-81e0-de79724223ef",
  "message_type" : "failure.start.scan",
  "group_id" : "2aa3b4c9-87de-4ad3-81e0-de79724223ef",
  "created" : 1629809644503914993,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.stop/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "8fdee537-f7f1-44d5-942f-6d4345feb773",
  "message_type" : "failure.stop.scan",
  "group_id" : "8fdee537-f7f1-44d5-942f-6d4345feb773",
  "created" : 1629809644503930633,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.create/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "b62084da-84b7-4e95-ba96-2f8616035158",
  "message_type" : "failure.create.scan",
  "group_id" : "b62084da-84b7-4e95-ba96-2f8616035158",
  "created" : 1629809644503945153,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.modify/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "6dbb534f-39ea-4a08-a509-44471939704d",
  "message_type" : "failure.modify.scan",
  "group_id" : "6dbb534f-39ea-4a08-a509-44471939704d",
  "created" : 1629809644503959505,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure.get/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "2ca8f207-b2db-4533-b319-a2185c0f70af",
  "message_type" : "failure.get.scan",
  "group_id" : "2ca8f207-b2db-4533-b319-a2185c0f70af",
  "created" : 1629809644503973464,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
## failure/scan

Topic: eulabeia/scan/info
```
{
  "message_id" : "39c3fe0f-c0bb-42c3-874f-bdf92a42a127",
  "message_type" : "failure.scan",
  "group_id" : "39c3fe0f-c0bb-42c3-874f-bdf92a42a127",
  "created" : 1629809644503987567,
  "id" : "example.id.scan",
  "error" : "some error description"
}
```
# target

## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "44d2d5f5-bbd7-4329-98cb-0d71e17bb1f4",
  "message_type" : "create.target",
  "group_id" : "44d2d5f5-bbd7-4329-98cb-0d71e17bb1f4",
  "created" : 1629809644504001909
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "818cb679-3c24-4c77-9d8a-bb7bc66244a2",
  "message_type" : "get.target",
  "group_id" : "818cb679-3c24-4c77-9d8a-bb7bc66244a2",
  "created" : 1629809644504025644,
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
  "message_id" : "e75e148b-7c0a-4a5e-a49a-4c7a15a5fb2b",
  "message_type" : "modify.target",
  "group_id" : "e75e148b-7c0a-4a5e-a49a-4c7a15a5fb2b",
  "created" : 1629809644504061966,
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
  "message_id" : "d3914ddd-ce33-4aa0-a35c-260f3c15e968",
  "message_type" : "created.target",
  "group_id" : "d3914ddd-ce33-4aa0-a35c-260f3c15e968",
  "created" : 1629809644504101612,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "428c6774-0328-46d0-ac74-41255ec0137b",
  "message_type" : "modified.target",
  "group_id" : "428c6774-0328-46d0-ac74-41255ec0137b",
  "created" : 1629809644504118455,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "2cbeba99-bdc5-49cc-845e-af4bf9cd0e3c",
  "message_type" : "got.target",
  "group_id" : "2cbeba99-bdc5-49cc-845e-af4bf9cd0e3c",
  "created" : 1629809644504132754,
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
  "message_id" : "d15f54bf-fe03-4bd4-b0e5-81dc12dd1fc7",
  "message_type" : "failure.create.target",
  "group_id" : "d15f54bf-fe03-4bd4-b0e5-81dc12dd1fc7",
  "created" : 1629809644504156042,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "726e99db-ed73-429d-84df-4da4c60a55e5",
  "message_type" : "failure.modify.target",
  "group_id" : "726e99db-ed73-429d-84df-4da4c60a55e5",
  "created" : 1629809644504171392,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "2b9e6bec-deab-4c48-bb3b-f6735d36b887",
  "message_type" : "failure.get.target",
  "group_id" : "2b9e6bec-deab-4c48-bb3b-f6735d36b887",
  "created" : 1629809644504188923,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "0df38efb-2142-4083-84d5-a053d108e0d4",
  "message_type" : "failure.target",
  "group_id" : "0df38efb-2142-4083-84d5-a053d108e0d4",
  "created" : 1629809644504203865,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "dc6a2c79-bef4-4060-87f4-47ad56f1c17c",
  "message_type" : "create.target",
  "group_id" : "dc6a2c79-bef4-4060-87f4-47ad56f1c17c",
  "created" : 1629809644504218865
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "72b16db2-2029-4cd8-9b6b-a831bc113828",
  "message_type" : "get.target",
  "group_id" : "72b16db2-2029-4cd8-9b6b-a831bc113828",
  "created" : 1629809644504242637,
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
  "message_id" : "27586272-93aa-47d5-9849-65228bd1db72",
  "message_type" : "modify.target",
  "group_id" : "27586272-93aa-47d5-9849-65228bd1db72",
  "created" : 1629809644504269765,
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
  "message_id" : "6351f549-42f0-4bf2-bd59-3874a5ec5e78",
  "message_type" : "created.target",
  "group_id" : "6351f549-42f0-4bf2-bd59-3874a5ec5e78",
  "created" : 1629809644504313406,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "a882a6a1-93f7-4cc5-9979-0e48935a7c81",
  "message_type" : "modified.target",
  "group_id" : "a882a6a1-93f7-4cc5-9979-0e48935a7c81",
  "created" : 1629809644504331544,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "69c6663c-7a01-4f92-b7f1-02e1e50034d1",
  "message_type" : "got.target",
  "group_id" : "69c6663c-7a01-4f92-b7f1-02e1e50034d1",
  "created" : 1629809644504345441,
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
  "message_id" : "c36192df-1d3a-4427-917c-d4d4945e01ed",
  "message_type" : "failure.create.target",
  "group_id" : "c36192df-1d3a-4427-917c-d4d4945e01ed",
  "created" : 1629809644504375760,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "91b11064-54a2-4e7d-8567-df1708b10a7b",
  "message_type" : "failure.modify.target",
  "group_id" : "91b11064-54a2-4e7d-8567-df1708b10a7b",
  "created" : 1629809644504392558,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "208cb43b-de13-462a-930d-12ca1d7c7cef",
  "message_type" : "failure.get.target",
  "group_id" : "208cb43b-de13-462a-930d-12ca1d7c7cef",
  "created" : 1629809644504407208,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1ea6e8e2-a85e-4c85-b486-c196e1b058f8",
  "message_type" : "failure.target",
  "group_id" : "1ea6e8e2-a85e-4c85-b486-c196e1b058f8",
  "created" : 1629809644504421995,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "ccca803c-a6ca-4698-8e4f-a9850c045eed",
  "message_type" : "create.target",
  "group_id" : "ccca803c-a6ca-4698-8e4f-a9850c045eed",
  "created" : 1629809644504441702
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "1ecddb0d-5970-444a-b0b9-3b03cf6538be",
  "message_type" : "get.target",
  "group_id" : "1ecddb0d-5970-444a-b0b9-3b03cf6538be",
  "created" : 1629809644504466707,
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
  "message_id" : "e8c8ab23-f8d5-46c4-bf2b-4594d5599522",
  "message_type" : "modify.target",
  "group_id" : "e8c8ab23-f8d5-46c4-bf2b-4594d5599522",
  "created" : 1629809644504490624,
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
  "message_id" : "a669c94f-763b-4b01-83a3-d946b12271e4",
  "message_type" : "created.target",
  "group_id" : "a669c94f-763b-4b01-83a3-d946b12271e4",
  "created" : 1629809644504524651,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "9dfa7cf2-d836-43f0-b326-57706a559781",
  "message_type" : "modified.target",
  "group_id" : "9dfa7cf2-d836-43f0-b326-57706a559781",
  "created" : 1629809644504547775,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "522214f8-4b43-48c7-a141-eaf979220340",
  "message_type" : "got.target",
  "group_id" : "522214f8-4b43-48c7-a141-eaf979220340",
  "created" : 1629809644504563025,
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
  "message_id" : "51995a27-778d-4179-ac69-cc651016bdfe",
  "message_type" : "failure.create.target",
  "group_id" : "51995a27-778d-4179-ac69-cc651016bdfe",
  "created" : 1629809644504585363,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "760b6056-fc4d-49c5-a8a3-ac3313eb3dd3",
  "message_type" : "failure.modify.target",
  "group_id" : "760b6056-fc4d-49c5-a8a3-ac3313eb3dd3",
  "created" : 1629809644504600475,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "f943de23-e01a-4e99-9e31-7cf35c56b516",
  "message_type" : "failure.get.target",
  "group_id" : "f943de23-e01a-4e99-9e31-7cf35c56b516",
  "created" : 1629809644504614780,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "aaf91c51-78fd-4ea9-977e-5eb99e252622",
  "message_type" : "failure.target",
  "group_id" : "aaf91c51-78fd-4ea9-977e-5eb99e252622",
  "created" : 1629809644504629298,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "ccf9d7aa-ae18-4938-9a06-8a0e60e394d5",
  "message_type" : "create.target",
  "group_id" : "ccf9d7aa-ae18-4938-9a06-8a0e60e394d5",
  "created" : 1629809644504643735
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "1c83a735-0388-4b4e-851c-f85dd51e1f2d",
  "message_type" : "get.target",
  "group_id" : "1c83a735-0388-4b4e-851c-f85dd51e1f2d",
  "created" : 1629809644504665549,
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
  "message_id" : "7fb13f01-0130-4acb-8683-a1e8de2330bd",
  "message_type" : "modify.target",
  "group_id" : "7fb13f01-0130-4acb-8683-a1e8de2330bd",
  "created" : 1629809644504691053,
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
  "message_id" : "102b6f33-c5b0-43a4-80ab-f73fb0a49418",
  "message_type" : "created.target",
  "group_id" : "102b6f33-c5b0-43a4-80ab-f73fb0a49418",
  "created" : 1629809644504728752,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "3bf50d75-5cd1-45c3-b323-7d6b0a6cd291",
  "message_type" : "modified.target",
  "group_id" : "3bf50d75-5cd1-45c3-b323-7d6b0a6cd291",
  "created" : 1629809644504747501,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "51b5c3c7-91d3-4eef-a8f7-d6f64b650b88",
  "message_type" : "got.target",
  "group_id" : "51b5c3c7-91d3-4eef-a8f7-d6f64b650b88",
  "created" : 1629809644504762350,
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
  "message_id" : "258fd5dc-9fc0-488f-9f15-4df4e5e2c448",
  "message_type" : "failure.create.target",
  "group_id" : "258fd5dc-9fc0-488f-9f15-4df4e5e2c448",
  "created" : 1629809644504792755,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "6508b976-326f-431e-a684-6782553fd2a2",
  "message_type" : "failure.modify.target",
  "group_id" : "6508b976-326f-431e-a684-6782553fd2a2",
  "created" : 1629809644504812957,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "21d89e6b-fd51-457d-8b3f-02e6993c5179",
  "message_type" : "failure.get.target",
  "group_id" : "21d89e6b-fd51-457d-8b3f-02e6993c5179",
  "created" : 1629809644504827862,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "3a83780a-f331-4fb4-934b-a1c4ebf9116c",
  "message_type" : "failure.target",
  "group_id" : "3a83780a-f331-4fb4-934b-a1c4ebf9116c",
  "created" : 1629809644504844904,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "aa9d4686-27bc-4880-a706-b92ed01dfab5",
  "message_type" : "create.target",
  "group_id" : "aa9d4686-27bc-4880-a706-b92ed01dfab5",
  "created" : 1629809644504859857
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "61b5695b-88d1-4738-a6d4-fc8315b210d3",
  "message_type" : "get.target",
  "group_id" : "61b5695b-88d1-4738-a6d4-fc8315b210d3",
  "created" : 1629809644504884222,
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
  "message_id" : "89d2b6eb-5220-4099-9f5f-98c8539d729b",
  "message_type" : "modify.target",
  "group_id" : "89d2b6eb-5220-4099-9f5f-98c8539d729b",
  "created" : 1629809644504907954,
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
  "message_id" : "fc9a4dbb-bc82-483e-931f-30a2e809c199",
  "message_type" : "created.target",
  "group_id" : "fc9a4dbb-bc82-483e-931f-30a2e809c199",
  "created" : 1629809644504938671,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "eb9ec703-4ecd-4cb2-8833-8395cd4548b9",
  "message_type" : "modified.target",
  "group_id" : "eb9ec703-4ecd-4cb2-8833-8395cd4548b9",
  "created" : 1629809644504953732,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "75d58cee-62f7-43d3-861c-cd17043983a4",
  "message_type" : "got.target",
  "group_id" : "75d58cee-62f7-43d3-861c-cd17043983a4",
  "created" : 1629809644504967298,
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
  "message_id" : "6c5a0973-82e0-4e3f-88cd-2d39bd5a75df",
  "message_type" : "failure.create.target",
  "group_id" : "6c5a0973-82e0-4e3f-88cd-2d39bd5a75df",
  "created" : 1629809644504991673,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "ed7081c4-d7ce-45cc-a18b-7ca519fb3e34",
  "message_type" : "failure.modify.target",
  "group_id" : "ed7081c4-d7ce-45cc-a18b-7ca519fb3e34",
  "created" : 1629809644505014218,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "0796d3b9-7d8c-47d1-8519-a08d2b189a3d",
  "message_type" : "failure.get.target",
  "group_id" : "0796d3b9-7d8c-47d1-8519-a08d2b189a3d",
  "created" : 1629809644505030192,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "55087f04-c2be-4e97-850f-88c2e2912347",
  "message_type" : "failure.target",
  "group_id" : "55087f04-c2be-4e97-850f-88c2e2912347",
  "created" : 1629809644505044754,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "5545d83f-a94c-4244-b5f4-1bcd0a417464",
  "message_type" : "create.target",
  "group_id" : "5545d83f-a94c-4244-b5f4-1bcd0a417464",
  "created" : 1629809644505059208
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "846b0ba8-42a1-41cb-ad97-871533443fde",
  "message_type" : "get.target",
  "group_id" : "846b0ba8-42a1-41cb-ad97-871533443fde",
  "created" : 1629809644505082051,
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
  "message_id" : "61f9e7ba-47d5-4188-950d-e1db528f7cea",
  "message_type" : "modify.target",
  "group_id" : "61f9e7ba-47d5-4188-950d-e1db528f7cea",
  "created" : 1629809644505109982,
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
  "message_id" : "bacaf194-bfe8-42c2-bc58-d65f663a606a",
  "message_type" : "created.target",
  "group_id" : "bacaf194-bfe8-42c2-bc58-d65f663a606a",
  "created" : 1629809644505147859,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "d8f547ab-faa2-4358-81f7-fee7bfeb4d09",
  "message_type" : "modified.target",
  "group_id" : "d8f547ab-faa2-4358-81f7-fee7bfeb4d09",
  "created" : 1629809644505165112,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "6fd3f0bf-8763-4ef4-b27f-000736eb1894",
  "message_type" : "got.target",
  "group_id" : "6fd3f0bf-8763-4ef4-b27f-000736eb1894",
  "created" : 1629809644505179144,
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
  "message_id" : "0e29ee33-10a8-4771-b8a4-8b7a2be7e06e",
  "message_type" : "failure.create.target",
  "group_id" : "0e29ee33-10a8-4771-b8a4-8b7a2be7e06e",
  "created" : 1629809644505200605,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "901ddcb5-be58-4efa-becd-93f106aca490",
  "message_type" : "failure.modify.target",
  "group_id" : "901ddcb5-be58-4efa-becd-93f106aca490",
  "created" : 1629809644505216733,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "e77d1c59-e218-4aa5-8fe4-d31cc4f5e232",
  "message_type" : "failure.get.target",
  "group_id" : "e77d1c59-e218-4aa5-8fe4-d31cc4f5e232",
  "created" : 1629809644505232353,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "ee615f7c-57ed-4bb3-85df-cbf553f77cd9",
  "message_type" : "failure.target",
  "group_id" : "ee615f7c-57ed-4bb3-85df-cbf553f77cd9",
  "created" : 1629809644505254203,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "dc09fbfd-fadb-458b-be19-e92e4f0e1d61",
  "message_type" : "create.target",
  "group_id" : "dc09fbfd-fadb-458b-be19-e92e4f0e1d61",
  "created" : 1629809644505270321
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "1a7f36b8-5065-462f-a2e9-293662ca9b40",
  "message_type" : "get.target",
  "group_id" : "1a7f36b8-5065-462f-a2e9-293662ca9b40",
  "created" : 1629809644505294809,
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
  "message_id" : "d9da2734-1c41-4bce-baf7-f00f57e54175",
  "message_type" : "modify.target",
  "group_id" : "d9da2734-1c41-4bce-baf7-f00f57e54175",
  "created" : 1629809644505320184,
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
  "message_id" : "582d8c00-2936-414f-bcea-1ac51834e6af",
  "message_type" : "created.target",
  "group_id" : "582d8c00-2936-414f-bcea-1ac51834e6af",
  "created" : 1629809644505357809,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "a4785443-e343-40bb-a1eb-f758dc02ced0",
  "message_type" : "modified.target",
  "group_id" : "a4785443-e343-40bb-a1eb-f758dc02ced0",
  "created" : 1629809644505375097,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "678f1e75-b305-440d-a196-c8a2a83c118c",
  "message_type" : "got.target",
  "group_id" : "678f1e75-b305-440d-a196-c8a2a83c118c",
  "created" : 1629809644505388954,
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
  "message_id" : "1e567de3-ed77-4e08-ad22-517bfe765f05",
  "message_type" : "failure.create.target",
  "group_id" : "1e567de3-ed77-4e08-ad22-517bfe765f05",
  "created" : 1629809644505410665,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "d4cf60ba-149e-4247-b77a-c048ce748cec",
  "message_type" : "failure.modify.target",
  "group_id" : "d4cf60ba-149e-4247-b77a-c048ce748cec",
  "created" : 1629809644505428096,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "28d1265b-b7cc-4fd8-9e92-5f3a4934d9ae",
  "message_type" : "failure.get.target",
  "group_id" : "28d1265b-b7cc-4fd8-9e92-5f3a4934d9ae",
  "created" : 1629809644505443659,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "19595bdc-5881-405c-a41c-013645b0f066",
  "message_type" : "failure.target",
  "group_id" : "19595bdc-5881-405c-a41c-013645b0f066",
  "created" : 1629809644505468376,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "114f7e9a-4776-456f-9a11-ca9ea83ba29e",
  "message_type" : "create.target",
  "group_id" : "114f7e9a-4776-456f-9a11-ca9ea83ba29e",
  "created" : 1629809644505492977
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "4276e0e8-0249-4bbe-84e3-d45aea1db2e9",
  "message_type" : "get.target",
  "group_id" : "4276e0e8-0249-4bbe-84e3-d45aea1db2e9",
  "created" : 1629809644505519942,
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
  "message_id" : "4d0faa5c-1c66-4354-938d-ab6d60ab4848",
  "message_type" : "modify.target",
  "group_id" : "4d0faa5c-1c66-4354-938d-ab6d60ab4848",
  "created" : 1629809644505543606,
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
  "message_id" : "66966e27-61d8-4f7a-bc40-8e0fe3300d5a",
  "message_type" : "created.target",
  "group_id" : "66966e27-61d8-4f7a-bc40-8e0fe3300d5a",
  "created" : 1629809644505574468,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "5dcbf7a2-4c58-4c51-8927-423e5203e478",
  "message_type" : "modified.target",
  "group_id" : "5dcbf7a2-4c58-4c51-8927-423e5203e478",
  "created" : 1629809644505589554,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "d1460eb9-1b02-4f0b-843b-96f7d0509a6d",
  "message_type" : "got.target",
  "group_id" : "d1460eb9-1b02-4f0b-843b-96f7d0509a6d",
  "created" : 1629809644505603393,
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
  "message_id" : "a8ec113b-07dc-41d2-a7d3-4ceebdc8281a",
  "message_type" : "failure.create.target",
  "group_id" : "a8ec113b-07dc-41d2-a7d3-4ceebdc8281a",
  "created" : 1629809644505628248,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "d3f8b62f-b4ea-4e47-aa8f-b2f306b3e415",
  "message_type" : "failure.modify.target",
  "group_id" : "d3f8b62f-b4ea-4e47-aa8f-b2f306b3e415",
  "created" : 1629809644505643250,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "cc994ac4-23cb-4bba-8b52-9a707f9ac5c6",
  "message_type" : "failure.get.target",
  "group_id" : "cc994ac4-23cb-4bba-8b52-9a707f9ac5c6",
  "created" : 1629809644505657841,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "fee1de25-913d-415f-b3bf-bad20b1f07d5",
  "message_type" : "failure.target",
  "group_id" : "fee1de25-913d-415f-b3bf-bad20b1f07d5",
  "created" : 1629809644505672316,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "c57b350b-780d-467c-a664-3bb0c470c529",
  "message_type" : "create.target",
  "group_id" : "c57b350b-780d-467c-a664-3bb0c470c529",
  "created" : 1629809644505686970
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "ae7aada9-9cc6-423d-a636-2d7a53012432",
  "message_type" : "get.target",
  "group_id" : "ae7aada9-9cc6-423d-a636-2d7a53012432",
  "created" : 1629809644505708874,
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
  "message_id" : "cea5d631-f859-4cda-9aea-b306dec79d86",
  "message_type" : "modify.target",
  "group_id" : "cea5d631-f859-4cda-9aea-b306dec79d86",
  "created" : 1629809644505741817,
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
  "message_id" : "2caf277c-8a92-439f-bdfa-e032a76305c7",
  "message_type" : "created.target",
  "group_id" : "2caf277c-8a92-439f-bdfa-e032a76305c7",
  "created" : 1629809644505778396,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "7e1d8756-16e0-404d-a216-8569e377503c",
  "message_type" : "modified.target",
  "group_id" : "7e1d8756-16e0-404d-a216-8569e377503c",
  "created" : 1629809644505795239,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "5b297444-c7a6-4ee5-9e44-22e6646799f9",
  "message_type" : "got.target",
  "group_id" : "5b297444-c7a6-4ee5-9e44-22e6646799f9",
  "created" : 1629809644505812050,
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
  "message_id" : "d9da456a-dc8e-4277-a31a-217ca9e667ec",
  "message_type" : "failure.create.target",
  "group_id" : "d9da456a-dc8e-4277-a31a-217ca9e667ec",
  "created" : 1629809644505834101,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "dda75ae2-48c0-4905-a3e0-75c32b682898",
  "message_type" : "failure.modify.target",
  "group_id" : "dda75ae2-48c0-4905-a3e0-75c32b682898",
  "created" : 1629809644505849101,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "612ef55e-2182-4da3-b5de-9bf9040e0af3",
  "message_type" : "failure.get.target",
  "group_id" : "612ef55e-2182-4da3-b5de-9bf9040e0af3",
  "created" : 1629809644505863270,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "1073045b-588b-4deb-8125-12f3a2ac9142",
  "message_type" : "failure.target",
  "group_id" : "1073045b-588b-4deb-8125-12f3a2ac9142",
  "created" : 1629809644505877987,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## create/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "c743ee8f-0e5c-42a1-9365-5c9b2e96b5ca",
  "message_type" : "create.target",
  "group_id" : "c743ee8f-0e5c-42a1-9365-5c9b2e96b5ca",
  "created" : 1629809644505892371
}
```
Responses:

- [created](#createdtarget)
- [failure.create](#failurecreatetarget)
## get/target

Topic: eulabeia/target/cmd/director
```
{
  "message_id" : "6a3b5bf8-80e9-4c85-a6f7-8a09d706e8a6",
  "message_type" : "get.target",
  "group_id" : "6a3b5bf8-80e9-4c85-a6f7-8a09d706e8a6",
  "created" : 1629809644505916686,
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
  "message_id" : "9d504ec7-61f3-45e4-b956-c452c0488ebd",
  "message_type" : "modify.target",
  "group_id" : "9d504ec7-61f3-45e4-b956-c452c0488ebd",
  "created" : 1629809644505941765,
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
  "message_id" : "5862c62d-b6ce-4710-9098-34190a197de4",
  "message_type" : "created.target",
  "group_id" : "5862c62d-b6ce-4710-9098-34190a197de4",
  "created" : 1629809644505984748,
  "id" : "example.id.target"
}
```
## modified/target

Topic: eulabeia/target/info
```
{
  "message_id" : "91e5bf7a-3c32-44ce-8e61-d72585643e60",
  "message_type" : "modified.target",
  "group_id" : "91e5bf7a-3c32-44ce-8e61-d72585643e60",
  "created" : 1629809644506002450,
  "id" : "example.id.target"
}
```
## got/target

Topic: eulabeia/target/info
```
{
  "message_id" : "76b66382-b39f-43e1-b8e6-77c4850ef8d2",
  "message_type" : "got.target",
  "group_id" : "76b66382-b39f-43e1-b8e6-77c4850ef8d2",
  "created" : 1629809644506016286,
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
  "message_id" : "be24f0bd-9099-4fec-a7c2-718e7df428b5",
  "message_type" : "failure.create.target",
  "group_id" : "be24f0bd-9099-4fec-a7c2-718e7df428b5",
  "created" : 1629809644506038089,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.modify/target

Topic: eulabeia/target/info
```
{
  "message_id" : "12f421c9-5b3d-4746-96cc-afd3e7dff690",
  "message_type" : "failure.modify.target",
  "group_id" : "12f421c9-5b3d-4746-96cc-afd3e7dff690",
  "created" : 1629809644506052706,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure.get/target

Topic: eulabeia/target/info
```
{
  "message_id" : "60f2023a-9c78-4e99-8841-9a66e023c015",
  "message_type" : "failure.get.target",
  "group_id" : "60f2023a-9c78-4e99-8841-9a66e023c015",
  "created" : 1629809644506067109,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
## failure/target

Topic: eulabeia/target/info
```
{
  "message_id" : "e378953e-c9ba-43c4-9831-a0f32cef2bf5",
  "message_type" : "failure.target",
  "group_id" : "e378953e-c9ba-43c4-9831-a0f32cef2bf5",
  "created" : 1629809644506083956,
  "id" : "example.id.target",
  "error" : "some error description"
}
```
