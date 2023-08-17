package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	log "github.com/chuqingq/simple-log"
)

var file = flag.String("file", "", "log file name")
var size = flag.Int("size", 20, "log file max size in MB")
var backups = flag.Int("backups", 2, "log max backups")
var disableStderr = flag.Bool("disableStderr", false, "disable stderr output")
var EnableMemory = flag.Bool("enableMemory", false, "enable keep log file in memory (/dev/shm/)")

func main() {
	flag.Parse()
	if *file == "" {
		flag.PrintDefaults()
		return
	}

	options := &log.Options{
		MaxSizeInMB:   *size,
		MaxBackups:    *backups,
		DisableStderr: *disableStderr,
		EnableMemory:  *EnableMemory,
	}
	fmt.Printf("options: %#v\n", options)

	logger := log.New(*file, options)

	io.Copy(logger.Out, os.Stdin)
}
