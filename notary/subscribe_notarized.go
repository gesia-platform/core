package notary

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/gesia-platform/core/types"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	blscommon "github.com/prysmaticlabs/prysm/v5/crypto/bls/common"
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
		fmt.Println("you are network owner")
		fmt.Println("start subscribe notarized event")
		notary.subscribeNotarized(ctx)
	}

}

func (notary *Notary) subscribeNotarized(ctx *context.Context) {
	config := ctx.Config()
	hostClient := ctx.ChainTree().Host.WSClient()

	notaryPublic, err := store.NewNotaryPublicStore(
		common.HexToAddress(config.ChainTree.Host.NotaryPublicAddress),
		hostClient,
	)
	if err != nil {
		panic(err)
	}

	var logs chan *store.NotaryPublicStoreNotarized

	sub, err := notaryPublic.WatchNotarized(
		&bind.WatchOpts{},
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
				fmt.Printf("watched notarized log: %d\n", &log)
				notary.checkAggreation(ctx, log, notaryPublic)
			}
		}
	}()

}

func (notary *Notary) checkAggreation(ctx *context.Context, log *store.NotaryPublicStoreNotarized, notaryPublic *store.NotaryPublicStore) error {
	var signers []common.Address
	var response types.JsonRPCResponse

	if err := ctx.ChainTree().Host.Client().Client().Call(
		&response,
		"clique_getSigners",
	); err != nil {
		return err
	} else {
		for _, signer := range response.Result {
			signers = append(signers, common.HexToAddress(signer))
		}
	}

	var signatures [][]byte
	var pubkeys []blscommon.PublicKey

	for _, signer := range signers {
		pubkeyBz, signature, err := notaryPublic.GetNotarization(
			&bind.CallOpts{Pending: true},
			log.Prefix,
			log.AppID,
			signer,
		)
		if err == nil && len(signature) >= 1 {
			signatures = append(signatures, signature)

			pubKey, err := blst.PublicKeyFromBytes(pubkeyBz)
			if err == nil {
				fmt.Println(fmt.Errorf("failed to blst public key from bytes: %d", err))

				pubkeys = append(pubkeys, pubKey)
			}
		}
	}

	if len(signers) == len(signatures) {
		aggreatedSiganture, err := blst.AggregateCompressedSignatures(signatures)
		if err != nil {
			return err
		}

		var message [32]byte
		switch log.Prefix {
		case types.NetworkAccessPermissionPrefix:
			message = types.GetNetwrokAccessPermissionMessage(*log.AppID)
		}

		if verified := aggreatedSiganture.FastAggregateVerify(pubkeys, message); !verified {
			return errors.New("invalid bls signatures")
		}

		return notary.responseNetworkAccessPermission(ctx, log.AppID, aggreatedSiganture.Marshal())
	}

	return nil
}
