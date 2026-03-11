package zapkratos

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yylego/zaplog"
)

// TestNewLogImp tests creating LogImp adapter and using it with log.Helper
// Verifies LogImp implements Kratos Logger and handles key-value data
//
// TestNewLogImp 测试创建 LogImp 适配器并与日志助手一起使用
// 验证 LogImp 实现 Kratos Logger 接口并处理键值对
func TestNewLogImp(t *testing.T) {
	logImp := NewLogImp(zaplog.LOG, "TEST-NEW-LOG-CAPTION")

	slog := log.NewHelper(logImp)

	slog.Infow("k", "v", "k1", "v2")
	slog.Infow("k", "v", "v2")
}
