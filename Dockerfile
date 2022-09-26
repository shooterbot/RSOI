FROM golang:latest as builder
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE=on

WORKDIR /app

ARG SERVICE
COPY ./backend/src/$SERVICE ./src/$SERVICE
COPY ./backend/src/database ./src/database
COPY ./backend/src/utility ./src/utility
COPY ./backend/src/utility ./src/config
RUN go mod init RSOI
RUN go mod tidy

WORKDIR ./src/$SERVICE
RUN go build -o main main.go

CMD ./main -docker
