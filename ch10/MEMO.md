# パッケージ命名について

衝突を避けるためにパッケージを称しているホストなどで始めるべき

# パッケージ規約の例外

実行可能なgoプログラムを定義しているパッケージはmain

_test

インポートパスにバージョンが付与されているもの。(gopkg.in/yaml.v2など)接尾辞は除去される

# 環境変数

GOPATH ワークスペースのroot
GOROOT 標準ライブラリのパッケージを提供しているroot,ユーザは設定不要

# docコメント

package宣言直前のコメントはパッケージのコメントとしてみなされる

どのファイルに書いてもいいが１つだけ

長くなる場合はdoc.goなどにかく

go docで表示

# internal package

internalディレクトリの親をルートとするツリー内の他のパッケージからのみimport可能

net/http/internal/chunkedは

net/http/httputil net/http などから参照可能