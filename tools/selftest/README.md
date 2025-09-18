# Self-test tool

This small tool runs an end-to-end self-test of the crowdfunding backend using an in-memory SQLite database.

It exercises the following flow:
- Register
- Login
- Create Project
- Fund Project
- Get Progress

Usage

Run locally:

```powershell
# from project root
go run tools/selftest/selftest.go
```

To force a file-based temporary DB instead of in-memory, set `DATABASE_URL` before running:

```powershell
$env:DATABASE_URL = "file:selftest.db?cache=shared&mode=rwc"
go run tools/selftest/selftest.go
```

CI

You can add the self-test to CI by running `go run tools/selftest/selftest.go` in your workflow.
