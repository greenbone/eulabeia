.PHONY: format check test build update
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

build:
	go build -o $(DESTDIR)bin/eulabeia-director cmd/eulabeia-director/main.go
	go build -o $(DESTDIR)bin/eulabeia-sensor cmd/eulabeia-sensor/main.go
	go build -o $(DESTDIR)bin/example-client cmd/example-client/main.go

update:
	go get -u all
