package helm

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"unicode"
)

// InstallHelm installs Helm if it is not already installed.
func CheckHelmInstalled() {
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

func DoHelmTemplating() {
	chartPath := os.Getenv("HELM_CHART_PATH")

	if chartPath == "" {
		log.Fatalf("HELM_CHART_PATH environment variable is not set")
	}

	cmd := exec.Command("helm", "template", chartPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error in running helm template for the Helm Chart: %v", err)
	}

	fmt.Println(cmd.Stdout)

	fmt.Println("Helm chart templating complete.")
}

// DoInstallHelmChart deploys the Helm chart using `helm install`.
func DoInstallHelmChart() {
	chartPath := os.Getenv("HELM_CHART_PATH")
	chartName := os.Getenv("HELM_CHART_NAME")

	if chartPath == "" {
		log.Fatalf("HELM_CHART_PATH environment variable is not set")
	}

	if chartName == "" {
		log.Fatalf("HELM_CHART_NAME environment variable is not set")
	}

	cmd := exec.Command("helm", "install", chartName, chartPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error installing Helm Chart: %v", err)
	}

	fmt.Printf("Successfully installed Helm chart: %s\n", chartPath)

	fmt.Println("Helm chart deployment complete.")
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

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isLetterOrDigitOrUnderscore(r rune) bool {
	return isLetter(r) || unicode.IsDigit(r)
}

func isValidEnvVarKey(key string) bool {
	if key == "" {
		return false
	}

	runes := []rune(key)

	// First character must be a letter or underscore
	if !isLetter(runes[0]) {
		return false
	}

	// Rest of the characters must be letters, digits, or underscores
	for _, r := range runes[1:] {
		if !isLetterOrDigitOrUnderscore(r) {
			return false
		}
	}

	return true
}
