# go-rest-clean-plane-wire
base project for go rest api

## ライブラリの検討
- wire を実装する

## todo
### README に載せたい情報
- swagger の使い方
- mock generator の使い方
- wire の使い方

### その他
- swagger gen code
- 

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
