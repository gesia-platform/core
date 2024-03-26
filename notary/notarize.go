package notary

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/gesia-platform/core/types"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
)

func notarize(ctx *context.Context, prefix [1]byte, appID *big.Int) error {
	fmt.Printf("notarize prefix: %d, appId: %d\n", prefix, appID.Int64())

	config := ctx.Config()
	hostClient := ctx.ChainTree().Host.Client()

	var message [32]byte

	switch prefix {
	case types.NetworkAccessPermissionPrefix:
		message = types.GetNetwrokAccessPermissionMessage(*appID)
	}

	notaryPublic, err := store.NewNotaryPublicStore(
		common.HexToAddress(config.ChainTree.Host.NotaryPublicAddress),
		hostClient,
	)
	if err != nil {
		return err
	}

	skbz, err := hex.DecodeString(config.ChainTree.BLSSecretKey)
	if err != nil {
		return err
	}

	sk, err := blst.SecretKeyFromBytes(skbz)
	if err != nil {
		return err
	}

	signature := sk.Sign(message[:])

	fmt.Printf("bls siganture: %d\n", signature.Marshal())

	chainId := new(big.Int)
	chainId.SetUint64(config.ChainTree.Host.ChainID)

	txOpts, err := genTransactOpts(hostClient, chainId, config.ChainTree.PrivateKey)
	if err != nil {
		return err
	}

	tx, err := notaryPublic.Notarize(
		txOpts,
		prefix,
		appID,
		signature.Marshal(),
	)
	if err != nil {
		return err
	}

	fmt.Printf("notarized success tx: %d\n", tx)

	return nil
}
