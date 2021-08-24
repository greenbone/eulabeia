FROM greenbone/eulabeia-c-lib AS build
COPY . /usr/local/src/
RUN apt-get update &&\
    apt-get remove -y libeulabeia-dev &&\
    apt-get install -y libeulabeia-dev 
WORKDIR /usr/local/src/
RUN make c-examples

FROM greenbone/eulabeia-c-lib
COPY --from=build /usr/local/src/c/example/build/message-json-overview-md /usr/local/bin/message-json-over-md
RUN chmod +x /usr/local/bin/message-json-over-md
CMD [ "/usr/local/bin/message-json-over-md" ]
