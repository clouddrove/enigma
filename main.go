package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/clouddrove/enigma/generate"
	"github.com/clouddrove/enigma/modules/docker"
	"github.com/joho/godotenv"
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

func main() {
	enigmaFile := flag.String("enigma", ".enigma", "Path to the .enigma file")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: enigma [-enigma=<file>] <command>")
		fmt.Println("Commands: init, bake, publish, build-publish")
		os.Exit(1)
	}

	command := flag.Args()[0]

	switch command {
	case "init":
		err := generate.GenerateEnigmaFile("modules/docker", *enigmaFile)
		if err != nil {
			fmt.Printf("Error generating %s file: %v\n", *enigmaFile, err)
			os.Exit(1)
		}
	case "bake":
		loadDockerEnv(*enigmaFile)
		docker.InstallBinfmt()
		docker.BuildDockerImage()
		docker.ScanDockerImage()
	case "publish":
		loadDockerEnv(*enigmaFile)
		docker.InstallBinfmt()
		docker.TagDockerImage()
		docker.PushDockerImage()
	case "build-publish":
		docker.LoadEnvFromEnigma(*enigmaFile)
		docker.InstallBinfmt()
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
		fmt.Println("Commands: init, bake, publish, build-publish")
	}
}