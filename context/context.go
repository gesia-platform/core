package context

import (
	"github.com/gesia-platform/core/config"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Context struct {
	echo.Context

	config *config.Config

	logger *logrus.Logger
}

func (c *Context) SetConfig(config *config.Config) {
	c.config = config
}

func (c *Context) Config() *config.Config {
	return c.config
}

func (c *Context) SetLogger(logger *logrus.Logger) {
	c.logger = logger
}

func (c *Context) Logger() *logrus.Logger {
	return c.logger
}
