package generate

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateEnigmaFile(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		outputPath  string
		wantErr     bool
		expectedLen int // Expected number of lines in the output file
	}{
		{
			name:        "Valid output path",
			outputPath:  "test_output.env",
			wantErr:     false,
			expectedLen: len(DOCKER_ENV_VARIABLES) + len(HELM_ENV_VARIABLES),
		},
		{
			name:        "Invalid directory path",
			outputPath:  "/nonexistent/directory/test.env",
			wantErr:     true,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the function
			err := GenerateEnigmaFile(tt.outputPath)

			// Check if error matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateEnigmaFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If we expect an error, no need to check the file
			if tt.wantErr {
				return
			}

			// Open and read the generated file
			file, err := os.Open(tt.outputPath)
			if err != nil {
				t.Fatalf("Failed to open generated file: %v", err)
			}
			defer file.Close()

			// Count lines and verify content
			scanner := bufio.NewScanner(file)
			lineCount := 0
			for scanner.Scan() {
				lineCount++
				line := scanner.Text()
				if line == "" {
					t.Errorf("Line %d is empty", lineCount)
				}
			}

			if err := scanner.Err(); err != nil {
				t.Fatalf("Error reading file: %v", err)
			}

			// Verify the number of lines matches expected
			if lineCount != tt.expectedLen {
				t.Errorf("Expected %d lines, got %d", tt.expectedLen, lineCount)
			}

			// Cleanup test file
			os.Remove(tt.outputPath)
		})
	}
}

func TestGenerateEnigmaFileContent(t *testing.T) {
	testPath := "test_content.env"

	// Generate the file
	err := GenerateEnigmaFile(testPath)
	if err != nil {
		t.Fatalf("Failed to generate file: %v", err)
	}
	defer os.Remove(testPath)

	// Read the generated file
	content, err := os.ReadFile(testPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Convert content to string
	contentStr := string(content)

	// Check if all Docker variables are present
	for _, env := range DOCKER_ENV_VARIABLES {
		if !contains(contentStr, env) {
			t.Errorf("Docker variable %s not found in generated file", env)
		}
	}

	// Check if all Helm variables are present
	for _, env := range HELM_ENV_VARIABLES {
		if !contains(contentStr, env) {
			t.Errorf("Helm variable %s not found in generated file", env)
		}
	}
}

// Helper function to check if a string contains another string
func contains(content, substring string) bool {
	return len(content) >= len(substring) && content[0:len(content)-(len(content)-len(substring))] == substring
}

func TestGenerateEnigmaFilePermissions(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "test_permissions.env")

	// Generate the file
	err := GenerateEnigmaFile(testPath)
	if err != nil {
		t.Fatalf("Failed to generate file: %v", err)
	}

	// Check file permissions
	info, err := os.Stat(testPath)
	if err != nil {
		t.Fatalf("Failed to stat generated file: %v", err)
	}

	// Check if file permissions are as expected (readable and writable)
	expectedPerm := os.FileMode(0644)
	if info.Mode().Perm() != expectedPerm {
		t.Errorf("Expected file permissions %v, got %v", expectedPerm, info.Mode().Perm())
	}
}
