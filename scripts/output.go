package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type OutputManager struct {
	output_path string
}

func (self OutputManager) ensureFilePath(path string) (string, error) {
	output_path := filepath.Join(self.output_path, path)
	output_directory := filepath.Dir(output_path)
	err := os.MkdirAll(output_directory, 0775)

	return output_path, err
}

func (self OutputManager) WriteFile(path, data string) error {
	output_path, err := self.ensureFilePath(path)
	if err != nil {
		return err
	}

	fmt.Printf("WriteFile %s \n", output_path)

	out_fd, err := os.OpenFile(
		output_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer out_fd.Close()

	out_fd.Write([]byte(data))

	return nil
}

func (self OutputManager) CopyDirectory(src, dst string) error {
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Make sure the output directory exists
	output_directory := filepath.Join(self.output_path, dst)

	fmt.Printf("CopyDirectory %s -> %s\n", src, output_directory)
	err = os.MkdirAll(output_directory, 0775)
	if err != nil {
		return err
	}

	for _, f := range files {
		filename := f.Name()
		srcPath := filepath.Join(src, filename)
		destPath := filepath.Join(output_directory, filename)
		if copy_regex.MatchString(filename) {
			err := Copy(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (self OutputManager) CopyFile(src, dst string) error {
	// Make sure the output directory exists
	output_path := filepath.Join(self.output_path, dst)
	fmt.Printf("CopyFile %s -> %s\n", src, output_path)

	output_directory := filepath.Dir(output_path)
	err := os.MkdirAll(output_directory, 0775)
	if err != nil {
		return err
	}

	return Copy("./"+src, output_path)
}

func Copy(srcFile, dstFile string) error {
	fmt.Printf("Copy %s -> %s\n", srcFile, dstFile)

	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		in, err = os.Open("./presentations/" + srcFile)
		if err != nil {
			return err
		}
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
