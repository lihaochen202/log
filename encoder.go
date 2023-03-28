package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Create a log encoder.
//
// 创建日志编码器。
func newEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(newEncoderConfig())
}

// Create a log ecoder configuation.
//
// 创建日志编码器配置。
func newEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeTime = timeEncoder
	config.EncodeDuration = durationEncoder
	return config
}

// 日志时间编码器（以更加友好的方式显示时间信息）。
func timeEncoder(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
	pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 日志时间差值编码器（以 ms 为单位显示时间差值）。
func durationEncoder(d time.Duration, pae zapcore.PrimitiveArrayEncoder) {
	pae.AppendFloat64(float64(d) / 1e6)
}
