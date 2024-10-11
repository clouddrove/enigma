package cmd

import (
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "To publish",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		enigmaFile, _ := cmd.Flags().GetString("enigmafile")
		dockerFlag, _ := cmd.Flags().GetBool("d")
		if dockerFlag {
			loadDockerEnv(*&enigmaFile)
			docker.CreateBuildxInstance()
			docker.InstallBinfmt()
			docker.TagDockerImage()
			docker.PushDockerImage()
		}
	},
}
