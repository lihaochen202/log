package log

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(config *Config) error {
	l, err := New(config)
	if err != nil {
		return err
	}
	std.z = l.z
	return nil
}

func New(config *Config) (*logger, error) {
	mergeDefaultConfig(config)

	cores := []zapcore.Core{}

	if !config.DisableConsole {
		level := normalizeLevel(config.ConsoleLevel)
		rankLevel := normalizeLevel(config.RankLevel)

		if config.DisableRank {
			cores = append(cores, newCore(newStdoutWriter(), level))
		} else {
			cores = append(
				cores,
				newRankCore(newStdoutWriter(), level, rankLevel),
				newCore(newStderrWriter(), rankLevel),
			)
		}
	}

	if !config.DisableOutput {
		level := normalizeLevel(config.OutputLevel)
		rankLevel := normalizeLevel(config.RankLevel)

		writer, err := newWriter(config.OutputFile, config)
		if err != nil {
			return nil, err
		}

		if config.DisableRank {
			cores = append(cores, newCore(writer, level))
		} else {
			cores = append(cores, newRankCore(writer, level, rankLevel))

			writer, err := newWriter(config.RankFile, config)
			if err != nil {
				return nil, err
			}
			cores = append(cores, newCore(writer, rankLevel))
		}
	}

	if len(cores) == 0 {
		return nil, errors.New("log core is empty")
	}

	opts := []zap.Option{}
	if !config.DisableCaller {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	if !config.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(normalizeLevel(config.StacktraceLevel)))
	}

	return &logger{z: zap.New(zapcore.NewTee(cores...), opts...)}, nil
}

var (
	Debug   = std.Debug
	Info    = std.Info
	Warn    = std.Warn
	Error   = std.Error
	DPanic  = std.DPanic
	Panic   = std.Panic
	Fatal   = std.Fatal
	Trace   = std.Trace
	Context = std.Context
	Clone   = std.Clone
	Sync    = std.Sync
)

var std = &logger{}

type logger struct {
	z      *zap.Logger
	traces []TraceFunc
}

func (l *logger) Debug(msg string, fields ...Field) {
	if l.z != nil {
		l.z.Debug(msg, fields...)
	}
}

func (l *logger) Info(msg string, fields ...Field) {
	if l.z != nil {
		l.z.Info(msg, fields...)
	}
}

func (l *logger) Warn(msg string, fields ...Field) {
	if l.z != nil {
		l.z.Warn(msg, fields...)
	}
}

func (l *logger) Error(msg string, fields ...Field) {
	if l.z != nil {
		l.z.Error(msg, fields...)
	}
}

func (l *logger) DPanic(msg string, fields ...Field) {
	if l.z != nil {
		l.z.DPanic(msg, fields...)
	}
}

func (l *logger) Panic(msg string, fields ...Field) {
	if l.z != nil {
		l.z.Panic(msg, fields...)
	}
}

func (l *logger) Fatal(msg string, fields ...Field) {
	if l.z != nil {
		l.z.Fatal(msg, fields...)
	}
}

func (l *logger) Trace(fns ...TraceFunc) *logger {
	l.traces = append(l.traces, fns...)
	return l
}

func (l *logger) Context(ctx context.Context, fields ...Field) *logger {
	if l.z != nil {
		lc := l.Clone(true)
		for _, trace := range lc.traces {
			lc.z = lc.z.With(trace(ctx)...)
		}
		lc.clearTraces()
		return lc
	}
	return l
}

func (l *logger) Clone(withTrace ...bool) *logger {
	_withTrace := false
	if len(withTrace) > 0 {
		_withTrace = withTrace[0]
	}

	lc := *l
	if !_withTrace {
		l.clearTraces()
	}
	return &lc
}

func (l *logger) clearTraces() {
	l.traces = nil
}

func (l *logger) Sync() error {
	if l.z != nil {
		return l.z.Sync()
	}
	return nil
}
