package main

import (
	"bytes"
	"fmt"
)

func main(){
	var buffer bytes.Buffer
	buffer.Write([]byte("Hello World\n"))
	buffer.Write([]byte("Hello World\n"))
	buffer.Write([]byte("Hello World\n"))
	fmt.Println(buffer.String())

	
	//bytes.Bufferにはbyteではなく文字列を受け取れる
	//WriteStringというメソッドがあるので次のようにも書ける
	//buffer.WriteString("bytes.Buffer example\n")

	//WriteString は io.Writer のメソッドではないため、
	//他の構造体では使えない。
	//代わりに、次の io.WriteString を使えばキャストは不要になる。
	//io.WriteString(buffer, "bytes.Buffer example\n")
	

}