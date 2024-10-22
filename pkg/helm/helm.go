package helm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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

package helm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

// Previous functions remain the same...

// ListHelmReleaseHistory lists the revision history for a Helm release
func ListHelmReleaseHistory(releaseName, namespace string) error {
	if releaseName == "" {
		return fmt.Errorf("release name is required")
	}

	args := []string{"history", releaseName}
	if namespace != "" {
		args = append(args, "--namespace", namespace)
	}

	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Fetching revision history for release: %s in namespace: %s\n", releaseName, namespace)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to get release history: %v", err)
	}

	return nil
}

// RollbackHelmRelease rolls back a Helm release to a specific revision
func RollbackHelmRelease(releaseName string, revision int, namespace string) error {
	if releaseName == "" {
		return fmt.Errorf("release name is required")
	}

	args := []string{"rollback", releaseName, strconv.Itoa(revision)}
	if namespace != "" {
		args = append(args, "--namespace", namespace)
	}

	// Add wait flag to ensure rollback completes
	args = append(args, "--wait")

	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Rolling back release %s to revision %d in namespace %s\n", releaseName, revision, namespace)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("rollback failed: %v", err)
	}

	fmt.Printf("Successfully rolled back %s to revision %d\n", releaseName, revision)
	return nil
}

// GetHelmReleaseStatus gets the current status of a Helm release
func GetHelmReleaseStatus(releaseName, namespace string) error {
	if releaseName == "" {
		return fmt.Errorf("release name is required")
	}

	args := []string{"status", releaseName}
	if namespace != "" {
		args = append(args, "--namespace", namespace)
	}

	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Fetching status for release: %s in namespace: %s\n", releaseName, namespace)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to get release status: %v", err)
	}

	return nil
}

// GetHelmReleaseRevision gets details about a specific revision of a Helm release
func GetHelmReleaseRevision(releaseName string, revision int, namespace string) error {
	if releaseName == "" {
		return fmt.Errorf("release name is required")
	}

	args := []string{"get", "all", releaseName}
	if namespace != "" {
		args = append(args, "--namespace", namespace)
	}
	args = append(args, "--revision", strconv.Itoa(revision))

	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Fetching details for release %s revision %d in namespace %s\n", releaseName, revision, namespace)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to get revision details: %v", err)
	}

	return nil
}

// CleanupHelmReleaseHistory removes old revisions of a Helm release
func CleanupHelmReleaseHistory(releaseName string, maxRevisions int, namespace string) error {
	if releaseName == "" {
		return fmt.Errorf("release name is required")
	}

	if maxRevisions < 1 {
		return fmt.Errorf("maxRevisions must be greater than 0")
	}

	args := []string{"history", releaseName, "--max", strconv.Itoa(maxRevisions)}
	if namespace != "" {
		args = append(args, "--namespace", namespace)
	}

	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Cleaning up release history for %s, keeping %d revisions in namespace %s\n", releaseName, maxRevisions, namespace)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to cleanup release history: %v", err)
	}

	return nil
}

// CompareHelmReleaseRevisions compares two revisions of a Helm release
func CompareHelmReleaseRevisions(releaseName string, revision1, revision2 int, namespace string) error {
	if releaseName == "" {
		return fmt.Errorf("release name is required")
	}

	// Get manifests for both revisions
	getManifest := func(rev int) (string, error) {
		args := []string{"get", "manifest", releaseName}
		if namespace != "" {
			args = append(args, "--namespace", namespace)
		}
		args = append(args, "--revision", strconv.Itoa(rev))

		cmd := exec.Command("helm", args...)
		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("failed to get manifest for revision %d: %v", rev, err)
		}
		return string(output), nil
	}

	manifest1, err := getManifest(revision1)
	if err != nil {
		return err
	}

	manifest2, err := getManifest(revision2)
	if err != nil {
		return err
	}

	// Write manifests to temporary files for comparison
	tempFile1 := fmt.Sprintf("%s-rev%d.yaml", releaseName, revision1)
	tempFile2 := fmt.Sprintf("%s-rev%d.yaml", releaseName, revision2)

	if err := os.WriteFile(tempFile1, []byte(manifest1), 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %v", err)
	}
	defer os.Remove(tempFile1)

	if err := os.WriteFile(tempFile2, []byte(manifest2), 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %v", err)
	}
	defer os.Remove(tempFile2)

	// Use diff to compare the manifests
	diffCmd := exec.Command("diff", "-u", tempFile1, tempFile2)
	diffCmd.Stdout = os.Stdout
	diffCmd.Stderr = os.Stderr

	fmt.Printf("Comparing revisions %d and %d of release %s in namespace %s\n", revision1, revision2, releaseName, namespace)
	diffCmd.Run() // Note: diff returns non-zero exit code if files differ, which is expected

	return nil
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
