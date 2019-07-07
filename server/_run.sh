#!/bin/sh

# docker run
docker-compose up -d

# dep ensure -update -v
dep ensure

export ENV=local
go run main.go
