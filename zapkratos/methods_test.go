package zapkratos

import (
	"testing"

	"github.com/yylego/zaplog"
)

// TestZapKratos_GetHelper tests getting log helper from ZapKratos
// Verifies GetHelper creates working helper with both Info and Infow
//
// TestZapKratos_GetHelper 测试从 ZapKratos 实例获取日志助手
// 验证 GetHelper 创建的助手具有 Info 和 Infow 方法且工作正常
func TestZapKratos_GetHelper(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	slog := zapKratos.GetHelper("TEST-GET-HELPER")

	slog.Info("woca", "[a]", "[b]", "[c]")
	slog.Infow("k", "v", "k1", "v2")
}

// TestZapKratos_NewHelper tests creating new log helper from ZapKratos
// Verifies NewHelper creates working helper with given caption
//
// TestZapKratos_NewHelper 测试从 ZapKratos 实例创建新日志助手
// 验证 NewHelper 使用消息说明创建工作正常的助手
func TestZapKratos_NewHelper(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	slog := zapKratos.NewHelper("TEST-NEW-HELPER")

	slog.Info("test-message")
	slog.Infow("test", "message")
}
