#!/bin/bash

go get golang.org/x/sys/unix
GOARCH=amd64 GOOS=linux go build -o ../devops/logger/bin/logger
