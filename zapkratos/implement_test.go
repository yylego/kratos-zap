package zapkratos

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// TestNewSlogLogger checks the slog entries reach Zap with the caption, the pairs and the code position, and ERROR comes with a stack trace
// TestNewSlogLogger 检查 slog 日志带着 caption、键值对和调用位置写进 Zap，ERROR 日志附带调用栈
func TestNewSlogLogger(t *testing.T) {
	core, logs := observer.New(zapcore.DebugLevel)
	slogLog := NewSlogLogger(zap.New(core), "TEST-CAPTION")

	slogLog.Info("abc", "k", "v")
	slogLog.Error("xyz")

	entries := logs.TakeAll()
	require.Len(t, entries, 2)

	require.Equal(t, zapcore.InfoLevel, entries[0].Level)
	require.Equal(t, "abc", entries[0].Message)
	require.Equal(t, map[string]any{"caption": "TEST-CAPTION", "k": "v"}, entries[0].ContextMap())
	require.Equal(t, "github.com/yylego/kratos-zap/zapkratos.TestNewSlogLogger", entries[0].Caller.Function)
	require.Empty(t, entries[0].Stack)

	require.Equal(t, zapcore.ErrorLevel, entries[1].Level)
	require.Equal(t, "xyz", entries[1].Message)
	require.NotEmpty(t, entries[1].Stack)
}
