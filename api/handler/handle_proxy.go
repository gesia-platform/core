package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler *APIHandler) HandleProxy(c echo.Context, url string) error {
	req, err := http.NewRequest(
		c.Request().Method,
		url,
		c.Request().Body,
	)
	if err != nil {
		return err
	}

	for name, headers := range c.Request().Header {
		for _, header := range headers {
			req.Header.Add(name, header)
		}
	}

	client := &http.Client{}

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
}
