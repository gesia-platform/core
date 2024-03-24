package context

import (
	"github.com/gesia-platform/core/chaintree"
	"github.com/gesia-platform/core/config"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context

	config *config.Config

	chainTree *chaintree.ChainTree
}

func (c *Context) SetConfig(config *config.Config) {
	c.config = config
}

func (c *Context) Config() *config.Config {
	return c.config
}

func (c *Context) SetChainTree(chainTree *chaintree.ChainTree) {
	c.chainTree = chainTree
}

func (c *Context) ChainTree() *chaintree.ChainTree {
	return c.chainTree
}
