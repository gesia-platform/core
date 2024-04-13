package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Port uint `toml:"port"`

	ChainTree ChainTree `toml:"chain_tree"`

	Keychain Keychain `toml:"keychain"`
}

func NewConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := &Config{}

	if err := toml.Unmarshal(data, config); err != nil {
		panic(err)
	}

	return config
}
