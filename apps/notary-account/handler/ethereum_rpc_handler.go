package handler

import (
	"github.com/labstack/echo/v4"
)

func EthereumRPCHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		// 1. Extract address from header
		// 2. Check notary address signed from all signer at snapshostBlockHeight
		// 3. check dns ip address and ip
		return err
	}
}
