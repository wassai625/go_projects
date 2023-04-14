# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgit・sqliteのインストール
RUN apk update && apk add git sqlite
# ワーキングディレクトリの設定
WORKDIR /todo_app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /todo_app
# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports