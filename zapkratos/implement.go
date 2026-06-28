package zapkratos

import (
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

// NewSlogLogger builds a *slog.Logger backed by the given Zap, tagging each entry
// with the caption so callers can tell apart logs from distinct call sites
// Bridges Uber Zap to the stdlib slog via the official zapslog handler
//
// NewSlogLogger 基于给定 Zap 构造 *slog.Logger，给每条日志打上 caption 标记
// 让调用方能区分不同来源的日志
// 通过官方 zapslog handler 把 Uber Zap 桥接到标准库 slog
func NewSlogLogger(zapLog *zap.Logger, msgCaption string) *slog.Logger {
	handler := zapslog.NewHandler(zapLog.Core())
	return slog.New(handler).With("caption", msgCaption)
}
