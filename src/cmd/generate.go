package cmd

import (
	"github.com/Velocidex/Training/src/generator"
	"github.com/spf13/cobra"
)

// Generates the HTML files for the site.
var (
	generateCmd = &cobra.Command{
		Use:   "generate [flags] output_directory",
		Short: "Generate the site",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generator.GenerateSite(args[0], Verbose)
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
}
