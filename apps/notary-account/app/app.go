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

func NewNotaryApplication(configPath string) *NotaryApplication {
	echo := echo.New()
	echo.Use(middleware.Logger())

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Out = os.Stdout

	config := config.NewConfig()
	config.Read(configPath)

	return &NotaryApplication{echo: echo, logger: logger, config: config}
}

func (app *NotaryApplication) Run() {
	app.setupHostChain()
	app.setupMasterChain()
}
