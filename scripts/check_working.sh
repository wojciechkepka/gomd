#!/bin/bash

./gomd &

PING="$(curl http://127.0.0.1:5001/ping)"

[[ ! $PING =~ "pong" ]] && exit 1
