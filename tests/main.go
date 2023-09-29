package main

import (
	"context"
	"os"
	"os/exec"
	"time"

	log "github.com/chuqingq/simple-log"
	"github.com/sirupsen/logrus"
)

func clean() {
	os.Remove("test.log")
	exec.Command("sh", "-c", "rm /dev/shm/test*.log*").Run()
}

func main() {
	defer clean()

	options := &log.Options{
		EnableMemory: true,
	}
	logger := log.New("test", options)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	go func() {
		time.Sleep(time.Second * 5)
		logger.SetLevel(logrus.ErrorLevel)
		time.Sleep(time.Second * 10)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			logger.Debugf("this is debug log")
			logger.Infof("this is info log")
			logger.Warnf("this is warn log")
			logger.Errorf("this is error log")
		}
	}
}
