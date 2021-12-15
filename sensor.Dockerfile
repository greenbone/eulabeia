ARG OPENVAS_REPOSITORY="git://github.com/greenbone/openvas-scanner"
ARG OPENVAS_BRANCH="middleware"
FROM greenbone/eulabeia-c-lib AS build
# we need our own copy of openvas to have the latest 
# eulabeia lib changes.
ARG OPENVAS_REPOSITORY
ARG OPENVAS_BRANCH
RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y \
    bison \
    build-essential \
    clang \
    clang-format \
    clang-tools \
    cmake \
    lcov \
    libcgreen1-dev \
    libgnutls28-dev \
    libgpgme-dev \
    libjson-glib-dev \
    libksba-dev \
    libpaho-mqtt-dev \
    libpcap-dev \
    libssh-gcrypt-dev \
    libnet1-dev \
	git \
    && rm -rf /var/lib/apt/lists/*
RUN echo "cloning: $OPENVAS_REPOSITORY"
RUN git clone --depth 1 --single-branch --branch $OPENVAS_BRANCH $OPENVAS_REPOSITORY /source
RUN cmake -DCMAKE_BUILD_TYPE=Release -B/build /source
RUN DESTDIR=/install cmake --build /build -- install 

FROM greenbone/eulabeia-c-lib
# installing openvas scanner into the sensor image
COPY --from=build /install/ /
COPY openvas_log.conf /etc/openvas/openvas_log.conf
COPY config.toml /usr/etc/eulabeia/config.toml
COPY bin/eulabeia-sensor /usr/bin/eulabeia-sensor
COPY plugins/* /var/lib/openvas/plugins/
RUN echo "mqtt_context = scanner" >> /etc/openvas/openvas.conf
RUN echo "mqtt_server_uri = broker:9138" >> /etc/openvas/openvas.conf
RUN ldconfig
CMD /usr/bin/eulabeia-sensor -clientid localhorst
