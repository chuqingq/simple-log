package log

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// myFormatter 自定义日志格式
type myFormatter struct {
}

// Format 格式化日志
func (f *myFormatter) Format(e *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %5.5s [%s:%v %s] %s\n",
		e.Time.Local().Format("01/02 15:04:05.000000"),
		e.Level.String(),
		splitAndGetLast(e.Caller.File, "/"),
		e.Caller.Line, splitAndGetLast(e.Caller.Function, "."),
		e.Message)), nil
}

// splitAndGetLast 分割字符串并返回最后一个元素
func splitAndGetLast(str string, sep string) string {
	slice := strings.Split(str, sep)
	return slice[len(slice)-1]
}
