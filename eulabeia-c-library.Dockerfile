ARG VERSION=unstable

FROM greenbone/gvm-libs:$VERSION as build
RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y \
    build-essential \
    cmake \
    pkg-config \
    libglib2.0-dev \
    libpaho-mqtt-dev \
    libjson-glib-dev \
    libcgreen1-dev \
    libgnutls28-dev \
    && rm -rf /var/lib/apt/lists/*
COPY c/libeulabeia /source
RUN cmake -DCMAKE_BUILD_TYPE=Release -B/build /source
RUN DESTDIR=/install cmake --build /build -- install test

FROM greenbone/gvm-libs:$VERSION
COPY --from=build /install/ / 
RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y \
    libjson-glib-1.0-0 \
    && rm -rf /var/lib/apt/lists/*
RUN ldconfig
