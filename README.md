# Discord API CLI Tool

このツールは、Discord API を操作するための高貴で洗練された CLI ツールです。
Go 1.26.1 の最新環境で構築されており、Bot Token を用いた認証と HTTP API Resource の操作に対応しています。

## 環境構築

### 1. Go のインストール
このプロジェクトは **Go 1.26.1** を前提としております。お手元の環境にインストールされているかご確認ください。

```bash
go version
# Output: go version go1.26.1 ...
```

### 2. リポジトリのビルド
プロジェクトのルートディレクトリで以下のコマンドを実行し、バイナリを生成します。

```bash
go build -o discord-cli ./cmd/discord-cli/main.go
```

## 使い方

### 1. 認証設定
Discord Developer Portal で取得した Bot Token を環境変数に設定してください。

```bash
export DISCORD_TOKEN=your_bot_token_here
```

### 2. コマンド実行例

#### 自身（Bot）の情報を取得
```bash
./discord-cli me
```

#### 自分が所属しているサーバー一覧を取得
```bash
./discord-cli me-guilds
```

#### メッセージの送信
```bash
./discord-cli message <channel_id> "皆様、ごきげんよう"
```

#### メッセージの編集
```bash
./discord-cli edit-message <channel_id> <message_id> "内容を更新いたしましたわ"
```

#### チャンネル名の変更
```bash
./discord-cli modify-channel <channel_id> "新しい社交場"
```

## テストの実行
全ての機能は `httptest` による Mock を用いて検証されています。

```bash
go test ./pkg/discord/... -v
```
