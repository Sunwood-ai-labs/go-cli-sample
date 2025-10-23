# Go開発環境用のDockerfile
# ベースイメージとしてGo公式イメージを使用
FROM golang:1.24.7-alpine

# 作業ディレクトリを設定
WORKDIR /app

# ビルドに必要なツールをインストール
RUN apk add --no-cache git

# go.modとgo.sumをコピー（依存関係のキャッシュ用）
COPY go.mod ./

# 依存関係をダウンロード
RUN go mod download

# プロジェクトのファイルをすべてコピー
COPY . .

# テストレポート用のディレクトリを作成
RUN mkdir -p test-reports

# アプリケーションをビルド
RUN go build -o go-cli-sample main.go

# デフォルトのコマンドを設定
CMD ["./go-cli-sample"]
