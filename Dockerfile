FROM golang:1.13.6-alpine

WORKDIR /go
COPY . /go

RUN apk update && apk add git && \
    go get github.com/labstack/echo

CMD ["go", "run", "main.go"]