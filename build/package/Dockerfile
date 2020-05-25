FROM golang:1.14-alpine3.11 AS build-engine

ARG BUILD_VERSION
ARG BUILD_TIME

WORKDIR /go/src/mayaleng.org/engine

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -v -ldflags="-X 'version.BuildVersion=${BUILD_VERSION}}' -X 'version.BuildTime=${BUILD_TIME}'" -o /tmp/httpd ./cmd/http/httpd.go

FROM mayaleng/linguakit:alpine3.11

WORKDIR /engine

COPY --from=build-engine /tmp/httpd .

ENTRYPOINT [ "./httpd" ]
