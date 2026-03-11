package zapkratos

import "github.com/go-kratos/kratos/v2/log"

// GetLogger creates Kratos log.Logger with given message caption
// Wraps underlying Zap and returns the created log instance
//
// GetLogger 使用给定的消息说明创建 Kratos log.Logger
// 包装底层 Zap 日志器并返回实例
func (A *ZapKratos) GetLogger(msgCaption string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgCaption)
}

// NewLogger creates Kratos log.Logger with given message caption
// Wraps underlying Zap and returns the created log instance
// Note: Same as GetLogger, provided to match naming patterns
//
// NewLogger 使用给定的消息说明创建 Kratos log.Logger
// 包装底层 Zap 日志器并返回实例
// 注意：与 GetLogger 相同，提供以匹配命名模式
func (A *ZapKratos) NewLogger(msgCaption string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgCaption)
}

// GetHelper creates Kratos log.Helper with given message caption
// Provides convenient methods and returns the created Helper instance
//
// GetHelper 使用给定的消息说明创建 Kratos log.Helper
// 提供便捷日志方法并返回实例
func (A *ZapKratos) GetHelper(msgCaption string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgCaption))
}

// NewHelper creates Kratos log.Helper with given message caption
// Provides convenient methods and returns the created Helper instance
// Note: Same as GetHelper, provided to match naming patterns
//
// NewHelper 使用给定的消息说明创建 Kratos log.Helper
// 提供便捷日志方法并返回实例
// 注意：与 GetHelper 相同，提供以匹配命名模式
func (A *ZapKratos) NewHelper(msgCaption string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgCaption))
}
