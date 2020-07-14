#!/bin/bash

./gomd &

echo "$(ip a)"

PING="$(curl http://localhost:5001/ping)"

echo "Got: '$PING' want 'pong'"

[[ ! $PING =~ pong ]] && exit 1
