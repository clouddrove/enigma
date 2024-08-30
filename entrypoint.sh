#!/bin/sh
set -e

# Login to AWS ECR using the provided credentials from environment variables
aws ecr get-login-password --region "$AWS_REGION" | docker login --username AWS --password-stdin "$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com"

# Execute the main application
exec "/go/src/app/enigma" "$@"