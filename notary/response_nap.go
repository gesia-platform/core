package notary

import (
	"fmt"
	"math/big"

	basectx "context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

	// 1. PoA Genesis Node는 appID에 맵핑된 EOA를 생성한다.
	pk, err := crypto.GenerateKey()
	if err != nil {
		return err
	}

	address := crypto.PubkeyToAddress(pk.PublicKey).Hex()
	if err := ctx.Keychain().Set(basectx.Background(), appID.String(), hexutil.Encode(crypto.FromECDSA(pk)), 0).Err(); err != nil {
		fmt.Println("failed to set notary account to redis")
		return err
	}

	fmt.Printf("notary account created appID: %s, address: %s\n", appID, address)

	tx, err := appPermission.ResponseNetworkAccess(txOpts, appID, notary.networkAccountAddress, aggreatedSiganture, true)
	if err != nil {
		return err
	}

	fmt.Printf("successfully response network access permission tx: %s\n", tx.Hash())

	return nil
}
