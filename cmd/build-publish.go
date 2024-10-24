package cmd

import (
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
	"os"
)

var bake_publishCmd = &cobra.Command{
	Use:   "bake-publish",
	Short: "To bake and publish",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		enigmaFile, _ := cmd.Flags().GetString("enigmafile")
		dockerFlag, _ := cmd.Flags().GetBool("d")
		if dockerFlag {
			if os.Getenv("DOCKER_MULTI_ARCH_BUILD") == "true" {
				loadDockerEnv(enigmaFile)
				docker.InstallBinfmt()
				docker.CreateBuildxInstance()
				docker.BuildDockerImageAndPublishMultiArch()
			} else {
				loadDockerEnv(enigmaFile)
				docker.InstallBinfmt()
				docker.BuildDockerImage()
				docker.ScanDockerImage()
				docker.TagDockerImage()
				docker.PushDockerImage()
			}
		}
	},
}
