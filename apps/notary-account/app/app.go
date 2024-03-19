package app

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gesia-platform/core/apps/notary-account/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type NotaryApplication struct {
	echo         *echo.Echo
	logger       *logrus.Logger
	config       *config.Config
	masterClient *ethclient.Client
	hostClient   *ethclient.Client
}

func (app *NotaryApplication) Run() {
	app.notarizeMissedAccounts()

	app.subscribeNotarize()
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

	app.setupLocalChain()
	app.setupMasterChain()

	return app
}
