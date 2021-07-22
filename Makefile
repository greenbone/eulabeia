.PHONY: sensor director format check test build update
all: format check test build

prepare:
	go install honnef.co/go/tools/cmd/staticcheck

format:
	go mod tidy
	go fmt ./...

check:
	go vet ./...
	staticcheck ./...

test:
	go test ./...

GO_BUILD = CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
DOCKER_BUILD = docker build --force-rm=true --compress=true
BROKER_IP = $(or $(shell docker container inspect -f '{{ .NetworkSettings.IPAddress }}' eulabeia_broker), $(echo ""))
MQTT_CONTAINER = docker run -e "MQTT_SERVER=$(call BROKER_IP):9138" --rm


start-broker:
	docker run --rm -d --name eulabeia_broker greenbone/eulabeia-broker:latest

stop-broker:
	docker kill eulabeia_broker

start-director:
	$(MQTT_CONTAINER) -d --name eulabeia_director greenbone/eulabeia-director

stop-director:
	docker stop eulabeia_director

start-sensor:
	docker volume create eulabeia_feed
	docker pull greenbone/community-feed-vts:latest
	docker run -d --rm -v eulabeia_feed:/opt/greenbone/feed/plugins greenbone/community-feed-vts:latest echo ""
	$(MQTT_CONTAINER) -d -v eulabeia_feed:/var/lib/openvas/feed/plugins --name eulabeia_sensor greenbone/eulabeia-sensor

stop-sensor:
	docker stop eulabeia_sensor
	docker volume rm eulabeia_feed

run-example-client:
	until test `docker inspect eulabeia_sensor --format='{{.State.Running}}'` = "true"; do echo "waiting for sensor"; sleep 1; done
	until test `docker inspect eulabeia_director --format='{{.State.Running}}'` = "true"; do echo "waiting for director"; sleep 1; done
	$(MQTT_CONTAINER) --name eulabeia_example greenbone/eulabeia-example-client

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
	$(DOCKER_BUILD) -t greenbone/eulabeia-broker -f broker.Dockerfile .

build-container-director: build-director
	$(DOCKER_BUILD) -t greenbone/eulabeia-director -f director.Dockerfile .

build-container-sensor: build-sensor
	$(DOCKER_BUILD) -t greenbone/eulabeia-sensor -f sensor.Dockerfile .

build-container-example: build-example
	$(DOCKER_BUILD) -t greenbone/eulabeia-example-client -f example-client.Dockerfile .

build-container-build-helper:
	$(DOCKER_BUILD) -t greenbone/eulabeia-build-helper -f .docker/build-helper.Dockerfile .docker/

build-container: build-container-build-helper build-container-broker build-container-director build-container-sensor build-container-example

update:
	go get -u all
