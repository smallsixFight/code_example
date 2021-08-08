FROM ubuntu:21.04

ENV HOME="/root"

WORKDIR /root/

COPY gorm_test /root/

RUN apt-get update

CMD nohup ./gorm_test; sleep infinity