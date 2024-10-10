package cmd

import (
	"fmt"
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
	"os"
)

var build_publishCmd = &cobra.Command{
	Use:   "build-publish",
	Short: "To build and publish",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		enigmaFile, _ := cmd.Flags().GetString("enigmafile")
		dockerFlag, _ := cmd.Flags().GetBool("d")
		if dockerFlag {
			loadDockerEnv(enigmaFile)
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
