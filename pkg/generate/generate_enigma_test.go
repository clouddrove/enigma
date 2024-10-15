package generate

import (
	"bufio"
	"os"
	"testing"
)

func TestGenerateEnigmaFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "enigma_test_*.env")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Test case for DOCKER environment type
	t.Run("DOCKER environment", func(t *testing.T) {
		err := GenerateEnigmaFile(tempFile.Name(), DOCKER)
		if err != nil {
			t.Fatalf("GenerateEnigmaFile failed: %v", err)
		}

		// Read the content of the generated file
		file, err := os.Open(tempFile.Name())
		if err != nil {
			t.Fatalf("Failed to open generated file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineCount := 0
		for scanner.Scan() {
			line := scanner.Text()
			if lineCount >= len(DOCKER_ENV_VARIABLES) {
				t.Errorf("Too many lines in the generated file")
				break
			}
			expectedLine := DOCKER_ENV_VARIABLES[lineCount]
			if line != expectedLine {
				t.Errorf("Line %d: expected %q, got %q", lineCount+1, expectedLine, line)
			}
			lineCount++
		}

		if err := scanner.Err(); err != nil {
			t.Fatalf("Error reading generated file: %v", err)
		}

		if lineCount != len(DOCKER_ENV_VARIABLES) {
			t.Errorf("Expected %d lines, got %d", len(DOCKER_ENV_VARIABLES), lineCount)
		}
	})

	// Test case for invalid file path
	t.Run("Invalid file path", func(t *testing.T) {
		err := GenerateEnigmaFile("/invalid/path/file.env", DOCKER)
		if err == nil {
			t.Error("Expected an error for invalid file path, but got nil")
		}
	})
}
