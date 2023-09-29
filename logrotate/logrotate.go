package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	log "github.com/chuqingq/simple-log"
)

var name = flag.String("Name", "", "log name, without '.log' suffix")
var size = flag.Int("MaxSizeInMB", 10, "log file max size in MB")
var backups = flag.Int("MaxBackups", 1, "log max backups")
var disableStderr = flag.Bool("DisableStderr", false, "disable stderr output")
var EnableMemory = flag.Bool("EnableMemory", false, "enable keep log file in memory (/dev/shm/)")

func main() {
	flag.Parse()
	if *name == "" {
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

	logger := log.New(*name, options)

	io.Copy(logger.Out, os.Stdin)
}
