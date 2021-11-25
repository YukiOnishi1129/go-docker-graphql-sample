# go-docker-graphql-sample

golang docker GraphQL のサンプル

## docker コマンド

```
// ビルド
docker-compose build

// コンテナ起動
docker-compose up

// コンテナ起動(バックグラウンド実行)
docker-compose up -d

// コンテナ停止
docker-compose down

// コンテナ停止&ボリューム削除(DBデータを削除)
docker-compose down -v

// appコンテナへログイン
docker exec -it 20211125_go_graphql_server sh

// dbコンテナへログイン
docker exec -it 20211125_go_graphql_db /bin/bash

```
