package configs

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Configuration struct {
	App AppConfig `toml:"app"`
}

type AppConfig struct {
	Host      string `toml:"host"`
	Port      int64  `toml:"port"`
	Env       string `toml:"env"`
	SecretKey string `toml:"secret_key"`
}

func InitConfig(configFile string, config *Configuration) error {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("configuration file %s not exists", configFile)
	}
	src, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("configuration file %s cannot be opened: %v", configFile, err)
	}
	defer src.Close()
	if _, err := toml.DecodeFile(configFile, config); err != nil {
		return fmt.Errorf("failed to parse the configuration file %s: %v", configFile, err)
	}
	return nil
}
