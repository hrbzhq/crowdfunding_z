# Autoupdater module

这是一个示例的可插拔自我升级（autoupdater）模块。它提供三种接口：

- `Fetcher`：从网络或其他来源抓取现有的众筹项目数据。
- `Analyzer`：对抓取的数据进行分析，产生建议或可执行项。
- `Updater`：将建议应用到本地系统（例如创建 issue、更新配置或触发训练/微调）。

当前提供的实现：
- `MockFetcher`：返回固定的示例项目列表（无网络请求）。
- `MockAnalyzer`：基于项目进度生成文本建议。
- `MockUpdater`：将建议打印到 stdout（示例行为）。

扩展建议：
- 实现 `HTTPFetcher`：真实抓取众筹平台（注意遵守 robots.txt 和平台 API 规则）。
- 实现更复杂的 `Analyzer`：可接入 ML 模型或统计分析，输出结构化建议。
- 实现 `Updater`：自动创建 GitHub issue、提交 PR 或更新配置文件。

测试：
- 运行 `go test ./tools/autoupdater` 将执行 scheduler 的一次周期（使用内置 mock 实现）。

安全与限制：本模块默认不执行网络请求；如果实现真实抓取器，请先在本地测试并遵守目标网站的使用条款。

启用与集成说明
----------------

本模块已经集成到 `main.go`，通过环境变量控制是否启动自动更新调度器：

- `ENABLE_AUTOUPDATE`：将其设置为 `true`（不区分大小写）以启用调度器。默认关闭。
- `AUTOFETCH_URLS`：可选，逗号分隔的 URL 列表。若提供则使用 `HTTPFetcher` 抓取这些 URL（每个 URL 预期返回 JSON 数组，格式与 `FetchedProject` 匹配）。示例：

	AUTOFETCH_URLS="https://example.com/projects.json,https://another/source.json"

- `GITHUB_TOKEN`：可选，若提供且 `GITHUB_REPO` 也设置，则使用 `GitHubUpdater` 将建议发布为 issue（注意权限）。
- `GITHUB_REPO`：可选，形如 `owner/repo`，与 `GITHUB_TOKEN` 配合使用。

默认行为（安全）：
- 若 `ENABLE_AUTOUPDATE=true` 且未设置 `AUTOFETCH_URLS`，调度器使用 `MockFetcher`（不会联网）。
- 若未提供 `GITHUB_TOKEN`/`GITHUB_REPO`，调度器使用 `MockUpdater`（只打印建议）。

示例（仅使用 MockFetcher）：

```powershell
$env:ENABLE_AUTOUPDATE = "true"
go run main.go
```

示例（远程抓取并创建 GitHub issue）：

```powershell
$env:ENABLE_AUTOUPDATE = "true"
$env:AUTOFETCH_URLS = "https://example.com/projects.json"
$env:GITHUB_TOKEN = "ghp_..."
$env:GITHUB_REPO = "youruser/yourrepo"
go run main.go
```

注意事项：
- 请把 `GITHUB_TOKEN` 存为 GitHub Secrets 或安全的环境变量，不要在仓库中明文存储。
- `HTTPFetcher` 期望目标 URL 返回符合 `FetchedProject` 结构的 JSON。如果目标 API 不兼容，考虑先写一个适配层来转换数据。

扩展点
- 如果需要对抓取数据进行更智能的分析，可以实现 `Analyzer` 接口并输出结构化建议（例如包含 action-type、priority、labels），便于自动化 `Updater` 根据这些元数据创建受控的 issue / PR。

欢迎贡献更多 `Fetcher`（例如带爬虫限速的抓取器）、`Analyzer`（基于 ML 的评分器）和 `Updater`（更丰富的发布策略）。