# Go システムプログラミングの練習
## 参考
<https://ascii.jp/serialarticles/1235262/>

## ./code
- **a.go** ファイル出力  
`os.File`のインスタンスは`os.Create()`（新規ファイルの場合）や`os.Open`（既存のファイルオープン）などの関数で作る

- **b.go** 画面出力  
`os.Stdout.Write`について`fmt.Println`は内部で最終的にこれを呼び出している

- **c.go**  書かれた内容を記憶しておくバッファ  
ファイルや画面出力のようなOSが提供する出力先に出すだけが`io.Writer`の機能ではなく。`Write()` メソッドで書き込まれた内容をためておける`bytes.Buffer`がある。

- **d.go** インターネットアクセス  
`net.Dial()` 関数を使うと、`net.Conn`という通信のコネクションを表すインタフェースが返される。
`net.Conn` は`io.Writer`と`io.Reader`のハイブリッドなインタフェースで、`io.Writer` としても使うことができる。
`net.Dial()` 関数が返す`net.Conn`インタフェースの実体は`net.TCPConn`構造体のポインタ。

- **e.go** インターネットアクセス  
`http.ResponseWriter`はウェブサーバーから、ブラウザに対してメッセージを書き込むのに使う

- **f.go** io.Writerのフィルタ  
`io.Writer`を受け取り、書き込まれたデータを加工して別の`io.Writer`に書き出す構造体について。`io.MultiWriter`は、複数の`io.Writer`を受け取り、書き込まれた内容をすべてに同時に書き込むフィルタ

- **g.go** io.Writerのフィルタ  
書き込まれたデータをgzipで圧縮して、予め渡されていた`os.File`に中継する。圧縮のフィルタ

- **h.go** io.Writerのフィルタ  
出力結果を一時的にためておいて、まとめて書き出す`bufio.Writer`について。`Flush()`メソッドを呼ぶと、後続の`io.Writer`に書き出せる。`Flush()`メソッドを呼ばないと、書き込まれたデータを腹に抱えたまま消滅してしまうので要注意。
`Flush()`を自動で呼び出す場合には、バッファサイズ指定の `bufio.NewWriterSize(os.Stdout, バッファサイズ)`関数で`bufio.Writer`を作成する。Go言語ではどの出力もバッファリングしない

- **i.go** フォーマットしてデータをio.Writerに書き出す  
`fmt.Fprintf()`は整形したデータを`io.Writer`へと書き出すC言語で言うところの`printf`のようなフォーマット出力のための関数。
フォーマット（２つ目の引数）にしたがって、`io.Writer`（最初の引数）にデータ（３つ目以降の引数）を書き出す。  
Goには何でも表示できる`%v`というフォーマット指定子があり、プリミティブ型でもそれ以外でも`String()`メソッドがあればそれを表示に使って出力してくれる。これも`fmt.Stringer`インターフェースとして定義されている。

- **j.go** フォーマットしてデータをio.Writerに書き出す  
JSONを整形して`io.Writer`に書き出すこともできる。j.go内ではコンソールに出力しているが、`io.Writer`と組み合わせれば、サーバーにJSONを送ったり、ブラウザにJSONを返すことも簡単にできる。

- **k.go** net/httpパッケージのRequest構造体  
`net\http`パッケージの`Request`構造体は、文字通りHTTPリクエストを取り扱う構造体 クライアント側のリクエストを送るときにも使えるほか、サーバ側でレスポンスを返すときにクライアントの情報をパースするのにも使える。`io.Writer` に書き出すのは前者の用途。  
`e.go`でのサーバへのリクエストの通信の例ではHTTPプロトコルを手書きしたが、この`Request`構造体を使え、ミスが減る。この構造体の `Write`メソッドを使わないといけないケースは実際には少ないが、`Transfer-Encoding: chunked`でチャンクに分けて送信したり、プロトコルのアップグレードで別のプロトコルと併用するようなHTTPリクエストを送るときには使うことになる。

- **l.go** 標準入力  
標準入力に対応するオブジェクトは`os.Stdin`これは`io.Writer`、`io.Closer`の各インターフェースを満たす。プログラムを実行すると入力待ちでブロックしてしまう（入力まで、プログラムが停止してしまう）。Go言語の`Read()`はタイムアウトのような仕組みがなく、このブロックを避けることができない。Goの場合はゴルーチンと呼ばれる軽量スレッドを使い、ノンブロッキングな処理を書く。  

- **m.go** ファイルの入力  
`os.File`構造体を使う。新規作成は`os.Create()`関数で行っていたが、`os.Open()`関数を使うと既存のファイルを開くことができる。内部的にはこの２つの関数は`os.OpenFile()`関数のフラグの違いのエイリアスで同じシステムコールが呼ばれている。
```go
func Open(name string) (*File, error) {
    return OpenFile(name, O_RDONLY, 0)
}

func Create(name string) (*File, error) {
    return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```


- **n.go** インターネット通信  
インターネット上でのデータのやり取りは、送信データを送信者側から見ると書き込みで、受信者側から見ると読み込みになる。`d.go`では
`net.Dial`で返される`conn`が`net.Conn`型で、これを`io.Copy`を使って標準出力にコピーすることでデータを一括で読み込んでいる。  
この場合読み込まれるのは生のHTTP通信内容そのもの。  
Go言語では、HTTPのレスポンスをパースする`http.ReadResponse()`関数が用意されている。この関数に`bufio.Reader`でラップした`net.Conn`を渡すと、`http.Response`構造体のオブジェクトが返される。   `bufio.Reader`でラップするには、`bufio.NewReader()`関数を呼ぶ。 このオブジェクトはHTTPのヘッダーやボディーなどに分解されているため、プログラムでの利用が簡単になる。

- o.go
- p.go
- q.go
- r.go
- s.go
- t.go
- u.go
- v.go
- w.go
- z.go

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