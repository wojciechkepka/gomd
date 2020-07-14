#!/bin/bash

./gomd &

sleep 2
PING="$(curl http://localhost:5001/ping)"
[[ ! $PING =~ pong ]] && echo '1'
