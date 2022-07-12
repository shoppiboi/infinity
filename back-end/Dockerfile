FROM golang:1.18-bullseye

RUN mkdir code

ADD ./src/ /code
WORKDIR /code/

ENV PORT=8080

RUN go build main.go

CMD ["./main"]