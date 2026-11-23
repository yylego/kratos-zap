package zapkratos

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// TestZapKratos_NewLogger checks the slog log writes to the underlying Zap with the caption
// TestZapKratos_NewLogger 检查 slog 日志写进底层 Zap，并带上 caption
func TestZapKratos_NewLogger(t *testing.T) {
	core, logs := observer.New(zapcore.DebugLevel)
	zapKratos := NewZapKratos(zaplog.NewZap(zap.New(core)), NewOptions())

	zapKratos.NewLogger("app").Info("abc", "stage", "demo")

	entries := logs.TakeAll()
	require.Len(t, entries, 1)
	require.Equal(t, "abc", entries[0].Message)
	require.Equal(t, map[string]any{"caption": "app", "stage": "demo"}, entries[0].ContextMap())
}

// subZapWrapped wraps SubZapSkip(1) to serve TestZapKratos_SubZapSkip
// It sits outside zapkratos_test.go so an incorrect skip yields a different module
// subZapWrapped 是 TestZapKratos_SubZapSkip 用的包装函数
// 它不放在 zapkratos_test.go 里，跳帧数写错时得到的模块名就会不同
func subZapWrapped(zapKratos *ZapKratos) *zaplog.Zap {
	return zapKratos.SubZapSkip(1)
}
