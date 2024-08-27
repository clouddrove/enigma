#!/bin/sh
set -e

# Pass environment variables to the Go application
export DOCKER_IMAGE
export DOCKER_TAG
export CLEANUP

# Check if Docker socket is available
if [ ! -e /var/run/docker.sock ]; then
    echo "Error: Docker socket (/var/run/docker.sock) is not available. Make sure to mount it when running the container."
    exit 1
fi

# Execute the command passed to the Docker container
exec "/go/src/app/enigma" "$@"