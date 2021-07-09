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
RUN DESTDIR="/usr/local/" make test build
