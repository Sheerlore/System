## ./quiz
- **q1.go** ファイルのコピー  
古いファイル(old.txt)を新しいファイル(new.txt)にコピーする

- **q2.go** テスト用の適当なサイズのファイルを作成  
ファイルを作成してランダムな内容で埋めてみる。  
`crypto/rand`パッケージの`rand.Reader`という`io.Reader`を使う。このReaderはランダムなバイトを延々と出力し続ける無限長のファイルのような動作をする。これを使って1024バイトの長さのバイナリファイルを作ってみる。`io.Copy()`はReaderの終了するまでコピーするので`rand.Reader`には終わらない。


- **q3.go** zipファイルの書き込み  
`io.Writer`や`io.Reader`は、1つのファイルやデバイスと1対1に対応している。複数ファイルを格納するアーカイブフォーマットのtarやzipファイルやインターネットのマルチパート形式（ブラウザのフォームによって作られるデータやファイルを格納するデータ構造）をサポートする`mime/multipart`パッケージの構造体は、中に格納される一つ一つの要素が`io.Writer`や`io.ReadCloser`になっている。  
`archive/zip`パッケージを使ってzipファイルを作成する。出力先のファイルのWriterをまず作って、それを`zip.NewWriter()`関数に渡すと、zipファイル書き込み用の構造体ができる。最後に`Close()`を呼ぶ。  
この構造体そのものは`io.Writer`ではないが、`Create()`メソッドを呼ぶと、個別のファイルを書き込むための`io.Writer`が返される。
ここでは、`strings.Reader`を使ってzipファイルを作成する。
```go
zipWriter := zip.NewWriter(file)
defer zipWriter.Close()
```

- **q4.go** zipファイルをウェブサーバーからダウンロード  
zipの出力先は単なる`io.Writer`なので、ウェブサーバーでzipファイルを作成して、そのままダウンロードさせることも可能。ウェブサーバーにブラウザでアクセスしたらファイルがダウンロードされるようにしてみる。  
この場合は`Content-Type`ヘッダーを使ってファイルの種類がzipファイルで有ることをブラウザに教えて上げる必要がある。ファイル名も指定できる。
```go
func handler(w http.ResposeWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/zip")
    w.Header().Set("Content-Disposition", "attachment: filename=websample.zip")
}
```