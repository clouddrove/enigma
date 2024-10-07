package cmd

import (
	"fmt"
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
	"os"
)

var build_publishCmd = &cobra.Command{
	Use:     "build-publish",
	Aliases: []string{"bdpb", "build-p"},
	Short:   "To build and publish",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		dockerFlag, _ := cmd.Flags().GetBool("d")
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
