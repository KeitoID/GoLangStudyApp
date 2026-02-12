# Go学習アプリ

インタラクティブなGo言語学習アプリケーション。10章34レッスンの体系的なカリキュラムで、コードエディタ・クイズ・進捗管理を備えています。

## 前提条件

- [Go](https://go.dev/dl/) 1.24 以上

## 環境構築

```bash
# リポジトリをクローン
git clone <リポジトリURL>
cd go-learning-app

# 依存パッケージをダウンロード
go mod download
```

## ローカルで実行

```bash
go run .
```

ブラウザで http://localhost:8080 にアクセスします。

初回アクセス時にユーザー名の入力を求められます。入力すると学習を開始でき、進捗は SQLite データベース（`go-learning.db`）に保存されます。

### ポート変更

デフォルトは `8080` です。変更する場合は環境変数 `PORT` を指定します。

```bash
PORT=3000 go run .
```

### データベースの保存先変更

デフォルトはカレントディレクトリの `go-learning.db` です。変更する場合は環境変数 `DB_PATH` を指定します。

```bash
DB_PATH=/tmp/learning.db go run .
```

## 停止

ターミナルで `Ctrl + C` を押すとサーバーが停止します。

進捗データは `go-learning.db` に永続化されているため、再起動しても保持されます。

---

## GCP Cloud Run へのデプロイ

### 前提条件

1. [Google Cloud CLI](https://cloud.google.com/sdk/docs/install) をインストール
2. GCPプロジェクトを作成済み
3. 課金が有効化されていること

### デプロイ手順

#### 1. gcloud CLIの設定

```bash
# ログイン
gcloud auth login

# プロジェクト設定
gcloud config set project YOUR_PROJECT_ID

# リージョン設定（東京）
gcloud config set run/region asia-northeast1
```

#### 2. Cloud Runへデプロイ（ソースから直接）

```bash
cd go-learning-app
gcloud run deploy go-learning-app --source .
```

初回実行時に以下を聞かれます：
- **Artifact Registry APIの有効化** → `y`
- **Cloud Build APIの有効化** → `y`
- **未認証アクセスの許可** → `y`（公開する場合）

#### 3. デプロイ完了

デプロイ完了後、URLが表示されます：
```
Service URL: https://go-learning-app-xxxxx-an.a.run.app
```

---

## 便利なコマンド

```bash
# ログを確認
gcloud run services logs read go-learning-app

# サービス一覧
gcloud run services list

# サービス削除
gcloud run services delete go-learning-app
```
