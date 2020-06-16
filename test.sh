#!/bin/bash
docker-compose up -d

go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

go tool cover -html=coverage.txt

docker-compose kill
