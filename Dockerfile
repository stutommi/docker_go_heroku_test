FROM golang:latest

WORKDIR /go/src
COPY ./main.go .

CMD go run main.go
