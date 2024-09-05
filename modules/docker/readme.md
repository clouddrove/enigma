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
```

#### When working on local keep SCAN=true to scan your image and generate report for it.
#### If no Dockerfile path passed it will use the root one.
#### To Build on Different ARCHITECTURE pass it in BUILD_ARCHITECTURE variable. It supports `amd64`, `arm64` and `arm`

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
### This GitHub Action builds docker image, tags and pushes to the registry you want. 

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
        uses: clouddrove/enigma@v0.0.8
        with:
          command: bake
          DOCKER_IMAGE: ${{ env.DOCKER_IMAGE }}
          DOCKER_TAG: ${{ env.DOCKER_TAG }}
          AWS_ACCOUNT_ID: ${{ env.AWS_ACCOUNT_ID }}
          AWS_REGION: ${{ env.AWS_REGION }}

      - name: Publish Docker Image
        uses: clouddrove/enigma@v0.0.8
        with:
          command: publish
          DOCKER_IMAGE: ${{ env.DOCKER_IMAGE }}
          DOCKER_TAG: ${{ env.DOCKER_TAG }}
          AWS_ACCOUNT_ID: ${{ env.AWS_ACCOUNT_ID }}
          AWS_REGION: ${{ env.AWS_REGION }}
```
