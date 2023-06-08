package main

import (
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type OverlayFS struct {
	search_path []string
}

func (self OverlayFS) Open(name string) (fs.File, error) {
	for _, prefix := range self.search_path {
		path := filepath.Join(prefix, name)
		res, err := os.Open(path)
		if err == nil {
			return res, nil
		}
	}
	return nil, errors.New("Not Found")
}

func main() {
	overlay := OverlayFS{
		search_path: []string{"./", "./presentations"},
	}
	fs := http.FileServer(http.FS(overlay))
	http.Handle("/", fs)

	log.Print("Listening on :1313...")
	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		log.Fatal(err)
	}
}
