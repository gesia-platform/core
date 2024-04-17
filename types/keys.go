package types

import "math/big"

const (
	EthereumProxyPath = "/ethereum"
	EthereumTxPath    = "/txs"
	IOAPath           = "/ioas"
)

var (
	NetworkAccessPermissionPrefix = [1]byte{0x01}
)

func GetNetwrokAccessPermissionMessage(appID big.Int) [32]byte {
	var bz [32]byte

	for i, b := range appID.Bytes() {
		bz[i] = b
	}

	return bz
}
