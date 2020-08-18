# Go システムプログラミングの練習
## 参考
<https://ascii.jp/serialarticles/1235262/>

## 各種ファイルに対してのメモ
- a.go ファイル出力  
`os.File`のインスタンスは`os.Create()`（新規ファイルの場合）や`os.Open`（既存のファイルオープン）などの関数で作る

- b.go 画面出力  
`os.Stdout.Write`について`fmt.Println`は内部で最終的にこれを呼び出している

- c.go  書かれた内容を記憶しておくバッファ  
ファイルや画面出力のようなOSが提供する出力先に出すだけが`io.Writer`の機能ではなく。`Write()` メソッドで書き込まれた内容をためておける`bytes.Buffer`がある。

- d.go インターネットアクセス  
`net.Dial()` 関数を使うと、`net.Conn`という通信のコネクションを表すインタフェースが返される。
`net.Conn` は`io.Writer`と`io.Reader`のハイブリッドなインタフェースで、`io.Writer` としても使うことができる。
`net.Dial()` 関数が返す`net.Conn`インタフェースの実体は`net.TCPConn`構造体のポインタ。

- e.go インターネットアクセス  
`http.ResponseWriter`はウェブサーバーから、ブラウザに対してメッセージを書き込むのに使う

- f.go io.Writerのフィルタ
`io.Writer`を受け取り、書き込まれたデータを加工して別の`io.Writer`に書き出す構造体について。`io.MultiWriter`は、複数の`io.Writer`を受け取り、書き込まれた内容をすべてに同時に書き込むフィルタ

- h.go io.Writerのフィルタ
書き込まれたデータをgzipで圧縮して、予め渡されていた`os.File`に中継する。圧縮のフィルタ



