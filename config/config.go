package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Port    uint   `toml:"port"`
	Chain   string `toml:"chain"`
	ChainID uint64 `toml:"chain_id"`

	SignerAddress    string `toml:"signer_address"`
	SignerPrivateKey string `toml:"signer_private_key"`

	ExternalURL string `toml:"external_url"`

	RPCURL    string `toml:"rpc_url"`
	ProxyPath string `toml:"proxy_path"`

	NotaryChainRPCURL   string `toml:"notary_chain_rpc_url"`
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

}
