#!/bin/bash

if [ -z "$MAIN_PORT" ]; then
    export MAIN_PORT="7777"
fi

if [ -z "$DB_NAME" ]; then
    export DB_NAME="logger"
fi

if [ -z "$DB_PORT" ]; then
    export DB_PORT="5875"
fi

docker-compose build && docker-compose up -d