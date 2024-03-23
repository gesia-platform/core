package notarycontext

import (
	"github.com/gesia-platform/core/apps/notary/config"
	"github.com/labstack/echo/v4"
)

type NotaryContext struct {
	echo.Context

	config *config.Config
}

func (c *NotaryContext) SetConfig(config *config.Config) {
	c.config = config
}

func (c *NotaryContext) Config() *config.Config {
	return c.config
}
