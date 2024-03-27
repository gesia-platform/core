package notary

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/gesia-platform/core/types"
)

func (notary *Notary) SubscribeNetworkAccessRequested(ctx *context.Context) {
	config := ctx.Config()
	rootWSClient := ctx.ChainTree().Root.WSClient()

	appStore, err := store.NewAppPermissionStore(
		common.HexToAddress(config.ChainTree.Root.AppStoreAddress),
		rootWSClient,
	)
	if err != nil {
		panic(err)
	}

	var logs chan *store.AppPermissionStoreNetworkAccessPermissionRequested

	sub, err := appStore.WatchNetworkAccessPermissionRequested(
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
				fmt.Println(fmt.Errorf("watch network access requested subscription err: %d", err))
			case log := <-logs:
				fmt.Printf("watched network access requested log: %d\n", &log)
				if bytes.Equal(
					log.NetworkAccount.Bytes(),
					common.HexToAddress(config.ChainTree.Root.NetworkAccountAddress).Bytes(),
				) {
					if err := notarize(ctx, types.NetworkAccessPermissionPrefix, log.AppID); err != nil {
						fmt.Println(fmt.Errorf("notarize network access requested subscription err: %d", err))
					}
				}
			}
		}
	}()

}
