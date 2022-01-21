# syntax=docker/dockerfile:1

# minimal golang image for building
FROM golang:1.16-alpine as build
WORKDIR /go/src

# install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy src and build
COPY *.go ./
RUN go build -o build/sonic


# alpine image for running
FROM alpine

# install dependencies
RUN apk add --no-cache ca-certificates && update-ca-certificates

# copy binary
COPY --from=build /go/src/build/sonic /usr/bin/sonic

# run on port 8080
EXPOSE 8080
ENTRYPOINT ["/usr/bin/sonic"]
