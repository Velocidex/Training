package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jessevdk/go-flags"
)

var (
	opts struct {
		Verbose  bool `short:"v" long:"verbose" description:"Show verbose debug information"`
		Generate *struct {
			Positional struct {
				Output string
			} `positional-args:"true" required:"true"`
		} `command:"generate"`

		Serve *struct {
			Port      int    `default:"1313"`
			Directory string `default:"."`
		} `command:"serve"`
	}

	copy_regex   = regexp.MustCompile("\\.(png|md|css|js|svg|woff2|ttf|woff|ttf|gif)$")
	asset_regex  = regexp.MustCompile(`src="([^"]+)"`)
	asset_regex2 = regexp.MustCompile(`!\[\]\(([^\)]+)\)`)
)

func getHeading(part string) string {
	for _, line := range strings.Split(part, "\n") {
		if strings.HasPrefix(line, "#") {
			return strings.Trim(line, "# ")
		}
	}
	return ""
}

func doIt(output_directory string, verbose bool) error {
	course, err := ParseCourse()
	if err != nil {
		return err
	}

	fmt.Printf("Loading course with %v\n", Stats(course))

	// Prepare the skeleton
	output_manager := OutputManager{output_directory, verbose}
	output_manager.CopyDirectory("./presentations/dist/", "dist")
	output_manager.CopyDirectory("./presentations/plugin/", "plugin")
	output_manager.CopyDirectory("./presentations/plugin/highlight", "plugin/highlight")
	output_manager.CopyDirectory("./presentations/plugin/markdown", "plugin/markdown")
	output_manager.CopyDirectory("./presentations/plugin/notes", "plugin/notes")
	output_manager.CopyDirectory("./presentations/plugin/zoom", "plugin/zoom")
	output_manager.CopyDirectory("./presentations/themes/workshop/", "themes/workshop")
	output_manager.CopyDirectory("./presentations/dist/theme", "dist/theme")
	output_manager.CopyDirectory("./presentations/resources", "resources")
	output_manager.CopyDirectory("./webfonts/", "webfonts")
	output_manager.CopyDirectory("./css", "css")
	output_manager.CopyDirectory("./js", "js")

	if verbose {
		Dump(course)
	}

	err = output_manager.WriteFile("index.html", buildCourseTOC(course))
	if err != nil {
		return err
	}

	// Copy module directories into the output
	for _, module := range course.Modules {
		// Create a HTML for the whole module
		output_manager.WriteFile(
			filepath.Join(module.Path, "index.html"),
			buildIndexHtml(module))

		// Copy all the files over
		output_manager.CopyDirectory("./"+module.Path, module.Path)

		// Check all the topics and merge them with this module.
		for _, topic := range module.Topics {
			// Copy the directory if it is absolute.
			if filepath.IsAbs(topic.Path) {
				output_manager.CopyDirectory("./"+filepath.Dir(topic.Path),
					filepath.Dir(topic.Link))
			}

			// Also create a html for each topic.
			if topic.Path != "index.md" {
				output_manager.WriteFile(topic.Link, buildIndexHtml(&Module{
					Topics: []*Topic{topic},
				}))
			}

			// Make sure all assets are copied over if they are not
			// already inside the module directory.
			for _, slide := range topic.Slides {
				for _, asset := range slide.Assets {
					if strings.HasPrefix(asset, "/") {
						output_manager.CopyFile(asset, asset)
					}
				}
			}
		}
	}

	return nil
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Generate != nil {
		err = doIt(opts.Generate.Positional.Output, opts.Verbose)

	} else if opts.Serve != nil {
		err = ServeStatic(opts.Serve.Directory, opts.Serve.Port)
	}

	if err != nil {
		log.Fatal(err)
	}
}
