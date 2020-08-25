package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("-----HEADER-----\n")
	content := bytes.NewBufferString("io.MultiReader sample \n")
	footer := bytes.NewBufferString("-----FOOTER-----\n")

	reader := io.MultiReader(header, content, footer)

	io.Copy(os.Stdout, reader)
}