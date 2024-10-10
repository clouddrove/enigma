package cmd

import (
	"fmt"
	"github.com/clouddrove/enigma/pkg/generate"
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "To init the command",
	Long: `
Used to initialize the .enigmafile

--d - for docker
--f - for specifying the file.

`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, err := cmd.Flags().GetString("f")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var fileName string
		if fileFlag != "" {
			fileName = fmt.Sprintf(".enigma.%s", fileFlag)
		} else {
			fileName = ".enigma"
		}
		dockerFlag, err := cmd.Flags().GetBool("d")
		if err != nil {
			fmt.Println("Please specify the tool to generate .enigma file.")
			os.Exit(1)
		}
		if dockerFlag {
			err := generate.GenerateEnigmaFile(fileName, generate.DOCKER)
			if err != nil {
				fmt.Printf("Error generating %s file: %v\n", fileName, err)
				os.Exit(1)
			}
			fmt.Printf("Generated %s file for Docker.", fileName)
		}
	},
}
