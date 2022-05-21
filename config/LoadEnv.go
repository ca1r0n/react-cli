package config

import "sync"

var conf Config
var once sync.Once

func LoadEnv() *Config {
	once.Do(func() {
		
	})

	return &conf
}
