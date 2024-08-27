#!/bin/sh

set -e

# Check required inputs
if [ -z "$INPUT_COMMAND" ]; then
  echo "Error: command is not set"
  exit 1
fi

if [ -z "$INPUT_DOCKER_IMAGE" ]; then
  echo "Error: docker_image is not set"
  exit 1
fi

if [ -z "$INPUT_DOCKER_TAG" ]; then
  echo "Error: docker_tag is not set"
  exit 1
fi

# Pass the arguments to the enigma binary
/go/src/app/enigma $INPUT_COMMAND --docker-image "$INPUT_DOCKER_IMAGE" --docker-tag "$INPUT_DOCKER_TAG" --cleanup "$INPUT_CLEANUP"
