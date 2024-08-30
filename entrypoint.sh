#!/bin/sh
set -e

# Add Google Cloud SDK to PATH
export PATH=$PATH:/google-cloud-sdk/bin

# Check if provider is AWS
if [ "$PROVIDER" = "aws" ]; then
  aws ecr get-login-password --region "$AWS_REGION" | docker login --username AWS --password-stdin "$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com"
# Check if provider is GCP
elif [ "$PROVIDER" = "gcp" ]; then
  echo "$GCLOUD_AUTH" | gcloud auth activate-service-account --key-file=-
  gcloud auth configure-docker
# Check if provider is Docker Hub
elif [ "$PROVIDER" = "dockerhub" ]; then
  echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin
# Check if provider is GitHub Container Registry
elif [ "$PROVIDER" = "github" ]; then
  echo "$TOKEN" | docker login ghcr.io -u "$GITHUB_USERNAME" --password-stdin
else
  echo "Please pass a valid provider"
  exit 1
fi

# Execute the application
exec "/go/src/app/enigma" "$@"
