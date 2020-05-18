FROM golang:1.14-alpine3.11 AS build-engine

WORKDIR /go/src/mayaleng.org/engine

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /tmp/engine ./cmd/engine/main.go


FROM mayaleng/linguakit:alpine3.11

WORKDIR /engine

COPY --from=build-engine /tmp/engine .

ENTRYPOINT [ "./engine" ]
