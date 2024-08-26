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

// TagDockerImage tags a Docker image with a new tag for Docker Hub.
// It uses the `docker tag` command to apply the new tag to the image.
func TagDockerImage() {
    dockerImage := os.Getenv("DOCKER_IMAGE")
    dockerTag := os.Getenv("DOCKER_TAG")
    dockerHubUsername := os.Getenv("DOCKERHUB_USERNAME")
    dockerHubRepo := os.Getenv("DOCKERHUB_REPO")
    
    sourceImage := fmt.Sprintf("%s:%s", dockerImage, dockerTag)
    targetImage := fmt.Sprintf("%s/%s:%s", dockerHubUsername, dockerHubRepo, dockerTag)

    cmdTag := exec.Command("docker", "tag", sourceImage, targetImage)
    cmdTag.Stdout = os.Stdout
    cmdTag.Stderr = os.Stderr

    fmt.Printf("Tagging Docker image: docker tag %s %s\n", sourceImage, targetImage)

    err := cmdTag.Run()
    if err != nil {
        log.Fatalf("Error tagging docker image: %v", err)
    }

    fmt.Println("Docker image tagged successfully.")
}

// PushDockerImage pushes the tagged Docker image to Docker Hub and optionally cleans up local images.
// It uses the `docker push` command to upload the image to the specified Docker Hub repository.
// Cleanup is performed by default or when explicitly set to "true". It's only disabled when set to "false".
func PushDockerImage() {
    dockerImage := os.Getenv("DOCKER_IMAGE")
    dockerTag := os.Getenv("DOCKER_TAG")
    dockerHubUsername := os.Getenv("DOCKERHUB_USERNAME")
    dockerHubRepo := os.Getenv("DOCKERHUB_REPO")
    cleanup := os.Getenv("CLEANUP")

    localImage := fmt.Sprintf("%s:%s", dockerImage, dockerTag)
    dockerHubImage := fmt.Sprintf("%s/%s:%s", dockerHubUsername, dockerHubRepo, dockerTag)

    cmdPush := exec.Command("docker", "push", dockerHubImage)
    cmdPush.Stdout = os.Stdout
    cmdPush.Stderr = os.Stderr

    fmt.Printf("Pushing Docker image to Docker Hub: docker push %s\n", dockerHubImage)

    err := cmdPush.Run()
    if err != nil {
        log.Fatalf("Error pushing docker image: %v", err)
    }

    fmt.Println("Docker image successfully pushed to Docker Hub.")

    if cleanup != "false" {
        fmt.Println("Cleanup is enabled. Removing local images.")

        cmdRmHub := exec.Command("docker", "rmi", dockerHubImage)
        cmdRmHub.Stdout = os.Stdout
        cmdRmHub.Stderr = os.Stderr
        fmt.Printf("Removing Docker Hub tagged image: docker rmi %s\n", dockerHubImage)
        if err := cmdRmHub.Run(); err != nil {
            log.Printf("Error removing Docker Hub tagged image: %v", err)
        }

        cmdRmLocal := exec.Command("docker", "rmi", localImage)
        cmdRmLocal.Stdout = os.Stdout
        cmdRmLocal.Stderr = os.Stderr
        fmt.Printf("Removing originally built image: docker rmi %s\n", localImage)
        if err := cmdRmLocal.Run(); err != nil {
            log.Printf("Error removing originally built image: %v", err)
        }

        fmt.Println("Cleanup complete.")
    } else {
        fmt.Println("Cleanup is disabled. Local images will not be removed.")
    }
}