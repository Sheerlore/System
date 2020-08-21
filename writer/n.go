package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)

	//ヘッダーを表示する
	fmt.Println(res.Header)

	//ボディを表示する。最後に必ずClose()する
	defer res.Body.Close()
	io.Copy(os.Stdout,res.Body)
}