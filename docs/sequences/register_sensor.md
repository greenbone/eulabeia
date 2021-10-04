# Register a new sensor

When a new sensor is spawned to be directly accessible to the client then it must register itself to the director.

This is done by sending a [modify.sensor](../message_examples.md#modifysensor) ideally when starting.
This modify must contain the identifier of the sensor.

A sensor must then subscribe to `scanner/cmd/+/sensor_identifier` where `+` should be set to the aggregates the sensor handles and `sensor_identifier` the actual identifier of the sensor.

When a sensor stops it should send a [delete.sensor](../message_examples.md#deletesensor) event to deregister itself.

A functioning sensor is expected to handle at least:

- [start.scan](../message_examples.md#startscan) - see [Step 2 of start.scan](./start_scan.md#start-scan-1)
