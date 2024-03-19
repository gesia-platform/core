package config

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Port string `toml:"port"`
	URL  string `toml:"url"`

	Address string `toml:"address"`

	MasterRPCURL                  string `toml:"master-rpc-url"`
	MasterWSURL                   string `toml:"master-ws-url"`
	MasterNotaryPublicAddress     string `toml:"master-notary-public-address"`
	MasterNotaryPublicChainPrefix string `toml:"master-notary-public-chain-prefix"`

	RPCURL       string `toml:"rpc-url"`
	RPCProxyPath string `toml:"rpc-proxy-path"`
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

	if !strings.HasPrefix(config.RPCProxyPath, "/") {
		panic("ethereum-rpc-proxy-path must be prefixed with '/'.")
	}
}
