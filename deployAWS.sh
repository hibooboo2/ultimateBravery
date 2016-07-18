#!/usr/bin/env bash

ssh -i ~/.ssh/AWS2Drunk.pem aws.jhrb.us RIOT_API_KEY=$RIOT_API_KEY /home/wizardofmath/go/src/github.com/hibooboo2/ultimateBravery/scripts/deployToDocker.sh
