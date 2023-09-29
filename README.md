# simple-log  [![GoDoc](https://pkg.go.dev/badge/github.com/chuqingq/simple-log)](https://pkg.go.dev/github.com/chuqingq/simple-log@v0.1.0) 

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
