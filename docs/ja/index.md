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
