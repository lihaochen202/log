package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

func newStdoutWriter() io.Writer {
	return os.Stdout
}

func newStderrWriter() io.Writer {
	return os.Stderr
}

func newWriter(filename string, config *Config) (io.Writer, error) {
	if config.DisableOutputRotate {
		return newUnrotateWriter(filename)
	}
	return newRotateWriter(filename, config), nil
}

func newUnrotateWriter(filename string) (io.Writer, error) {
	return openFile(filename)
}

func openFile(filename string) (io.Writer, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return openNewFile(filename)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("can't open log file that already exists: %s", err)
	}
	return file, nil
}

func openNewFile(filename string) (io.Writer, error) {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("can't make directories for new log file: %s", err)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("can't create log file: %s", err)
	}
	return file, nil
}

func newRotateWriter(filename string, config *Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    config.OutputRotateMaxSize,
		MaxAge:     config.OutputRotateMaxAge,
		MaxBackups: config.OutputRotateMaxBackups,
		Compress:   config.OutputRotateCompress,
	}
}
