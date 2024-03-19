package app

import (
	"fmt"
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *NotaryApplication) setupHostChain() {
	ethereumRpcUrl, err := url.Parse(app.config.RPCURL)

	if err != nil {
		panic("failed to parse ethereum rpc url.")
	}

	ethereum := app.echo.Group(app.config.RPCProxyPath)

	ethereum.Use(GuardNotary)

	ethereum.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: ethereumRpcUrl,
		},
	})))

	app.logger.Info("Successfully initlized ethereum rpc proxy server.")

	client, err := ethclient.Dial(app.config.RPCURL)
	if err != nil {
		panic(fmt.Sprintf("failed to dial host ethclient. %d", err))
	}

	app.hostClient = client
}

func GuardNotary(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		// 1. Extract address from header
		// 2. Check notary address signed from all signer at snapshostBlockHeight
		// 3. check dns ip address and ip
		return err
	}
}
