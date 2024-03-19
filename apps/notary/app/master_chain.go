package app

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (app *NotaryApplication) setupMasterChain() {
	client, err := ethclient.Dial(app.config.MasterWSURL)
	if err != nil {
		panic(fmt.Sprintf("failed to dial ethclient. %d", err))
	}

	app.masterClient = client

	app.notarizeMissedAccounts()
	app.subscribeNotarize()
}
