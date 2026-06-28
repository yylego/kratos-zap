package zapkratos

import "log/slog"

// GetLogger builds a *slog.Logger with the given caption
// Wraps the underlying Zap and returns the created logger
//
// GetLogger 使用给定的消息说明构造 *slog.Logger
// 包装底层 Zap 日志器并返回实例
func (A *ZapKratos) GetLogger(msgCaption string) *slog.Logger {
	return NewSlogLogger(A.GetZap().LOG, msgCaption)
}

// NewLogger builds a *slog.Logger with the given caption
// Wraps the underlying Zap and returns the created logger
// Note: Same as GetLogger, provided to match naming patterns
//
// NewLogger 使用给定的消息说明构造 *slog.Logger
// 包装底层 Zap 日志器并返回实例
// 注意：与 GetLogger 相同，提供以匹配命名模式
func (A *ZapKratos) NewLogger(msgCaption string) *slog.Logger {
	return NewSlogLogger(A.GetZap().LOG, msgCaption)
}

// GetSlogLogger builds a *slog.Logger with the given caption
// Explicit-named accessor mirroring the NewSlogLogger package func
//
// GetSlogLogger 使用给定的消息说明构造 *slog.Logger
// 跟包级 NewSlogLogger 对应的显式命名访问方法
func (A *ZapKratos) GetSlogLogger(msgCaption string) *slog.Logger {
	return NewSlogLogger(A.GetZap().LOG, msgCaption)
}
