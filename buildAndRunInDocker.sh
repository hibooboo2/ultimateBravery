#!/usr/bin/env bash
set -e

ubVersion=tmp/ultimatebravery:$(git rev-parse --verify HEAD)

docker build -t ${ubVersion} .
docker run --restart=always -d -p 9001:8000 -e RIOT_API_KEY=${RIOT_API_KEY} ${ubVersion}
