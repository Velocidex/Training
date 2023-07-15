package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// Serve a static directory
var (
	serve_port int64

	serveCmd = &cobra.Command{
		Use:   "serve [flags] directory",
		Short: "serve a directory as static HTML",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			output_directory := args[0]
			logger := log.New(os.Stdout, "http: ", log.LstdFlags)

			fs := http.FileServer(http.Dir(output_directory))
			http.Handle("/", logging(logger)(fs))

			fmt.Printf("Listening on :%d while serving %v\n",
				serve_port, output_directory)
			return http.ListenAndServe(
				fmt.Sprintf(":%d", serve_port), nil)
		},
	}
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &StatusRecorder{
				ResponseWriter: w,
				Status:         200,
			}
			next.ServeHTTP(w, r)
			logger.Println(r.Method, r.URL.Path, r.RemoteAddr,
				r.UserAgent(), recorder.Status)
		})
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().Int64VarP(
		&serve_port, "port", "p", 1313, "Port to serve ")
}
