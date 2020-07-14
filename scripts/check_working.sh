#!/bin/bash

./gomd &

sleep 2
PING="$(curl http://localhost:5001/ping)"
if [[ ! $PING =~ pong ]]
then
    exit 1
fi

