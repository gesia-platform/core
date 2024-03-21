package app

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gesia-platform/core/apps/notary-account/config"
	"github.com/gesia-platform/core/apps/notary-account/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type NotaryApplication struct {
	echo   *echo.Echo
	logger *logrus.Logger
	config *config.Config
	client *ethclient.Client
}

func NewNotaryApplication(configPath string) *NotaryApplication {
	app := &NotaryApplication{
		echo:   echo.New(),
		logger: logrus.New(),
		config: config.NewConfig(),
	}

	app.echo.Use(middleware.Logger())

	app.logger.Formatter = &logrus.JSONFormatter{}
	app.logger.Out = os.Stdout

	app.config.Read(configPath)

	app.setup()

	return app
}

func (app *NotaryApplication) Run() {
	app.notarizeMissedAccounts()

	app.subscribeNotarize()
}

func (app *NotaryApplication) setup() {
	ethereumRPCURL, err := url.Parse(app.config.RPCURL)
	if err != nil {
		panic("failed to parse ethereum rpc url.")
	}

	ethereum := app.echo.Group(app.config.ProxyPath)

	ethereum.Use(handler.EthereumRPCHandler)
	ethereum.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: ethereumRPCURL,
		},
	})))

	app.logger.Info("Successfully initlized ethereum rpc proxy server.")

	client, err := ethclient.Dial(app.config.RPCURL)
	if err != nil {
		panic(fmt.Sprintf("failed to dial ethclient. %d", err))
	}

	app.client = client
}
