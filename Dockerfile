FROM golang:1.17.1

RUN apt update
RUN apt install -y vim
RUN apt install -y curl

COPY . /app
RUN /app/build-app.sh

WORKDIR /app-dev

CMD /app/my-app
