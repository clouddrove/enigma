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

# Ensure Docker is using the correct config
export DOCKER_CONFIG="/github/home/.docker"

# Re-authenticate Docker with ECR (in case the authentication has expired)
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin 924144197303.dkr.ecr.us-east-1.amazonaws.com

exec "/go/src/app/enigma" "$@"