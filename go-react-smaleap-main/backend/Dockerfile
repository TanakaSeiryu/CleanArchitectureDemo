# ベースイメージとして公式のGoイメージを使用
FROM golang:1.22.5-alpine

# 作業ディレクトリを設定
WORKDIR /app

# モジュールファイルを作業ディレクトリにコピー
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Airをインストール

RUN go install github.com/air-verse/air@latest
# 残りのソースコードをコピー
COPY . .

# ポート8080を公開
EXPOSE 8080

# コンテナが起動する際に実行するコマンドを設定
CMD ["air", "-c", ".air.toml"]
