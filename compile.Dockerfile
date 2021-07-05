FROM debian:stable-slim
RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y \
    ca-certificates \
    golang \
    build-essential \
    git
RUN apt-get remove --purge --auto-remove -y &&\
	rm -rf /var/lib/apt/lists/*
COPY . /usr/local/src
COPY config.toml /usr/etc/eulabeia/config.toml
WORKDIR /usr/local/src
RUN go mod tidy
RUN go mod download
RUN go build --race -o /usr/local/bin/example-client cmd/example-client/main.go
RUN go build --race -o /usr/local/bin/eulabia-sensor cmd/eulabia-sensor/main.go
RUN go build --race -o /usr/local/bin/eulabia-director cmd/eulabia-director/main.go
