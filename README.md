# Question points

## Quickstart

```console
make init
go run .
curl -X POST \
  -H 'content-type: application/json' \
  -d '{"message": "hello, world!"}'
  'http://localhost:8080/hello '
```

## API 定義を変更した場合

1. コードを再生成する

   ```console
   make generate/openapi
   ```

1. Server 実装のコンパイルエラーを解消する

## Project structure

| Directory      | Detail                               |
| -------------- | ------------------------------------ |
| configs        | 設定ファイル等の配置する             |
| docs           | ドキュメントを配置する               |
| domain         | ドメインコードを定義する             |
| infrastructure | インフラストラクチャコードを定義する |
| presentation   | プレゼンテーションコードを定義する   |
| internal       | アプリケーショの共通処理系を定義する |
