package app

import (
	"fmt"
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gesia-platform/core/apps/notary-account/handler"
	"github.com/labstack/echo/v4/middleware"
)

func (app *NotaryApplication) setupLocalChain() {
	ethereumRpcUrl, err := url.Parse(app.config.Chain.Local.RPCURL)

	if err != nil {
		panic("failed to parse ethereum rpc url.")
	}

	ethereum := app.echo.Group(app.config.Chain.Local.RPCProxyPath)

	ethereum.Use(handler.ETHRPCHandler)

	ethereum.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: ethereumRpcUrl,
		},
	})))

	app.logger.Info("Successfully initlized ethereum rpc proxy server.")

	client, err := ethclient.Dial(app.config.Chain.Local.RPCURL)
	if err != nil {
		panic(fmt.Sprintf("failed to dial host ethclient. %d", err))
	}

	app.hostClient = client
}

func (app *NotaryApplication) setupMasterChain() {
	client, err := ethclient.Dial(app.config.Chain.Master.RPCURL)
	if err != nil {
		panic(fmt.Sprintf("failed to dial ethclient. %d", err))
	}

	app.masterClient = client
}
