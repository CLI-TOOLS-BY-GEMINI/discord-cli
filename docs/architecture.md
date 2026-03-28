# アーキテクチャ構成

本プロジェクトは、クリーンで拡張性の高い構成を目指して構築されておりますわ。

## ディレクトリ構成

```text
/home/minatosingull/repo/discord-cli/
├───cmd/
│   └───discord-cli/
│       └───main.go        # エントリポイント。CLIのフラグ解析とコマンド呼び出しを担いますわ。
└───pkg/
    └───discord/
        ├───client.go      # APIクライアントの実装。HTTPリクエストの送信を管理いたします。
        ├───types.go       # Discord APIのリソース（User, Channel, Message等）の構造体定義ですわ。
        └───client_test.go # クライアントのユニットテスト。httptestによるモックを使用しております。
```

## パッケージの責務

- **main パッケージ (`cmd/discord-cli/`)**:
  - `flag` パッケージを用いて、環境変数や引数からトークンやコマンドを読み取りますわ。
  - ユーザーからの入力を解釈し、適切な `pkg/discord` のメソッドを呼び出します。
- **discord パッケージ (`pkg/discord/`)**:
  - Discord REST API V10 との通信を抽象化しておりますわ。
  - 各種リソースの取得・作成・更新・削除をメソッドとして提供し、低レベルなHTTP通信の詳細を隠蔽いたします。
