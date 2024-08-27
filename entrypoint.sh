#!/bin/sh
set -e

export DOCKER_IMAGE
export DOCKER_TAG
export CLEANUP
export AWS_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY
export AWS_REGION

exec "/go/src/app/enigma" "$@"