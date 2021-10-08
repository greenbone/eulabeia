# Requirements

This document describes the requirements for installing and running eulabeia

Eulabeia can be installed in two methods:

- [docker](#Docker), describes how to setup docker and use the provided makefile
- [local](#Local), describes how to setup and install the requirements without the provided images

This document will use Debian 11 (bullseye) as a basis and is targeting developer machines with relatively lax security requirements. 

For production usage please harden docker, redis and mosquitto accordingly.

## Docker

To use eulabeia with docker you need to install `go`, `docker` as well as `gnu-make`. For easier docker usage you can add your user to the docker group:

```
sudo apt install docker.io make golang
sudo usermod -a -G docker $USER
```

To reload the newly assign group, login to your user account again

```
$ su - $USER
$ docker ps
```

To start Eulabeia based on the previously build images you can use `make start-container`.
This will start a useable sensor, director and broker.

To build images based on your changes you can use:

| cmd | image tag | description |
| -- | -- | -- |
| make build-container-broker | greenbone/eulabeia-broker:latest | builds broker  |
| make build-container-redis | greenbone/eulabeia-redis:latest | builds redis  |
| make build-container-director | greenbone/eulabeia-director:latest | builds the director based on local changes  |
| make build-container-sensor | greenbone/eulabeia-sensor:latest | builds the openvas-sensor based on local changes  |
| make build-container-example | greenbone/eulabeia-example-client:latest | builds example based on local changes |
| make build-container | | builds each container based on local changes |

If you want to verify the setup you can simply use `make run-example-client`.

If you have `mosquitto-clients` installed, you can validate that the local connection works. You can subscribe to any topic of the scanner context and monitore it:
```
$ mosquitto_sub -L 'mqtt://localhost:9138/scanner/#'
```

Afterwards you can publish an irregular message on a topic that is used and processed by the director:

```
$ mosquitto_pub -L 'mqtt://localhost:9138/scanner/scan/cmd/director' -m "hi"
```

### Broker

To run a broker you can execute `make start-broker`. This will run mosquitto and opens the local port `9138`.

### Sensor

To run a sensor you can execute `make start-sensor`. This will start the mosquitto broker and opens the local port `9138` as well as start a sensor containing a test `.nasl` file (oid: `1.3.6.1.4.1.25623.1.0.90022`) used by [openvas](https://github.com/greenbone/openvas-scanner).

To communicate with the sensor use MQTT as described previously.

### Director

To run a director you can execute `make start-director`. This will run the mosquitto broker and opens the local port `9138` as well as start a director.

To communicate with the director use mqtt as described previously.

## Local

### Broker (mosquitto)

To install and configure mosquitto:

```
sudo apt install mosquitto
sudo echo "allow_anonymous true" > /etc/mosquitto/conf.d/allow_anonymous.conf
sudo systemctl restart mosquitto
```

This allows anonymous connections into the running mqtt broker.

### Redis

Redis is used by [openvas](https://github.com/greenbone/openvas-scanner) as well as sensor.

To install and configure redis:

```
sudo apt install redis
sudo sed -i 's/\# unixsocket\([ p]\)/unixsocket\1/' /etc/redis/redis.conf
sudo systemctl restart redis
sudo usermod -a -G redis $USER
```

This should create a redis socket in `/run/redis/redis-server.sock`

### gvm-libs

Is used by `libeulabeia` to establish mqtt-connection and [openvas](https://github.com/greenbone/openvas-scanner).
```
sudo apt install libglib2.0-dev \
    cmake \
    lcov \
    libcgreen1-dev \
    libgnutls28-dev \
    libgpgme-dev \
    libhiredis-dev \
    libical-dev \
    libnet1-dev \
    libnet1-dev \
    libpaho-mqtt-dev \
    libpcap-dev \
    libpq-dev \
    libssh-gcrypt-dev \
    libssl-dev \
    libxml2-dev \
    make \
    pkg-config \
    postgresql-server-dev-all \
    uuid-dev \
    xsltproc
```

```

git clone https://github.com/greenbone/gvm-libs.git -b middleware
cmake -Bbuild -DREDIS_SOCKET_PATH=/run/redis/redis-server.sock
sudo cmake --build build --target install
sudo ldconfig
```

### libeulabeia

Is used by [openvas](https://github.com/greenbone/openvas-scanner) to communicate with eulabeia. For the installation you will need to be in the root directory of the eulabeia repository.

```
sudo apt install libjson-glib-dev
```

```
make c-library
sudo cmake --build c/libeulabeia/build -- install
sudo ldconfig
```


### openvas

[openvas](https://github.com/greenbone/openvas-scanner) is used by our sensor.

```
sudo apt install libksba-dev bison
```

```
git clone https://github.com/greenbone/openvas-scanner.git -b middleware
cmake -Bbuild
sudo cmake --build build --target install
echo "db_address = /run/redis/redis-server.sock" | sudo tee /etc/openvas/openvas.conf
echo "mqtt_server_uri = localhost:1883" | sudo tee -a /etc/openvas/openvas.conf
echo "$USER ALL = NOPASSWD: /usr/local/sbin/openvas" | sudo tee /etc/sudoers.d/allow_openvas
```

### director

Installing the required packages with `apt` and configure the director as follows:

```
sudo apt install golang
```

```
make build-director
mkdir -p ~/.config/eulabeia
sed 's/server = ".*/server = "localhost:1883"/' config.toml | tee ~/.config/eulabeia/config.toml
sed -i 's/StoragePath = ".*/StoragePath = "\/tmp\/"/' ~/.config/eulabeia/config.toml
sed -i 's/KeyFile = ".*/KeyFile = "\/tmp\/private.key"/' ~/.config/eulabeia/config.toml
./bin/eulabeia-director
```

This changes the MQTT server address to `localhost:1883`, as well as the storage and key file paths to `/tmp`.

### sensor

Installing the required packages with `apt` and configure the sensor as follows:

```
sudo apt install golang
```

```
make build-sensor
sed -i 's/RedisDbAddress = ".*/RedisDbAddress = "\/run\/redis\/redis-server.sock"/' ~/.config/eulabeia/config.toml
sudo cp plugins/* /var/lib/openvas/plugins/
./bin/eulabeia-sensor
```

This changes the MQTT server address to `localhost:1883` and the redis socket path to `/run/redis/redis-server.sock`.
