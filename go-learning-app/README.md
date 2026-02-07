# Go学習アプリ

インタラクティブなGo言語学習アプリケーション

## ローカルで実行

```bash
cd go-learning-app
go run .
```

ブラウザで http://localhost:8080 にアクセス

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
