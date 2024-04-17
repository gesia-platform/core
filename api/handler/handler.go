package handler

import (
	"net/http"
	"strings"

	apimiddleware "github.com/gesia-platform/core/api/middleware"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/notary"
	"github.com/gesia-platform/core/types"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pretty66/websocketproxy"
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

	wp, err := websocketproxy.NewProxy(hostConfig.WSURL, func(r *http.Request) error {
		return nil
	})
	if err != nil {
		panic(err)
	}

	ethereum := apiHandler.Mux.Group(types.EthereumProxyPath)

	// Guard
	ethereum.Use(apimiddleware.MiddlewareNetworkAccess)

	// Notary Call
	ethereum.POST(types.EthereumTxPath, apiHandler.NotaryCall)

	ethereum.GET(types.EthereumIOAPath, apiHandler.ListIOAs)

	// JSON-RPC
	ethereum.Any("", func(c echo.Context) error {
		if strings.EqualFold(c.Request().Header.Get("Upgrade"), "websocket") {
			wp.Proxy(c.Response().Writer, c.Request())
			return nil
		} else {
			return apiHandler.HandleProxy(c, hostConfig.RPCURL)
		}
	})

}
