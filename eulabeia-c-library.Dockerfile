ARG VERSION=middleware

from greenbone/eulabeia-build-helper as lib-gvm
arg VERSION
run sed -i 's/deb.debian.org/ftp.de.debian.org/' /etc/apt/sources.list
copy .docker/descriptions/gvm-libs /usr/local/src/gvm-libs
run /usr/local/bin/clone.sh gvm-libs $VERSION
run /usr/local/bin/build.sh gvm-libs $VERSION

from greenbone/eulabeia-build-helper as lib-eulabeia
copy --from=lib-gvm /usr/local/src/docker.list /etc/apt/sources.list.d/docker.list
copy --from=lib-gvm /usr/local/src/packages /usr/local/src/packages
run mkdir /usr/local/src/libeulabeia
copy c/debian /usr/local/src/libeulabeia/debian
copy c/libeulabeia /usr/local/src/libeulabeia/libeulabeia
run /usr/local/bin/build.sh libeulabeia $VERSION "1.0.0"

from debian:stable-slim as openvas
arg VERSION
run sed -i 's/deb.debian.org/ftp.de.debian.org/' /etc/apt/sources.list
copy .docker/descriptions/openvas /usr/local/src/openvas-scanner
copy --from=lib-eulabeia /usr/local/src/docker.list /etc/apt/sources.list.d/docker.list
copy --from=lib-eulabeia /usr/local/src/packages /usr/local/src/packages
RUN apt-get update
RUN apt-get install --no-install-recommends -y libeulabeia-dev
