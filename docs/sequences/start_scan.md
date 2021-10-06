# Start Scan

Before a scan can be started it should be created, this has the advantage that a scan can be executed as previously defined without having to resend it.

Although it is possible to create a temporary scan, that will be automatically deleted, when finished, it should not be seen as the default case.

The flow is:
1. client creates a target (either via create or directly via modify when it has directly all the data at hand)
1. client creates a scan with the target_id
1. client sends start.scan with the ID of the scan
1. director informs the corresponding sensor
1. sensor verify if it can start the scan
1. sensor starts the actual scan
1. while scanning result.scan and status.scan are send back to the client
1. sensor may inform the client about failures

Currently we only support OpenVAS which is not written as a daemon and therefore there is a differentiation between sensor and scanner which may not be the case on different scanner.

The difference between target and scan is that target contains of information which usually won't change that often and if changed it can affect different scan.

Although the scan contains the target data directly implementation wise they are separated. 


## Create Scan
<!---
render with: plantuml -tsvg start_scan_sequence.md
@startuml create_scan
skinparam monochrome reverse
autonumber
participant Client
participant "Eulabeia Director" as director
Client <-> director : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#createtarget create.target]], [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#createdtarget created.target]]
Client <-> director : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#modifytarget modify.target]], [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#modifiedtarget modified.target]]
Client <-> director : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#modifyscan modify.scan]], [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#modifiedscan modified.scan]]
@enduml

--->
![create a scan](./create_scan.svg)

### Messages

- [create.target](../message_examples.md#createtarget)
- [created.target](../message_examples.md#createdtarget)
- [modify.target](../message_examples.md#modifytarget)
- [modified.target](../message_examples.md#modifiedtarget)
- [modify.scan](../message_examples.md#modifyscan)
- [modified.scan](../message_examples.md#modifiedscan)

## Start Scan
<!---
render with: plantuml -tsvg start_scan_sequence.md
@startuml start_scan
skinparam monochrome reverse
autonumber
participant Client
participant "Eulabeia Director" as director
collections "Eulabeia OpenVAS Sensor" as sensor
collections "OpenVAS Scanner" as openvas
Client -> director : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#startscan start.scan]]
director -> sensor : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#startscan start.scan]]
loop until sufficient memory
    sensor -> sensor : check memory
end
sensor -> openvas: start scanner
openvas <-> director : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#getscan get.scan]], [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#gotscan got.scan]]
loop until openvas finishes scan
    openvas -> Client : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#resultscan result.scan]]
    openvas -> Client : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#statusscan status.scan]]
end
break on wrong openvas exit code
    sensor -> Client : [[https://github.com/greenbone/eulabeia/blob/main/docs/message_examples.md#failurestartscan failure.start.scan]]
end
@enduml

--->

![start a scan](./start_scan.svg)

### Messages

- [start.scan](../message_examples.md#startscan)
- [get.scan](../message_examples.md#getscan)
- [got.scan](../message_examples.md#gotscan)
- [result.scan](../message_examples.md#resultscan)
- [status.scan](../message_examples.md#statusscan)
- [failure.start.scan](../message_examples.md#failurestartscan)
