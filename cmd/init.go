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

--f - for specifying the enigma file.
--d - to generate enigma file for docker.
--h - to generate enigma file for helm.

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

		helmFlag, err := cmd.Flags().GetBool("helm")
		if err != nil {
			fmt.Println("Please specify the tool to generate .enigma file.")
			os.Exit(1)
		}

		switch {
		case dockerFlag:
			err := generate.GenerateEnigmaFile(fileName, generate.DOCKER)
			if err != nil {
				fmt.Printf("Error generating %s file: %v\n", fileName, err)
				os.Exit(1)
			}
			fmt.Printf("Generated %s file for Docker.\n", fileName)

		case helmFlag:
			err := generate.GenerateEnigmaFile(fileName, generate.HELM)
			if err != nil {
				fmt.Printf("Error generating %s file: %v\n", fileName, err)
				os.Exit(1)
			}
			fmt.Printf("Generated %s file for Helm.\n", fileName)

		default:
			fmt.Println("Please specify a valid tool (Docker or Helm) to generate the .enigma file.")
			os.Exit(1)
		}
	},
}
