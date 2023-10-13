package log

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func clean() {
	os.Remove("test.log")
	os.Remove("/dev/shm/test.log")
	os.RemoveAll("./logs")
}

func TestLogLevel(t *testing.T) {
	defer clean()

	logger := New("test.log")
	logger.Formatter = &logrus.JSONFormatter{}
	go func() {
		time.Sleep(time.Second * 5)
		logger.SetLevel(logrus.ErrorLevel)
	}()
	for {
		logger.Debugf("this is debug log")
		logger.Infof("this is info log")
		logger.Warnf("this is warn log")
		logger.Errorf("this is error log")

		time.Sleep(time.Second)
	}
}

func TestLogFileDir(t *testing.T) {
	const filename = "./logs/test.log"
	defer clean()

	logger := New(filename)
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Errorf("this is error log")
	logger.Writer().Close()

	// 验证有文件
	f, err := os.Open(filename)
	assert.Nil(t, err)
	f.Close()
}

func TestDirAndMemory(t *testing.T) {
	const filename = "./logs/test.log"
	defer clean()

	options := &Options{
		EnableMemory: true,
	}

	logger := New(filename, options)
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Errorf("this is error log")
	logger.Writer().Close()

	// 验证有文件
	f, err := os.Open("/dev/shm/test.log")
	assert.Nil(t, err)
	f.Close()
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestAppendOutput(t *testing.T) {
	defer clean()
	logger := New("test.log")

	// 打开文件，如果不存在则创建
	filename := filepath.Join(getCurrentPath(), "test.log")
	log.Printf("filename: %v", filename)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	assert.Equal(t, nil, err)
	defer file.Close()
	defer os.Remove(filename)

	go func() {
		time.Sleep(time.Second * 5)
		AppendOutput(logger, file)
	}()
	for i := 0; i < 10; i++ {
		logger.Debugf("this is debug log")
		logger.Infof("this is info log")
		logger.Warnf("this is warn log")
		logger.Errorf("this is error log")

		time.Sleep(time.Second)
	}
}
