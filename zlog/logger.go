package zlog

import (
	"github.com/omaralsoudanii/http-srv/config"
)

func NewLoggerInstance(lc *config.Log) *Logger {
	logOpts := &LoggerConfiguration{
		EnableConsole: lc.Console.Enabled,
		ConsoleJson:   lc.Console.Json,
		ConsoleLevel:  lc.Console.Level,
		EnableFile:    lc.File.Enabled,
		FileJson:      lc.File.Json,
		FileLevel:     lc.File.Level,
		FileLocation:  lc.File.Location,
	}
	log := NewLogger(logOpts)
	log.Info("Logger loaded")
	return log
}
