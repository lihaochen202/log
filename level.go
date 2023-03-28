package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// By exporting package-level constants, the log levels
// within the zap package are made accessible in the form
// of log.InfoLevel, for ease of access to these values.
//
// 通过包级别常量导出 zap 包内部的日志级别，方便以 log.InfoLevel
// 的形式访问这些值。
const (
	DebugLevel  = zap.DebugLevel
	InfoLevel   = zap.InfoLevel
	WarnLevel   = zap.WarnLevel
	ErrorLevel  = zap.ErrorLevel
	DPanicLevel = zap.DPanicLevel
	PanicLevel  = zap.PanicLevel
	FatalLevel  = zap.FatalLevel
)

// When the parsing of the log level in string
// format fails, the default log level is used.
//
// 解析字符串格式的日志级别失败时使用的默认日志级别
const defaultLevel = InfoLevel

// Parsing the log level in string format
// into the corresponding Level type.
//
// 将字符串格式的日志级别解析为对应的 Level 类型
func normalizeLevel(text string) zapcore.Level {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(text)); err != nil {
		level = defaultLevel
	}
	return level
}

// Return a function used to construct a logger level
// filtering function for zapcore. The filtering rule
// specifies that logs at or higher than the specified
// level will be recorded.
//
// 返回一个用于构建 zapcore 的日志级别过滤函数。
// 过滤规则为：高于指定级别（包含此级别）的日志将被记录。
func levelFunc(level zapcore.Level) func(level zapcore.Level) bool {
	return func(_level zapcore.Level) bool {
		return _level >= level
	}
}

// Return a function used to construct a logger level
// filtering function for zapcore. The filtering rule
// specifies that logs within the specified level range
// (left-closed and right-open) will be recorded.
//
// 返回一个用于构建 zapcore 的日志级别过滤函数。
// 过滤规则为：位于指定级别区间（左闭右开原则）的日志将被记录。
func rankLevelFunc(level, rankLevel zapcore.Level) func(level zapcore.Level) bool {
	return func(_level zapcore.Level) bool {
		return _level >= level && _level < rankLevel
	}
}
