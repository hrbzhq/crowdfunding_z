# CrowdfundingZ

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go CI](https://github.com/hrbzhq/crowdfunding_z/actions/workflows/go.yml/badge.svg)](https://github.com/hrbzhq/crowdfunding_z/actions)
[![Release](https://img.shields.io/github/v/release/hrbzhq/crowdfunding_z?label=release)](https://github.com/hrbzhq/crowdfunding_z/releases/latest)

A modern, scalable crowdfunding platform built with Go, featuring user authentication, real-time updates, and blockchain integration for decentralized funding.

> There is no doubt that in the new era of AI, every person and every endeavor can become the center of the world. Every voice, every hope, deserves to be heard and cared for. Technology is no longer a barrier — it effortlessly serves people. Let the power of a community support a single hope. Start now: everyone can crowdfund.

> 新しいAIの時代において、誰もが、どんな取り組みも世界の中心になり得ます。一人ひとりの声や希望は、聞かれ、大切にされるべきです。技術はもはや障壁ではなく、人々に自然に寄り添います。コミュニティの力で、ひとつの希望を支えましょう。今すぐ始めてください：誰でもクラウドファンディングができます。

## Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Market Analysis](#market-analysis)
- [Development Roadmap](#development-roadmap)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Code of Conduct](#code-of-conduct)
- [Security](#security)

## Features

### Core Functionality
- **User Management**: Secure registration and JWT-based authentication
- **Project Creation**: Easy project setup with goal setting and deadlines
- **Funding System**: Seamless contribution tracking and progress monitoring
- **Real-time Updates**: WebSocket-powered live progress notifications
- **Modular Architecture**: Clean separation of concerns for easy maintenance

### Advanced Features
- **Blockchain Integration**: Ethereum-based smart contracts for transparent funding
- **Scalable Database**: SQLite with GORM for efficient data management
- **RESTful API**: Well-documented endpoints for easy integration
- **WebSocket Support**: Real-time communication for dynamic user experiences

## Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Gin (HTTP router and middleware)
- **Database**: SQLite with GORM ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Real-time**: Gorilla WebSocket
- **Blockchain**: Go-Ethereum client

### Infrastructure
- **Version Control**: Git
- **Package Management**: Go Modules
- **Testing**: Go's built-in testing framework
- **Documentation**: GoDoc

### Development Tools
- **IDE**: Visual Studio Code
- **Linting**: golangci-lint
- **Dependency Management**: go mod

## Market Analysis

### Industry Overview
The global crowdfunding market is experiencing rapid growth, with projections reaching $300 billion by 2025. Traditional crowdfunding platforms face challenges in transparency, fees, and global accessibility.

### Target Market
- **Creators**: Artists, entrepreneurs, and innovators seeking funding
- **Investors**: Individuals and organizations looking for impact investments
- **Developers**: Tech enthusiasts interested in decentralized finance
- **Enterprises**: Companies exploring alternative funding models

### Competitive Advantages
- **Decentralized Option**: Blockchain integration for trustless transactions
- **Low Fees**: Open-source model reduces operational costs
- **Global Accessibility**: No geographical restrictions
- **Transparency**: Public ledger for all transactions
- **Extensibility**: Modular design for easy feature additions

### Market Opportunities
- **Emerging Markets**: High growth potential in developing economies
- **Niche Crowdfunding**: Specialized platforms for specific industries
- **DeFi Integration**: Combining traditional crowdfunding with decentralized finance
- **Social Impact**: Funding for charitable and social enterprises

## Development Roadmap

### Phase 1: Core Platform (Current)
- [x] User authentication system
- [x] Basic project creation and funding
- [x] Real-time progress updates
- [x] SQLite database integration
- [x] RESTful API implementation

### Phase 2: Enhanced Features (Q4 2025)
- [ ] Payment gateway integration (Stripe, PayPal)
- [ ] Email/SMS notifications
- [ ] Advanced project analytics
- [ ] Mobile app development
- [ ] Multi-language support

### Phase 3: Blockchain Integration (Q1 2026)
- [ ] Smart contract deployment
- [ ] Token-based rewards system
- [ ] Decentralized identity verification
- [ ] Cross-chain compatibility
- [ ] NFT integration for backers

### Phase 4: Enterprise Features (Q2 2026)
- [ ] Advanced analytics dashboard
- [ ] API for third-party integrations
- [ ] White-label solutions
- [ ] Compliance and regulatory features
- [ ] Scalability optimizations

### Long-term Vision
- AI-powered project recommendations
- Global regulatory compliance
- Integration with major blockchains
- Decentralized autonomous organization (DAO) features

## Quick Start

### Prerequisites
- Go 1.21 or later
- GCC compiler (for SQLite CGO support)
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/crowdfunding_z.git
   cd crowdfunding_z
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the application**
   ```bash
   CGO_ENABLED=1 go run main.go
   ```

4. **Access the API**
   The server will start on `http://localhost:8080`

### Configuration
- Database: SQLite (crowdfunding.db)
- Port: 8080 (configurable)
- JWT Secret: Configurable via environment variables

### Autoupdate (CLI)

English

Starting from v1, the application supports the `--autoupdate` command-line flag to run the autoupdater scheduler. The CLI flag takes precedence over the `ENABLE_AUTOUPDATE` environment variable and is convenient for one-off runs such as in CI jobs.

Example (one-off run via CLI):

```powershell
go run main.go --autoupdate
```

You can configure the scheduler interval using the `--autoupdate-interval` flag or the `AUTOSCHED_INTERVAL` environment variable (examples: `30m`, `1h`). Command-line flags override environment variables:

```powershell
go run main.go --autoupdate --autoupdate-interval=30m
```

For CI usage, you can create a manual GitHub Actions workflow that runs the autoupdater once; see `.github/workflows/run_autoupdate.yml` in this repository for an example.

日本語

v1 以降、このアプリケーションは autoupdater スケジューラを実行するためのコマンドラインフラグ `--autoupdate` をサポートしています。コマンドラインフラグは環境変数 `ENABLE_AUTOUPDATE` より優先され、CI ジョブなどの一時的な実行に便利です。

例（CLI による一回実行）:

```powershell
go run main.go --autoupdate
```

スケジューラの間隔は `--autoupdate-interval` フラグまたは環境変数 `AUTOSCHED_INTERVAL`（例: `30m`, `1h`）で設定できます。コマンドラインフラグが環境変数より優先されます:

```powershell
go run main.go --autoupdate --autoupdate-interval=30m
```

CI で利用する場合、autoupdater を一度だけ実行する手動トリガー型の GitHub Actions ワークフローを作成できます。リポジトリ内の `.github/workflows/run_autoupdate.yml` を参照してください。

### Setting up GitHub Actions for Autoupdate

English

To safely enable autoupdate in GitHub Actions and optionally create GitHub issues, follow these steps.

1) Add repository secrets

Go to your repository's Settings → Secrets and variables → Actions and add the following secrets:

- `JWT_SECRET`: a long random string used to sign JWTs (e.g. `openssl rand -hex 32`).
- `GITHUB_TOKEN`: a Personal Access Token (PAT) with `repo:issues` permission if you want the updater to create issues. Prefer using the built-in `secrets.GITHUB_TOKEN` where possible; if a PAT is needed, create one with the minimum scope required.
- `GITHUB_REPO`: the target repository in `owner/repo` format (e.g. `yourusername/yourrepo`). This tells the updater where to create issues.

Security note: Do not hard-code these values in code. Rotate PATs regularly and remove them when no longer needed.

2) Trigger the workflow

- Go to the repository's Actions tab.
- Select the `Run Autoupdate (manual)` workflow.
- Click **Run workflow**.
- By default the workflow runs a safe (mock) updater (`create_issues=no`). To enable real issue creation, re-run with `create_issues=yes` and ensure required secrets are configured.

3) Safety test steps

Before enabling issue creation on a production repository, validate the workflow in a test repository:

1. Create a new private repository for testing.
2. Push the code to the test repository.
3. Add the secrets listed above to the test repository (use a test PAT if needed).
4. Run the workflow with `create_issues=yes` and verify that issues are created as expected.
5. After verification, delete the test repository or clean up test issues.
6. Once confident, repeat the steps in the production repository.

If the workflow fails, check the Actions logs for details (common causes: insufficient token scopes or network access).

日本語

GitHub Actions で autoupdate を安全に実行し、必要に応じて GitHub Issue を作成するには、次の手順に従ってください。

1) リポジトリシークレットの追加

リポジトリの Settings → Secrets and variables → Actions に移動し、次のシークレットを追加します：

- `JWT_SECRET`: JWT の署名に使用する長いランダム文字列（例: `openssl rand -hex 32`）。
- `GITHUB_TOKEN`: Issue 作成が必要な場合、`repo:issues` 権限を持つ Personal Access Token (PAT)。可能であれば `secrets.GITHUB_TOKEN`（ワークフロー組み込みトークン）を利用してください。PAT を作成する場合は必要最小限のスコープにしてください。
- `GITHUB_REPO`: 対象リポジトリを `owner/repo` 形式で指定（例: `yourusername/yourrepo`）。

セキュリティ注意: これらの値をコードにハードコーディングしないでください。PAT は定期的にローテーションし、不要時は削除してください。

2) ワークフローのトリガー

- リポジトリの Actions タブを開きます。
- `Run Autoupdate (manual)` ワークフローを選択します。
- **Run workflow** を押します。
- デフォルトでは `create_issues=no`（安全なモック実行）で動作します。Issue 作成を有効にするには `create_issues=yes` で再実行し、必要なシークレットが設定されていることを確認してください。

3) 安全テスト手順

本番リポジトリで直接 Issue を作成する前に、テスト用リポジトリで動作確認を行ってください：

1. テスト用のプライベートリポジトリを作成します。
2. コードをテストリポジトリにプッシュします。
3. 上述のシークレットをテストリポジトリに追加します（テスト用 PAT を使用）。
4. `create_issues=yes` でワークフローを実行し、Issue が正しく作成されるか確認します。
5. 検証後、テスト用リポジトリを削除するか、作成した Issue をクリーンアップします。
6. 問題なければ本番リポジトリで同じ手順を実行します。

ワークフローが失敗する場合、Actions のログを確認してください。よくある原因はトークンのスコープ不足やネットワークアクセスです。

### JWT secret

The application reads the `JWT_SECRET` environment variable at startup. If not set, a default (insecure) secret will be used and a warning will be logged. For CI and production, set `JWT_SECRET` to a long random value.

Example (PowerShell):

```powershell
$env:JWT_SECRET = "a-very-secret-value"
go run main.go
```

## Development - Seeding & WebSocket testing

During development it's convenient to populate the local database with example users, projects and fundings, and to test real-time broadcasts.

1) Dev-only HTTP seeder endpoint

- The server exposes a development-only endpoint `POST /dev/seed` which runs the seeder. This endpoint is only registered when the environment variable `ENABLE_DEV_ENDPOINTS` is set to `true`.

Example (PowerShell):

```powershell
$env:ENABLE_DEV_ENDPOINTS = 'true'
go run .

# then in another shell:
Invoke-RestMethod -Method Post -Uri http://localhost:8080/dev/seed
```

The `Seed()` function is idempotent and safe to call multiple times — it checks for existing users and projects by email/title before creating new records.

2) CLI seeder

There is also a small CLI wrapper you can run directly:

```powershell
go run tools/seed/cmd/seed_data
```

3) WebSocket selftest

To test real-time broadcasts (project published / project funded), a simple test client is provided at `tools/selftest/ws_client.go`.

Run it like this (new terminal):

```powershell
go run tools/selftest/ws_client.go
```

Now trigger a funding or publish event (use a valid JWT):

```powershell
#$token = (Invoke-RestMethod -Method Post -Uri http://localhost:8080/login -Body (@{email='carol@example.com'; password='password3'} | ConvertTo-Json) -ContentType 'application/json').token
#$headers = @{ Authorization = "Bearer $token" }
Invoke-RestMethod -Method Post -Uri http://localhost:8080/projects/9/fund -Body (@{amount=50} | ConvertTo-Json) -ContentType 'application/json' -Headers $headers
```

The WebSocket client will print received JSON messages such as:

```
{"type":"project_funded","project_id":9,"raised":200,"goal":2000}
```

4) Notes

- Keep `ENABLE_DEV_ENDPOINTS` disabled in production. The dev seeder is intended only for local development.
- If port `8080` is in use, stop the process using it or change the server port in `main.go`.

## API Documentation
# CrowdfundingZ — Quick Start & API (English / 日本語)

English: A compact local-first crowdfunding backend built with Go (Gin + GORM). The server exposes a REST API and a WebSocket endpoint. This README provides quick start instructions, API examples and notes for developers.

日本語: Go（Gin + GORM）で構築された軽量なローカル優先クラウドファンディングバックエンド。REST API と WebSocket を公開します。本 README はクイックスタート、API 例、開発者向け注記を含みます。

---

## Quick Start (English)

- Requirements: Go 1.21+, Git
- Run locally:

```powershell
git clone https://github.com/hrbzhq/crowdfunding_z.git
cd crowdfunding_z
go mod tidy
$env:JWT_SECRET = "a-secure-random-secret"
CGO_ENABLED=1 go run main.go
```

The server listens on `http://localhost:8080`.

## Quick Start (日本語)

- 前提: Go 1.21+, Git
- ローカルで実行:

```powershell
git clone https://github.com/hrbzhq/crowdfunding_z.git
cd crowdfunding_z
go mod tidy
$env:JWT_SECRET = "a-secure-random-secret"
CGO_ENABLED=1 go run main.go
```

サーバーは `http://localhost:8080` で待ち受けます。

---

## API overview / API 概要

Authentication / 認証

- Register / 登録
  - POST `/register`
  - Body: `{ "username":"alice", "email":"alice@example.com", "password":"pass" }`

- Login / ログイン
  - POST `/login`
  - Body: `{ "email":"alice@example.com", "password":"pass" }`
  - Response contains `token` (JWT) — use `Authorization: Bearer <token>` for protected endpoints.

Projects / プロジェクト

- GET `/projects` — list public (published) projects; authenticated request returns owner drafts as well.
- POST `/projects` — create project (authenticated). Default status = `draft`.
- POST `/projects/:id/publish` — publish project (owner only).
- POST `/projects/:id/fund` — fund a published project (authenticated).
- GET `/projects/:id/progress` — get progress summary.

WebSocket / リアルタイム

- `ws://localhost:8080/ws` — connect to receive live updates.

---

## Self-test (demo) / 自動デモ

The repo contains `tools/selftest/selftest.go` which runs an end-to-end demo using an in-memory DB. To run:

```powershell
go run tools/selftest/selftest.go
```

---

## Developer notes / 開発メモ

- CLI: `--autoupdate` enables the autoupdater scheduler; `--autoupdate-interval` or `AUTOSCHED_INTERVAL` configure the run interval.
- Config via env: `DATABASE_URL`, `JWT_SECRET`, `ENABLE_AUTOUPDATE`, `AUTOFETCH_URLS`, `GITHUB_TOKEN`, `GITHUB_REPO`.

---

## Contributing / 貢献

Please open issues or PRs on the repository. Run tests with `go test ./...`.

---

## License / ライセンス

MIT

---

日本語版詳細について翻訳や追加が必要ならお知らせください。
