FROM scratch
COPY config.toml /usr/etc/eulabeia/config.toml
COPY bin/eulabeia-sensor /
CMD [ "/eulabeia-sensor" ]
