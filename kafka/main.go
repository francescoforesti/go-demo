package main

import (
	"github.com/francescoforesti/go-demo/kafka/services/kafka"
	"github.com/francescoforesti/go-demo/logging"
)

func main() {
	logging.InitializeLoggers()
	logging.Debug("this is debug")
	logging.Info("this is info")
	logging.Warn("this is warn")
	logging.Error("this is error")

	kafka.InitializeHandlers()
}
