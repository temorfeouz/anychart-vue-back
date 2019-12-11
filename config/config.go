package config

import (
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Service struct {
		Host string
		Post int
	}
}

var config *Config

func Get() *Config {
	return config
}

// Init init service config
func Init() {
	// read config file
	var tmpConfig Config
	if err := gcfg.ReadFileInto(&tmpConfig, "config.gcfg"); err != nil {
		panic(err)
	}
	config = &tmpConfig
}
