package main

import (
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logger.log",
		MaxSize:    1,
		MaxBackups: 2,
	})
	log.SetLevel(log.DebugLevel)
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(Formatter)
	// default formatter is log.TextFormatter
	//log.SetFormatter(&log.TextFormatter{})
}

func main() {
	for {
		log.Debug("debug message")
		log.Info("info message")
		log.Warn("warning message")
		log.Error("debug message")
	}
}


// output
// 从时间及文件名上看已经切割了日志
// ➜  hello git:(master) ✗ ls -ltrh|grep log
// -rw-r--r--  1 dbo dbo 1.0M 3月  12 16:04 logger-2018-03-12T08-04-57.991.log
// -rw-r--r--  1 dbo dbo 1.0M 3月  12 16:04 logger-2018-03-12T08-04-58.099.log
// -rw-r--r--  1 dbo dbo 377K 3月  12 16:04 logger.log
// ➜  hello git:(master) ✗
// ➜  hello git:(master) ✗
// ➜  hello git:(master) ✗ ls -ltrh|grep log
// -rw-r--r--  1 dbo dbo 1.0M 3月  12 16:05 logger-2018-03-12T08-05-04.089.log
// -rw-r--r--  1 dbo dbo 1.0M 3月  12 16:05 logger-2018-03-12T08-05-04.164.log
// -rw-r--r--  1 dbo dbo 1.0M 3月  12 16:05 logger.log
// ➜  hello git:(master) ✗
// ➜  hello git:(master) ✗ tail -10 logger.log
// time="12-03-2018 16:05:14" level=error msg="debug message"
// time="12-03-2018 16:05:14" level=debug msg="debug message"
// time="12-03-2018 16:05:14" level=info msg="info message"
// time="12-03-2018 16:05:14" level=warning msg="warning message"
// time="12-03-2018 16:05:14" level=error msg="debug message"
// time="12-03-2018 16:05:14" level=debug msg="debug message"
// time="12-03-2018 16:05:14" level=info msg="info message"
// time="12-03-2018 16:05:14" level=warning msg="warning message"
// time="12-03-2018 16:05:14" level=error msg="debug message"
// time="12-03-2018 16:05:14" level=debug msg="debug message"
// ➜  hello git:(master) ✗
