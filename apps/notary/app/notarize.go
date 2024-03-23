package app

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gesia-platform/core/apps/notary/store"
)

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func (app *NotaryApplication) notarize(applicationID uint64) error {
	notaryPublicAddress := common.HexToAddress(app.config.NotaryPublicAddress)
	notaryPublic, err := store.NewNotaryPublicStore(notaryPublicAddress, app.chainClient.Notary)
	if err != nil {
		panic(err)
	}

	bn := new(big.Int)
	bn.SetUint64(applicationID)

	privateKey, err := crypto.HexToECDSA(app.config.SignerPrivateKey)
	if err != nil {
		return err
	}

	stringT, _ := abi.NewType("string", "", nil)
	uint256T, _ := abi.NewType("uint256", "", nil)

	args := abi.Arguments{{Type: stringT}, {Type: uint256T}}

	packed, err := args.Pack([]byte("notarize_application"), bn)
	if err != nil {
		return err
	}

	messageHash := crypto.Keccak256Hash(packed)

	prefixedHash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(messageHash))),
		messageHash.Bytes(),
	)

	signature, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
	signature2, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return err
	}

	fmt.Println(signature, signature2)

	nonce, err := app.chainClient.Notary.PendingNonceAt(context.Background(), common.HexToAddress(app.config.SignerAddress))
	if err != nil {
		return err
	}

	gasPrice, err := app.chainClient.Notary.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(app.config.ChainID)))
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := notaryPublic.NotarizeApplication(auth, bn, app.config.Chain, app.config.ExternalURL, signature)
	if err != nil {
		return err
	}

	fmt.Println(tx.Hash().Hex())

	return nil
}

func (app *NotaryApplication) notarizePendingAccounts() {
	app.logger.Info("start notarize pending accounts")
	notaryPublicAddress := common.HexToAddress(app.config.NotaryPublicAddress)
	signerAddress := common.HexToAddress(app.config.SignerAddress)

	notaryPublic, err := store.NewNotaryPublicStore(notaryPublicAddress, app.chainClient.Notary)
	if err != nil {
		panic(err)
	}

	iterator, err := notaryPublic.FilterApplicationCreated(&bind.FilterOpts{})
	if err != nil {
		panic(err)

	}
	defer iterator.Close()

	for iterator.Next() {
		notarizedIterator2, err := notaryPublic.FilterApplicationNotarized(&bind.FilterOpts{Start: iterator.Event.Raw.BlockNumber})
		if err != nil {
			panic(err)
		}

		defer notarizedIterator2.Close()

		var signed bool
		for notarizedIterator2.Next() {
			if notarizedIterator2.Event.Notary == signerAddress {
				signed = true
			}
		}

		if signed {
			continue
		}

		app.logger.Info("found unnotarized domain: ", iterator.Event.Domain, ", applicationID: ", iterator.Event.ApplicationID)

		if err := app.notarize(iterator.Event.ApplicationID.Uint64()); err != nil {
			panic(err)
		}
	}

	/*
	   - createdEventSignature := crypto.Keccak256Hash([]byte("NotaryAccountUpdated(string,address,address)"))
	     notarizedEventSignature := crypto.Keccak256Hash([]byte("RPCNotarized(address,string)"))

	     notaryPublicAbi, err := abi.JSON(strings.NewReader(NotaryPublicABI))
	     if err != nil {
	     panic(fmt.Sprintf("failed to read notary public abi json. %d", err))
	     }

	     notaryAccountAbi, err := abi.JSON(strings.NewReader(NotaryAccountABI))
	     if err != nil {
	     panic(fmt.Sprintf("failed to read notary account abi json. %d", err))
	     }

	     publicLogs, err := app.client.FilterLogs(context.TODO(), ethereum.FilterQuery{
	     Addresses: []common.Address{notaryPublicAddress},
	     Topics:    [][]common.Hash{{createdEventSignature}},
	     })

	     if err != nil {
	     panic(fmt.Sprintf("failed to missed notary account filter. %d", err))
	     }

	     var chainAddresses []common.Address

	     for _, log := range publicLogs {
	     var event EventNotaryAccountUpdated

	     err := notaryPublicAbi.UnpackIntoInterface(&event, "NotaryAccountUpdated", log.Data)
	     if err != nil {
	     panic(fmt.Sprintf("failed to unpack notary public event. %d", err))
	     }

	     chainAddresses = append(chainAddresses, event.NotaryAccount)
	     }

	     accountLogs, err := app.client.FilterLogs(context.TODO(), ethereum.FilterQuery{
	     Addresses: chainAddresses,
	     Topics:    [][]common.Hash{{notarizedEventSignature}},
	     })

	     if err != nil {
	     panic(fmt.Sprintf("failed to missed notary account filter. %d", err))
	     }

	     var pendingAccounts []common.Address

	     for _, log := range accountLogs {
	     var event EventNotarized

	     err := notaryAccountAbi.UnpackIntoInterface(&event, "Notarized", log.Data)
	     if err != nil {
	     panic(fmt.Sprintf("failed to unpack notary account event. %d", err))
	     }

	     if !strings.EqualFold(event.Notary.Hex(), app.config.SignerAddress) {
	     pendingAccounts = append(pendingAccounts, event.Notary)
	     }
	     }

	     for _, account := range pendingAccounts {
	     if err := app.notarize(account); err != nil {
	     panic(fmt.Sprintf("failed to notarize: %s. %d", account.Hex(), err))
	     }
	     }

	     app.logger.Info(fmt.Sprintf("finished notarizing for missed accounts."))
	*/
}

/*

func (app *NotaryApplication) notarize(notaryAccount common.Address) error {
	var response RPCResponse
	var domain string

	signer := common.HexToAddress(app.config.SignerAddress)

	if err := app.client.Client().Call(
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

	notaryPublicAddress := common.HexToAddress(app.config.NotaryPublicAddress)
	createdEventSignature := crypto.Keccak256Hash([]byte("NotaryAccountUpdated(string,address,address)"))
	notarizedEventSignature := crypto.Keccak256Hash([]byte("RPCNotarized(address,string)"))

	notaryPublicAbi, err := abi.JSON(strings.NewReader(NotaryPublicABI))
	if err != nil {
		panic(fmt.Sprintf("failed to read notary public abi json. %d", err))
	}

	notaryAccountAbi, err := abi.JSON(strings.NewReader(NotaryAccountABI))
	if err != nil {
		panic(fmt.Sprintf("failed to read notary account abi json. %d", err))
	}

	publicLogs, err := app.client.FilterLogs(context.TODO(), ethereum.FilterQuery{
		Addresses: []common.Address{notaryPublicAddress},
		Topics:    [][]common.Hash{{createdEventSignature}},
	})

	if err != nil {
		panic(fmt.Sprintf("failed to missed notary account filter. %d", err))
	}

	var chainAddresses []common.Address

	for _, log := range publicLogs {
		var event EventNotaryAccountUpdated

		err := notaryPublicAbi.UnpackIntoInterface(&event, "NotaryAccountUpdated", log.Data)
		if err != nil {
			panic(fmt.Sprintf("failed to unpack notary public event. %d", err))
		}

		chainAddresses = append(chainAddresses, event.NotaryAccount)
	}

	accountLogs, err := app.client.FilterLogs(context.TODO(), ethereum.FilterQuery{
		Addresses: chainAddresses,
		Topics:    [][]common.Hash{{notarizedEventSignature}},
	})

	if err != nil {
		panic(fmt.Sprintf("failed to missed notary account filter. %d", err))
	}

	var pendingAccounts []common.Address

	for _, log := range accountLogs {
		var event EventNotarized

		err := notaryAccountAbi.UnpackIntoInterface(&event, "Notarized", log.Data)
		if err != nil {
			panic(fmt.Sprintf("failed to unpack notary account event. %d", err))
		}

		if !strings.EqualFold(event.Notary.Hex(), app.config.SignerAddress) {
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

	notaryPublicAddress := common.HexToAddress(app.config.NotaryPublicAddress)
	createdEventSignature := crypto.Keccak256Hash([]byte("NotaryAccountUpdated(string,string,address)"))

	notaryPublicAbi, err := abi.JSON(strings.NewReader(NotaryPublicABI))
	if err != nil {
		panic(fmt.Sprintf("failed to read notary public abi json. %d", err))
	}

	publicLogs := make(chan coretypes.Log)

	sub, err := app.client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
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
			var event EventNotaryAccountUpdated

			if err := notaryPublicAbi.UnpackIntoInterface(&event, "NotaryAccountUpdated", log.Data); err != nil {
				app.logger.Error(fmt.Sprintf("failed to unpack notary public event. %d", err))
			} else {
				app.notarize(event.NotaryAccount)
			}
		}
	}
}
*/
