package notary

import (
	"encoding/hex"
	"fmt"
	"math/big"

	basectx "context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
)

func (notary *Notary) responseNetworkAccessPermission(ctx *context.Context, appID *big.Int, aggreatedSiganture []byte) error {
	config := ctx.Config()
	rootClient := ctx.ChainTree().Root.Client()

	chainId := new(big.Int)
	chainId.SetUint64(config.ChainTree.Root.ChainID)

	hostChainId := new(big.Int)
	hostChainId.SetUint64(config.ChainTree.Host.ChainID)

	appPermission, err := store.NewAppPermissionStore(
		common.HexToAddress(ctx.Config().ChainTree.Root.AppPermissionAddress),
		rootClient,
	)
	if err != nil {
		panic(err)
	}

	// 1. PoA Genesis Node는 appID에 맵핑된 IOA를 생성하고 키 체인에 저장한다
	pk, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	ioaAddress := crypto.PubkeyToAddress(pk.PublicKey)

	if err := ctx.Keychain().Set(basectx.Background(), appID.String(), hex.EncodeToString(crypto.FromECDSA(pk)), 0).Err(); err != nil {
		fmt.Println("failed to set notary account to redis")
		return err
	}
	fmt.Printf("IOA for appID: %s created. address: %s\n", appID, ioaAddress.Hex())

	/*// 2. IOA에 선택된 Node가 기본 이더를 미리 송금해놓는다. 가스비 대납
	hostClient := ctx.ChainTree().Host.Client()
	sendETHTxOpt, err := genTransactOpts(hostClient, hostChainId, config.ChainTree.PrivateKey)
	if err != nil {
		// 실패 시 삭제
		ctx.Keychain().Del(basectx.Background(), appID.String())
		return err
	}

	// in wei (1 eth)
	value := big.NewInt(1000000000000000000)
	myPk, err := crypto.HexToECDSA(config.ChainTree.PrivateKey)
	if err != nil {
		ctx.Keychain().Del(basectx.Background(), appID.String())
		return err
	}

	tx, err := types.SignTx(
		types.NewTransaction(sendETHTxOpt.Nonce.Uint64(), ioaAddress, value, 0, sendETHTxOpt.GasPrice, nil),
		types.NewEIP155Signer(hostChainId),
		myPk,
	)
	if err != nil {
		// 실패 시 삭제
		fmt.Println("failed to send default coin to tx signing")
		ctx.Keychain().Del(basectx.Background(), appID.String())
		return err
	}

	if err := hostClient.SendTransaction(basectx.Background(), tx); err != nil {
		ctx.Keychain().Del(basectx.Background(), appID.String())
		fmt.Println("failed to send default coin to ioa")
		return err
	} else {
		fmt.Printf("successfully send default coin to ioa tx: %s\n", tx.Hash())
	}*/

	// 3. app의 IOA 접근 권한을 허용한다.
	txOpts, err := genTransactOpts(rootClient, chainId, config.ChainTree.PrivateKey)
	if err != nil {
		return err
	}

	tx, err := appPermission.ResponseNetworkAccess(txOpts, appID, notary.networkAccountAddress, aggreatedSiganture, true)
	if err != nil {
		// 실패 시 삭제
		ctx.Keychain().Del(basectx.Background(), appID.String())
		return err
	}

	fmt.Printf("successfully response network access permission tx: %s\n", tx.Hash())

	return nil
}
