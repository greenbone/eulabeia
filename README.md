# eulabeia

Is management layer with the goal to make control of multiple scanner based on multiple machines easier and more predictable by introducing a small set of commands, orchestration based on a director and usage of a broker.

There are four [roles](docs/roles/roles-and-relationship.md) within eulabeia:

- a director - manages multiple sensors and is the main communication partner of a client
- sensor - manages (or is) a scanner
- scanner - a program to scan for vulnerabilities
- client - is the using party of eulabeia

The communication between each party is done via [normalized messages](docs/messaging.md) over a MQTTv5.

For a detailed overview about eulabeia please checkout [docs](docs/README.md).

![overview participants](./docs/roles/relationship.svg)

## Included projects

- director, the director implementation
- sensor, the sensor implementation for [openvas-scanner](https://github.com/greenbone/openvas-scanner/)
- libeulabeia, the c library to use eulabeia

## Requirements

The minimal requirements to use eulabeia are:
- go
- docker
- make

If you want to build and run eulabeia without container you need:

- a mqtt broker (e.g. mosquitto)
- openvas-scanner (currently middleware branch)
- redis

If you want to build libeulabeia you need:
- [gvm-libs](https://github.com/greenbone/gvm-libs)
- [libpaho](https://www.eclipse.org/paho/files/mqttdoc/MQTTClient/html/index.html)
- [cgreen](https://cgreen-devs.github.io/)


installed and configured.

## Installation

If you just want to test eulabeia without installing it on your machine you should use the provided docker images by running `make start-smoke-test` .

Please follow the [requirements](docs/requirements.md) instruction before installing eulabeia.

When all requirements are met then you can build the sensor and director binaries by executing:

```
make build
```

This will create the binaries:
- `bin/eulabeia-director`
- `bin/eulabeia-sensor`
- `bin/example-client`

Before you can use the director or sensor you need to [configure it based on your MQTT address and redis socket](docs/requirements.md#director-1).

The next step is to start the director (`./bin/eulabeia-director`) and sensor (`./bin/eulabeia-sensor`).

To test if the setup is correct you can execute the example client: `./bin/example-client`.

For a more detailed explanation about the requirements please refer to [requirements.md](docs/requirements.md).

## Usage

Since eulabeia is not indented to be used as a program but is a rather a service environment and meant to be used via a library or specified workflows by another service / program please refer to:

- [docs](./docs/README.md)
- [docs/messaging](./docs/messaging.md)
- [docs/examples](./docs/message_examples.md)
- [docs/error](./docs/error-handling.md)
- [docs/start-scan](./docs/sequences/start_scan.md)

## Contributing

Your contributions are highly appreciated.

<!---
After making yourself familiar with

- [coding-style](./docs/coding-style.md)
- [testing](./docs/testing.md)
- [commit-structure](./docs/commits.md)

please create a pull request.
--->
