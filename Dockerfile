FROM ubuntu:focal

RUN apt update
RUN apt install -y vim
RUN apt install -y curl

COPY install-go.sh /app/install-go.sh 
RUN source /app/install-go.sh

ENV PATH="/usr/local/go/bin:${PATH}"

#COPY . /app
WORKDIR /app
