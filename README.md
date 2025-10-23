# Go CLI Sample

これはGo言語で作ったシンプルなCLIツールのサンプルだよ〜！✨

## 概要

このプロジェクトは、Go言語でCLIツールを作成する基本的な例を提供します。
コマンドライン引数の処理や標準出力の使い方を学ぶことができます。

## 必要要件

- Go 1.16以上

## セットアップ

```bash
# リポジトリをクローン
git clone https://github.com/Sunwood-ai-labs/go-cli-sample.git
cd go-cli-sample

# 依存関係をダウンロード（必要に応じて）
go mod download
```

## 使い方

### 基本的な実行

```bash
# 引数なしで実行
go run main.go

# 引数ありで実行
go run main.go "引数だよん"
```

### ビルドして実行

```bash
# バイナリをビルド
go build

# ビルドしたバイナリを実行
./go-cli-sample "引数だよん"
```

## テスト

### テストの実行

```bash
# すべてのテストを実行
go test -v

# カバレッジ付きでテストを実行
go test -v -cover

# カバレッジレポートを生成
go test -coverprofile=test-reports/coverage.out
go tool cover -html=test-reports/coverage.out -o test-reports/coverage.html
```

### ベンチマークの実行

```bash
# ベンチマークテストを実行
go test -bench=. -benchmem
```

## テストレポート

テスト結果とカバレッジレポートは `test-reports/` ディレクトリに生成されます：

- `coverage.out` - カバレッジデータ
- `coverage.html` - HTMLカバレッジレポート
- `test-results.txt` - テスト実行結果

## プロジェクト構成

```
.
├── main.go          # メインのCLIアプリケーション
├── main_test.go     # テストコード
├── README.md        # このファイル
└── test-reports/    # テストレポート出力ディレクトリ
```

## 機能

- コマンドライン引数の受け取り
- 標準出力への出力
- ユニットテスト
- ベンチマークテスト

イケてるっしょ？😉
