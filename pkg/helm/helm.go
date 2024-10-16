package helm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

// InstallHelm installs Helm if it is not already installed.
func InstallHelm() {
	fmt.Println("Checking Helm installation...")
	cmd := exec.Command("helm", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Helm is not installed or not found in the path: %v", err)
	}
	fmt.Println("Helm is installed.")
}

// BuildHelmChart builds a Helm chart package from the provided chart directory.
func BuildHelmChart() {
	chartPath := os.Getenv("HELM_CHART_PATH")
	if chartPath == "" {
		log.Fatalf("HELM_CHART_PATH environment variable is not set")
	}

	args := []string{"package", chartPath}

	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Building Helm chart from path: %s\n", chartPath)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error building Helm chart: %v", err)
	}

	fmt.Println("Helm chart build complete.")
	TagHelmChart()
}

// LintHelmChart lints the Helm chart to check for issues using `helm lint`.
func LintHelmChart() {
	chartPath := os.Getenv("HELM_CHART_PATH")
	if chartPath == "" {
		log.Fatalf("HELM_CHART_PATH environment variable is not set")
	}

	cmd := exec.Command("helm", "lint", chartPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Linting Helm chart: %s\n", chartPath)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Helm chart linting failed: %v", err)
	}

	fmt.Println("Helm chart linting complete. No issues found.")
}

// ScanHelmChart scans the Helm chart using `helm lint` for chart issues.
func ScanHelmChart() {
	chartPath := os.Getenv("HELM_CHART_PATH")

	if chartPath == "" {
		log.Fatalf("HELM_CHART_PATH environment variable is not set")
	}

	cmd := exec.Command("helm", "lint", chartPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Scanning Helm chart: %s\n", chartPath)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running Helm lint: %v", err)
	}

	fmt.Println("Helm chart scan complete.")
}

// TagHelmChart tags the Helm chart with a specific version or label.
func TagHelmChart() {
	chartName := os.Getenv("HELM_CHART_NAME")
	chartVersion := os.Getenv("HELM_CHART_VERSION")
	finalChartTag := fmt.Sprintf("%s-%s.tgz", chartName, chartVersion)

	if chartName == "" || chartVersion == "" {
		log.Fatalf("HELM_CHART_NAME or HELM_CHART_VERSION environment variable is not set")
	}

	fmt.Printf("Tagging Helm chart: %s with version: %s\n", chartName, chartVersion)

	// Typically, tagging in Helm is done using versioning, so this is a placeholder
	// In Helm, packages are tarred with the version.
	fmt.Printf("Helm chart tagged as %s\n", finalChartTag)
}

// PushHelmChart pushes the Helm chart to the specified Helm registry.
func PushHelmChart() {
	chartVersion := os.Getenv("HELM_CHART_VERSION")
	chartName := os.Getenv("HELM_CHART_NAME")
	registry := os.Getenv("HELM_REGISTRY")
	chartTag := fmt.Sprintf("%s-%s.tgz", chartName, chartVersion)

	if chartName == "" || chartVersion == "" {
		log.Fatalf("HELM_CHART_NAME or HELM_CHART_VERSION environment variable is not set")
	}

	if registry == "" {
		log.Fatalf("HELM_REGISTRY environment variable is not set")
	}

	cmdPush := exec.Command("helm", "push", chartTag, registry)
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr

	fmt.Printf("Pushing Helm chart: %s to registry: %s\n", chartTag, registry)
	if err := cmdPush.Run(); err != nil {
		log.Fatalf("Error pushing Helm chart: %v", err)
	}

	fmt.Println("Helm chart pushed successfully.")
}

// LoadEnvFromHelmFile loads environment variables from a Helm-specific .helm file.
func LoadEnvFromHelmFile(filename string) {
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
			if err := os.Setenv(key, value); err != nil {
				log.Printf("Failed to set environment variable %s: %v", key, err)
			} else {
				fmt.Printf("Set %s to %s\n", key, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .helm file: %v", err)
	}
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isLetterOrDigitOrUnderscore(r rune) bool {
	return isLetter(r) || unicode.IsDigit(r)
}
