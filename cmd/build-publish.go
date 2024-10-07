package cmd

import (
	"fmt"
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
	"os"
)

// This is to define -d functionality

var build_publishCmd = &cobra.Command{
	Use:     "build-publish",
	Aliases: []string{"bdpb", "build-p"},
	Short:   "To build and publish",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if dockerFlag {
			docker.LoadEnvFromEnigma(*&enigmaFile)
			docker.InstallBinfmt()
			docker.BuildDockerImage()
			docker.ScanDockerImage()
			if os.Getenv("PUBLISH") == "true" {
				docker.TagDockerImage()
				docker.PushDockerImage()
			} else {
				fmt.Println("Publish is set to false. Skipping publish step.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
