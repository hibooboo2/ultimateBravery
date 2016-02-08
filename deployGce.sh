#!/usr/bin/env bash
git cm "Auto commit from auto deploy $(date)"
git push

ssh gce.jhrb.us RIOT_API_KEY=$RIOT_API_KEY /home/wizardofmath/code/ultimateBravery/scripts/deployToDocker.sh
