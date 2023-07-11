package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	Verbose bool

	rootCmd = &cobra.Command{
		Use:   "course",
		Short: "Manage the Velociraptor: Digging Deeper course",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose,
		"verbose", "v", false, "verbose output")
}
