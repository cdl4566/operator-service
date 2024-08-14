FROM centos:centos7.9.2009

WORKDIR /root/operator-service/

RUN mkdir -p /root/operator-service/bin/
RUN mkdir -p /root/operator-service/config/

COPY bin/manager bin/
