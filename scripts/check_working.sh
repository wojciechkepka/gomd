#!/bin/bash

go test -v ./...
./build/gomd &

sleep 2
PING="$(curl http://localhost:5001/ping)"

[[ ! $PING =~ pong ]] && exit 1

