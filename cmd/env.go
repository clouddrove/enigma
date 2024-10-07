package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadDockerEnv(enigmaFile string) {
	isCICD := os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != ""
	if !isCICD {
		err := godotenv.Load(enigmaFile)
		if err != nil {
			log.Fatalf("Error loading %s file: %v", enigmaFile, err)
		}
	} else {
		fmt.Println("Running in CI/CD environment; skipping .enigma file load.")
	}
}
