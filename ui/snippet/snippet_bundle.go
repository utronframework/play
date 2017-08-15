package snippet

//go:generate go run generate.go

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

var (
	assets     = map[string][]string{}
	assetFiles = map[string]fileData{}
)

type fileData struct {
	path string
	root string
	data []byte
}

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
	return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) io.Reader {
	reader, err := FindFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return reader
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, error) {
	item, ok := assetFiles[path]
	if !ok {
		return nil, fmt.Errorf("File %q not found in file system", path)
	}

	if !doGzip {
		return bytes.NewReader(item.data), nil
	}

	return gzip.NewReader(bytes.NewReader(item.data))
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
	body, err := ReadFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error) {
	body, err := ReadFileByte(path, doGzip)
	return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
	body, err := ReadFileByte(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error) {
	reader, err := FindFile(path, doGzip)
	if err != nil {
		return nil, err
	}

	var bu bytes.Buffer

	if _, err := io.Copy(&bu, reader); err != nil && err != io.EOF {
		return nil, err
	}

	return bu.Bytes(), nil
}
