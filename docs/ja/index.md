````markdown
# CrowdfundingZ — クイックスタート & ドキュメント (日本語)

このリポジトリは Go（Gin + GORM）で構築されたローカル優先のクラウドファンディングバックエンドを含みます。簡潔な日本語リファレンスとしてご利用ください。

## クイックスタート

- 前提: Go 1.21+, Git

```powershell
git clone https://github.com/hrbzhq/crowdfunding_z.git
cd crowdfunding_z
go mod tidy
$env:JWT_SECRET = "a-secure-random-secret"
CGO_ENABLED=1 go run main.go
```

サーバーは `http://localhost:8080` で待ち受けます。

## Autoupdate (CLI)

v1 以降、`--autoupdate` コマンドラインフラグで autoupdater スケジューラを実行できます。コマンドラインフラグは `ENABLE_AUTOUPDATE` 環境変数より優先され、CI ジョブ等での一時実行に便利です。

例（CLI による一回実行）:

```powershell
go run main.go --autoupdate
```

スケジューラの間隔は `--autoupdate-interval` または `AUTOSCHED_INTERVAL` 環境変数（例: `30m`, `1h`）で設定できます。

## GitHub Actions — Autoupdate (manual)

`.github/workflows/run_autoupdate.yml` に手動トリガーのワークフローが用意されています。デフォルトは安全なモック実行です。Issue 作成を有効にする場合は必要なシークレットを設定し、`create_issues=yes` で実行してください。

- 必要なシークレット（Settings → Secrets → Actions に追加）:
  - `JWT_SECRET` — JWT 署名に使う長いランダム文字列
  - `GITHUB_TOKEN` — `repo:issues` 権限を持つ PAT（可能であれば `secrets.GITHUB_TOKEN` を優先してください）
  - `GITHUB_REPO` — Issue 作成先の `owner/repo`

セキュリティ注意: 可能な限り `secrets.GITHUB_TOKEN` を利用し、PAT を作成する場合は最小権限に留めてください。

## シードと WebSocket テスト

ローカルシーダー: `POST /dev/seed`（`ENABLE_DEV_ENDPOINTS=true` のとき有効）

CLI シーダー: `go run tools/seed/cmd/seed_data`

WebSocket テストクライアント: `go run tools/selftest/ws_client.go`

---

詳細な API はリポジトリルートの `README.md` を参照してください。

````
# CrowdfundingZ — 日本語ドキュメント

こちらは日本語のドキュメントです。クイックスタートと autoupdate の説明は `README_ja.md` を参照してください。

## クイックスタート

```powershell
git clone https://github.com/hrbzhq/crowdfunding_z.git
cd crowdfunding_z
go mod tidy
$env:JWT_SECRET = "a-secure-random-secret"
CGO_ENABLED=1 go run main.go
```

サーバーは `http://localhost:8080` で待ち受けます。

(詳しい日本語ドキュメントは `README_ja.md` を参照してください)
