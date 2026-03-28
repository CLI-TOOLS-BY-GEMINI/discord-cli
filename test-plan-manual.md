# 実環境テスト (E2E) 準備・実施マニュアル

## 概要
実環境テストを安全かつ確実に実施するため、Discord Developer Portal および Discord クライアントでの手動操作、および環境変数の設定手順を解説いたしますわ。

## 1. Discord 側での事前準備

### 1-1. テスト用アプリケーション(Bot)の作成
1. [Discord Developer Portal](https://discord.com/developers/applications) にアクセスいたしますわ。
2. **"New Application"** をクリックし、適当な名前（例: `Discord-CLI-Tester`）を付けて作成いたします。
3. 左メニューの **"Bot"** を選択し、**"Reset Token"**（または初回は "Copy Token"）をクリックして、Bot トークンを控えておきますわ。これが `DISCORD_E2E_TOKEN` となります。
4. 同じく **"Bot"** ページ内にある **"Privileged Gateway Intents"** にて、以下の権限が必要に応じてオンになっているかご確認ください。
   - `MESSAGE CONTENT INTENT` (メッセージ内容の読み書きを行う場合)

### 1-2. テスト用サーバー(Guild)の準備
1. Discord のデスクトップ版またはウェブ版クライアントを開きます。
2. **"+" (サーバーを追加)** をクリックし、**"オリジナルの作成"** からテスト専用のサーバーを作成いたします。
3. 作成後、サーバー名を右クリックして **"IDをコピー"** を選択しますわ。これが `DISCORD_E2E_GUILD_ID` です。
   - *注意: IDをコピーできない場合は、ユーザー設定 > 詳細 > 開発者モード をオンになさってくださいませ。*

### 1-3. テスト用Botの招待
1. Developer Portal の **"OAuth2"** > **"URL Generator"** を選択します。
2. **"Scopes"** で `bot` にチェックを入れます。
3. **"Bot Permissions"** で `Administrator`（または `Manage Channels`, `Send Messages` 等の適切な権限）にチェックを入れます。
4. 生成された URL をブラウザで開き、先ほど作成したテスト用サーバーに招待いたしますわ。

### 1-4. テスト用チャンネルの作成
1. サーバー内にテスト用のテキストチャンネル（例: `#e2e-test`）を作成します。
2. チャンネル名を右クリックし、**"IDをコピー"** を選択します。これが `DISCORD_E2E_CHANNEL_ID` ですわ。

---

## 2. ローカル環境でのテスト実行手順

### 2-1. 環境変数の設定
ターミナルで以下のコマンドを実行し、先ほど控えた ID 群を環境変数として設定いたします。

```bash
export DISCORD_E2E_TOKEN="あなたのBotトークン"
export DISCORD_E2E_GUILD_ID="あなたのサーバーID"
export DISCORD_E2E_CHANNEL_ID="あなたのチャンネルID"
```

### 2-2. テストの実行
環境変数が設定された状態で `go test` を実行いたします。

```bash
# E2Eテストのみを実行する場合 (将来的に e2e_test.go が作成された場合)
go test -v ./pkg/discord -run TestE2E
```

## 3. 手動による動作確認 (CLIバイナリ使用)
自動テストだけでなく、手動で CLI の挙動を確認する場合は以下の通りですわ。

1. **ビルド**:
   ```bash
   go build -o discord-cli ./cmd/discord-cli/main.go
   ```
2. **Bot情報の確認**:
   ```bash
   DISCORD_TOKEN=$DISCORD_E2E_TOKEN ./discord-cli me
   ```
3. **メッセージ送信**:
   ```bash
   DISCORD_TOKEN=$DISCORD_E2E_TOKEN ./discord-cli message $DISCORD_E2E_CHANNEL_ID "手動テストですわ"
   ```
4. **チャンネル名の変更**:
   ```bash
   DISCORD_TOKEN=$DISCORD_E2E_TOKEN ./discord-cli modify-channel $DISCORD_E2E_CHANNEL_ID "麗しの社交場"
   ```
