# eulabeia

Is a project to control various greenbone sensor.

It is separated in 
- a director to channel various instruction to a sensor
- a sensor to channel the instructions of a director to an actual scanner (e.g. openvas)

The communication between client (e.g. gvmd) and director as well as director to sensor is done via mqtt.

This project is in a very early state and is not ready for usage yet.

