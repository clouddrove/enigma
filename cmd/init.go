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

		err = generate.GenerateEnigmaFile(fileName)
		if err != nil {
			fmt.Printf("Error generating %s file: %v\n", fileName, err)
			os.Exit(1)
		}
	},
}
