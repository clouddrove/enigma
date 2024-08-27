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

    if dockerImage == "" || dockerTag == "" {
        log.Fatalf("DOCKER_IMAGE or DOCKER_TAG environment variable is not set")
    }

    // Build the image using the base name (e.g., "aws")
    cmd := exec.Command("docker", "build", "-t", dockerImage, ".")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    fmt.Println("Building Docker image:", dockerImage)

    err := cmd.Run()
    if err != nil {
        log.Fatalf("Error running docker build: %v", err)
    }

    fmt.Println("Build complete.")

    TagDockerImage()
}

// // ScanDockerImage performs a security scan of the Docker image and saves the report in SARIF format.
// // It uses the `docker scout` command to scan the image for vulnerabilities.
// func ScanDockerImage() {
//     dockerTag := os.Getenv("DOCKER_TAG")

//     if dockerTag == "" {
//         log.Fatalf("DOCKER_TAG environment variable is not set")
//     }

//     // sarifFile := "sarif.output.json"

//     cmd := exec.Command("docker", "scout", "cves", dockerTag)
//     cmd.Stdout = os.Stdout
//     cmd.Stderr = os.Stderr

//     fmt.Println("Scanning Docker image:", dockerTag)

//     err := cmd.Run()
//     if err != nil {
//         log.Fatalf("Error running docker scout scan: %v", err)
//     }
    
// }

// TagDockerImage tags the built Docker image for the specified registry.
// It uses the `docker tag` command to apply the new tag to the image.
func TagDockerImage() {
    dockerImage := os.Getenv("DOCKER_IMAGE")
    dockerTag := os.Getenv("DOCKER_TAG")
    
    if dockerImage == "" || dockerTag == "" {
        log.Fatalf("DOCKER_IMAGE or DOCKER_TAG environment variable is not set")
    }

    cmdTag := exec.Command("docker", "tag", dockerImage, dockerTag)
    cmdTag.Stdout = os.Stdout
    cmdTag.Stderr = os.Stderr

    fmt.Printf("Tagging Docker image: %s as %s\n", dockerImage, dockerTag)

    err := cmdTag.Run()
    if err != nil {
        log.Fatalf("Error tagging docker image: %v", err)
    }

    fmt.Println("Docker image tagged successfully.")
}


// PushDockerImage pushes the tagged Docker image to the specified registry and optionally cleans up local images.
// It uses the `docker push` command to upload the image to the registry specified in DOCKER_TAG.
// Cleanup is performed by default or when explicitly set to "true". It's only disabled when set to "false".
func PushDockerImage() {
    dockerTag := os.Getenv("DOCKER_TAG")
    cleanup := os.Getenv("CLEANUP")
    
    if dockerTag == "" {
        log.Fatalf("DOCKER_TAG environment variable is not set")
    }
    
    cmdPush := exec.Command("docker", "push", dockerTag)
    cmdPush.Stdout = os.Stdout
    cmdPush.Stderr = os.Stderr
    fmt.Printf("Pushing Docker image: docker push %s\n", dockerTag)
    err := cmdPush.Run()
    if err != nil {
        log.Fatalf("Error pushing docker image: %v", err)
    }
    fmt.Println("Docker image successfully pushed to the specified registry.")
    
    if cleanup != "false" {
        fmt.Println("Cleanup is enabled. Removing tagged image.")
        cmdRm := exec.Command("docker", "rmi", dockerTag)
        cmdRm.Stdout = os.Stdout
        cmdRm.Stderr = os.Stderr
        fmt.Printf("Removing tagged image: docker rmi %s\n", dockerTag)
        if err := cmdRm.Run(); err != nil {
            log.Printf("Error removing tagged image: %v", err)
        }
        fmt.Println("Cleanup complete.")
    } else {
        fmt.Println("Cleanup is disabled. Tagged image will not be removed.")
    }
}