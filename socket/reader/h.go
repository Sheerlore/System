package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("io.TeeReader sample\n")
	teeReader := io.TeeReader(reader, &buffer)
	//データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)

	fmt.Println(buffer.String())
}