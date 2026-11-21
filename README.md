<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/kratos-zap/release.yml?branch=main&label=BUILD)](https://github.com/yylego/kratos-zap/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/kratos-zap)](https://pkg.go.dev/github.com/yylego/kratos-zap)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/kratos-zap/main.svg)](https://coveralls.io/github/yylego/kratos-zap?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/kratos-zap.svg)](https://github.com/yylego/kratos-zap/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/kratos-zap)](https://goreportcard.com/report/github.com/yylego/kratos-zap)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# kratos-zap

Zap integration that connects Uber Zap with Kratos microservice framework, providing structured logging with high performance.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

## Main Features

- 🚀 Simple Integration - Replace Kratos default logging with just a few lines
- 📊 Structured Logging - Use Uber Zap's fast structured logging
- ⚡ High Performance - Benefit from Zap's zero-allocation design
- 🔧 Flexible Config - Build config with custom options
- 🎯 Module Tracking - Auto add module info to logs
- ⚙️ Adaptation Pattern - Bridge Zap and Kratos with ease

## Installation

```bash
go get github.com/yylego/kratos-zap/zapkratos
```

## Quick Start

```go
package main

import (
    "github.com/go-kratos/kratos/v3"
    "github.com/yylego/kratos-zap/zapkratos"
    "github.com/yylego/zaplog"
    "go.uber.org/zap"
)

func main() {
    // Create ZapKratos instance
    zapKratos := zapkratos.NewZapKratos(
        zaplog.LOGGER,
        zapkratos.NewOptions(),
    )

    // Get logging with module context
    zapLog := zapKratos.SubZap()
    zapLog.LOG.Info("application starting...")

    // Use in Kratos app
    app := kratos.New(
        kratos.Name("my-service"),
        kratos.Logger(zapKratos.NewLogger("app")),
    )

    if err := app.Run(); err != nil {
        zapLog.LOG.Fatal("app run failed", zap.Error(err))
    }
}
```

## Complete Examples

See [kratos-examples/zap](https://github.com/kratos-examples/zap) to view complete integration inside Kratos projects:

- **[demo1kratos](https://github.com/kratos-examples/zap/tree/main/demo1kratos)** - Basic integration with HTTP and gRPC
- **[demo2kratos](https://github.com/kratos-examples/zap/tree/main/demo2kratos)** - Advanced usage with Wire DI

## API Reference

### ZapKratos

Main struct that wraps Zap logging and provides Kratos-compatible interfaces.

```go
type ZapKratos struct {
    // Contains filtered or unexported fields
}
```

#### Constructor

```go
func NewZapKratos(zap *zaplog.Zap, options *Options) *ZapKratos
```

Creates new ZapKratos instance with the given Zap and options.

#### Methods

```go
func (A *ZapKratos) Zap() *zaplog.Zap
```

Returns the underlying Zap instance.

```go
func (A *ZapKratos) SubZap() *zaplog.Zap
```

Creates child Zap with calling module info.

```go
func (A *ZapKratos) NewLogger(msgCaption string) *slog.Logger
```

Builds a \*slog.Logger with the given caption, pass it to `kratos.Logger(...)`. Each invocation builds a new one on top of the underlying Zap.

### Options

Config options to ZapKratos.

```go
type Options struct {
    ModuleKeyName string // Module field key name in log output
}
```

#### Constructor

```go
func NewOptions() *Options
```

Creates Options with default settings (module field = "module").

#### Methods

```go
func (T *Options) WithModuleKeyName(moduleKeyName string) *Options
```

Sets custom module field name in a chainable fashion.

### NewSlogLogger

Builds a \*slog.Logger backed by Zap via the official zapslog handler.

#### Constructor

```go
func NewSlogLogger(zapLog *zap.Logger, msgCaption string) *slog.Logger
```

Bridges the given Zap to stdlib slog, tagging each log with the caption. zapslog takes the Zap core alone, so the code position of each log and the stack trace from ERROR are switched on in the bridge itself.

## Dependencies

- `go.uber.org/zap` - Uber Zap structured logging
- `go.uber.org/zap/exp/zapslog` - Official Zap-to-slog bridge
- `github.com/yylego/zaplog` - Zap management package

Targets the Kratos v3 framework: the returned \*slog.Logger plugs straight into `kratos.Logger(...)`.

## Related Projects

**Frameworks:**

- [Kratos](https://github.com/go-kratos/kratos) - Go microservice framework
- [Zap](https://github.com/uber-go/zap) - Uber's structured logging

**kratos-zap Ecosystem:**

- [kratos-zap](https://github.com/yylego/kratos-zap) - This project
- [kratos-examples/zap](https://github.com/kratos-examples/zap) - Demo projects
- [kratos-zapzh](https://github.com/yylego/kratos-zapzh) - Chinese version with Chinese function names
- [kratos-examples/zapzh](https://github.com/kratos-examples/zapzh) - Chinese version demos

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
