package log

import "github.com/imdario/mergo"

// Config is a set of optional configuration options for
// creating a log instance, and the fields contained within
// are all optional. There is no need to worry about missing
// values for any of the fields, as the log package will
// provide default values for any unspecified fields.
// The optional values for log levels are:
// "debug", "info", "warn", "error", "dpanic", "panic", "fatal".
//
// Config 是创建 log 实例的配置选项，其中包含的字段都是可选的。无须
// 担心某些字段值缺失而导致包功能错误的情况，包内部的默认配置会出手。
// 日志级别可选值为：
// "debug", "info", "warn", "error", "dpanic", "panic", "fatal"。
type Config struct {
	// Disable logging to Stdout and Stderr,
	// the default value is false.
	//
	// 禁止 Stdout、Stderr 输出日志，默认值为 false。
	DisableConsole bool
	// The log level that is allowed to be logged to Stdout and Stderr,
	// the default value is "debug".
	//
	// 允许 Stdout、Stderr 输出的日志级别，默认值为 "debug"。
	ConsoleLevel string

	// Disable writing logs to files,
	// the default value is false.
	//
	// 禁止日志写入文件，默认值为 false。
	DisableOutput bool
	// The log level that allows for writing to a file,
	// the default value is "info".
	//
	// 允许写入文件的日志级别，默认值为 "info"。
	OutputLevel string
	// Specifies the name of the file to which the logs will be written,
	// the default value is "log/access.log".
	//
	// 指定日志写入文件的名称，默认值为 "log/access.log"。
	OutputFile string

	// Disable logging categorization,
	// the default value is false.
	//
	// 禁用日志归类，默认值为 false。
	DisableRank bool
	// Specify the logging categorization level,
	// the default value is "error".
	//
	// 日志归类级别，默认值为 "error"。
	RankLevel string
	// Specify the name of the file where categorized logs should be written,
	// the default value is "log/error.log".
	//
	// 指定归类日志写入文件的名称，默认值为 "log/error.log"。
	RankFile string

	// Disable automatic archiving of log files，
	// the default value is false.
	//
	// 是否禁用日志文件自动归档，默认值为 false。
	DisableOutputRotate bool
	// The size in megabytes at which log files should be automatically archived,
	// the default value is 5.
	//
	// 日志文件自动归档大小（MB），默认值为 5。
	OutputRotateMaxSize int
	// Retention period of archived log files (in days),
	// the default value is 7.
	//
	// 已归档日志文件保存时间（Day），默认值为 7。
	OutputRotateMaxAge int
	// Maximum number of archived log files,
	// default is unlimited.
	//
	// 已归档日志文件最大数量，默认无限制。
	OutputRotateMaxBackups int
	// Enable compression of archived log files,
	// the default value is false.
	//
	// 是否开启归档日志文件压缩，默认值为 false。
	OutputRotateCompress bool

	// Disable logging of caller position in the log file,
	// the default value is false.
	//
	// 是否禁止在日志中记录调用者位置，默认值为 false。
	DisableCaller bool
	// Disable logging of stack trace in the log file,
	// the default value is false.
	//
	// 是否禁止在日志中记录堆栈信息，默认值为 false。
	DisableStacktrace bool
	// The log level that allows for stack trace,
	// the default value is "panic".
	//
	// 允许记录堆栈信息的日志级别，默认值为 "panic"。
	StacktraceLevel string
}

// The default configuation for creating log instance.
//
// 创建 log 实例的默认配置。
var defaultConfig = &Config{
	ConsoleLevel: "debug",

	OutputLevel: "info",
	OutputFile:  "log/access.log",

	RankLevel: "error",
	RankFile:  "log/error.log",

	OutputRotateMaxSize: 5,
	OutputRotateMaxAge:  7,

	StacktraceLevel: "panic",
}

func mergeDefaultConfig(config *Config) {
	mergo.Merge(config, defaultConfig)
}
