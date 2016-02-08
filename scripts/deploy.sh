#!/usr/bin/env bash

set -e

cd $(dirname $0)/..

git fetch --all
git checkout origin/master

[[ -z "$(docker ps -a | grep ultimateBravery)" ]] && docker rm -fv ultimateBravery
docker build -t ultimate-bravery .
docker run --restart=always -d -p 9001:8000 -e RIOT_API_KEY=$RIOT_API_KEY --name=ultimateBravery ultimate-bravery
