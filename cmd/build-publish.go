package cmd

import (
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
)

var bake_publishCmd = &cobra.Command{
	Use:   "bake-publish",
	Short: "To bake and publish",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		enigmaFile, _ := cmd.Flags().GetString("enigmafile")
		dockerFlag, _ := cmd.Flags().GetBool("d")
		if dockerFlag {
			loadDockerEnv(enigmaFile)
			docker.CreateBuildxInstance()
			docker.InstallBinfmt()
			docker.BuildDockerImage()
			docker.ScanDockerImage()
			docker.TagDockerImage()
			docker.PushDockerImage()
		}
	},
}
