#!/bin/sh
set -e

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 924144197303.dkr.ecr.us-east-1.amazonaws.com

exec "/go/src/app/enigma" "$@"