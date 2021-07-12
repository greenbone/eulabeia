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


BROKER_IP = $(or $(shell docker container inspect -f '{{ .NetworkSettings.IPAddress }}' eulabeia_broker), $(echo ""))
MQTT_CONTAINER = docker run -e "MQTT_SERVER=$(call BROKER_IP):9138" --rm

# change to build specific production ready container instead of a compile image
build-container:
	docker build -t eulabeia/broker -f broker.Dockerfile .
	docker build -t eulabeia/compile -f compile.Dockerfile .

start-broker:
	docker run --rm -d --name eulabeia_broker eulabeia/broker:latest

stop-broker:
	docker kill eulabeia_broker

start-director:
	$(MQTT_CONTAINER) -d --name eulabeia_director eulabeia/compile eulabeia-director --clientid director

stop-director:
	docker stop eulabeia_director

start-sensor:
	$(MQTT_CONTAINER) -d --name eulabeia_sensor eulabeia/compile eulabeia-sensor

stop-sensor:
	docker stop eulabeia_sensor

run-example-client:
	until test `docker inspect eulabeia_sensor --format='{{.State.Running}}'` = "true"; do echo "waiting for sensor"; sleep 1; done
	until test `docker inspect eulabeia_director --format='{{.State.Running}}'` = "true"; do echo "waiting for director"; sleep 1; done
	docker ps
	$(MQTT_CONTAINER) --name eulabeia_example --rm eulabeia/compile example-client --clientid example || ( docker logs eulabeia_director && docker logs eulabeia_sensor && exit 1) 

start-smoke-test: start-container run-example-client

start-container: start-broker start-director start-sensor
stop-container: stop-broker stop-director stop-sensor

smoke-test: build-container start-smoke-test stop-container

build:
	go build -o $(DESTDIR)bin/eulabeia-director cmd/eulabeia-director/main.go
	go build -o $(DESTDIR)bin/eulabeia-sensor cmd/eulabeia-sensor/main.go
	go build -o $(DESTDIR)bin/example-client cmd/example-client/main.go

update:
	go get -u all
