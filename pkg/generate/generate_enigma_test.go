package generate

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestGenerateEnigmaFile_DOCKER tests the GenerateEnigmaFile function with a valid DOCKER envType.
func TestGenerateEnigmaFile_DOCKER(t *testing.T) {
	// Create a temporary file to write the environment variables
	tmpFile, err := ioutil.TempFile("", "enigma_docker_test")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the file after test

	// Call the GenerateEnigmaFile function with DOCKER envType
	err = GenerateEnigmaFile(tmpFile.Name(), DOCKER)
	if err != nil {
		t.Errorf("GenerateEnigmaFile failed: %v", err)
	}

	// Read the content of the file
	content, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	// Check if the content matches the expected DOCKER_ENV_VARIABLES
	expected := "DOCKER_TAG\nNO_CACHE\nBUILD_ARCHITECTURE\nSCAN\nDOCKERFILE_PATH\nBUILD_ARGS\nDOCKER_IMAGE\nCLEANUP\nGITHUB_REF_NAME\n"
	if string(content) != expected {
		t.Errorf("expected content %q, got %q", expected, string(content))
	}
}

// TestGenerateEnigmaFile_InvalidEnvType tests the GenerateEnigmaFile function with an invalid envType.
func TestGenerateEnigmaFile_InvalidEnvType(t *testing.T) {
	// Create a temporary file to write the environment variables
	tmpFile, err := ioutil.TempFile("", "enigma_invalid_test")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the file after test

	// Call the GenerateEnigmaFile function with an invalid envType (empty string in this case)
	err = GenerateEnigmaFile(tmpFile.Name(), EnvType("INVALID"))
	if err != nil {
		t.Errorf("GenerateEnigmaFile failed with invalid envType: %v", err)
	}

	// Read the content of the file
	content, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	// Check that no environment variables were written
	if len(content) != 0 {
		t.Errorf("expected empty file, got %q", string(content))
	}
}

// TestGenerateEnigmaFile_FileCreationError tests the GenerateEnigmaFile function's handling of file creation errors.
func TestGenerateEnigmaFile_FileCreationError(t *testing.T) {
	// Pass an invalid file path to induce an error
	invalidPath := "/invalid/path/enigma_docker_test"

	// Call the GenerateEnigmaFile function with DOCKER envType and an invalid file path
	err := GenerateEnigmaFile(invalidPath, DOCKER)
	if err == nil {
		t.Error("expected error when writing to an invalid file path, but got none")
	}
}
