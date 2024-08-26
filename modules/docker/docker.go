package docker

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// BuildDockerImage builds a Docker image based on the environment variables.
// It uses the `docker build` command to create the image with a specified tag.
func BuildDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	cmd := exec.Command("docker", "build", "-t", dockerImageName, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building Docker image:", dockerImageName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker build: %v", err)
	}

	fmt.Println("Build and Tag complete.")
}

// RunDockerContainer runs a Docker container from a specified image.
// It uses the `docker run` command with options to run in detached mode and map ports.
func RunDockerContainer() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)
	containerName := os.Getenv("CONTAINER_NAME")
	hostPort := os.Getenv("HOST_PORT")
	containerPort := os.Getenv("CONTAINER_PORT")

	cmd := exec.Command("docker", "run", "-d", "-p", fmt.Sprintf("%s:%s", hostPort, containerPort), "--name", containerName, dockerImageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running Docker container from image:", dockerImageName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker container: %v", err)
	}

	fmt.Println("Container is running.")
}

// StopDockerContainer stops a running Docker container.
// It uses the `docker stop` command to stop the container by name.
func StopDockerContainer() {
	containerName := os.Getenv("CONTAINER_NAME")

	cmd := exec.Command("docker", "stop", containerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Stopping Docker container:", containerName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error stopping docker container: %v", err)
	}

	fmt.Println("Container stopped.")
}

// RemoveDockerContainer removes a stopped Docker container.
// It uses the `docker rm` command to delete the container by name.
func RemoveDockerContainer() {
	containerName := os.Getenv("CONTAINER_NAME")

	cmd := exec.Command("docker", "rm", containerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Removing Docker container:", containerName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error removing docker container: %v", err)
	}

	fmt.Println("Container removed.")
}

// RemoveDockerImage removes a Docker image.
// It uses the `docker rmi` command to delete the image by name and tag.
func RemoveDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	cmd := exec.Command("docker", "rmi", dockerImageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Removing Docker image:", dockerImageName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error removing docker image: %v", err)
	}

	fmt.Println("Image removed.")
}

// ScanDockerImage performs a security scan of the Docker image and saves the report in SARIF format.
// It uses the `docker scout` command to scan the image for vulnerabilities.
func ScanDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	sarifFile := "sarif.output.json"

	cmd := exec.Command("docker", "scout", "cves", dockerImageName, "--output", sarifFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Scanning Docker image:", dockerImageName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker scout scan: %v", err)
	}

	fmt.Printf("Scan complete. Report saved to %s\n", sarifFile)
	
}