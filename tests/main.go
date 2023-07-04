package main

import (
	"os"
	"os/exec"
	"time"

	log "github.com/chuqingq/simple-log"
	"github.com/sirupsen/logrus"
)

func clean() {
	os.Remove("test.log")
	exec.Command("sh", "-c", "rm /dev/shm/test*.log").Run()
}

func main() {
	defer clean()

	logger := log.New("test.log")
	log.AppendOutput(logger, os.Stderr)

	go func() {
		time.Sleep(time.Second * 5)
		logger.SetLevel(logrus.ErrorLevel)
	}()

	for i := 0; i < 8; i++ {
		logger.Debugf("this is debug log")
		logger.Infof("this is info log")
		logger.Warnf("this is warn log")
		logger.Errorf("this is error log")

		time.Sleep(time.Second)
	}
}
