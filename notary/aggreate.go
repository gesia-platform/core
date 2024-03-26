package notary

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
)

func (notary *Notary) aggreate(ctx *context.Context, appID *big.Int, signatures [][]byte) error {
	rootClient := ctx.ChainTree().Root.Client()
	config := ctx.Config()

	chainId := new(big.Int)
	chainId.SetUint64(config.ChainTree.Root.ChainID)

	appPermission, err := store.NewAppPermissionStore(
		common.HexToAddress(ctx.Config().ChainTree.Root.AppPermissionAddress),
		rootClient,
	)
	if err != nil {
		panic(err)
	}

	txOpts, err := genTransactOpts(rootClient, chainId, config.ChainTree.Root.PrivateKey)
	if err != nil {
		return err
	}

	aggreatedSiganture, err := blst.AggregateCompressedSignatures(signatures)
	if err != nil {
		return err
	}

	tx, err := appPermission.ResponseNetworkAccess(txOpts, appID, notary.networkAccountAddress, aggreatedSiganture.Marshal(), true)
	if err != nil {
		return err
	}

	fmt.Printf("response network access success tx: %d\n", tx)

	return nil
}
