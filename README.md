# CrowdfundingZ

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go CI](https://github.com/hrbzhq/crowdfunding_z/actions/workflows/go.yml/badge.svg)](https://github.com/hrbzhq/crowdfunding_z/actions)
[![Release](https://img.shields.io/github/v/release/hrbzhq/crowdfunding_z?label=release)](https://github.com/hrbzhq/crowdfunding_z/releases/latest)

A modern, scalable crowdfunding platform built with Go, featuring user authentication, real-time updates, and blockchain integration for decentralized funding.

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

在 GitHub Actions 中，你可以创建一个手动触发的 workflow 来运行一次 autoupdate：见 `.github/workflows/run_autoupdate.yml`（仓库内示例）。

### JWT secret

The application reads the `JWT_SECRET` environment variable at startup. If not set, a default (insecure) secret will be used and a warning will be logged. For CI and production, set `JWT_SECRET` to a long random value.

Example (PowerShell):

```powershell
$env:JWT_SECRET = "a-very-secret-value"
go run main.go
```

## API Documentation

### Authentication Endpoints

#### Register User
```http
POST /register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

#### Login User
```http
POST /login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "securepassword"
}
```

### Project Endpoints

#### Get All Projects
```http
GET /projects
Authorization: Bearer <jwt_token>
```

#### Create Project
```http
POST /projects
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "title": "My Awesome Project",
  "description": "A detailed description of the project",
  "goal": 10000.00,
  "deadline": "2025-12-31"
}
```

#### Fund Project
```http
POST /projects/{id}/fund
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "amount": 100.00,
  "user_id": 1
}
```

#### Get Project Progress
```http
GET /projects/{id}/progress
```

### WebSocket Endpoint

#### Real-time Updates
```javascript
const ws = new WebSocket('ws://localhost:8080/ws');
// Listen for messages
ws.onmessage = (event) => {
  console.log('Received:', event.data);
};
```

## Contributing

We welcome contributions from the community! Here's how you can help:

### Development Setup
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes and add tests
4. Run tests: `go test ./...`
5. Commit your changes: `git commit -am 'Add new feature'`
6. Push to the branch: `git push origin feature/your-feature`
7. Submit a pull request

### Code Standards
- Follow Go naming conventions
- Add comments for exported functions
- Write unit tests for new features
- Update documentation as needed

### Reporting Issues
- Use GitHub Issues for bug reports and feature requests
- Provide detailed steps to reproduce bugs
- Include relevant code snippets and error messages

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

- **Project Lead**: Your Name
- **Email**: your.email@example.com
- **GitHub**: [@yourusername](https://github.com/yourusername)
- **LinkedIn**: [Your LinkedIn](https://linkedin.com/in/yourprofile)

## Code of Conduct

Please read our [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) before contributing.

## Security

If you discover a security vulnerability, please report it privately via the contact email in [SECURITY.md](SECURITY.md).

### Community
- **Discord**: [Join our Discord](https://discord.gg/crowdfundingz)
- **Twitter**: [@CrowdfundingZ](https://twitter.com/CrowdfundingZ)
- **Forum**: [Community Forum](https://forum.crowdfundingz.com)

---

**CrowdfundingZ** - Democratizing funding through technology and transparency.

## GitHub SSH 推送指南

本项目支持通过 SSH 推送代码到 GitHub，以下是配置流程：

### 1. 生成 SSH 密钥
```
ssh-keygen -t ed25519 -C "hrbzhq@163.com" -f ~/.ssh/id_ed25519 -N ""
```

### 2. 添加公钥到 GitHub
复制并执行：
```
cat ~/.ssh/id_ed25519.pub
```
然后访问 GitHub → Settings → SSH and GPG keys → New SSH key，把复制的公钥粘贴进去并保存。

### 3. 测试连接
```
ssh -T git@github.com
```

### 4. 设置远程仓库地址并推送
```
git remote set-url origin git@github.com:hrbzhq/crowdfunding_z.git
git add .
git commit -m "Add init script and SSH guide"
git push origin master
```

### 一键初始化脚本（可选）
仓库根含 `init_git.sh`，运行它会生成 SSH key、显示公钥并尝试初始化/推送仓库：
```
bash init_git.sh
```

注意：在 CI/生产环境请使用更安全的密钥管理方式，勿在公共场合暴露私钥。

## GitHub SSH Push Guide (English)

This project supports pushing to GitHub over SSH. Follow the steps below to configure SSH and push the repository using SSH authentication.

### 1. Generate an SSH key
Run this command (replace the email if needed):
```bash
ssh-keygen -t ed25519 -C "hrbzhq@163.com" -f ~/.ssh/id_ed25519 -N ""
```

### 2. Add the public key to GitHub
Display and copy your public key:
```bash
cat ~/.ssh/id_ed25519.pub
```
Then go to GitHub → Settings → SSH and GPG keys → New SSH key, paste the public key and save.

### 3. Test the SSH connection
```bash
ssh -T git@github.com
```
You should see a success message like:
```
Hi <username>! You've successfully authenticated, but GitHub does not provide shell access.
```

### 4. Set the remote URL and push
```bash
git remote set-url origin git@github.com:hrbzhq/crowdfunding_z.git
git add .
git commit -m "Add init script and SSH guide"
git push origin master
```

### One-step init script (optional)
There is a convenience script `init_git.sh` in the repo root which will generate an SSH key, print the public key, and attempt to initialize/push the repository:
```bash
bash init_git.sh
```

Note: For CI and production use secure key management; never expose your private key publicly.
