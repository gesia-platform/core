package config

type ChainTree struct {
	Root ChainTreeRoot `toml:"root"`

	Host ChainTreeHost `toml:"host"`
}

type ChainTreeRoot struct {
	ChainID uint64 `toml:"chain_id"`

	Address string `toml:"address"`

	PrivateKey string `toml:"private_key"`

	RPCURL string `toml:"rpc_url"`

	NetworkAccountAddress string `toml:"network_account_address"`

	AppStoreAddress string `toml:"app_store_address"`

	AppPermissionAddress string `toml:"app_permission_address"`
}

type ChainTreeHost struct {
	ChainID uint64 `toml:"chain_id"`

	Address string `toml:"address"`

	NotaryPublicAddress string `toml:"notary_public_address"`

	PrivateKey string `toml:"private_key"`

	RPCURL string `toml:"rpc_url"`

	ProxyPath string `toml:"proxy_path"`
}
