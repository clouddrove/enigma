package main

import (
	"fmt"
	"log"
	"os"

	"github.com/clouddrove/enigma/generate"
	"github.com/clouddrove/enigma/modules/docker"
	"github.com/joho/godotenv"
)

// loadDockerEnv loads environment variables from the .enigma file located in the docker module.
// This function sets up necessary environment variables for Docker operations.
func loadDockerEnv() {
	isCICD := os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != ""

	if !isCICD {
		err := godotenv.Load(".enigma")
		if err != nil {
			log.Fatalf("Error loading .enigma file: %v", err)
		}
	} else {
		fmt.Println("Running in CI/CD environment; skipping .enigma file load.")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: enigma <command>")
		fmt.Println("Commands: init, bake, publish")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		err := generate.GenerateEnigmaFile("modules/docker", ".enigma")
		if err != nil {
			fmt.Printf("Error generating .enigma file: %v\n", err)
			os.Exit(1)
		}
	case "bake":
		loadDockerEnv()
		docker.BuildDockerImage()
		docker.ScanDockerImage()
	case "publish":
		loadDockerEnv()
		docker.TagDockerImage()
		docker.PushDockerImage()
	case "build-publish":
		loadDockerEnv()
		docker.BuildDockerImage()
		docker.ScanDockerImage()
		if os.Getenv("PUBLISH") == "true" {
			docker.TagDockerImage()
			docker.PushDockerImage()
		} else {
			fmt.Println("Publish is set to false. Skipping publish step.")
		}
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Commands: init, bake, publish")
	}
}
