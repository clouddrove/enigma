package docker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

// BuildDockerImage builds a Docker image based on environment variables.
// It supports dynamic build arguments, optional no-cache, and specifying the build architecture.
func BuildDockerImage() {
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerfilePath := os.Getenv("DOCKERFILE_PATH")
	noCache := os.Getenv("DOCKER_NO_CACHE") == "true"
	buildArgs := os.Getenv("DOCKER_BUILD_ARGS")
	dockerImage := os.Getenv("DOCKER_IMAGE")
	buildArchitecture := os.Getenv("DOCKER_BUILD_ARCHITECTURE")

	if dockerImage == "" {
		log.Fatalf("DOCKER_IMAGE environment variable is not set")
	}

	if dockerfilePath == "" {
		dockerfilePath = "Dockerfile"
	}

	args := []string{"build", "-f", dockerfilePath, "-t", dockerImage, "."}

	if noCache {
		args = append(args, "--no-cache")
	}

	if buildArchitecture != "" {
		platform := getPlatform(buildArchitecture)
		if platform != "" {
			args = append(args, "--platform", platform)
		}
	}

	if buildArgs != "" {
		for _, arg := range strings.Split(buildArgs, ",") {
			args = append(args, "--build-arg", strings.TrimSpace(arg))
		}
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Building Docker image: %s:%s\n", dockerImage, dockerTag)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error building Docker image: %v", err)
	}

	fmt.Println("Build complete.")
	TagDockerImage()
}

func getPlatform(architecture string) string {
	switch strings.ToLower(architecture) {
	case "amd64":
		return "linux/amd64"
	case "arm64":
		return "linux/arm64"
	case "arm":
		return "linux/arm/v7"
	default:
		log.Printf("Unsupported architecture: %s. Using default platform.", architecture)
		return ""
	}
}

// ScanDockerImage performs a security scan of the Docker image and saves the report in SARIF format.
// It uses the `docker scout` command to scan the image for vulnerabilities.
func ScanDockerImage() {
	scan := os.Getenv("DOCKER_SCAN")

	if scan != "true" {
		fmt.Println("SCAN is not set to true. Skipping Docker image scan.")
		return
	}

	dockerTag := os.Getenv("DOCKER_TAG")

	if dockerTag == "" {
		log.Fatalf("DOCKER_TAG environment variable is not set")
	}

	sarifFile := "sarif.output.json"

	cmd := exec.Command("docker", "scout", "cves", dockerTag, "--output", sarifFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Scanning Docker image:", dockerTag)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker scout scan: %v", err)
	}

	fmt.Println("Docker image scan complete.")
	fmt.Printf("Scan complete. Report saved to %s\n", sarifFile)
}

// TagDockerImage tags the built Docker image for the specified registry.
// It uses the `docker tag` command to apply the new tag to the image.
func TagDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	finalDockerTag := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	if dockerImage == "" || dockerTag == "" {
		log.Fatalf("DOCKER_IMAGE or DOCKER_TAG environment variable is not set")
	}

	fmt.Printf("Tagging Docker image: %s as %s\n", dockerImage, dockerTag)

	cmdTag := exec.Command("docker", "tag", dockerImage, finalDockerTag)
	cmdTag.Stdout = os.Stdout
	cmdTag.Stderr = os.Stderr

	if err := cmdTag.Run(); err != nil {
		log.Fatalf("Error tagging Docker image: %v", err)
	}

	fmt.Println("Docker image tagged successfully.")
}

// PushDockerImage pushes the tagged Docker image to the specified registry and optionally cleans up local images.
// It uses the `docker push` command to upload the image to the registry specified in DOCKER_TAG.
// Cleanup is performed by default or when explicitly set to "true". It's only disabled when set to "false".
func PushDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	cleanup := os.Getenv("DOCKER_CLEANUP")
	finalDockerTag := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	if dockerTag == "" {
		log.Fatalf("DOCKER_TAG environment variable is not set")
	}

	cmdPush := exec.Command("docker", "push", finalDockerTag)
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	fmt.Printf("Pushing Docker image: docker push %s\n", finalDockerTag)
	err := cmdPush.Run()
	if err != nil {
		log.Fatalf("Error pushing Docker image: %v", err)
	}
	fmt.Println("Docker image successfully pushed to the specified registry.")

	if cleanup != "false" {
		fmt.Println("Cleanup is enabled. Removing tagged image.")
		cmdRm := exec.Command("docker", "rmi", finalDockerTag)
		cmdRm.Stdout = os.Stdout
		cmdRm.Stderr = os.Stderr
		fmt.Printf("Removing tagged image: docker rmi %s\n", finalDockerTag)
		if err := cmdRm.Run(); err != nil {
			log.Printf("Error removing tagged image: %v", err)
		}
		fmt.Println("Cleanup complete.")
	} else {
		fmt.Println("Cleanup is disabled. Tagged image will not be removed.")
	}
}

// InstallBinfmt installs the binfmt support for multi-platform builds.
func InstallBinfmt() {
	fmt.Println("Installing binfmt for multi-platform builds...")

	cmd := exec.Command("docker", "run", "--privileged", "--rm", "tonistiigi/binfmt", "--install", "all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing binfmt: %v", err)
	}

	fmt.Println("binfmt installation complete.")
}

// LoadEnvFromEnigma loads environment variables from the .enigma file.
func LoadEnvFromEnigma(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("%s file not found. No variables set.\n", filename)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening %s file: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if isValidEnvVarKey(key) {
			value = strings.ReplaceAll(value, "${{ github.ref_name }}", os.Getenv("GITHUB_REF_NAME"))

			if err := os.Setenv(key, value); err != nil {
				log.Printf("Failed to set environment variable %s: %v", key, err)
			} else {
				fmt.Printf("Set %s to %s\n", key, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .enigma file: %v", err)
	}
}

func isValidEnvVarKey(key string) bool {
	if len(key) == 0 {
		return false
	}
	for i, r := range key {
		if (i == 0 && !isLetter(r)) || (i > 0 && !isLetterOrDigitOrUnderscore(r)) {
			return false
		}
	}
	return true
}

func CreateBuildxInstance() error {
	cmd := exec.Command("docker", "buildx", "create", "--use")

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create Buildx instance: %w\nOutput: %s", err, string(output))
	}

	fmt.Println("Buildx instance created and set to use.")
	return nil
}

func BuildDockerImageAndPublishMultiArch() {
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerfilePath := os.Getenv("DOCKERFILE_PATH")
	noCache := os.Getenv("DOCKER_NO_CACHE") == "true"
	buildArgs := os.Getenv("DOCKER_BUILD_ARGS")
	dockerImage := os.Getenv("DOCKER_IMAGE")

	if dockerImage == "" {
		log.Fatalf("DOCKER_IMAGE environment variable is not set")
	}

	if dockerfilePath == "" {
		dockerfilePath = "Dockerfile"
	}

	fullImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	args := []string{"buildx", "build", "--push", "-f", dockerfilePath, "-t", fullImageName, "--platform linux/amd64,linux/arm64", "."}

	if noCache {
		args = append(args, "--no-cache")
	}

	if buildArgs != "" {
		for _, arg := range strings.Split(buildArgs, ",") {
			args = append(args, "--build-arg", strings.TrimSpace(arg))
		}
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Building and Pushing Docker image: %s:%s\n", dockerImage, dockerTag)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error building Docker image: %v", err)
	}

	fmt.Println("Building and Pushing completed.")
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isLetterOrDigitOrUnderscore(r rune) bool {
	return isLetter(r) || unicode.IsDigit(r)
}
