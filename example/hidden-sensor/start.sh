#!/bin/bash
trap "redis-cli -s /var/run/redis/redis.sock shutdown" SIGTERM
sudo mosquitto -c /etc/mosquitto.conf &
sudo redis-server /etc/redis/redis.conf
sudo /usr/sbin/sshd
/usr/bin/eulabeia-sensor -clientid hiddenhorst
