package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Port uint   `toml:"port"`
	URL  string `toml:"url"`

	Address string `toml:"address"`

	Chain struct {
		Local struct {
			RPCURL       string `toml:"rpc_url"`
			RPCProxyPath string `toml:"rpc_proxy_path"`
		} `toml:"local"`

		Master struct {
			RPCURL                  string `toml:"rpc_url"`
			NotaryPublicAddress     string `toml:"notary_public_address"`
			NotaryPublicChainPrefix string `toml:"notary_public_chain_prefix"`
		} `toml:"master"`
	} `toml:"chain"`
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
