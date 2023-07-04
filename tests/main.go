package main

import (
	"os"
	"time"

	log "github.com/chuqingq/simple-log"
	"github.com/sirupsen/logrus"
)

func clean() {
	os.Remove("test.log")
	os.Remove("/dev/shm/test.log")
}

func main() {
	defer clean()

	logger := log.New("test.log")
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
