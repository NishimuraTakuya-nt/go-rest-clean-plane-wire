# go-rest-clean-plane-wire
base project for go rest api

## 使用ライブラリ
- wire
- viper
- swag
- gomock
- go-playground validator

## 実装内容
- 標準ライブラリでのルーディング
- Middleware
  - context value
    - request id
    - request info
  - CORS
  - logging
  - エラーハンドリング（recover）
  - タイムアウト
  - 認証
- カスタムロガー
- バリデーター
- wire ジェネレート
- swagger定義のジェネレート
- mock ジェネレート
- CI
  - lint
  - test
  - Dockerfile/docker-compose

### README に載せたい情報
- swagger の使い方
- mock generator の使い方
- wire の使い方

## このプロジェクトは以下のディレクトリ構造に基づいています：
```
.
├── cmd
│     └── api : アプリケーションのエントリーポイント
├── docs
│     └── swagger
├── internal : プロジェクト固有のパッケージ
│     ├── adapters : 外部システムとのインターフェース
│     │     ├── primary
│     │     │     └── http
│     │     │         ├── handlers
│     │     │         ├── middleware
│     │     │         └── routes
│     │     └── secondary
│     │         ├── aws
│     │         ├── db
│     │         └── graphql
│     ├── core : ビジネスロジック
│     │     ├── domain
│     │     ├── services
│     │     └── usecases
│     ├── errors
│     ├── infrastructure : 技術的な実装（ロギングなど）
│     │     ├── auth
│     │     ├── config
│     │     └── logger
│     └── utils
├── pkg
└── scripts
```
