package main

import "io"

// FileReaderExtension extends file reader
type FileReaderExtension struct {
	io.Reader
	total    int64
	length   int64
	progress float64
}
