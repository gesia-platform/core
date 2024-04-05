package notary

import (
	"bytes"
	basectx "context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/gesia-platform/core/types"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	blscommon "github.com/prysmaticlabs/prysm/v5/crypto/bls/common"
)

var (
	notaryPublicABI abi.ABI
)

func (notary *Notary) SubscribeNotarizedWithCondition(ctx *context.Context) {
	config := ctx.Config()
	rootClient := ctx.ChainTree().Root.Client()

	notary.networkAccountAddress = common.HexToAddress(config.ChainTree.Root.NetworkAccountAddress)

	networkAccount, err := store.NewNetworkAccountStore(
		notary.networkAccountAddress,
		rootClient,
	)
	if err != nil {
		panic(err)
	}

	owner, err := networkAccount.Owner(&bind.CallOpts{Pending: true})
	if err != nil {
		panic(err)
	}

	if bytes.Equal(owner.Bytes(), common.HexToAddress(config.ChainTree.Address).Bytes()) {
		fmt.Println("start subscribe notarized event for bls aggreation (you're network account owner)")
		if err := notary.subscribeNotarized(ctx); err != nil {
			panic(err)
		}
	}

}

func (notary *Notary) subscribeNotarized(ctx *context.Context) error {
	config := ctx.Config()
	hostWSClient := ctx.ChainTree().Host.WSClient()

	notaryPublicABI, err := abi.JSON(strings.NewReader(string(store.NotaryPublicStoreABI)))
	if err != nil {
		panic(err)
	}

	logs := make(chan coretypes.Log)

	sub, err := hostWSClient.SubscribeFilterLogs(
		basectx.Background(),
		ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(config.ChainTree.Host.NotaryPublicAddress)},
		},
		logs,
	)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				fmt.Println(fmt.Errorf("watch notarized subscription err: %d", err))
			case log := <-logs:
				if bytes.Equal(
					crypto.Keccak256Hash([]byte("Notarized(bytes1,uint256,bytes,address)")).Bytes(),
					log.Topics[0].Bytes(),
				) {
					var event store.NotaryPublicStoreNotarized

					if err := notaryPublicABI.UnpackIntoInterface(&event, "Notarized", log.Data); err != nil {
						fmt.Println(fmt.Errorf("faild to unpack event: %d", err))
					} else {
						if err := notary.checkAggreation(ctx, &event); err != nil {
							fmt.Println(fmt.Errorf("failed aggreation: %d", err))
						}
					}
				}

			}
		}
	}()

	return nil
}

func (notary *Notary) checkAggreation(ctx *context.Context, log *store.NotaryPublicStoreNotarized) error {
	hostClient := ctx.ChainTree().Host.Client()

	notaryPublic, err := store.NewNotaryPublicStore(
		common.HexToAddress(ctx.Config().ChainTree.Host.NotaryPublicAddress),
		hostClient,
	)
	if err != nil {
		return err
	}

	var signers []common.Address
	var response []string

	blockNumber, err := hostClient.BlockNumber(basectx.TODO())
	if err != nil {
		return err
	}

	if err := hostClient.Client().Call(
		&response,
		"clique_getSigners",
		fmt.Sprintf("0x%x", blockNumber),
	); err != nil {
		return err
	} else {
		for _, signer := range response {
			signers = append(signers, common.HexToAddress(signer))
		}
	}

	var signatures [][]byte
	var pubkeys []blscommon.PublicKey

	for _, signer := range signers {
		if pubkeyBz, signature, err := notaryPublic.GetNotarization(
			&bind.CallOpts{Pending: true},
			log.Prefix,
			log.AppID,
			signer,
		); err != nil {
			return err
		} else {
			if len(signature) >= 1 {
				if len(pubkeyBz) == 0 {
					return fmt.Errorf("bls public key is not registered: %s", signer.Hex())
				} else {
					if pubKey, err := blst.PublicKeyFromBytes(pubkeyBz); err != nil {
						return errors.Join(err, fmt.Errorf("failed to gen bls pubkey from bytes: %s", signer.Hex()))
					} else {
						pubkeys = append(pubkeys, pubKey)
						signatures = append(signatures, signature)
					}
				}
			}
		}
	}

	if len(signers) == len(signatures) {
		fmt.Println("ready to bls aggreation")
		for i, siganture := range signatures {
			fmt.Printf("%s's BLS siganture: %s\n", signers[i].Hex(), hex.EncodeToString(siganture))
		}

		aggreatedSiganture, err := blst.AggregateCompressedSignatures(signatures)
		if err != nil {
			return err
		}
		fmt.Printf("aggreated BLS siganture: %s\n", hex.EncodeToString(aggreatedSiganture.Marshal()))

		var message [32]byte
		switch log.Prefix {
		case types.NetworkAccessPermissionPrefix:
			message = types.GetNetwrokAccessPermissionMessage(*log.AppID)
		}

		if verified := aggreatedSiganture.FastAggregateVerify(pubkeys, message); !verified {
			return fmt.Errorf("bls verify failed")
		}

		cato, err := genTransactOpts(hostClient, big.NewInt(int64(ctx.Config().ChainTree.Host.ChainID)), ctx.Config().ChainTree.PrivateKey)
		if err != nil {
			return err
		}

		cnatx, err := notaryPublic.CreateNotaryAccount(cato, log.AppID)
		if err != nil {
			return err
		}

		ctx.Logger().Infof("notary account created tx: %s", cnatx.Hash())

		rc, err := hostClient.TransactionReceipt(basectx.TODO(), cnatx.Hash())
		if err != nil {
			return err
		}

		ctx.Logger().Infof("notary account created receipt found: %d", rc)

		var nacEvent store.NotaryPublicStoreNotaryAccountCreated

		if err := notaryPublicABI.UnpackIntoInterface(&nacEvent, "NotaryAccountCreated", rc.Logs[0].Data); err != nil {
			return err
		}

		return notary.responseNetworkAccessPermission(ctx, log.AppID, aggreatedSiganture.Marshal(), nacEvent.NotaryAccount)
	} // else wait

	return nil
}

func init() {
	var err error
	notaryPublicABI, err = abi.JSON(strings.NewReader(string(store.NotaryPublicStoreABI)))
	if err != nil {
		panic(err)
	}
}
