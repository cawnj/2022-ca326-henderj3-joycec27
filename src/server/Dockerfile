# syntax=docker/dockerfile:1

# minimal golang image for building
FROM golang:1.17-alpine as build
WORKDIR /go/src/sonic-server

# copy src and install dependencies
COPY . .
RUN go get -d -v ./...

# build
RUN go build -o build/sonic-server

# alpine image for running
FROM alpine

# install dependencies
RUN apk add --no-cache ca-certificates && update-ca-certificates

# copy binary
COPY --from=build /go/src/sonic-server/build/sonic-server /usr/bin/sonic-server

# run on port 8080
EXPOSE 8080
ENTRYPOINT ["/usr/bin/sonic-server"]
