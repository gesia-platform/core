package notary

import (
	"bytes"
	basectx "context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/gesia-platform/core/types"
)

func (notary *Notary) SubscribeNetworkAccessRequested(ctx *context.Context) {
	config := ctx.Config()
	rootWSClient := ctx.ChainTree().Root.WSClient()

	appPermissionABI, err := abi.JSON(strings.NewReader(string(store.AppPermissionStoreABI)))
	if err != nil {
		panic(err)
	}

	logs := make(chan coretypes.Log)

	sub, err := rootWSClient.SubscribeFilterLogs(
		basectx.Background(),
		ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(config.ChainTree.Root.AppPermissionAddress)},
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
				fmt.Println(fmt.Errorf("watch network access requested subscription err: %d", err))
			case log := <-logs:
				var event store.AppPermissionStoreNetworkAccessPermissionRequested

				if err := appPermissionABI.UnpackIntoInterface(&event, "NetworkAccessPermissionRequested", log.Data); err != nil {
					fmt.Println(fmt.Errorf("faild to unpack event: %d", err))
				} else {
					if bytes.Equal(
						event.NetworkAccount.Bytes(),
						common.HexToAddress(config.ChainTree.Root.NetworkAccountAddress).Bytes(),
					) {
						if err := notarize(ctx, types.NetworkAccessPermissionPrefix, event.AppID); err != nil {
							fmt.Println(fmt.Errorf("notarize network access requested subscription err: %d", err))
						}
					}
				}
			}
		}
	}()

}
