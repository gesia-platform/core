package handler

import (
	"net/url"

	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/notary"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIHandler struct {
	Mux    *echo.Echo
	Notary *notary.Notary
}

func NewAPIHandler(ctx *context.Context, notary *notary.Notary) *APIHandler {
	mux := echo.New()

	apiHandler := &APIHandler{Mux: mux, Notary: notary}

	mux.Use(middleware.Logger())
	mux.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := ctx
			cc.Context = c
			return next(cc)
		}
	})

	ethereum := mux.Group(ctx.Config().ChainTree.Host.ProxyPath)

	ethereum.Use(apiHandler.EthereumRPCHandler)

	proxyURL, err := url.Parse(ctx.Config().ChainTree.Host.RPCURL)
	if err != nil {
		panic("failed to parse ethereum rpc url.")
	}

	ethereum.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: proxyURL,
		},
	})))

	return apiHandler
}
