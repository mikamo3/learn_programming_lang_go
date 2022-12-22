# テスト

ファイル名は_test.goで終わる必要がある

テスト関数はTestName(t *testing.T)とする必要がある

ベンチマーク関数はBenchmarkで始まる

コード例関数はExampleで始まる

t.Errorfはテストを止めない

# Example関数

Exampleで始まる関数はパラメータも結果も持たない。ドキュメンテーション、go testで実行可能、実際の動作確認につかう