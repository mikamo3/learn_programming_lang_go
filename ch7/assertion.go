package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f, ok := w.(*os.File)
	b, ok := w.(*bytes.Buffer)
}
