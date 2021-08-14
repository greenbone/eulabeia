# we use debian:testing due to paho otherwise we would need to install
# manually
FROM debian:testing-slim as core 
RUN sed -i 's/deb.debian.org/ftp.de.debian.org/' /etc/apt/sources.list
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y \
	vim \
	git \
	gnupg2 \
	build-essential\
	equivs \
	python3-pip\
	python3-setuptools\
	python3-dev\
	libssl-dev\
	libffi-dev \
	devscripts &&\
	apt-get remove --purge --auto-remove -y &&\
	rm -rf /var/lib/apt/lists/*
RUN pip3 install pontos poetry
COPY clone.sh /usr/local/bin/clone.sh
COPY build.sh /usr/local/bin/build.sh

WORKDIR /usr/local/src/
