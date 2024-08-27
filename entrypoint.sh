#!/bin/sh
set -e

echo "AWS_ACCESS_KEY_ID is set: ${AWS_ACCESS_KEY_ID:+true}"
echo "AWS_SECRET_ACCESS_KEY is set: ${AWS_SECRET_ACCESS_KEY:+true}"
echo "AWS_REGION is set: ${AWS_REGION:+true}"

export DOCKER_IMAGE
export DOCKER_TAG
export CLEANUP
export AWS_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY
export AWS_REGION

exec "/go/src/app/enigma" "$@"