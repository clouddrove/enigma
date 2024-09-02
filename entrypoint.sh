#!/bin/sh
set -e

# Check if provider is AWS
if [ "$PROVIDER" = "aws" ]; then
  aws ecr get-login-password --region "$AWS_REGION" | docker login --username AWS --password-stdin "$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com"

# Check if provider is GCP
elif [ "$PROVIDER" = "gcp" ]; then
  gcloud auth configure-docker "$GCP_REGION-docker.pkg.dev" --quiet

# Check if provider is Docker Hub
elif [ "$PROVIDER" = "dockerhub" ]; then
  echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin

# Check if provider is GitHub Container Registry
elif [ "$PROVIDER" = "github" ]; then
  echo "$TOKEN" | docker login ghcr.io -u "$GITHUB_USERNAME" --password-stdin

# Check if provider is Azure
elif [ "$PROVIDER" = "azure" ]; then
  az login --service-principal -u "$AZURE_CLIENT_ID" -p "$AZURE_CLIENT_SECRET" --tenant "$AZURE_TENANT_ID" 
  az acr login --name "$AZURE_REGISTRY_NAME"

else
  echo "Please pass a valid provider"
fi

# Execute the application
exec "/go/src/app/enigma" "$@"