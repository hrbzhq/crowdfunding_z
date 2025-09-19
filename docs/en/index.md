# CrowdfundingZ — English Docs

This is the English documentation. For quick start and autoupdate notes, see the repository `README_en.md`.

## Quick Start

```powershell
git clone https://github.com/hrbzhq/crowdfunding_z.git
cd crowdfunding_z
go mod tidy
$env:JWT_SECRET = "a-secure-random-secret"
CGO_ENABLED=1 go run main.go
```

The server listens on `http://localhost:8080`.

(Full English docs are available in `README_en.md`)
