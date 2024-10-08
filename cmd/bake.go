package cmd

import (
	"github.com/clouddrove/enigma/pkg/docker"
	"github.com/spf13/cobra"
)

// This is to define -d functionality

var bakeCmd = &cobra.Command{
	Use:   "bake",
	Short: "To Bake the command",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		enigmaFile, _ := cmd.Flags().GetString("enigmafile")
		dockerFlag, _ := cmd.Flags().GetBool("d")
		if dockerFlag {
			loadDockerEnv(enigmaFile)
			docker.InstallBinfmt()
			docker.BuildDockerImage()
			docker.ScanDockerImage()
		}
	},
}
