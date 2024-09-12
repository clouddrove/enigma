## To use Enigma for Docker

1. Clone the repository
`git clone https://github.com/clouddrove/enigma.git`

2. Set your values in `.enigma` file. To set values in `.enigma` just pass the values. This is an sample-
```
DOCKER_IMAGE=nginx
DOCKER_TAG=xyz:v1
CLEANUP=true
SCAN=false
DOCKERFILE_PATH=test/Dockerfile
BUILD_ARCHITECTURE=amd64
NO_CACHE="true"        
BUILD_ARGS="APP_VERSION=0.0.0,APP_ENV=test"
SCAN=false
```

#### When working on local keep SCAN=true to scan your image and generate report for it.
#### If no Dockerfile path passed it will use the root one.
#### To Build on Different ARCHITECTURE pass it in BUILD_ARCHITECTURE variable. It supports `amd64`, `arm64` and `arm`
### Pass values in BUILD_ARGS to give args values in Dockerfile
### Set NO_CACHE=true for fresh Docker build without using cached layers.

3. Add your Dockerfile.

4. Now from root of the folder run:

At first run-
```
go build -o enigma main.go
```

### To work with Docker commands run-
- To Build and Tag:
  ```
  ./enigma bake
  ```
 
- To Push Image to Registry and cleanup Image at end(Cleanup will be only done if in `.enigma` CLEANUP is set true or by default it will take true):
  ```
  ./enigma publish
  ```

## Usage in GitHub Actions
### This GitHub Action builds docker image, tags and pushes to the registry you want. It also provies Build architecture feature to build on platform you want

```yaml
name: Enigma Docker

on:
  push:
    branches: main

jobs:
  login:
    runs-on: ubuntu-latest
    steps:
 
      - name: Build Docker Image
        uses: clouddrove/enigma@v0.0.10
        with:
          command: bake
          DOCKER_IMAGE: ${{ env.DOCKER_IMAGE }}
          DOCKER_TAG: ${{ env.DOCKER_TAG }}
          AWS_ACCOUNT_ID: ${{ env.AWS_ACCOUNT_ID }}
          AWS_REGION: ${{ env.AWS_REGION }}
          BUILD_ARCHITECTURE: ${{ inputs.BUILD_ARCHITECTURE }}

      - name: Publish Docker Image
        uses: clouddrove/enigma@v0.0.10
        with:
          command: publish
          DOCKER_IMAGE: ${{ env.DOCKER_IMAGE }}
          DOCKER_TAG: ${{ env.DOCKER_TAG }}
          AWS_ACCOUNT_ID: ${{ env.AWS_ACCOUNT_ID }}
          AWS_REGION: ${{ env.AWS_REGION }}
          BUILD_ARCHITECTURE: ${{ inputs.BUILD_ARCHITECTURE }}
```
