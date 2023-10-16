package log

// Logger logger interface
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

// NullLogger An empty Logger type, print nothing.
type NullLogger struct {
}

func (nl NullLogger) Debugf(format string, args ...interface{}) {
}

func (nl NullLogger) Infof(format string, args ...interface{}) {
}

func (nl NullLogger) Warnf(format string, args ...interface{}) {
}

func (nl NullLogger) Errorf(format string, args ...interface{}) {
}
func (nl NullLogger) Fatalf(format string, args ...interface{}) {
}

func (nl NullLogger) Panicf(format string, args ...interface{}) {
}
