[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/kratos-zap/release.yml?branch=main&label=BUILD)](https://github.com/yylego/kratos-zap/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/kratos-zap)](https://pkg.go.dev/github.com/yylego/kratos-zap)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/kratos-zap/main.svg)](https://coveralls.io/github/yylego/kratos-zap?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/kratos-zap.svg)](https://github.com/yylego/kratos-zap/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/kratos-zap)](https://goreportcard.com/report/github.com/yylego/kratos-zap)

# kratos-zap

将 Uber Zap 与 Kratos 微服务框架集成的日志适配器，提供高性能的结构化日志记录。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 主要特性

- 🚀 简单集成 - 只需几行代码即可替换 Kratos 默认日志功能
- 📊 结构化日志 - 使用 Uber Zap 快速的结构化日志能力
- ⚡ 高性能 - 受益于 Zap 的零内存分配设计
- 🔧 灵活配置 - 构建配置使用自定义选项
- 🎯 模块追踪 - 自动添加模块信息到日志
- ⚙️ 适配模式 - 轻松桥接 Zap 和 Kratos

## 安装

```bash
go get github.com/yylego/kratos-zap/zapkratos
```

## 快速开始

```go
package main

import (
    "github.com/go-kratos/kratos/v3"
    "github.com/yylego/kratos-zap/zapkratos"
    "github.com/yylego/zaplog"
)

func main() {
    // 创建 ZapKratos 实例
    zapKratos := zapkratos.NewZapKratos(
        zaplog.LOGGER,
        zapkratos.NewOptions(),
    )

    // 获取带模块上下文的日志
    zapLog := zapKratos.SubZap()
    zapLog.LOG.Info("application starting...")

    // 在 Kratos 应用中使用
    app := kratos.New(
        kratos.Name("my-service"),
        kratos.Logger(zapKratos.NewLogger("app")),
    )

    if err := app.Run(); err != nil {
        zapLog.LOG.Fatal("app run failed", zap.Error(err))
    }
}
```

## 完整示例

查看 [kratos-zap-demos](https://github.com/yylego/kratos-zap-demos) 了解实际 Kratos 项目中的完整集成：

- **[demo1kratos](https://github.com/yylego/kratos-zap-demos/tree/main/demo1kratos)** - HTTP 和 gRPC 基础集成
- **[demo2kratos](https://github.com/yylego/kratos-zap-demos/tree/main/demo2kratos)** - Wire 依赖注入高级用法

## API 参考

### ZapKratos

包装 Zap 日志并提供 Kratos 兼容接口的主要结构体。

```go
type ZapKratos struct {
    // 包含已过滤或未导出的字段
}
```

#### 构造函数

```go
func NewZapKratos(zap *zaplog.Zap, options *Options) *ZapKratos
```

使用给定的 Zap 和选项创建新的 ZapKratos 实例。

#### 方法

```go
func (A *ZapKratos) GetZap() *zaplog.Zap
```

返回底层的 Zap 实例。

```go
func (A *ZapKratos) SubZap() *zaplog.Zap
```

创建带有调用模块信息的子 Zap。

```go
func (A *ZapKratos) GetLogger(msgCaption string) *slog.Logger
```

使用给定的消息说明构造 \*slog.Logger。

```go
func (A *ZapKratos) NewLogger(msgCaption string) *slog.Logger
```

与 GetLogger 相同，构造 \*slog.Logger。

```go
func (A *ZapKratos) GetSlogLogger(msgCaption string) *slog.Logger
```

跟包级 NewSlogLogger 对应的显式命名访问方法。

### Options

ZapKratos 的配置选项。

```go
type Options struct {
    ModuleKeyName string // 日志输出中的模块字段键名
}
```

#### 构造函数

```go
func NewOptions() *Options
```

创建具有默认设置的 Options（模块键 = "module"）。

#### 方法

```go
func (T *Options) WithModuleKeyName(moduleKeyName string) *Options
```

以可链式调用的方式设置自定义模块字段键名。

### NewSlogLogger

通过官方 zapslog handler 基于 Zap 构造 \*slog.Logger。

#### 构造函数

```go
func NewSlogLogger(zapLog *zap.Logger, msgCaption string) *slog.Logger
```

把给定 Zap 桥接到标准库 slog，并给每条日志打上 caption 标记。

## 依赖项

- `go.uber.org/zap` - Uber Zap 结构化日志
- `go.uber.org/zap/exp/zapslog` - 官方 Zap 转 slog 桥接
- `github.com/yylego/zaplog` - Zap 管理包
- `github.com/yylego/runpath` - 运行时路径工具

面向 Kratos v3 框架：返回的 \*slog.Logger 可直接传给 `kratos.Logger(...)`。

## 相关项目

**框架：**

- [Kratos](https://github.com/go-kratos/kratos) - Go 微服务框架
- [Zap](https://github.com/uber-go/zap) - Uber 的结构化日志

**kratos-zap 生态：**

- [kratos-zap](https://github.com/yylego/kratos-zap) - 本项目
- [kratos-zap-demos](https://github.com/yylego/kratos-zap-demos) - 演示项目
- [kratos-zapzh](https://github.com/yylego/kratos-zapzh) - 使用中文函数名的中文版本
- [kratos-zapzh-demos](https://github.com/yylego/kratos-zapzh-demos) - 中文版演示项目

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证

MIT License - 查看 [LICENSE](LICENSE) 文件

---

## 💬 联系反馈

**问题和反馈：**

- 🐛 **Bug 报告？** 打开 issue 并描述问题和复现步骤
- ✨ **功能想法？** 打开 issue 讨论实现方案
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/kratos-zap.git`）
3. **导航**：进入克隆的项目（`cd kratos-zap`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细说明

请确保测试通过并包含相关的文档更新。

---

## 🌟 支持

欢迎通过提交 Merge Request 和报告 Issue 为此项目做贡献。

**项目支持：**

- ⭐ **给 GitHub 星标** 如果这个项目帮助了你
- 🤝 **分享给队友** 和（golang）编程朋友
- 📝 **写技术博客** 关于开发工具和工作流程 - 我们提供内容写作支持
- 🌟 **加入生态系统** - 致力于支持开源和（golang）开发场景

**用这个包快乐编程！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![标星点赞](https://starchart.cc/yylego/kratos-zap.svg?variant=adaptive)](https://starchart.cc/yylego/kratos-zap)
