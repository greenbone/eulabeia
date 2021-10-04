GO_BUILD = CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
DOCKER_BUILD = docker build --no-cache --force-rm=true --compress=true
BROKER_IP = $(or $(shell docker container inspect -f '{{ .NetworkSettings.IPAddress }}' eulabeia_broker), $(echo ""))
MQTT_CONTAINER = docker run -e "MQTT_SERVER=$(call BROKER_IP):9138" --rm
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)

ifndef REPOSITORY
	REPOSITORY := "greenbone"
endif

.PHONY: sensor director format check test build update
all: format prepare check test build

prepare:
	@if [ $(GO_MINOR_VERSION) -gt 15 ]; then\
		go install honnef.co/go/tools/cmd/staticcheck@latest;\
	else\
		go install honnef.co/go/tools/cmd/staticcheck;\
	fi

format:
	go mod tidy
	go fmt ./...
	clang-format -i --style=file ./c/libeulabeia/src/*.c
	clang-format -i --style=file ./c/libeulabeia/test/src/*.c
	clang-format -i --style=file ./c/libeulabeia/include/eulabeia/*.h
	clang-format -i --style=file ./c/example/*.c

c-library:
	cmake -S c/libeulabeia/ -B c/libeulabeia/build
	cmake --build c/libeulabeia/build
	cmake --build c/libeulabeia/build -- test

c-examples:
	cmake -S c/example/ -B c/example/build
	cmake --build c/example/build

check:
	go vet ./...
	staticcheck ./...

test:
	go test ./...

start-broker:
	docker run --rm -d -p 9138:9138 --name eulabeia_broker $(REPOSITORY)/eulabeia-broker:latest

stop-broker:
	docker kill eulabeia_broker

start-director:
	$(MQTT_CONTAINER) -d --name eulabeia_director $(REPOSITORY)/eulabeia-director

stop-director:
	docker stop eulabeia_director

start-sensor:
	docker volume create eulabeia_redis_socket
	docker run -d --rm -v eulabeia_redis_socket:/run/redis --name eulabeia_redis $(REPOSITORY)/eulabeia-redis
	docker volume create eulabeia_feed
	$(MQTT_CONTAINER) -d -v eulabeia_feed:/var/lib/openvas/feed/plugins -v eulabeia_redis_socket:/run/redis --name eulabeia_sensor $(REPOSITORY)/eulabeia-sensor
	docker exec eulabeia_sensor mkdir -p /etc/openvas
	docker exec eulabeia_sensor bash -c 'echo "mqtt_server_uri = $(BROKER_IP):9138" >> /etc/openvas/openvas.conf'
	docker exec eulabeia_sensor openvas -u

stop-sensor:
	docker stop eulabeia_sensor
	docker volume rm eulabeia_feed
	docker kill eulabeia_redis
	docker volume rm eulabeia_redis_socket

run-example-client:
	until test `docker inspect eulabeia_sensor --format='{{.State.Running}}'` = "true"; do echo "waiting for sensor"; sleep 1; done
	until test `docker inspect eulabeia_director --format='{{.State.Running}}'` = "true"; do echo "waiting for director"; sleep 1; done
	$(MQTT_CONTAINER) --name eulabeia_example $(REPOSITORY)/eulabeia-example-client

start-smoke-test: start-container run-example-client

start-container: start-broker start-director start-sensor
stop-container: stop-director stop-sensor stop-broker

smoke-test: build-container start-smoke-test stop-container

build-director:
	$(GO_BUILD) -o $(DESTDIR)bin/eulabeia-director cmd/eulabeia-director/main.go

build-sensor:
	$(GO_BUILD) -o $(DESTDIR)bin/eulabeia-sensor cmd/eulabeia-sensor/main.go

build-example:
	$(GO_BUILD) -o $(DESTDIR)bin/example-client cmd/example-client/main.go


build: build-director build-sensor build-example

build-container-broker:
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-broker -f broker.Dockerfile .

build-container-redis:
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-redis -f redis.Dockerfile .

build-container-director: build-director
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-director -f director.Dockerfile .

build-container-clib:
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-c-lib -f eulabeia-c-library.Dockerfile .

build-container-c-example: build-container-clib
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-message-json-overview -f ./message-json-overview-md.Dockerfile .

build-container-sensor: build-container-clib build-sensor
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-sensor -f sensor.Dockerfile .

build-container-example: build-example
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-example-client -f example-client.Dockerfile .

build-container-build-helper:
	docker pull debian:stable-slim
	$(DOCKER_BUILD) -t $(REPOSITORY)/eulabeia-build-helper -f .docker/build-helper.Dockerfile .docker/

build-container: build-container-build-helper build-container-redis build-container-broker build-container-director build-container-sensor build-container-example

update:
	go get -u all
