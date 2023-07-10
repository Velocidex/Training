package main

import (
	"fmt"
	"net/http"
)

func ServeStatic(output_directory string, port int) error {
	fs := http.FileServer(http.Dir(output_directory))
	http.Handle("/", fs)

	fmt.Printf("Listening on :%d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return err
	}

	return nil
}
