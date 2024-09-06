package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"io/fs"
)

// GenerateEnigmaFile extracts environment variables from .go files in the specified directory and writes them to the specified output file path.
func GenerateEnigmaFile(dir string, outputPath string) error {
	envVars := make(map[string]bool)

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	err = filepath.WalkDir(absDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			re := regexp.MustCompile(`os\.Getenv\(["']([^"']+)["']\)`)
			matches := re.FindAllStringSubmatch(string(data), -1)

			for _, match := range matches {
				envVars[match[1]] = true
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %v", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating .enigma file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString("## Docker Variables\n")
	if err != nil {
		return fmt.Errorf("error writing to .enigma file: %v", err)
	}

	for envVar := range envVars {
		value := os.Getenv(envVar)
		file.WriteString(fmt.Sprintf("%s=%s\n", envVar, value))
	}

	fmt.Println(".enigma file generated.")
	return nil
}
