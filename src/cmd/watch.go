package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/Velocidex/Training/src/generator"
	"github.com/radovskyb/watcher"
	"github.com/spf13/cobra"
)

// Generates the HTML files for the site.
var (
	watchCmd = &cobra.Command{
		Use:   "watch [flags] output_directory",
		Short: "Watch the repo for change and regenerate the site",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			w := watcher.New()
			w.SetMaxEvents(1)

			go func() {
				for {
					select {
					case event := <-w.Event:
						fmt.Printf("Detected change: %v\n", event)
						err := generator.GenerateSite(args[0], Verbose)
						if err != nil {
							log.Fatalln(err)
						}
					case err := <-w.Error:
						log.Fatalln(err)
					case <-w.Closed:
						return
					}
				}
			}()

			// Watch this folder for changes.
			err := w.AddRecursive("./modules")
			if err != nil {
				return err
			}

			return w.Start(time.Second)
		},
	}
)

func init() {
	rootCmd.AddCommand(watchCmd)
}
