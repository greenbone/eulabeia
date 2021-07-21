eulabeia
========

Is a project to control various greenbone sensor.

It is separated in

-	a director to channel various instruction to a sensor
-	a sensor to channel the instructions of a director to an actual scanner (e.g. openvas)

The communication between client (e.g. gvmd) and director as well as director to sensor is done via mqtt.

![overview participants](docs/pictures/overview_participants.svg?raw=true)

This project is in a very early state and is not ready for usage yet.
=====================================================================

make commands
-------------

| cmd                             | description                                                               |
|---------------------------------|---------------------------------------------------------------------------|
| `make build-container-broker`   | build broker image                                                        |
| `make build-container-director` | build director image                                                      |
| `make build-container-example`  | build example image                                                       |
| `make build-container-sensor`   | build sensor image                                                        |
| `make build-container`          | build broker, sensor, director and example-client image                   |
| `make build`                    | builds the director, sensor and example client binary                     |
| `make check`                    | runs static checks                                                        |
| `make format`                   | formats go code                                                           |
| `make prepare`                  | installs staticcheck                                                      |
| `make run-example-client`       | runs example client                                                       |
| `make smoke-test`               | runs the smoke tests (builds image, start smoke tests and stop container) |
| `make start-broker`             | starts a broker as a container                                            |
| `make start-container`          | starts the broker, director as well as sensor container                   |
| `make start-director`           | starts a director container                                               |
| `make start-sensor`             | starts sensor container                                                   |
| `make start-smoke-test`         | starts the smoke tests (broker, director, sensor and runs example client) |
| `make stop-broker`              | stops broker container                                                    |
| `make stop-container`           | stops running container                                                   |
| `make stop-director`            | stops director container                                                  |
| `make stop-sensor`              | stops a sensor container                                                  |
| `make test`                     | runs unittests                                                            |
| `make update`                   | updates dependencies                                                      |
| `make`                          | default commands on make without target are: format, check, test, build   |
