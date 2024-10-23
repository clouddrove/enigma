package generate

import (
	"bufio"
	"os"
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
