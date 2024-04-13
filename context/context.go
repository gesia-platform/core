package context

import (
	"math/big"

	"github.com/gesia-platform/core/chaintree"
	"github.com/gesia-platform/core/config"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type Context struct {
	echo.Context

	config *config.Config

	chainTree *chaintree.ChainTree

	appID *big.Int

	keychain *redis.Client
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

func (c *Context) SetAppID(appID *big.Int) {
	c.appID = appID
}

func (c *Context) AppID() *big.Int {
	return c.appID
}

func (c *Context) SetKeychain(keychain *redis.Client) {
	c.keychain = keychain
}

func (c *Context) Keychain() *redis.Client {
	return c.keychain
}
