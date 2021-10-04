# Requirements

This document describes the requirements. 

The favored usage of eulabeia is via [docker](#Docker).

It is divided between two methods:
- [docker](#Docker), will describe how to setup docker and use the provided makefile
- [local](#Local), will describe how to setup and install the requirements without provided images

This document will use Debian 11 as a basis and is targeting developer machines with relatively lax security requirements. 

For production usage please harden docker, redis and mosquitto accordingly.

## Docker

To use eulabeia with docker you need to install `go`, `docker` as well as `gnu-make` to make it easier to use you can add your user to the docker group.

```
# apt-get install docker.io make golang
# usermod -a -G docker your_username
```

To reload the newly assign group, login to your user account again

```
$ su - $USER
$ docker ps
```

to build the images based on your changes you can use `make start-container` with that you have a useable sensor, director and broker up and running.

If you want to verify the setup you can simply use `make run-example-client`.

If you have `mosquitto-clients` installed, you can can verify that the local connection works:
```
$ mosquitto_sub -L 'mqtt://localhost:9138/scanner/#'
```

to subscribe to any message inside the scanner context. 

Afterwards you can send a irregular message

```
$ mosquitto_pub -L 'mqtt://localhost:9138/scanner/scan/cmd/director' -m "hi"
```

### Broker

To run a broker you can execute `make start-broker` this will run mosquitto and open a local port 9138.

### Sensor

To run a sensor you can execute `make start-sensor` this will run mosquitto and open a local port 9138 as well as start a sensor containing a test nasl file (oid: `1.3.6.1.4.1.25623.1.0.90022`) used by openvas.

To communicate with the sensor use mqtt as described previously.

### Director

To run a director you can execute `make start-director` this will run mosquitto and open a local port 9138 as well as start a director.

To communicate with the director use mqtt as described previously.

## Local

### Broker (mosquitto)

To install and configure mosquitto:

```
# apt-get install mosquitto
# echo "allow_anonymous true" > /etc/mosquitto/conf.d/allow_anonymous.conf
# systemctl restart mosquitto
```

This allows anonymous connections into the running mqtt broker.

### Redis

Redis is used by openvas as well as sensor.

To install and configure redis:

```
# apt install redis
# sed -i 's/\# unixsocket\([ p]\)/unixsocket\1/' /etc/redis/redis.conf
# systemctl restart redis
# usermod -a -G redis your_username
```

This should create a redis socket in `/run/redis/redis-server.sock`

### gvm-libs

Is used by libeulabeia to establish mqtt-connection and openvas.

```
# apt install libglib2.0-dev \
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

$ git clone https://github.com/greenbone/gvm-libs.git
$ git checkout -t origin/middleware
$ cmake -Bbuild -DREDIS_SOCKET_PATH=/run/redis/redis-server.sock
$ sudo 'cmake --build build --target install'
$ sudo 'ldconfig'
```

### libeulabeia

Is used by openvas to communicate with eulabeia.

```
# apt install libjson-glib-dev
```

```
$ make c-library
$ sudo 'cmake --build c/libeulabeia/build -- install'
$ sudo 'ldconfig'
```


### openvas

openvas is used by our sensor.

```
# apt install libksba-dev bison
```

```
$ git clone https://github.com/greenbone/openvas-scanner.git
$ cmake -Bbuild
$ sudo 'cmake --build build --target install'
$ echo "db_address = /run/redis/redis-server.sock" | sudo tee /etc/openvas/openvas.conf
$ echo "mqtt_server_uri = localhost:1883" | sudo tee -a /etc/openvas/openvas.conf
$ echo "$USER ALL = NOPASSWD: /usr/local/sbin/openvas" | sudo tee /etc/sudoers.d/allow_openvas
```

### director

```
# apt install golang
```

```
$ make build-director
$ mkdir -p ~/.config/eulabeia
$ sed 's/server = ".*/server = "localhost:1883"/' config.toml | tee ~/.config/eulabeia/config.toml
$ sed -i 's/StoragePath = ".*/StoragePath = "\/tmp\/"/' ~/.config/eulabeia/config.toml
$ sed -i 's/KeyFile = ".*/KeyFile = "\/tmp\/private.key"/' ~/.config/eulabeia/config.toml
$ ./bin/eulabeia-director
```

This changes the mqtt server address to `localhost:1883`, the storage and key file path to `/tmp`.

### sensor

```
# apt install golang
```

```
$ make build-sensor
$ sed -i 's/RedisDbAddress = ".*/RedisDbAddress = "\/run\/redis\/redis-server.sock"/' ~/.config/eulabeia/config.toml
$ sudo cp plugins/* /var/lib/openvas/plugins/
$ ./bin/eulabeia-sensor
```

This changes the mqtt server address to `localhost:1883` and the redis socket path to `/run/redis/redis-server.sock`.
