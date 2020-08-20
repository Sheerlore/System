package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	//zipの内容を書き込むファイル
	file, err := os.Create("../sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//zipファイル
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()


	//ファイルの数だけ書き込み
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("text1"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("text2"))
}