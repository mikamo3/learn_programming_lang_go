# 配列同士の比較

すべての対応する要素が等しいかを報告する(指しているポインタではない)

# 配列をパラメータに渡されたとき

goでは配列のコピーを渡す

# スライス同士の比較

配列のように==での比較はできない

bytes.Equalはあるが他の型は自作しないといけない

スライス型のゼロ値はnilでその比較はできる。

空であるかを確かめるなら

```go

len(s) == 0
```

# mapの並び順

定義されていない。


# mapのアドレス

取得できない。ハッシングされて別のメモリ位置に移動する可能性があるから

# mapのゼロ値

nil. この状態でマップへ保存しようとするとパニック

# mapの比較

スライス同様直接比較はできない

# 構造体

フィールドの名前は大文字で公開、フィールド順は逆にしたりすると異なる構造体を定義したという扱いになる

# 構造体リテラル

すべてのフィールドに正しい順に値を指定する必要がある。

変更に脆弱。パッケージ内の利用科順序が明らかな小さな構造体で使うことが多い

# JSONのマーシャリング
公開されているフィールドのみが対象,フィールドタグによりコンパイル時の構造体のフィールドに関連付ける