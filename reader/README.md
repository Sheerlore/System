## ./reader
### バイナリ解析

- **a.go** 必要な部位を切り出す  
  - `io.LimitReader`  - 先頭から指定して切り出す。
  - `io.SectionReader` - スタート位置を指定して切り出す。  
  ただし`io.SectionReader`は`io.Reader`が使えず、代わりに`io.ReaderAt`インターフェースを使う。`os.File`は`io.ReaderAt`を満たすがそれ以外の`io.Reader`を満たす型から`io.SectionReader`で読み込むことはできない。文字列やバイト列にいったん書き出し、`strings.Reader`や`bytes.Reader`でラップしてから`io.SectionReader`に渡す。

- **b.go** エンディアン変換  
任意のエンディアンの数値を、現在の実行環境のエンディアン数値に変換するには、`encoding/binary`パッケージを使う。このパッケージの`binary.Read()`メソッドに、`io.Reader`とデータのエンディアン、それに変換結果を格納する変数のポインタを渡せば、エンディアンが修正されたデータが得られる。

- **c.go** PNGファイルの分析  
PNGファイルはバイナリフォーマット。先頭8バイトがシグニチャ（固定のバイト列）となっている。それ以降はチャンクと呼ばれるブロック（データの塊）で構成されている。コード内では、`readChunks()`関数でチャンクごとに`io.SectionReader`を作って配列に格納して返す。それをチャンクを表示する関数`dumpChunk()`で表示する。
```
長さ：4バイト
種類：4バイト
データ：長さで指定されたバイト列
CRC（誤り検知符号）：4バイト
```

- **d.go** PNG画像に秘密のテキストを入れる  
PNGには、テキストを追加するための`tEXt`というチャンクがある。また、それに圧縮をかけた`tTXt`というチャンクもある。これらはデータの中に埋め込まれるだけで画像としては表示されない。

### テキスト解析
- **e.go** 改行/単語で区切る  
全部読み込んでから文字列処理で改行に分割する方法があるが、`io.Reader`による入力では`bufio.Reader`を使うとシンプルになる。`ReadString()`、`ReadBytes()`を使うと任意の文字で分割することができる。また、読み込んだ文字を戻すこともできる。
終端を気にしないで短く書くなら`bufio.Scanner`を使う方法もある。しかし、`bufio.Reader`の結果には末尾に改行記号が残っているがこっちの方法だと分割文字が削除されている。`bufio.Scanner`のデフォルトは改行区切りだが、分割関数を指定することで任意の分割が行える。
```go
func main() {
  scanner := bufio.NewScanner(strings.NewReader(source))
  for scanner.Scan() {
    fmt.Printf("%#v/n", scanner.Text())
  }
}
```
```go
//分割処理を単語区切りに指定
scanner.Split(bufio.ScanWords)
```

- **f.go** データ型を指定して解析  
`io.Reader`のデータを整数や浮動小数点に変換するには、`fmt.Fscan`を使う。
1つめの引数に`io.Reader`を渡し、それ以降に変数のポインタを渡すと、その変数にデータが書き込まれる。`fmt.Fscan`はデータがスペース区切りであることを前提としている。`fmt.Fscanln`は改行区切りのときに用いる。
`fmt.Fscanf`を使うと任意のデータ区切りをフォーマット文字列として指定できる。Go言語は型情報をデータが持っているため、すべて`%v`と書いておけばOK。
```go
fmt.Fscanf(reader, "%v, %v, %v, %v", &i, &f, &g, &s)
```