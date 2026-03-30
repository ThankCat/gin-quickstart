# gin-quickstart

一个基于 [Gin](https://github.com/gin-gonic/gin) 的最小可运行示例，包含：

- 路由分组（`/v1`、`/v2`）
- 简单登录接口（`POST /v1/login`、`POST /v2/login`）
- 中间件鉴权示例（`/v2` 分组）
- 使用 [Air](https://github.com/air-verse/air) 的本地热重载开发

## 环境要求

- Go `1.25.6`（与 `go.mod` 对齐）

## 快速开始

1. 安装依赖

```bash
go mod tidy
```

2. 直接运行

```bash
go run main.go
```

默认监听：`http://127.0.0.1:8080`

## 接口说明

### `POST /v1/login`

无鉴权，返回：

```json
{"action":"login"}
```

### `POST /v2/login`

经过 `AuthRequired` 中间件，当前示例会直接返回 403：

```json
{"msg":"你没有权限"}
```

## 使用 Air 热重载（推荐开发时）

本项目已包含 `.air.toml` 配置，关键行为：

- 编译命令：`go build -o ./tmp/main .`
- 运行入口：`./tmp/main`
- 构建日志：`tmp/build-errors.log`
- 监听扩展：`go`, `tpl`, `tmpl`, `html`

### 1) 安装 Air

```bash
go install github.com/air-verse/air@latest
```

确保你的 `$GOPATH/bin`（或 `$(go env GOPATH)/bin`）在 `PATH` 中。

### 2) 启动热重载

```bash
air
```

Air 启动后会根据 `.air.toml` 自动编译并运行项目，代码变更后自动重启。

## 上传到 GitHub 前建议

- 已通过 `.gitignore` 忽略本地产物（如 `tmp/`、日志、环境文件）
- 建议在首次提交前执行：

```bash
git init
git add .
git commit -m "init: gin quickstart with air"
```

然后在 GitHub 创建仓库并推送即可。
