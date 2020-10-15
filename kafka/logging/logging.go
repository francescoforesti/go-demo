package logging

import (
	"github.com/francescoforesti/go-demo/goka/utils"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	errTextLogger error

	levels = map[string]logrus.Level{
		"DEBUG": logrus.DebugLevel,
		"INFO":  logrus.InfoLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
	}
	logStdout = logrus.New()
)

func Debug(message string) {
	logStdout.Debug(message)
}

func Info(message string) {
	logStdout.Info(message)
}

func Warn(message string) {
	logStdout.Warn(message)
}

func Error(message string) {
	logStdout.Error(message)
}

func InitializeLoggers() {
	customTextFormatter := new(logrus.TextFormatter)
	customTextFormatter.FullTimestamp = true
	logStdout.Out = os.Stdout
	logStdout.Formatter = customTextFormatter //&logrus.JSONFormatter{}
	logStdout.Level = getLogLevel()
	Info(utils.LOG_MSG_LOGGERS_INITIALIZED)
}

func getLogLevel() logrus.Level {
	var logLevel = os.Getenv(utils.LOG_LEVEL_ENV_VAR)
	if len(logLevel) == 0 {
		logLevel = utils.LOG_LEVEL
	}
	return levels[logLevel]
}
