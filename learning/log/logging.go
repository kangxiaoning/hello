package main

import (
	"log"
	"github.com/op/go-logging"
)

func main() {
	var logger = logging.MustGetLogger("root")
	backend := logging.NewLogBackend(&lumberjack.Logger{
		Filename:   "logging.log",
		MaxSize:    1,
		MaxBackups: 2,
	}, "", log.LstdFlags|log.Lshortfile)
	//logging.SetBackend(backend)
	backend1Leveled := logging.AddModuleLevel(backend)
	backend1Leveled.SetLevel(logging.DEBUG, "")
	logger.SetBackend(backend1Leveled)
	for {

		logger.Debug("debug")
		logger.Info("info")
		logger.Error("error")
		logger.Critical("critical")
	}
}