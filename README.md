[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/kratos-zap/release.yml?branch=main&label=BUILD)](https://github.com/yylego/kratos-zap/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/kratos-zap)](https://pkg.go.dev/github.com/yylego/kratos-zap)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/kratos-zap/main.svg)](https://coveralls.io/github/yylego/kratos-zap?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/kratos-zap.svg)](https://github.com/yylego/kratos-zap/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/kratos-zap)](https://goreportcard.com/report/github.com/yylego/kratos-zap)

# kratos-zap

Zap integration that connects Uber Zap with Kratos microservice framework, providing structured logging with high performance.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

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

See [kratos-zap-demos](https://github.com/yylego/kratos-zap-demos) to view complete integration in actual Kratos projects:

- **[demo1kratos](https://github.com/yylego/kratos-zap-demos/tree/main/demo1kratos)** - Basic integration with HTTP and gRPC
- **[demo2kratos](https://github.com/yylego/kratos-zap-demos/tree/main/demo2kratos)** - Advanced usage with Wire DI

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
func (A *ZapKratos) GetZap() *zaplog.Zap
```

Returns the underlying Zap instance.

```go
func (A *ZapKratos) SubZap() *zaplog.Zap
```

Creates child Zap with calling module info.

```go
func (A *ZapKratos) GetLogger(msgCaption string) *slog.Logger
```

Builds a \*slog.Logger with the given caption.

```go
func (A *ZapKratos) NewLogger(msgCaption string) *slog.Logger
```

Same as GetLogger, builds the \*slog.Logger.

```go
func (A *ZapKratos) GetSlogLogger(msgCaption string) *slog.Logger
```

Explicit-named accessor mirroring the NewSlogLogger package func.

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

Bridges the given Zap to stdlib slog, tagging each entry with the caption.

## Dependencies

- `go.uber.org/zap` - Uber Zap structured logging
- `go.uber.org/zap/exp/zapslog` - Official Zap-to-slog bridge
- `github.com/yylego/zaplog` - Zap management package
- `github.com/yylego/runpath` - Runtime path utilities

Targets the Kratos v3 framework: the returned \*slog.Logger plugs straight into `kratos.Logger(...)`.

## Related Projects

**Frameworks:**

- [Kratos](https://github.com/go-kratos/kratos) - Go microservice framework
- [Zap](https://github.com/uber-go/zap) - Uber's structured logging

**kratos-zap Ecosystem:**

- [kratos-zap](https://github.com/yylego/kratos-zap) - This project
- [kratos-zap-demos](https://github.com/yylego/kratos-zap-demos) - Demo projects
- [kratos-zapzh](https://github.com/yylego/kratos-zapzh) - Chinese version with Chinese function names
- [kratos-zapzh-demos](https://github.com/yylego/kratos-zapzh-demos) - Chinese version demos

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE) file

---

## 💬 Contact & Feedback

**Issues & Feedback:**

- 🐛 **Bug reports?** Open an issue and describe the problem with reproduction steps
- ✨ **Feature ideas?** Open an issue to discuss the implementation approach
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
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/kratos-zap.git`).
3. **Navigate**: Navigate to the cloned project (`cd kratos-zap`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
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

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/kratos-zap.svg?variant=adaptive)](https://starchart.cc/yylego/kratos-zap)
