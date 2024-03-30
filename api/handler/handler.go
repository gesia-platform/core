package handler

import (
	"net/http"
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

	proxyPath := ctx.Config().ChainTree.Host.ProxyPath
	proxyURL, err := url.Parse(ctx.Config().ChainTree.Host.RPCURL)
	if err != nil {
		panic(err)
	}

	ethereum := mux.Group(proxyPath)

	ethereum.Use(apiHandler.EthereumRPCHandler)

	// json RPC
	ethereum.POST("", func(c echo.Context) error {
		client := &http.Client{}

		req, err := http.NewRequest(c.Request().Method, proxyURL.String(), c.Request().Body)
		if err != nil {
			return err
		}

		for name, headers := range c.Request().Header {
			for _, header := range headers {
				req.Header.Add(name, header)
			}
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		for name, headers := range resp.Header {
			for _, header := range headers {
				c.Response().Header().Add(name, header)
			}
		}

		return c.Stream(
			resp.StatusCode,
			resp.Header.Get("Content-Type"),
			resp.Body,
		)
	})

	return apiHandler
}
