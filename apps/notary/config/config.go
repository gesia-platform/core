package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Port uint `toml:"port"`

	SignerAddress string `toml:"signer_address"`

	ExternalURL string `toml:"external_url"`

	RPCURL    string `toml:"rpc_url"`
	ProxyPath string `toml:"proxy_path"`

	NotaryPublicAddress string `toml:"notary_public_address"`
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) Read(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(data, config); err != nil {
		panic(err)
	}
	fmt.Print(config)

}
