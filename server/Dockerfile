# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY go.mod go.sum ./

COPY serv_user.go ./

RUN go build -o ./user-app

EXPOSE 8080

CMD ["/go/src/app/user-app"]