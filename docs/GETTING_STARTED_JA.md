# Go言語 初心者向けガイド 🎓

このドキュメントでは、このプロジェクトで使用されているGo言語の基本とDocker環境について説明します。

## 目次

1. [プロジェクト構造](#プロジェクト構造)
2. [各ファイルの説明](#各ファイルの説明)
3. [Go言語の基本](#go言語の基本)
4. [Docker環境の使い方](#docker環境の使い方)
5. [よくある質問](#よくある質問)

---

## プロジェクト構造

```
go-cli-sample/
├── main.go              # メインプログラム（エントリーポイント）
├── main_test.go         # テストコード
├── go.mod               # Go モジュール定義ファイル
├── Dockerfile           # Docker イメージの定義
├── docker-compose.yml   # Docker サービスの構成
├── .dockerignore        # Docker ビルドで除外するファイル
├── test-reports/        # テスト結果の出力先
└── docs/                # ドキュメント
    └── GETTING_STARTED_JA.md  # このファイル
```

---

## 各ファイルの説明

### main.go

**役割**: アプリケーションのメインプログラム

```go
package main  // このファイルがmainパッケージに属することを宣言

import (      // 必要なパッケージをインポート
    "fmt"     // フォーマット出力用
    "os"      // OS関連の機能（コマンドライン引数など）
)
```

**主な機能**:
- `printBanner()`: ASCII アートバナーを表示する関数
- `main()`: プログラムのエントリーポイント（最初に実行される関数）
- `os.Args`: コマンドライン引数を取得

**重要なポイント**:
- `package main` と `func main()` は実行可能なGoプログラムに必須
- `import` で標準ライブラリや外部パッケージを読み込む

### main_test.go

**役割**: プログラムの動作を自動でテストするコード

```go
package main

import (
    "testing"  // テスト用の標準ライブラリ
)

func TestMainOutput(t *testing.T) {
    // テストコード
}
```

**主な機能**:
- `TestXxx` 形式の関数がテストケース
- `BenchmarkXxx` 形式の関数がベンチマーク（性能測定）
- `*testing.T` を使ってテストの成功/失敗を報告

**テストの種類**:
1. **ユニットテスト**: 個々の関数が正しく動作するか確認
2. **ベンチマーク**: 性能（実行速度、メモリ使用量）を測定

### go.mod

**役割**: Goモジュールの定義ファイル（プロジェクトの設定）

```go
module github.com/Sunwood-ai-labs/go-cli-sample  // モジュール名

go 1.24.7  // 使用するGoのバージョン
```

**重要なポイント**:
- プロジェクトの依存関係を管理
- `go mod download` で依存パッケージをダウンロード
- 新しい依存関係は `go get` で追加

### Dockerfile

**役割**: Docker イメージの構築手順を定義

```dockerfile
FROM golang:1.24.7-alpine  # ベースとなるイメージ

WORKDIR /app  # 作業ディレクトリを設定

COPY go.mod ./  # ファイルをコピー
RUN go mod download  # 依存関係をダウンロード

COPY . .  # すべてのファイルをコピー
RUN go build -o go-cli-sample main.go  # ビルド

CMD ["./go-cli-sample"]  # デフォルトで実行するコマンド
```

**重要な命令**:
- `FROM`: ベースイメージを指定
- `WORKDIR`: 作業ディレクトリを設定
- `COPY`: ファイルをコンテナにコピー
- `RUN`: コマンドを実行
- `CMD`: コンテナ起動時に実行するコマンド

### docker-compose.yml

**役割**: 複数のDockerサービスを定義・管理

```yaml
version: '3.8'

services:
  app:         # アプリケーション実行用
  test:        # テスト実行用
  benchmark:   # ベンチマーク実行用
  dev:         # 開発用シェル
```

**主な設定**:
- `build`: Dockerfile の場所
- `volumes`: ホストとコンテナ間のファイル共有
- `command`: コンテナ起動時のコマンド
- `environment`: 環境変数

### .dockerignore

**役割**: Docker ビルド時に除外するファイルを指定

```
.git
*.md
test-reports/
```

**重要なポイント**:
- ビルド時間を短縮
- イメージサイズを削減
- 不要なファイルの混入を防止

---

## Go言語の基本

### パッケージ

```go
package main  // パッケージ名を宣言
```

- すべてのGoファイルは何らかのパッケージに属する
- `main` パッケージは特別（実行可能なプログラム）
- その他のパッケージはライブラリとして使用

### インポート

```go
import "fmt"  // 単一パッケージ

import (      // 複数パッケージ
    "fmt"
    "os"
)
```

**標準ライブラリの例**:
- `fmt`: フォーマット入出力
- `os`: OS関連機能
- `strings`: 文字列操作
- `testing`: テスト機能

### 関数

```go
func 関数名(引数名 型) 戻り値の型 {
    // 処理
    return 値
}

// 例
func add(x int, y int) int {
    return x + y
}
```

### 変数

```go
// 明示的な型宣言
var name string = "太郎"

// 型推論
var age = 20

// 短い宣言（関数内のみ）
message := "Hello"
```

### 配列とスライス

```go
// 配列（固定長）
var arr [5]int

// スライス（可変長）
slice := []string{"a", "b", "c"}
```

### 制御構文

```go
// if文
if x > 10 {
    fmt.Println("大きい")
} else {
    fmt.Println("小さい")
}

// forループ
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// 範囲for
for index, value := range slice {
    fmt.Println(index, value)
}
```

---

## Docker環境の使い方

### 前提条件

- Docker がインストールされていること
- Docker Compose がインストールされていること

### 基本コマンド

#### 1. アプリケーションを実行

```bash
docker-compose up app
```

**動作**:
- Go アプリケーションをビルドして実行
- 引数 "Docker環境から実行だよ〜" で起動

#### 2. テストを実行

```bash
docker-compose up test
```

**動作**:
- すべてのテストを実行
- カバレッジレポートを生成
- 結果を `test-reports/` に保存

#### 3. ベンチマークを実行

```bash
docker-compose up benchmark
```

**動作**:
- ベンチマークテストを実行
- 性能測定結果を `test-reports/benchmark-results.txt` に保存

#### 4. 開発用シェルを起動

```bash
docker-compose run --rm dev
```

**動作**:
- コンテナ内でシェルを起動
- Go コマンドを直接実行可能

**シェル内でのコマンド例**:
```bash
# プログラムを実行
go run main.go

# テストを実行
go test -v

# ビルド
go build

# フォーマット
go fmt ./...
```

#### 5. コンテナを停止

```bash
docker-compose down
```

#### 6. イメージを再ビルド

```bash
docker-compose build
```

### 開発ワークフロー

1. **コードを編集**
   ```bash
   vim main.go  # または好きなエディタで編集
   ```

2. **テストで確認**
   ```bash
   docker-compose up test
   ```

3. **アプリケーションを実行**
   ```bash
   docker-compose up app
   ```

4. **ベンチマークで性能確認**
   ```bash
   docker-compose up benchmark
   ```

### ボリュームマウントについて

```yaml
volumes:
  - .:/app  # 現在のディレクトリをコンテナの/appにマウント
```

**メリット**:
- ホスト側でファイルを編集すると、即座にコンテナに反映
- テスト結果などをホスト側で確認可能
- コンテナを再起動せずに開発可能

---

## よくある質問

### Q1. `go run` と `go build` の違いは？

**A**:
- `go run main.go`: コンパイルして即座に実行（一時的）
- `go build`: 実行可能ファイルを生成（永続的）

### Q2. テストが失敗したらどうする？

**A**:
```bash
# 詳細なログを確認
docker-compose up test

# test-reports/test-results.txt を確認
cat test-reports/test-results.txt
```

### Q3. 新しいパッケージを追加するには？

**A**:
```bash
# 開発用シェルで
docker-compose run --rm dev

# シェル内で
go get パッケージ名
```

### Q4. Dockerを使わずに開発できる？

**A**: できます！
```bash
# ローカルでGoがインストールされている場合
go run main.go
go test -v
go build
```

### Q5. カバレッジレポートの見方は？

**A**:
```bash
# test-reports/coverage.html をブラウザで開く
# 緑色: テスト済み
# 赤色: 未テスト
# グレー: テスト対象外
```

### Q6. コンテナ内のファイルを確認したい

**A**:
```bash
# シェルを起動
docker-compose run --rm dev

# ls, cat などのコマンドを使用
ls -la
cat main.go
```

### Q7. エラー "port already in use" が出る

**A**:
```bash
# 実行中のコンテナを停止
docker-compose down

# それでもダメなら、すべてのコンテナを停止
docker stop $(docker ps -aq)
```

---

## 次のステップ

1. **main.go を改造してみる**
   - 新しい関数を追加
   - コマンドライン引数を増やす
   - ファイル入出力を追加

2. **テストを追加してみる**
   - 新しいテストケースを書く
   - エッジケースをテスト

3. **外部パッケージを使ってみる**
   - CLI フレームワーク（cobra など）
   - ロギング（logrus など）

4. **Go言語の学習リソース**
   - [Go by Example](https://gobyexample.com/)
   - [A Tour of Go](https://go.dev/tour/)
   - [Effective Go](https://go.dev/doc/effective_go)

---

## まとめ

- `main.go`: プログラムの本体
- `main_test.go`: テストコード
- `go.mod`: プロジェクト設定
- `Dockerfile`: 環境構築手順
- `docker-compose.yml`: サービス定義

Dockerを使えば、環境構築不要で開発を始められます！

頑張ってね！ 🚀
