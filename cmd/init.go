package cmd

import (
	"fmt"
	"github.com/clouddrove/enigma/pkg/generate"
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"init"},
	Short:   "To init the command",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if dockerFlag {
			err := generate.GenerateEnigmaFile("pkg/docker", *&enigmaFile)
			if err != nil {
				fmt.Printf("Error generating %s file: %v\n", *&enigmaFile, err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
