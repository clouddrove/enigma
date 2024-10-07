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
		dockerFlag, _ := cmd.Flags().GetBool("d")
		if dockerFlag {
			err := generate.GenerateEnigmaFile(enigmaFile, generate.DOCKER)
			if err != nil {
				fmt.Printf("Error generating %s file: %v\n", *&enigmaFile, err)
				os.Exit(1)
			}
		}
	},
}
