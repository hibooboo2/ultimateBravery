#!/usr/bin/env bash
set -e

ubVersion="tmp/ultimateBravery-$(git rev-parse --verify HEAD)"

docker build -t ${ubVersion} .
docker run --restart=always -d -p 9001:8000 -e RIOT_API_KEY=4f77067c-ca74-4bdf-8b19-b7cd199f4a05 ${ubVersion}
