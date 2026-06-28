package zapkratos

import (
	"testing"

	"github.com/yylego/zaplog"
)

// TestZapKratos_GetLogger tests getting a *slog.Logger from ZapKratos
// Verifies the logger writes structured data without panic
//
// TestZapKratos_GetLogger 测试从 ZapKratos 实例获取 *slog.Logger
// 验证返回的 logger 能正常写结构化数据而不 panic
func TestZapKratos_GetLogger(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	applog := zapKratos.GetLogger("TEST-GET-LOGGER")

	applog.Info("test message", "k", "v", "k1", "v2")
}

// TestZapKratos_NewLogger tests creating a *slog.Logger from ZapKratos
// Verifies the logger built with the given caption works
//
// TestZapKratos_NewLogger 测试从 ZapKratos 实例创建 *slog.Logger
// 验证使用消息说明构造的 logger 工作正常
func TestZapKratos_NewLogger(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	applog := zapKratos.NewLogger("TEST-NEW-LOGGER")

	applog.Info("test message", "stage", "demo")
}

// TestZapKratos_GetSlogLogger tests the explicit-named *slog.Logger accessor
// Verifies the logger built with the given caption works
//
// TestZapKratos_GetSlogLogger 测试显式命名的 *slog.Logger 访问方法
// 验证使用消息说明构造的 logger 工作正常
func TestZapKratos_GetSlogLogger(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	applog := zapKratos.GetSlogLogger("TEST-GET-SLOG-LOGGER")

	applog.Info("test message", "stage", "demo")
}
