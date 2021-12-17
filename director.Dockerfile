FROM scratch
COPY config.toml /usr/etc/eulabeia/config.toml
COPY LICENSE /var/lib/eulabeia/director/storage/LICENSE
COPY --chmod=600 example/hidden-sensor/user /var/lib/eulabeia/director/user
COPY example/hidden-sensor/ssh_host_ed25519_key.pub /var/lib/eulabeia/director/host.pub
COPY bin/eulabeia-director /
CMD [ "/eulabeia-director" ]
