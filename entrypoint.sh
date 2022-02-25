#!/bin/sh

if [ ! -f .env ]; then
    cp .env.example .env
fi

# go run main.go
tail -f /dev/null