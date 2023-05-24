# syntax=docker/dockerfile:1
FROM golang:1.20-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./ ./

RUN go build -o /jenk

EXPOSE 9090

CMD [ "/jenk" ]
