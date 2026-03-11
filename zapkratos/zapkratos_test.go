package zapkratos

import (
	"testing"

	"github.com/yylego/zaplog"
)

// TestNewZapKratos tests creating ZapKratos and getting sub logging
// Verifies basic initialization and that sub logging can write log messages
//
// TestNewZapKratos 测试创建 ZapKratos 实例和获取子日志器
// 验证基本初始化工作正常且子日志器可以写入日志
func TestNewZapKratos(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	subLog := zapKratos.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}

// TestOptions_WithModuleKeyName tests custom module key name configuration
// Verifies WithModuleKeyName sets custom field name and sub logging works as expected
//
// TestOptions_WithModuleKeyName 测试自定义模块键名配置
// 验证 WithModuleKeyName 设置自定义字段名且子日志器工作正常
func TestOptions_WithModuleKeyName(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions().WithModuleKeyName("module"))

	subLog := zapKratos.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}
