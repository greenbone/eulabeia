FROM scratch
COPY config.toml /usr/etc/eulabeia/config.toml
COPY bin/example-client /
CMD [ "/example-client" ]
