#!/bin/bash

trap on_exit INT
trap on_exit EXIT


function on_exit() {
  docker-compose kill
}

docker-compose up -d

export $(cat .env | xargs) && go run ./cmd/http/httpd.go
