package middleware

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/labstack/echo/v4"
)

func MiddlewareNetworkAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*context.Context)

		appPermission, err := store.NewAppPermissionStore(
			common.HexToAddress(ctx.Config().ChainTree.Root.AppPermissionAddress),
			ctx.ChainTree().Root.Client(),
		)
		if err != nil {
			return err
		}

		appExists, appID, err := appPermission.GetAppID(
			&bind.CallOpts{Pending: true},
			hexutil.Encode([]byte(ctx.Context.RealIP())),
		)
		if err != nil {
			return err
		} else if !appExists {
			return errors.New("does not exist app at request ip")
		}

		if _, _, isGranted, err := appPermission.GetNetworkAccessResponse(
			&bind.CallOpts{Pending: true},
			appID,
			common.HexToAddress(ctx.Config().ChainTree.Root.NetworkAccountAddress),
		); err != nil {
			return err
		} else if !isGranted {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
