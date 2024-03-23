package app

import (
	"net/url"
	"os"

	"github.com/gesia-platform/core/apps/notary/chainclient"
	"github.com/gesia-platform/core/apps/notary/config"
	"github.com/gesia-platform/core/apps/notary/handler"
	"github.com/gesia-platform/core/apps/notary/notarycontext"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type NotaryApplication struct {
	echo   *echo.Echo
	logger *logrus.Logger

	config *config.Config
}

func NewNotaryApplication(configPath string) *NotaryApplication {
	app := &NotaryApplication{
		echo:   echo.New(),
		logger: logrus.New(),
		config: config.NewConfig(),
	}

	app.echo.Use(middleware.Logger())

	app.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &notarycontext.NotaryContext{Context: c}
			cc.SetConfig(app.config)
			return next(cc)
		}
	})

	app.logger.Formatter = &logrus.JSONFormatter{}
	app.logger.Out = os.Stdout

	app.config.Read(configPath)

	return app
}

func (app *NotaryApplication) Setup() {
	client := chainclient.NewChainClient(app.config.RPCURL, app.config.NotaryChainRPCURL)

	handler := handler.NewHandler(client, app.logger)

	proxyURL, err := url.Parse(app.config.RPCURL)
	if err != nil {
		panic("failed to parse ethereum rpc url.")
	}

	ethereum := app.echo.Group(app.config.ProxyPath)

	ethereum.Use(handler.EthereumRPCHandler)

	ethereum.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: proxyURL,
		},
	})))

}

func (app *NotaryApplication) Run() {
	// app.notarizeMissedAccounts()

	// app.subscribeNotarize()
}
