#!/bin/ash

./Client2

go test ./...
tcpdump -i lo -w
sleep 500
