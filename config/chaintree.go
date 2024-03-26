package config

type ChainTree struct {
	Root ChainTreeRoot `toml:"root"`

	Host ChainTreeHost `toml:"host"`

	Address string `toml:"address"`

	PrivateKey string `toml:"private_key"`

	BLSSecretKey string `toml:"bls_secret_key"`
}

type ChainTreeRoot struct {
	ChainID uint64 `toml:"chain_id"`

	RPCURL string `toml:"rpc_url"`

	NetworkAccountAddress string `toml:"network_account_address"`

	AppStoreAddress string `toml:"app_store_address"`

	AppPermissionAddress string `toml:"app_permission_address"`
}

type ChainTreeHost struct {
	ChainID uint64 `toml:"chain_id"`

	NotaryPublicAddress string `toml:"notary_public_address"`

	RPCURL string `toml:"rpc_url"`

	ProxyPath string `toml:"proxy_path"`
}
