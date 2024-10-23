package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "enigma",
	Short: "enigma - The comprehensive DevOps toolkit",
	Long: `Enigma is a tool designed to simplify the DevOps lifecycle, offering a seamless way to manage tools environments, 
	build, scan, and publish. Below is a quick guide to getting started with Enigma and using its core commands.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
Enigma is a tool designed to simplify the DevOps lifecycle, offering a seamless way to manage tools environments, 
        build, scan, and publish. Below is a quick guide to getting started with Enigma and using its core commands.

Usage:
  enigma [flags]
  enigma [command]

Available Commands:
  bake          To Bake the command
  bake-publish  To bake and publish
  completion    Generate the autocompletion script for the specified shell
  help          Help about any command
  init          To init the command
  publish       To publish

Flags:
      --enigmafile string   Path to the .enigma file (default ".enigma")
  -h, --help                help for enigma

Use "enigma [command] --help" for more information about a command.`)
	},
}

func addDockerFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("d", false, "use commands for docker")
}

func addHelmFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("hl", true, "use commands for helm")
}

func addFilenameForInitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String("f", "", "to init .enigma with custom filename")
}

func addEnigmaFileFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String("enigmafile", ".enigma", "Path to the .enigma file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	// Add to the root command
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(bakeCmd)
	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(bake_publishCmd)

	// Add Optional --enigmafile to the following commands
	addEnigmaFileFlag(bakeCmd)
	addEnigmaFileFlag(publishCmd)
	addEnigmaFileFlag(bake_publishCmd)

	// Add Optional --f to the following commands
	addFilenameForInitFlag(initCmd)

	// Add dockerflag --d to the following commands
	addDockerFlag(bakeCmd)
	addDockerFlag(publishCmd)
	addDockerFlag(bake_publishCmd)

	// Add helmFlag --h to the following commands
	addHelmFlag(bakeCmd)
	addHelmFlag(bake_publishCmd)
}
