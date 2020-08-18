# Go システムプログラミングの練習
## 参考
<https://ascii.jp/serialarticles/1235262/>

## 各種ファイルに対してのメモ
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


- k.go
- l.go
- m.go
- n.go
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




