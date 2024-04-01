package handler

import (
	"strings"

	apimiddleware "github.com/gesia-platform/core/api/middleware"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/notary"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIHandler struct {
	Mux     *echo.Echo
	Notary  *notary.Notary
	Context *context.Context
}

func NewAPIHandler(ctx *context.Context, notary *notary.Notary) *APIHandler {
	mux := echo.New()

	apiHandler := &APIHandler{Mux: mux, Notary: notary, Context: ctx}

	mux.Use(middleware.Logger())
	mux.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := ctx
			cc.Context = c
			return next(cc)
		}
	})

	apiHandler.registerEthereum()

	return apiHandler
}

func (apiHandler *APIHandler) registerEthereum() {
	hostConfig := apiHandler.Context.Config().ChainTree.Host

	proxyPath := hostConfig.ProxyPath

	ethereum := apiHandler.Mux.Group(proxyPath)

	ethereum.Use(apimiddleware.MiddlewareNetworkAccess)

	ethereum.Any("", func(c echo.Context) error {
		var url string

		if strings.HasPrefix(c.Request().URL.Scheme, "ws") {
			url = hostConfig.WSURL
		} else {
			url = hostConfig.RPCURL
		}

		return apiHandler.HandleProxy(c, url)
	})
}
