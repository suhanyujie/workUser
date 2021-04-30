#!/usr/bin/env bash
#
trap "rm server; kill 0" EXIT

go build -o server

./server -f=etc/workuser-api.yaml &

wait
