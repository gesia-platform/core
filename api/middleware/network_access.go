package middleware

import (
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/labstack/echo/v4"
)

var cahced = map[string](struct {
	appID *big.Int
	time  time.Time
}){}

func MiddlewareNetworkAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := c.(*context.Context)

		ctx.SetAppID(big.NewInt(6))

		return next(ctx)

		ip := ctx.Context.RealIP()

		cachedApp := cahced[ip]

		if cachedApp.appID == nil || strings.EqualFold(c.Request().URL.Query().Get("cache"), "update") || time.Now().After(cachedApp.time.Add(time.Duration(24)*time.Hour)) {
			appPermission, err := store.NewAppPermissionStore(
				common.HexToAddress(ctx.Config().ChainTree.Root.AppPermissionAddress),
				ctx.ChainTree().Root.Client(),
			)
			if err != nil {
				return err
			}

			appExists, appID, err := appPermission.GetAppID(
				&bind.CallOpts{Pending: true},
				hexutil.Encode([]byte(ip)),
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

			cahced[ip] = struct {
				appID *big.Int
				time  time.Time
			}{
				appID: appID,
				time:  time.Now(),
			}
			cachedApp.appID = appID
		}

		ctx.SetAppID(cachedApp.appID)

		return next(ctx)
	}
}
