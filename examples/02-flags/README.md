# 02. フラグの使い方 🚩

## 📖 概要

Go の標準ライブラリ `flag` パッケージを使って、コマンドラインフラグを処理する方法を学ぶよ〜！
フラグは `-name value` みたいな形式で引数を渡せて超便利！✨

## 🎯 学べること

- `flag.String()`, `flag.Int()`, `flag.Bool()` の使い方 📝
- フラグのデフォルト値の設定 🎁
- `flag.Parse()` でフラグをパースする方法 🔄
- フラグ以外の引数を取得する方法 (`flag.Args()`) 📋

## 💻 使い方

```bash
# このディレクトリに移動
cd examples/02-flags

# デフォルト値で実行
go run main.go

# 名前を指定して実行
go run main.go -name "あなたの名前"

# 繰り返し回数を指定
go run main.go -count 3

# 詳細表示モードで実行
go run main.go -verbose

# 全部組み合わせて実行！🔥
go run main.go -name "ギャル" -count 5 -verbose -emoji "🎀"

# ヘルプを表示
go run main.go -help
```

## 🔥 ポイント

- フラグは **ポインタ** で返ってくるから、使うときは `*name` みたいにアスタリスクつけてね！
- `flag.Parse()` を呼ぶ前にフラグを定義する必要があるよ 📌
- `-help` フラグは自動で追加されるから便利〜！
- ブール型フラグは `-verbose` だけで true になるよ（`-verbose=true` も OK）

めっちゃ実用的でしょ？CLI ツール作るときの必須テクニックだよ〜💪✨
