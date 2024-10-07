package cmd

import (
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
)

// This is to define -d functionality

var publishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"publish"},
	Short:   "To publish",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if dockerFlag {
			loadDockerEnv(*&enigmaFile)
			docker.InstallBinfmt()
			docker.TagDockerImage()
			docker.PushDockerImage()
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
