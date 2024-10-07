package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "enigma",
	Short: "enigma - The comprehensive DevOps toolkit",
	Long: `enigma is the tool that helps you 
   
One can use stringer to modify or inspect strings straight from the terminal`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// This will run before any subcommand
		fmt.Printf("Using enigma file: %s\n", enigmaFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing root command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	// Add --enigmafile flag to rootCmd
	rootCmd.PersistentFlags().StringVar(&enigmaFile, "enigmafile", ".enigma", "Path to the .enigma file")
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().Bool("d", true, "Init for dockerfile")
}
