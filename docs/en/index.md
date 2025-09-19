````markdown
# CrowdfundingZ — Quick Start & Docs (English)

This repository contains a local-first crowdfunding backend written in Go (Gin + GORM). Use this document for a concise English reference.

## Quick Start

- Requirements: Go 1.21+, Git

```powershell
git clone https://github.com/hrbzhq/crowdfunding_z.git
cd crowdfunding_z
go mod tidy
$env:JWT_SECRET = "a-secure-random-secret"
CGO_ENABLED=1 go run main.go
```

The server listens on `http://localhost:8080`.

## Autoupdate (CLI)

Starting from v1, the application supports the `--autoupdate` command-line flag to run the autoupdater scheduler. The CLI flag takes precedence over the `ENABLE_AUTOUPDATE` environment variable and is convenient for one-off runs such as in CI jobs.

Example (one-off run via CLI):

```powershell
go run main.go --autoupdate
```

Configure the scheduler interval using `--autoupdate-interval` or `AUTOSCHED_INTERVAL` (examples: `30m`, `1h`).

## GitHub Actions — Autoupdate (manual)

A manual workflow is provided at `.github/workflows/run_autoupdate.yml` which runs the updater in a safe mock mode by default. To enable real GitHub issue creation, provide the necessary secrets and run the workflow with `create_issues=yes`.

- Required secrets (add in Settings → Secrets → Actions):
  - `JWT_SECRET` — a long random string for signing JWTs.
  - `GITHUB_TOKEN` — a PAT with `repo:issues` permission (prefer `secrets.GITHUB_TOKEN` where possible).
  - `GITHUB_REPO` — `owner/repo` to tell the updater where to create issues.

Security note: prefer `secrets.GITHUB_TOKEN` (the repository's built-in token) when it suffices, otherwise create a PAT with the least privileges.

## Seeding & WebSocket testing

Local seeder: `POST /dev/seed` (enabled when `ENABLE_DEV_ENDPOINTS=true`)

CLI seeder: `go run tools/seed/cmd/seed_data`

WebSocket test client: `go run tools/selftest/ws_client.go`

---

For full API details, see `README.md` in the repository root.

````
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
