# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine3.15

ADD ./src/ /code/
WORKDIR /code

RUN go build main.go

EXPOSE 8080

CMD ["./main"]