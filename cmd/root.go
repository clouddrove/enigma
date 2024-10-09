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
	//PersistentPreRun: func(cmd *cobra.Command, args []string) {
	//	// This will run before any subcommand
	//	fmt.Printf("Using enigma file: %s\n", enigmaFile)
	//},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing root command")
	},
}

func addDockerFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("d", true, "Init for dockerfile")
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
	// Add --enigmafile flag to rootCmd [usuage] enigma --enigmafile [default .enigma ] can be set to any new.
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(bakeCmd)
	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(build_publishCmd)

	// Add Optional --enigmafile to the following commands
	addEnigmaFileFlag(bakeCmd)
	addEnigmaFileFlag(publishCmd)
	addEnigmaFileFlag(build_publishCmd)

	// Add Optional --f to the following commands
	addFilenameForInitFlag(initCmd)

	// Add dockerflag --d to the following commands
	addDockerFlag(initCmd)
	addDockerFlag(bakeCmd)
	addDockerFlag(publishCmd)
	addDockerFlag(build_publishCmd)
}
