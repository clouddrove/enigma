#!/bin/sh
set -e

export DOCKER_IMAGE
export DOCKER_TAG
export CLEANUP

exec "/go/src/app/enigma" "$@"