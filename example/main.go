package main

import (
	"errors"
	"time"

	"github.com/seerx/logo"
	"github.com/seerx/logo/log"
)

func main() {
	logger := log.GetDefaultLogger()
	logger.SetLevel(logo.LevelDebug)
	// logger.SetFormatter(logo.JSONFormatter)
	logger.SetLogFileLine(true)
	log.SetLogErrorCallStacks(true)
	// logger.SetColorLog(false)
	log.WithData(123).Debug("111")
	// runtime.Compiler
	log.Infof("os.Getegid()")
	log.WithData("aaa").WithError(errors.New("122222")).Warn("111")
	logger.WithError(errors.New("122222")).Error("000")

	time.Sleep(time.Second)
}
