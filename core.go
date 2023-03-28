package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Create a log core.
//
// 创建日志核心。
func newCore(w io.Writer, level zapcore.Level) zapcore.Core {
	return zapcore.NewCore(
		newEncoder(),
		zapcore.AddSync(w),
		zap.LevelEnablerFunc(levelFunc(level)),
	)
}

// Create a rank log core.
//
// 创建分类日志核心。
func newRankCore(w io.Writer, level, rankLevel zapcore.Level) zapcore.Core {
	return zapcore.NewCore(
		newEncoder(),
		zapcore.AddSync(w),
		zap.LevelEnablerFunc(rankLevelFunc(level, rankLevel)),
	)
}
