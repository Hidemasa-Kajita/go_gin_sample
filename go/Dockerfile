FROM golang:1.16

COPY ./src /go/src

WORKDIR /go/src

RUN go mod download && \
    go get -u \
    github.com/gin-gonic/gin \
    github.com/cosmtrek/air

RUN go build

# CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]
