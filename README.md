# simple-log

A simple log module in Go.

## Features

1. support log rotating by MaxSize, MaxBackups etc.(using lumberjack)
2. support dynamicly setting log level.
3. support appending multiple output writers.
4. support multiple print interfaces: Debugf, Infof, Warnf, Errorf, Fatalf, Panicf.(using logrus)
5. support multiple log format: json, text, console, json-pretty, logfmt, logstash.(using logrus)
6. support log file in memory, to reduce disk writing.(using /dev/shm)

## TODO

- [ ] use type alias to replace logrus.Loggerpackage log // import "github.com/chuqingq/simple-log"


## go doc

```
package log // import "github.com/chuqingq/simple-log"

func AppendOutput(logger *logrus.Logger, output io.Writer)
func New(filename string, option ...*Options) *logrus.Logger
type Options struct{ ... }
package log // import "github.com/chuqingq/simple-log"


FUNCTIONS

func AppendOutput(logger *logrus.Logger, output io.Writer)
    AppendOutput 添加日志输出

func New(filename string, option ...*Options) *logrus.Logger

TYPES

type Options struct {
	// FileName    string // 日志文件名，不包含路径
	MaxSizeInMB int // 日志文件大小，单位MB，>=1
	MaxBackups  int // 日志文件最大备份数，>=1
	Formatter   logrus.Formatter
}
```
