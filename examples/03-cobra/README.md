# 03. Cobra でサブコマンド 🐍

## 📖 概要

人気の CLI ライブラリ **Cobra** を使って、サブコマンドに対応した本格的な CLI ツールを作るよ〜！
Docker、Kubernetes、GitHub CLI などでも使われてる超有名ライブラリだよ✨

## 🎯 学べること

- Cobra の基本的な使い方 📝
- サブコマンドの作り方（`mycli greet`, `mycli info` みたいな）🌳
- グローバルフラグとローカルフラグの違い 🚩
- 自動ヘルプ生成 💡

## 💻 使い方

```bash
# このディレクトリに移動
cd examples/03-cobra

# 依存関係をインストール
go mod download

# ヘルプを表示
go run main.go --help

# greet サブコマンドを実行
go run main.go greet

# 名前を指定して実行
go run main.go greet "あなたの名前"

# フラグを使って実行
go run main.go greet "ギャル" --emoji "🎀" --times 3

# 短縮フラグで実行（-e と -e を使用）
go run main.go greet "ギャル" -e "💕" -t 5

# 詳細モードで実行（グローバルフラグ）
go run main.go greet --verbose

# info サブコマンドを実行
go run main.go info

# 詳細モードで info を実行
go run main.go info -v
```

## 🔥 ポイント

- **サブコマンド**：`greet` と `info` みたいに機能ごとにコマンドを分けられる！
- **グローバルフラグ**：`--verbose` はすべてのサブコマンドで使える 🌍
- **ローカルフラグ**：`--emoji` や `--times` は `greet` コマンド専用 🎯
- **自動ヘルプ**：`--help` が自動で生成されるから楽チン！
- **短縮フラグ**：`-e` や `-t` みたいに短く書けるよ 📝

## 📦 ビルドして使う

```bash
# バイナリをビルド
go build -o mycli

# ビルドしたバイナリを実行
./mycli greet "ギャルエンジニア" -e "✨" -t 3
```

プロっぽい CLI ツールが作れちゃうね！🔥✨
