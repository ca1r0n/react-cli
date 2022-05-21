package config

import (
	"os"
	"sync"
)

var conf Config
var once sync.Once

func LoadEnv() *Config {
	once.Do(func() {
		conf.REACT_CLI = os.Getenv("REACT_CLI")
	})

	return &conf
}
