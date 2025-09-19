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

从 v1 起，应用支持一个命令行标志 `--autoupdate`，用于启用 autoupdater 调度器。命令行标志优先于 `ENABLE_AUTOUPDATE` 环境变量，适合在 CI/一次性运行中使用。

示例（通过 CLI 在 CI 中一次性运行 autoupdater）：

```powershell
go run main.go --autoupdate
```

可以通过 `--autoupdate-interval` 或 `AUTOSCHED_INTERVAL` 环境变量配置调度器间隔（例如 `30m`、`1h`）。命令行标志优先于环境变量：

```powershell
go run main.go --autoupdate --autoupdate-interval=30m
```

在 GitHub Actions 中，你可以创建一个手动触发的 workflow 来运行一次 autoupdate：见 `.github/workflows/run_autoupdate.yml`（仓库内示例）。

### Setting up GitHub Actions for Autoupdate

要安全地在 GitHub Actions 中启用 autoupdate 并可选创建 GitHub issues，请按以下步骤配置：

#### 1. 添加仓库 Secrets
在你的 GitHub 仓库中，导航到 **Settings** → **Secrets and variables** → **Actions**，添加以下 secrets：

- `JWT_SECRET`: 一个长随机字符串，用于 JWT 签名（例如，使用 `openssl rand -hex 32` 生成）。
- `GITHUB_TOKEN`: 一个具有 `repo:issues` 权限的 personal access token (PAT)。创建 PAT 时，选择最小权限：勾选 "repo" 下的 "issues"。
- `GITHUB_REPO`: 目标仓库的名称，格式为 `owner/repo`（例如 `yourusername/yourrepo`）。这指定在哪里创建 issues。

**安全提示**: 不要在代码中硬编码这些值。PAT 应定期轮换，并在不需要时删除。

#### 2. 触发 Workflow
- 转到仓库的 **Actions** 标签页。
- 选择 `Run Autoupdate (manual)` workflow。
- 点击 **Run workflow**。
- 默认情况下，`create_issues` 设置为 `no`，这将运行 autoupdater 但不会创建真实 issues（使用 MockUpdater）。
- 如果你想启用真实 issue 创建，将 `create_issues` 设置为 `yes`（仅当 secrets 已配置时有效）。

#### 3. 安全测试步骤
为了避免在生产仓库中意外创建 issues，先在测试仓库中验证：

1. 创建一个新的私有 GitHub 仓库作为测试环境。
2. 将代码推送到测试仓库。
3. 在测试仓库中添加上述 secrets（使用测试 PAT 和测试仓库名称）。
4. 触发 workflow 并设置 `create_issues=yes`，观察是否成功创建 issues。
5. 验证后，删除测试仓库或清理创建的 issues。
6. 然后在生产仓库中重复步骤 1-4。

如果 workflow 失败，检查日志以获取错误详情（例如，token 权限不足或网络问题）。

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
