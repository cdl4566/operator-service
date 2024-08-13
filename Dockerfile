FROM golang:1.18.7-alpine3.16

WORKDIR /root/operator-service/

RUN mkdir -p /root/operator-service/bin/
RUN mkdir -p /root/operator-service/config/

COPY bin/manager bin/
