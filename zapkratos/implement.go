package zapkratos

import (
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yylego/erero"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogImp adapts zap.Logger to implement the Kratos log.Logger interface
// Bridges Uber Zap with Kratos using structured logging approach
//
// LogImp 适配 zap.Logger 以实现 Kratos log.Logger 接口
// 将 Uber Zap 结构化日志桥接到 Kratos 日志接口
type LogImp struct {
	zapLog     *zap.Logger // Underlying Zap logger instance // 底层 Zap 日志器实例
	msgCaption string      // Message caption added to logs // 每条日志条目的消息说明
}

// NewLogImp creates LogImp adapter instance from the given Zap
// Returns adapter that implements the Kratos log.Logger interface
//
// NewLogImp 创建包装给定 Zap 日志器的 LogImp 适配器实例
// 实现 Kratos log.Logger 接口并返回新适配器
func NewLogImp(zapLog *zap.Logger, msgCaption string) log.Logger {
	return &LogImp{
		zapLog:     zapLog,
		msgCaption: msgCaption,
	}
}

// Log implements the Kratos log.Logger interface
// Converts Kratos log level to Zap level and writes structured log with key-value data
// Generates default keys to handle odd key-value argument count
//
// Log 实现 Kratos log.Logger 接口
// 将 Kratos 日志级别转换为 Zap 级别，并使用键值对写入结构化日志
// 在需要时生成默认键来处理奇数个键值对
func (a *LogImp) Log(logLevel log.Level, keyvals ...interface{}) error {
	var zapLevel zapcore.Level
	switch logLevel {
	case log.LevelDebug:
		zapLevel = zap.DebugLevel
	case log.LevelInfo:
		zapLevel = zap.InfoLevel
	case log.LevelWarn:
		zapLevel = zap.WarnLevel
	case log.LevelError:
		zapLevel = zap.ErrorLevel
	case log.LevelFatal:
		zapLevel = zap.DPanicLevel
	default:
		zapLevel = zap.DebugLevel // Use debug when unknown // 未知时使用 debug
	}

	// Check if logging at this level is enabled
	// 检查此级别的日志是否已启用
	zapInstance := a.zapLog.Check(zapLevel, a.msgCaption)
	if zapInstance == nil {
		return erero.Errorf("WRONG-LOG-LEVEL-PARAM zap=%v arg=%v", zapLevel, logLevel)
	}

	// Pre-allocate fields slice, handle odd number of arguments
	// 预分配字段切片，处理奇数个参数
	var fields = make([]zap.Field, 0, (len(keyvals)+1)/2)
	for idx := 0; idx < len(keyvals); idx += 2 {
		if idx+1 < len(keyvals) {
			kes, ok := keyvals[idx].(string)
			if !ok {
				kes = newBK(idx) // Generate default key when key is not string // 当键不是字符串时生成默认键
			}
			fields = append(fields, zap.Any(kes, keyvals[idx+1]))
		}
	}
	if len(keyvals)%2 == 1 { // Handle trailing value with no key // 处理没有键的尾随值
		fields = append(fields, zap.Any(newBK(len(keyvals)-1), keyvals[(len(keyvals)-1)]))
	}

	zapInstance.Write(fields...)
	return nil
}

// newBK generates default key name using BAD_KEY prefix and position
// Used when key-value data is irregular or when keys are absent
//
// newBK 使用 BAD_KEY 前缀和索引生成默认键名
// 用于键值对不规则或键缺失时
func newBK(idx int) string {
	return "BAD_KEY_" + strconv.Itoa(idx)
}
