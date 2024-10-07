package generate

import (
	"fmt"
	"os"
)

type EnvType string

const (
	DOCKER EnvType = "DOCKER"
)

var DOCKER_ENV_VARIABLES = [9]string{
	"DOCKER_TAG=",
	"NO_CACHE=",
	"BUILD_ARCHITECTURE=",
	"SCAN=",
	"DOCKERFILE_PATH=",
	"BUILD_ARGS=",
	"DOCKER_IMAGE=",
	"CLEANUP=",
	"GITHUB_REF_NAME=",
}

// GenerateEnigmaFile writes the environment variables to the specified output file.
func GenerateEnigmaFile(outputPath string, envType EnvType) error {
	// Create or open the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Loop through the environment variables and write each to the file
	if envType == "DOCKER" {
		for _, env := range DOCKER_ENV_VARIABLES {
			_, err := file.WriteString(env + "\n")
			if err != nil {
				return fmt.Errorf("failed to write to file: %w", err)
			}
		}
	}
	fmt.Println("Environment variables successfully written to", outputPath)
	return nil
}
