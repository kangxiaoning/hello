package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("")

// InitLogger initializes logger.
func InitLogger(debug bool) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	logfile := path.Join(dir, "sinkmetrics.log")
	logFile := &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    1,
		MaxBackups: 2,
	}

	format := logging.MustStringFormatter(`%{level:.8s} %{message}`)
	backend := logging.NewLogBackend(logFile, "", log.LstdFlags|log.Lshortfile)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	if debug == true {
		backendLeveled.SetLevel(logging.DEBUG, "")
	} else {
		backendLeveled.SetLevel(logging.INFO, "")
	}

	logging.SetBackend(backendLeveled)
}

func NewLogger() *logging.Logger {
	return l
}