package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Log struct {
	Console ConsoleLogger
	File    FileLogger
}

type ConsoleLogger struct {
	Enabled bool
	Json    bool
	Level   string
}

type FileLogger struct {
	Enabled  bool
	Json     bool
	Level    string
	Location string
}

func NewLoggerConfig(v *viper.Viper) *Log {
	var config Log

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("%s \n", err.Error())
		panic("Exiting server...")
	}

	if err := v.UnmarshalKey("logger", &config); err != nil {
		fmt.Printf("%s \n", err.Error())
		panic("Exiting server...")
	}

	return &config
}
