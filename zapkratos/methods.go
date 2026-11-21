package zapkratos

import "log/slog"

// NewLogger builds a *slog.Logger with the given caption, pass it to kratos.Logger(...)
// Each time it builds a new one on top of the underlying Zap, see NewSlogLogger
//
// NewLogger 使用给定的消息说明构造 *slog.Logger，传给 kratos.Logger(...) 使用
// 每次调用都基于底层 Zap 新建一个，见 NewSlogLogger
func (A *ZapKratos) NewLogger(msgCaption string) *slog.Logger {
	return NewSlogLogger(A.Zap().LOG, msgCaption)
}
