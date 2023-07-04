package log

import (
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Options struct {
	// FileName    string // 日志文件名，不包含路径
	MaxSizeInMB   int // 日志文件大小，单位MB，>=1
	MaxBackups    int // 日志文件最大备份数，>=1
	Formatter     logrus.Formatter
	DisableStderr bool // 设置后不再打印到标准错误
}

func New(filename string, option ...*Options) *logrus.Logger {
	if filename == "" {
		filename = "test.log"
	}

	options := &Options{}
	if len(option) > 0 {
		options = option[0]
	}

	// options
	maxsize := 1
	if options.MaxSizeInMB > 1 {
		maxsize = options.MaxSizeInMB
	}

	maxbackups := 1
	if options.MaxBackups > 1 {
		maxbackups = options.MaxBackups
	}

	var formatter logrus.Formatter = &myFormatter{}
	if options.Formatter != nil {
		formatter = options.Formatter
	}

	var filenamepath string
	if runtime.GOOS == "linux" {
		filenamepath = filepath.Join("/dev/shm/", filename)
	} else {
		filenamepath = filepath.Join("./", filename)
	}

	// lumberjack logger作为logrus的输出
	output := &lumberjack.Logger{
		Filename:   filenamepath, // in memory
		MaxSize:    maxsize,      // megabytes
		MaxBackups: maxbackups,   // reserve 1 backup
		// MaxAge:     28, //days
	}

	logger := &logrus.Logger{
		Out: output,
		// Formatter: &logrus.TextFormatter{},
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	logger.SetReportCaller(true)

	// 设置后不再打印到标准错误
	if !options.DisableStderr {
		AppendOutput(logger, os.Stderr)
	}

	// 在当前目录创建链接
	if runtime.GOOS == "linux" {
		os.Symlink(filenamepath, filename)
	}

	return logger
}

// SetLogLevel 设置日志级别
// logger.SetLevel(logrus.DebugLevel)

// AppendOutput 添加日志输出
func AppendOutput(logger *logrus.Logger, output io.Writer) {
	logger.SetOutput(&logOutput{cur: logger.Out, next: output})
}
