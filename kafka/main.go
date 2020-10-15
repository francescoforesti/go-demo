package main

import (
	"github.com/francescoforesti/go-demo/kafka/logging"
	"github.com/francescoforesti/go-demo/kafka/services/kafka"
)

func main() {
	logging.InitializeLoggers()
	logging.Debug("this is debug")
	logging.Info("this is info")
	logging.Warn("this is warn")
	logging.Error("this is error")

	kafka.InitializeHandlers()
}
