package config

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg  *TomlConfig
	once sync.Once
)

// GetConfig get config
func GetConfig() *TomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("./src/config/toml/local.toml")
		if err != nil {
			panic(err)
		}
		// fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	fmt.Printf("cfg = %v\n", cfg)
	return cfg
}
