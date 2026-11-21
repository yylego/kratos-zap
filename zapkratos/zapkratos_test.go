package zapkratos

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// TestNewOptions checks the default module field name is "module"
// TestNewOptions 检查默认的模块键名是 "module"
func TestNewOptions(t *testing.T) {
	require.Equal(t, &Options{ModuleKeyName: "module"}, NewOptions())
}

// TestOptions_WithModuleKeyName checks the sub Zap writes the file name as the custom field name
// TestOptions_WithModuleKeyName 检查子 Zap 用自定义的键名写出文件名
func TestOptions_WithModuleKeyName(t *testing.T) {
	core, logs := observer.New(zapcore.DebugLevel)
	zapKratos := NewZapKratos(zaplog.NewZap(zap.New(core)), NewOptions().WithModuleKeyName("source"))

	zapKratos.SubZap().LOG.Info("abc")

	entries := logs.TakeAll()
	require.Len(t, entries, 1)
	require.Equal(t, map[string]any{"source": "zapkratos_test.go"}, entries[0].ContextMap())
}

// TestNewZapKratos checks Zap gives back the Zap passed in
// TestNewZapKratos 检查 Zap 返回传入的 Zap
func TestNewZapKratos(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())
	require.Same(t, zaplog.LOGGER, zapKratos.Zap())
}

// TestZapKratos_SubZap checks the sub Zap adds the file name of the code that calls it, LOG and SUG alike
// TestZapKratos_SubZap 检查子 Zap 加上调用方的文件名，LOG 和 SUG 都一样
func TestZapKratos_SubZap(t *testing.T) {
	core, logs := observer.New(zapcore.DebugLevel)
	zapKratos := NewZapKratos(zaplog.NewZap(zap.New(core)), NewOptions())

	sub := zapKratos.SubZap()
	sub.LOG.Info("abc")
	sub.SUG.Info("xyz")

	entries := logs.TakeAll()
	require.Len(t, entries, 2)
	for _, item := range entries {
		require.Equal(t, map[string]any{"module": "zapkratos_test.go"}, item.ContextMap())
	}
}
