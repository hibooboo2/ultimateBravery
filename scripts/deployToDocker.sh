#!/usr/bin/env bash

cd $(dirname $0)/..

NUM_GIT_CHANGES=$(($(git status -sb |wc -l)-1))
[[ "${NUM_GIT_CHANGES}" != "0" ]] && git add --all && git stash > /dev/null
GIT_LOCATION=$(git rev-parse --abbrev-ref HEAD 2> /dev/null)
[[ -z ${GIT_LOCATION} ]] && echo "Failed to find a branch" && exit 1

[[ -z ${RIOT_API_KEY} ]] && echo Need riot api key to deply && exit 1
function unStash {
    git checkout ${GIT_LOCATION}
    [[ "${NUM_GIT_CHANGES}" != "0" ]] && git stash pop
}
trap unStash EXIT

git stash


git checkout origin/master
PRE_GITCOMMIT=`git rev-parse --short HEAD`

git fetch --all
git checkout origin/master
GITCOMMIT=`git rev-parse --short HEAD`

deploy_container() {
    docker rm -fv ultimateBravery > /dev/null || echo "No container for ultimateBravery".
    docker build -t ultimate-bravery . > /dev/null
    docker run --restart=always -d -p 9000:8000 -e RIOT_API_KEY=${RIOT_API_KEY} --name=ultimateBravery ultimate-bravery > /dev/null
    echo Deployed to docker.
}
[[ ${PRE_GITCOMMIT} != ${GITCOMMIT} ]] && deploy_container
