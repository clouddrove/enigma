## To use Enigma for Docker

1. Clone the repository
`git clone https://github.com/clouddrove/enigma.git`

2. Set your values in `.enigma` file which is in `modules/docker`

3. Add Dockerfile of your in root of the folder

4. Now from root of the folder run:

At first run-
```
go build -o enigma main.go
```

### To work with Docker commands run-
- To Build, Scan and Tag:
  ```
  ./enigma bake
  ```
  
- To Build Docker Image:
  ```
  ./enigma docker/build
  ```
- To Run Docker Container:
  ```
  ./enigma docker/run
  ```
- To Stop Docker Container:
  ```
  ./enigma docker/stop
  ```
- To Remove Docker Container:
  ```
  ./enigma docker/remove
  ```
- To Remove Docker Image:
  ```
  ./enigma docker/remove-image
  ```