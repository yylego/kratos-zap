// Package zapkratos provides Zap logger integration for the Kratos v3 microservice framework
// Bridges Uber Zap to the stdlib slog via the official zapslog handler, returning a *slog.Logger
// Enables structured logging with module tracking and flexible configuration
//
// zapkratos 包为 Kratos v3 微服务框架提供 Zap 日志集成
// 通过官方 zapslog handler 把 Uber Zap 桥接到标准库 slog，返回 *slog.Logger
// 支持结构化日志记录、模块追踪和灵活的配置选项
package zapkratos

import (
	"path/filepath"

	"github.com/yylego/runpath"
	"github.com/yylego/zaplog"
)

// Options holds config settings of the ZapKratos
// Controls module field naming and logging behavior settings
//
// Options 保存 ZapKratos 实例的配置选项
// 控制模块字段命名和其他日志行为
type Options struct {
	ModuleKeyName string // Module field key name in log output // 日志输出中的模块字段键名
}

// NewOptions creates Options instance with default settings
// Sets module key name to "module" and returns Options
//
// NewOptions 创建具有默认设置的 Options 实例
// 设置模块键名为 "module" 并返回 Options
func NewOptions() *Options {
	return &Options{
		ModuleKeyName: "module",
	}
}

// WithModuleKeyName sets custom module field key name using builder pattern
// Returns same Options instance to enable chaining of methods
//
// WithModuleKeyName 使用构建器模式设置自定义模块字段键名
// 返回相同的 Options 实例以支持方法链式调用
func (T *Options) WithModuleKeyName(moduleKeyName string) *Options {
	T.ModuleKeyName = moduleKeyName
	return T
}

// ZapKratos wraps Zap to provide logging compatible with Kratos
// Holds underlying Zap and config to create logger instances as needed
//
// ZapKratos 包装 Zap 日志器以提供 Kratos 兼容的日志功能
// 持有底层 Zap 实例和配置选项用于创建日志器实例
type ZapKratos struct {
	zap     *zaplog.Zap // Underlying Zap logger instance // 底层 Zap 日志器实例
	options *Options    // Config options // 配置选项
}

// NewZapKratos creates ZapKratos with given Zap logger and config
// Returns new ZapKratos that can provide *slog.Logger instances for Kratos v3
//
// NewZapKratos 使用给定的 Zap 日志器和选项创建 ZapKratos 实例
// 返回准备好为 Kratos v3 提供 *slog.Logger 实例的新 ZapKratos
func NewZapKratos(zap *zaplog.Zap, options *Options) *ZapKratos {
	return &ZapKratos{
		zap:     zap,
		options: options,
	}
}

// GetZap returns the underlying Zap logging instance
// Provides access to raw Zap as needed
//
// GetZap 返回底层的 Zap 日志器实例
// 在需要时提供对原始 Zap 日志器的访问
func (A *ZapKratos) GetZap() *zaplog.Zap {
	return A.zap
}

// SubZap creates child Zap with calling module information
// Auto adds module field containing file basename from calling context
// Returns new Zap with module context included
//
// SubZap 创建带有调用者模块信息的子 Zap 日志器
// 使用 runpath 自动添加包含调用者文件基本名的模块字段
// 返回附加了模块上下文的新 Zap 实例
func (A *ZapKratos) SubZap() *zaplog.Zap {
	return A.GetZap().NewZap(A.options.ModuleKeyName, filepath.Base(runpath.Skip(1)))
}
