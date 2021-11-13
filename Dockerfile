FROM golang:1.17.1

RUN apt update
RUN apt install -y vim
RUN apt install -y curl

COPY . /app
WORKDIR /app
RUN ./build-app.sh

CMD ./burgh
