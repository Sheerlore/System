package main

import (
	"io"
	"os"
	"strings"
)

//文字列からSectionの部分だけを切り出したReaderを作成し、それをすべて
//os.Stdoutに書き出す。実際には文字列を分けるためにio.SectionReaderを
//使うことはない
func main(){
	reader := strings.NewReader("io.SectionReader sample \n")
	sectionReader := io.NewSectionReader(reader, 3 ,4)
	io.Copy(os.Stdout, sectionReader)
}