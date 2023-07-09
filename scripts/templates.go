package main

import (
	"fmt"
	"strings"
)

const (
	index_html = `
<!doctype html>
<html lang="en">

    <head>
        <meta charset="utf-8">

        <title>Velociraptor Deployment</title>

        <link rel="stylesheet" href="{Base}/dist/reveal.css">
        <link rel="stylesheet" href="{Base}/dist/theme/serif.css" id="theme">
        <link rel="stylesheet" href="{Base}/css/velo.css">
        <link rel="stylesheet" href="{Base}/plugin/highlight/vs.css">
    </head>
    <body>
        <div class="reveal">
            <div class="slides">
%v
            </div>
        </div>
        <script src="{Base}/dist/reveal.js"></script>
        <script src="{Base}/plugin/markdown/markdown.js"></script>
        <script src="{Base}/plugin/highlight/highlight.js"></script>
        <script src="{Base}/plugin/notes/notes.js"></script>
        <script src="{Base}/plugin/zoom/zoom.js"></script>
        <script src="{Base}/js/jquery-3.3.1.min.js?1688344844"></script>
        <script src="{Base}/js/slides.js"></script>
        <script>
            Reveal.initialize({
                controls: true,
                progress: true,
                history: false,
                hash: true,
                center: false,
                slideNumber: true,

                plugins: [ RevealMarkdown, RevealHighlight, RevealNotes, RevealZoom ]
            }).then(initializeSlides);

        </script>

    </body>
</html>
`
	section_template = `
<section data-markdown
  data-transition="fade"
  data-separator="^---+\n\n"
  data-separator-vertical="^>+\n\n">
<textarea data-template>
%v
</textarea>
</section>
`
)

func buildIndexHtml(module *Module) string {
	sections := ""
	for _, topic := range module.Topics {
		data, err := readFile("./" + topic.AbsPath)
		if err != nil {
			continue
		}
		sections += fmt.Sprintf(section_template,
			adjustSectionText(string(data)))
	}

	index := adjustSectionText(index_html)
	return fmt.Sprintf(index, sections)
}

func adjustSectionText(in string) string {
	in = strings.ReplaceAll(in, "/modules/", "../../modules/")
	return strings.ReplaceAll(in, "{Base}", "../..")
}
