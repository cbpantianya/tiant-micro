package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	App struct {
		Name    string `toml:"name"`
		Version string `toml:"version"`
	} `toml:"app"`

	Engine struct {
		Address string `toml:"address"`
		Port    int `toml:"port"`
	} `toml:"engine"`

	Database struct{
		Address string `toml:"address"`
		Port int `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		DBName string `toml:"dbname"`
	} `toml:"database"`

}

func NewConfig(path string) (cfg *Config, err error) {
	// Decode
	cfg = &Config{}
	_, err = toml.DecodeFile(path, cfg)
	return
}
