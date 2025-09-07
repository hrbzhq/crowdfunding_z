# Crowdfunding Platform

A Go-based crowdfunding platform with user authentication, project management, payment integration, and real-time updates.

## Features

- User registration and login with JWT
- Project creation and funding
- Real-time progress updates via WebSocket
- Database integration with SQLite
- Modular architecture

## Setup

1. Install Go 1.21 or later from https://go.dev/dl/

2. Install MinGW or GCC for CGO support (required for SQLite)

3. Run `go mod tidy` to download dependencies

4. Run `go run main.go` to start the server

## API Endpoints

- POST /register - Register a new user
- POST /login - Login user
- GET /projects - Get all projects
- POST /projects - Create a new project
- POST /projects/:id/fund - Fund a project
- GET /projects/:id/progress - Get project progress
- GET /ws - WebSocket for real-time updates

## Architecture

- `main.go` - Entry point
- `handlers/` - HTTP handlers
- `models/` - Database models
- `database/` - Database initialization
- `blockchain/` - Blockchain integration

## Future Enhancements

- Payment gateway integration (Stripe, PayPal)
- Notification system with RabbitMQ
- Blockchain support for decentralized crowdfunding
- Data analysis and reporting

## Running the Application

To run the application, use:

```bash
CGO_ENABLED=1 go run main.go
```

If you encounter CGO issues, install MinGW GCC:

```bash
winget install BrechtSanders.WinLibs.POSIX.UCRT
```

Then restart your terminal and run the command above.
