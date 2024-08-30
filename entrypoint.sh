#!/bin/sh
set -e

# Check if provider is AWS
if [ "$PROVIDER" = "aws" ]; then
  aws ecr get-login-password --region "$AWS_REGION" | docker login --username AWS --password-stdin "$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com"
fi

exec "/go/src/app/enigma" "$@"
