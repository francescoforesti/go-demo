package logging

import (
	"github.com/francescoforesti/go-demo/goka/utils"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	logStdout = logrus.New()
	levels    = map[string]logrus.Level{
		"DEBUG": logrus.DebugLevel,
		"INFO":  logrus.InfoLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
	}
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
	logStdout.Formatter = customTextFormatter
	logStdout.Level = getLogLevel()
	Info(utils.LOG_MSG_LOGGERS_INITIALIZED)
}

func getLogLevel() logrus.Level {
	var logLevel, available = os.LookupEnv(utils.LOG_LEVEL_ENV_VAR)
	if !available {
		logLevel = utils.LOG_LEVEL
	}
	return levels[logLevel]
}
