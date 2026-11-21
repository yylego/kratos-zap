package zapkratos

import (
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

// NewSlogLogger builds a *slog.Logger backed by the given Zap, tagging each entry
// with the caption so callers can tell apart logs from distinct call sites
// Bridges Uber Zap to the stdlib slog via the official zapslog handler
// zapslog takes the Zap core alone, so the code position and the ERROR stack trace are switched on here again
//
// NewSlogLogger 基于给定 Zap 构造 *slog.Logger，给每条日志打上 caption 标记
// 让调用方能区分不同来源的日志
// 通过官方 zapslog handler 把 Uber Zap 桥接到标准库 slog
// zapslog 只拿到 Zap 的 core，拿不到 Zap 上的选项，所以在这里重新开启调用位置和 ERROR 起的调用栈
func NewSlogLogger(zapLog *zap.Logger, msgCaption string) *slog.Logger {
	return slog.New(zapslog.NewHandler(zapLog.Core(), zapslog.WithCaller(true), zapslog.AddStacktraceAt(slog.LevelError))).With("caption", msgCaption)
}
