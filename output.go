package log

import "io"

type logOutput struct {
	cur  io.Writer
	next io.Writer
}

func (o *logOutput) Write(p []byte) (n int, err error) {
	o.cur.Write(p)
	return o.next.Write(p)
}
