package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (app *NotaryApplication) notarize(notaryAccount common.Address) error {
	var response RPCResponse
	var domain string

	signer := common.HexToAddress(app.config.Address)

	if err := app.hostClient.Client().Call(
		&response,
		"account_signData",
		accounts.MimetypeTextPlain,
		&signer,
		hexutil.Encode([]byte(domain))); err != nil {
		return err
	}

	// 1. signature
	// app.logger.Info("Successfully signed to account. %d", response)

	// 2. submit to account

	return nil
}

func (app *NotaryApplication) notarizeMissedAccounts() {
	app.logger.Info(fmt.Sprintf("start notarizing for missed accounts."))

	notaryPublicAddress := common.HexToAddress(app.config.Chain.Master.NotaryPublicAddress)
	createdEventSignature := crypto.Keccak256Hash([]byte("NotaryAccountCreated(string,string,address)"))
	notarizedEventSignature := crypto.Keccak256Hash([]byte("RPCNotarized(address,string)"))

	notaryPublicAbi, err := abi.JSON(strings.NewReader(NotaryPublicABI))
	if err != nil {
		panic(fmt.Sprintf("failed to read notary public abi json. %d", err))
	}

	notaryAccountAbi, err := abi.JSON(strings.NewReader(NotaryAccountABI))
	if err != nil {
		panic(fmt.Sprintf("failed to read notary account abi json. %d", err))
	}

	publicLogs, err := app.masterClient.FilterLogs(context.TODO(), ethereum.FilterQuery{
		Addresses: []common.Address{notaryPublicAddress},
		Topics:    [][]common.Hash{{createdEventSignature}},
	})

	if err != nil {
		panic(fmt.Sprintf("failed to missed notary account filter. %d", err))
	}

	var chainAddresses []common.Address

	for _, log := range publicLogs {
		var event EventNotaryAccountCreated

		err := notaryPublicAbi.UnpackIntoInterface(&event, "NotaryAccountCreated", log.Data)
		if err != nil {
			panic(fmt.Sprintf("failed to unpack notary public event. %d", err))
		}

		if !strings.EqualFold(event.Chain, app.config.Chain.Master.NotaryPublicChainPrefix) {
			continue
		}

		chainAddresses = append(chainAddresses, event.NotaryAccount)
	}

	accountLogs, err := app.masterClient.FilterLogs(context.TODO(), ethereum.FilterQuery{
		Addresses: chainAddresses,
		Topics:    [][]common.Hash{{notarizedEventSignature}},
	})

	var pendingAccounts []common.Address

	for _, log := range accountLogs {
		var event EventRPCNotarized

		err := notaryAccountAbi.UnpackIntoInterface(&event, "RPCNotarized", log.Data)
		if err != nil {
			panic(fmt.Sprintf("failed to unpack notary account event. %d", err))
		}

		if !strings.EqualFold(event.Notary.Hex(), app.config.Address) {
			pendingAccounts = append(pendingAccounts, event.Notary)
		}
	}

	for _, account := range pendingAccounts {
		if err := app.notarize(account); err != nil {
			panic(fmt.Sprintf("failed to notarize: %s. %d", account.Hex(), err))
		}
	}

	app.logger.Info(fmt.Sprintf("finished notarizing for missed accounts."))
}

func (app *NotaryApplication) subscribeNotarize() {
	app.logger.Info(fmt.Sprintf("start subscribing notarize."))

	notaryPublicAddress := common.HexToAddress(app.config.Chain.Master.NotaryPublicAddress)
	createdEventSignature := crypto.Keccak256Hash([]byte("NotaryAccountCreated(string,string,address)"))

	notaryPublicAbi, err := abi.JSON(strings.NewReader(NotaryPublicABI))
	if err != nil {
		panic(fmt.Sprintf("failed to read notary public abi json. %d", err))
	}

	publicLogs := make(chan coretypes.Log)

	sub, err := app.masterClient.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{notaryPublicAddress},
		Topics:    [][]common.Hash{{createdEventSignature}},
	}, publicLogs)
	if err != nil {
		panic(fmt.Sprintf("failed to subscribe filter logs. %d", err))
	}

	for {
		select {
		case err := <-sub.Err():
			app.logger.Error(fmt.Sprintf("error from filter logs channel. %d", err))
		case log := <-publicLogs:
			var event EventNotaryAccountCreated

			if err := notaryPublicAbi.UnpackIntoInterface(&event, "NotaryAccountCreated", log.Data); err != nil {
				app.logger.Error(fmt.Sprintf("failed to unpack notary public event. %d", err))
			} else {
				if strings.EqualFold(event.Chain, app.config.Chain.Master.NotaryPublicChainPrefix) {
					app.notarize(event.NotaryAccount)
				}
			}
		}
	}
}
