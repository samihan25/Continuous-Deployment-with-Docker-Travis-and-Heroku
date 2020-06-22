# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.14-alpine

RUN apk update

COPY ./server.go /

WORKDIR /

RUN go build server.go

ENTRYPOINT [ "./server" ]

# Document that the service listens on port 8080.
EXPOSE 8080