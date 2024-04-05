package notary

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
)

func (notary *Notary) responseNetworkAccessPermission(ctx *context.Context, appID *big.Int, aggreatedSiganture []byte) error {
	config := ctx.Config()
	rootClient := ctx.ChainTree().Root.Client()

	chainId := new(big.Int)
	chainId.SetUint64(config.ChainTree.Root.ChainID)

	appPermission, err := store.NewAppPermissionStore(
		common.HexToAddress(ctx.Config().ChainTree.Root.AppPermissionAddress),
		rootClient,
	)
	if err != nil {
		panic(err)
	}

	txOpts, err := genTransactOpts(rootClient, chainId, config.ChainTree.PrivateKey)
	if err != nil {
		return err
	}

	tx, err := appPermission.ResponseNetworkAccess(txOpts, appID, notary.networkAccountAddress, aggreatedSiganture, true)
	if err != nil {
		return err
	}

	fmt.Printf("successfully response network access permission tx: %s\n", tx.Hash())

	return nil
}
