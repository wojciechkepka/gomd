#!/bin/bash

./gomd &

echo "$(ip a)"

PING="$(curl http://localhost:5001/ping)"

[[ ! $PING =~ "pong" ]] && exit 1
