## To use Enigma for Docker

1. Clone the repository
`git clone https://github.com/clouddrove/enigma.git`

2. Set your values in `.enigma` file. To set values in `.enigma` just pass the values. This is an sample-
```
DOCKER_IMAGE=nginx
DOCKER_TAG=xyz:v1
CLEANUP=true
```

3. Add Dockerfile of your in root of the folder

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