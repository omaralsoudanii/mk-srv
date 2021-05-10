package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Log *Log
}

func LoadConfig() (conf *Config) {
	fmt.Println("Loading configuration started...")
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	logger := NewLoggerConfig(v)

	cf := &Config{
		Log: logger,
	}

	return cf
}
