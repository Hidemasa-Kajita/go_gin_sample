FROM golang:1.16

COPY ./src /go/src

WORKDIR /go/src

RUN go build

CMD ["go", "run", "main.go"]
