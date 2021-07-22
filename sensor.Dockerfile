ARG VERSION=middleware
FROM greenbone/eulabeia-build-helper AS lib-gvm
ARG VERSION
COPY .docker/descriptions/gvm-libs /usr/local/src/gvm-libs
RUN /usr/local/bin/clone.sh gvm-libs $VERSION
RUN /usr/local/bin/build.sh gvm-libs $VERSION

FROM greenbone/eulabeia-build-helper AS openvas
ARG VERSION
COPY .docker/descriptions/openvas /usr/local/src/openvas
COPY --from=lib-gvm /usr/local/src/docker.list /etc/apt/sources.list.d/docker.list
COPY --from=lib-gvm /usr/local/src/packages /usr/local/src/packages
RUN /usr/local/bin/clone.sh openvas $VERSION
RUN /usr/local/bin/build.sh openvas $VERSION

# we use debian:testing due to paho otherwise we would need to install
# manually
FROM debian:testing-slim
# installing openvas scanner into the sensor image
COPY --from=openvas /usr/local/src/docker.list /etc/apt/sources.list.d/docker.list
COPY --from=openvas /usr/local/src/packages /usr/local/src/packages
RUN apt-get update &&\
	apt-get install --no-install-recommends -y openvas &&\
	apt-get remove --purge --auto-remove -y &&\
	rm -rf /var/lib/apt/lists/*
RUN rm -rf /usr/local/src/*
RUN rm /etc/apt/sources.list.d/docker.list
COPY config.toml /usr/etc/eulabeia/config.toml
COPY bin/eulabeia-sensor /usr/bin/eulabeia-sensor
CMD [ "/usr/bin/eulabeia-sensor" ]
