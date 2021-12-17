
FROM greenbone/eulabeia-sensor

RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y \
    mosquitto \
    mosquitto-clients \
    redis \
    openssh-server \
    sudo &&\
	apt-get remove --purge --auto-remove -y &&\
	rm -rf /var/lib/apt/lists/*

RUN useradd -rm -d /home/gvm -s /bin/bash -g root -g mosquitto -G sudo -u 1000 gvm
RUN  echo 'gvm:test' | chpasswd
RUN echo "gvm ALL = NOPASSWD: /usr/sbin/openvas" > /etc/sudoers.d/allow_openvas
RUN echo "gvm ALL = NOPASSWD: /usr/sbin/sshd" > /etc/sudoers.d/allow_sshd
RUN echo "gvm ALL = NOPASSWD: /usr/bin/redis-server" > /etc/sudoers.d/allow_redis-server
RUN echo "gvm ALL = NOPASSWD: /usr/sbin/mosquitto" > /etc/sudoers.d/allow_mosquitto
RUN mkdir /run/mosquitto
#RUN echo "listener 0 /run/mosquitto/mosquitto.sock" > /etc/mosquitto.conf
RUN echo "allow_anonymous true" >> /etc/mosquitto.conf
RUN echo "log_dest file /tmp/mosquitto.log" >> /etc/mosquitto.conf
RUN echo "persistence true" >> /etc/mosquitto.conf
RUN echo "mqtt_server_uri = localhost:1883" >> /etc/openvas/openvas.conf
COPY config.toml /usr/etc/eulabeia/config.toml
COPY start.sh /usr/local/bin/start.sh
COPY redis.conf /etc/redis/redis.conf
RUN mkdir /run/redis
RUN chown redis:redis /run/redis
RUN chown mosquitto:mosquitto /run/mosquitto
RUN mkdir -p /home/gvm/.ssh
RUN chown gvm:mosquitto /home/gvm/.ssh
COPY --chown=gvm:mosquitto --chmod=0600 user.pub /home/gvm/.ssh/authorized_keys
COPY --chown=root:root --chmod=0600 ssh_host_ed25519_key /etc/ssh/ssh_host_ed25519_key 
COPY --chown=root:root --chmod=0640 ssh_host_ed25519_key.pub /etc/ssh/ssh_host_ed25519_key.pub
COPY sshd_config /etc/ssh/sshd_config
EXPOSE 22
RUN mkdir /run/sshd
RUN mkdir -p /var/log/mosquitto/
RUN chmod 777 /var/log/mosquitto
RUN chown -R gvm /var/log/gvm
USER gvm
WORKDIR /home/gvm
CMD [ "/usr/local/bin/start.sh" ]
