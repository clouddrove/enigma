package cmd

import (
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"publish"},
	Short:   "To publish",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		dockerFlag, _ := cmd.Flags().GetBool("d")
		if dockerFlag {
			loadDockerEnv(*&enigmaFile)
			docker.InstallBinfmt()
			docker.TagDockerImage()
			docker.PushDockerImage()
		}
	},
}
