package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// Serve a static directory
var (
	output_directory string
	serve_port       int64

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "serve the current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			fs := http.FileServer(http.Dir(output_directory))
			http.Handle("/", fs)

			fmt.Printf("Listening on :%d while serving %v\n",
				serve_port, output_directory)
			return http.ListenAndServe(
				fmt.Sprintf(":%d", serve_port), nil)
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(
		&output_directory, "directory", "o", ".", "The directory to serve ")
	serveCmd.Flags().Int64VarP(
		&serve_port, "port", "p", 1313, "Port to serve ")
}
