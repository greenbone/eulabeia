ARG VERSION=middleware
FROM greenbone/eulabeia-c-lib AS lib-eulabeia

FROM greenbone/eulabeia-build-helper AS openvas
ARG VERSION
COPY .docker/descriptions/openvas /usr/local/src/openvas-scanner
COPY --from=lib-eulabeia /etc/apt/sources.list.d/docker.list /etc/apt/sources.list.d/docker.list
COPY --from=lib-eulabeia /usr/local/src/packages /usr/local/src/packages
RUN /usr/local/bin/clone.sh openvas-scanner $VERSION
RUN /usr/local/bin/build.sh openvas-scanner $VERSION

FROM debian:stable-slim
# installing openvas scanner into the sensor image
COPY --from=openvas /usr/local/src/docker.list /etc/apt/sources.list.d/docker.list
COPY --from=openvas /usr/local/src/packages /usr/local/src/packages
RUN apt-get update &&\
	apt-get install --no-install-recommends -y openvas &&\
	apt-get remove --purge --auto-remove -y &&\
	rm -rf /var/lib/apt/lists/*
RUN rm -rf /usr/local/src/*
RUN rm /etc/apt/sources.list.d/docker.list
COPY openvas_log.conf /etc/openvas/openvas_log.conf
COPY config.toml /usr/etc/eulabeia/config.toml
COPY bin/eulabeia-sensor /usr/bin/eulabeia-sensor
RUN echo "mqtt_context = scanner" >> /etc/openvas/openvas.conf
CMD [ "/usr/bin/eulabeia-sensor" ]
