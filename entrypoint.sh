#!/bin/sh
set -e

# Pass environment variables to the Go application
export DOCKER_IMAGE
export DOCKER_TAG
export CLEANUP

# Execute the command passed to the Docker container
exec "/go/src/app/enigma" "$@"