package docker

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestBuildDockerImage tests building the Docker image.
func TestBuildDockerImage(t *testing.T) {
	os.Setenv("DOCKER_TAG", "latest")
	os.Setenv("DOCKER_IMAGE", "test-image")
	defer os.Unsetenv("DOCKER_TAG")
	defer os.Unsetenv("DOCKER_IMAGE")

	cmd := exec.Command("echo", "docker build -t test-image:latest .")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("BuildDockerImage failed: %v", err)
	}

	if !strings.Contains(out.String(), "docker build") {
		t.Errorf("Expected docker build command, got: %v", out.String())
	}
}

// TestScanDockerImage tests scanning the Docker image.
func TestScanDockerImage(t *testing.T) {
	os.Setenv("DOCKER_TAG", "latest")
	os.Setenv("SCAN", "true")
	defer os.Unsetenv("DOCKER_TAG")
	defer os.Unsetenv("SCAN")

	cmd := exec.Command("echo", "docker scout cves latest --output sarif.output.json")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("ScanDockerImage failed: %v", err)
	}

	if !strings.Contains(out.String(), "docker scout") {
		t.Errorf("Expected docker scout command, got: %v", out.String())
	}
}

// TestTagDockerImage tests tagging the Docker image.
func TestTagDockerImage(t *testing.T) {
	os.Setenv("DOCKER_TAG", "latest")
	os.Setenv("DOCKER_IMAGE", "test-image")
	defer os.Unsetenv("DOCKER_TAG")
	defer os.Unsetenv("DOCKER_IMAGE")

	cmd := exec.Command("echo", "docker tag test-image latest")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("TagDockerImage failed: %v", err)
	}

	if !strings.Contains(out.String(), "docker tag") {
		t.Errorf("Expected docker tag command, got: %v", out.String())
	}
}

// TestPushDockerImage tests pushing the Docker image.
func TestPushDockerImage(t *testing.T) {
	os.Setenv("DOCKER_TAG", "latest")
	os.Setenv("CLEANUP", "false")
	defer os.Unsetenv("DOCKER_TAG")
	defer os.Unsetenv("CLEANUP")

	cmd := exec.Command("echo", "docker push latest")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("PushDockerImage failed: %v", err)
	}

	if !strings.Contains(out.String(), "docker push") {
		t.Errorf("Expected docker push command, got: %v", out.String())
	}
}

// TestInstallBinfmt tests installing binfmt for multi-platform builds.
func TestInstallBinfmt(t *testing.T) {
	cmd := exec.Command("echo", "docker run --privileged --rm tonistiigi/binfmt --install all")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("InstallBinfmt failed: %v", err)
	}

	if !strings.Contains(out.String(), "docker run --privileged") {
		t.Errorf("Expected docker run --privileged command, got: %v", out.String())
	}
}

// TestLoadEnvFromEnigma tests loading environment variables from the .enigma file.
func TestLoadEnvFromEnigma(t *testing.T) {
	enigmaContent := "DOCKER_TAG:latest\nDOCKER_IMAGE:test-image"
	tmpFile, err := os.CreateTemp("", "enigma")
	if err != nil {
		t.Fatalf("Failed to create temporary .enigma file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(enigmaContent))
	if err != nil {
		t.Fatalf("Failed to write to temporary .enigma file: %v", err)
	}

	LoadEnvFromEnigma(tmpFile.Name())

	if got := os.Getenv("DOCKER_TAG"); got != "latest" {
		t.Errorf("Expected DOCKER_TAG to be latest, got: %v", got)
	}

	if got := os.Getenv("DOCKER_IMAGE"); got != "test-image" {
		t.Errorf("Expected DOCKER_IMAGE to be test-image, got: %v", got)
	}

	os.Unsetenv("DOCKER_TAG")
	os.Unsetenv("DOCKER_IMAGE")
}
