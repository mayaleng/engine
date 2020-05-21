#!/bin/bash

trap on_exit INT
trap on_exit EXIT


function on_exit() {
  docker-compose kill
}

docker-compose up -d

go run ./cmd/http/httpd.go
