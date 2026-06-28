package zapkratos

import (
	"testing"

	"github.com/yylego/zaplog"
)

// TestNewSlogLogger tests building a *slog.Logger from the given Zap
// Verifies the bridged logger writes key-value data without panic
//
// TestNewSlogLogger 测试基于给定 Zap 构造 *slog.Logger
// 验证桥接出的 logger 能正常写键值对而不 panic
func TestNewSlogLogger(t *testing.T) {
	applog := NewSlogLogger(zaplog.LOG, "TEST-NEW-LOG-CAPTION")

	applog.Info("test message", "k", "v", "k1", "v2")
	applog.Debug("debug message", "stage", "demo")
}
