FROM debian:stable-slim 
RUN apt-get -y update &&\
	apt-get -y --no-install-recommends install redis
COPY redis.conf /etc/redis/redis.conf
RUN mkdir /run/redis
RUN chown redis:redis /run/redis
USER redis
CMD redis-server /etc/redis/redis.conf
