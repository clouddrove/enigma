package generate

import (
	"fmt"
	"os"
)

type EnvType string

const (
	DOCKER EnvType = "DOCKER"
	HELM   EnvType = "HELM"
)

var DOCKER_ENV_VARIABLES = []string{
	"DOCKERFILE_PATH=",
	"DOCKER_IMAGE=",
	"DOCKER_TAG=",
	"DOCKER_BUILD_ARCHITECTURE=",
	"DOCKER_SCAN=",
	"DOCKER_BUILD_ARGS=",
	"DOCKER_CLEANUP=",
	"DOCKER_MULTI_ARCH_BUILD=",
}

var HELM_ENV_VARIABLES = []string{
	"HELM_CHART_PATH=",
	"HELM_CHART_NAME=",
	"HELM_CHART_VERSION=",
	"HELM_REGISTRY=",
}

// GenerateEnigmaFile writes the environment variables to the specified output file.
func GenerateEnigmaFile(outputPath string) error {
	// Create or open the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write Docker environment variables
	for _, env := range DOCKER_ENV_VARIABLES {
		_, err := file.WriteString(env + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	// Write Helm environment variables
	for _, env := range HELM_ENV_VARIABLES {
		_, err := file.WriteString(env + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	fmt.Println("Environment variables successfully written to", outputPath)
	return nil
}
