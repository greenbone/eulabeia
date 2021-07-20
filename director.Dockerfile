FROM scratch
COPY config.toml /usr/etc/eulabeia/config.toml
COPY LICENSE /var/lib/eulabeia/director/storage/LICENSE
COPY bin/eulabeia-director /
CMD [ "/eulabeia-director" ]
