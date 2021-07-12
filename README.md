eulabeia
========

Is a project to control various greenbone sensor.

It is separated in

-	a director to channel various instruction to a sensor
-	a sensor to channel the instructions of a director to an actual scanner (e.g. openvas)

The communication between client (e.g. gvmd) and director as well as director to sensor is done via mqtt.

![overview participants](docs/pictures/overview_participants.svg?raw=true)

This project is in a very early state and is not ready for usage yet.
=======

make commands
-------------

| name               | example                   | description                                                               |
|--------------------|---------------------------|---------------------------------------------------------------------------|
| all                | `make`                    | default commands on make without target are: format, check, test, build   |
| build-container    | `make build-container`    | build broker, sensor, director and example-client image                   |
| build              | `make build`              | builds the director, sensor and example client binary                     |
| check              | `make check`              | runs static checks                                                        |
| format             | `make format`             | formats go code                                                           |
| prepare            | `make prepare`            | installs staticcheck                                                      |
| run-example-client | `make run-example-client` | runs example client                                                       |
| smoke-test         | `make smoke-test`         | runs the smoke tests (builds image, start smoke tests and stop container) |
| start-broker       | `make start-broker`       | starts a broker as a container                                            |
| start-container    | `make start-container`    | starts the broker, director as well as sensor container                   |
| start-director     | `make start-director`     | starts a director container                                               |
| start-sensor       | `make start-sensor`       | starts sensor container                                                   |
| start-smoke-test   | `make start-smoke-test`   | starts the smoke tests (broker, director, sensor and runs example client) |
| stop-broker        | `make stop-broker`        | stops broker container                                                    |
| stop-container     | `make stop-container`     | stops running container                                                   |
| stop-director      | `make stop-director`      | stops director container                                                  |
| stop-sensor        | `make stop-sensor`        | stops a sensor container                                                  |
| test               | `make test`               | runs unittests                                                            |
| update             | `make update`             | updates dependencies                                                      |
