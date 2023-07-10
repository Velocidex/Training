package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func ParseCourse() (*Course, error) {
	root := &Course{}
	serialized, err := readFile("toc.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(serialized, root)
	if err != nil {
		return nil, fmt.Errorf("While processing %v: %w", "toc.yaml", err)
	}

	for _, module := range root.Modules {
		// Load each module
		serialized, err := readFile(filepath.Join(module.Path, "toc.yaml"))
		if err != nil {
			return nil, err
		}

		tmp := &Module{}
		err = yaml.Unmarshal(serialized, tmp)
		if err != nil {
			return nil, fmt.Errorf("While processing %v: %w", module.Path, err)
		}

		module.Topics = nil

		for _, topic := range tmp.Topics {

			// If a topic is an absolute path then it is a reference
			// to another module.
			if filepath.IsAbs(topic.Path) {
				serialized, err := readFile(filepath.Join(topic.Path, "toc.yaml"))
				if err != nil {
					return nil, err
				}

				tmp := &Module{}
				err = yaml.Unmarshal(serialized, tmp)
				if err != nil {
					return nil, fmt.Errorf("While processing %v: %w", module.Path, err)
				}
				for idx, sub_topic := range tmp.Topics {
					if idx == 0 && sub_topic.Name == "" {
						sub_topic.Name = topic.Name
					}
					sub_topic.Path = filepath.Join(topic.Path, sub_topic.Path)
					module.Topics = append(module.Topics, sub_topic)
				}

			} else {
				module.Topics = append(module.Topics, topic)
			}
		}

		// Load the slides for each topic
		for _, topic := range module.Topics {
			md_path := topic.Path
			if !filepath.IsAbs(md_path) {
				md_path = filepath.Join(module.Path, topic.Path)
			}

			topic.AbsPath = md_path
			topic.Link = strings.TrimSuffix(md_path, ".md") + ".html"
			topic.Link = strings.TrimPrefix(topic.Link, "/presentations")

			data, err := readFile(md_path)
			if err != nil {
				return nil, err
			}

			parts := strings.Split(string(data), "---")
			for idx, part := range parts {
				slide := &Slide{
					Title: getHeading(part),
					Index: idx,
				}
				for _, hit := range asset_regex.FindAllStringSubmatch(part, -1) {
					slide.Assets = append(slide.Assets, hit[1])
				}

				for _, hit := range asset_regex2.FindAllStringSubmatch(part, -1) {
					slide.Assets = append(slide.Assets, hit[1])
				}

				topic.Slides = append(topic.Slides, slide)
			}
		}
	}

	return root, nil
}

// Assemble the main TOC page.
func buildCourseTOC(course *Course) string {
	res := strings.ReplaceAll(`
<html lang="en" data-bs-theme="light">
<head>
   <meta charset="utf-8">
   <meta name="viewport" content="width=device-width, initial-scale=1">
   <link href="{Base}/css/fontawesome-all.min.css?1688344844" rel="stylesheet">
   <link href="{Base}/css/toc.css?1688344844" rel="stylesheet">
   <link href="{Base}/css/bootstrap.min.css" rel="stylesheet">
   <script src="{Base}/js/jquery-3.3.1.min.js?1688344844"></script>
   <script src="{Base}/js/toc.js?1688344844"></script>
   <script src="{Base}/js/bootstrap.bundle.min.js"></script>
</head>
<body>

<div class="px-4 py-5 my-5 text-center">
    <img class="d-block mx-auto mb-4"
         src="{Base}/css/velo_workshop.svg"
         alt="" height="100px">
    <h1 class="display-5 fw-bold text-body-emphasis">
      Velociraptor: Digging Deeper
    </h1>
    <div class="col-lg-6 mx-auto">
      <p class="lead mb-4">
        Welcome to the Velociraptor: Digging Deeper Course.
      </p>
      <div class="d-grid gap-2 d-sm-flex justify-content-sm-center">
        <button type="button" onClick="toggleAll()"
              class="btn btn-primary">Toggle All</button>

        <button type="button" onClick="setTheme('light')" id="light-btn"
              class="btn btn-primary">Light</button>

        <button type="button" onClick="setTheme('dark')" id="dark-btn"
              class="btn btn-primary">Dark</button>

      </div>
    </div>
  </div>

<div class="container">
  <ul class="toc">
`, "{Base}", ".")

	for _, module := range course.Modules {
		module_path := path.Clean(module.Path)
		res += fmt.Sprintf(`
       <li class="toc_close fs-2">
          <span onClick="toggleLeaf(this)">
             <i class="fa fa-sm category-icon fa-angle-right"></i>
          </span>
          <a href="./%v/index.html">
           %v
         </a>
         <a class="btn btn-link print-link" role="button"
                 href="./%v/index.html?print-pdf">
            <i class="fa fa-sm fa-print"></i>
         </a>
         %v
       </li>
`,
			module_path, module.Name,
			module_path, getCourseTopics(module))
	}

	res += `
  </ul>
</div>

</body>
</html>
`

	return res

}

// Provide direct links to specific slides inside one of the md files.
func getCourseSlides(topic *Topic, topic_link string) string {
	res := `
 <ul>
`
	last_title := ""
	for _, slide := range topic.Slides {
		if slide.Title == last_title || slide.Title == "" {
			continue
		}
		last_title = slide.Title

		res += fmt.Sprintf(`
       <li class="fs-4">
         <i class="fa fa-sm category-icon fa-chalkboard"></i>
         <a href="./%v#/%v" >
           %v
         </a>
       </li>
`, topic_link, slide.Index, slide.Title)
	}
	res += `
 </ul>
`
	return res
}

// A topic is inside a single .md file containing a self contained
// group of slides.
func getCourseTopics(module *Module) string {
	res := `
 <ul>
`
	for _, topic := range module.Topics {
		if topic.Path == "index.md" {
			continue
		}

		res += fmt.Sprintf(`
       <li class="toc_close fs-3">
         <span onClick="toggleLeaf(this)">
            <i class="fa fa-sm category-icon fa-angle-right"></i>
         </span>
         <a href="./%v">
           %v
         </a>
         %v
       </li>
`, topic.Link, topic.Name, getCourseSlides(topic, topic.Link))
	}
	res += `
 </ul>
`
	return res
}

func readFile(path string) ([]byte, error) {
	path = "./" + path
	fd, err := os.Open(path)
	if err != nil {
		fd, err = os.Open("./presentations/" + path)
		if err != nil {
			return nil, nil
		}
	}
	defer fd.Close()

	return ioutil.ReadAll(fd)
}
