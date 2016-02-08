#!/usr/bin/env bash

set -e

cd $(dirname $0)/..

docker rm -fv ultimateBravery || echo "No container for ultimateBravery".
docker build -t ultimate-bravery .
docker run --restart=always -d -p 9001:8000 -e RIOT_API_KEY=$RIOT_API_KEY --name=ultimateBravery ultimate-bravery
