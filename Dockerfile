FROM golang:alpine
COPY hello.go /tmp
CMD go run /tmp/hello.go
