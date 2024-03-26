package types

import "math/big"

var (
	NetworkAccessPermissionPrefix = [1]byte{0x01}
	networkAccessPermissionKey    = []byte("network_access_permission")
)

func GetNetwrokAccessPermissionMessage(appID big.Int) [32]byte {
	var bz [32]byte

	ki := 0
	for _, b := range networkAccessPermissionKey {
		bz[ki] = b
		ki++
	}

	for i, b := range appID.Bytes() {
		bz[ki+i] = b
	}

	return bz
}
