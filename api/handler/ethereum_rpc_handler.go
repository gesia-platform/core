package handler

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/labstack/echo/v4"
)

func (handler *APIHandler) EthereumRPCHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*context.Context)

		appID, err := strconv.ParseInt(c.Request().Header.Get("x-application-id"), 10, 64)
		if err != nil {
			return err
		}

		appPermission, err := store.NewAppPermissionStore(
			common.HexToAddress(ctx.Config().ChainTree.Root.AppPermissionAddress),
			ctx.ChainTree().Root.Client(),
		)
		if err != nil {
			return err
		}

		requested, iphex, err := appPermission.GetNetworkAccessRequest(
			&bind.CallOpts{Pending: true},
			big.NewInt(appID),
			common.HexToAddress(ctx.Config().ChainTree.Root.NetworkAccountAddress),
		)
		if err != nil {
			return err
		}

		ipbz, err := hexutil.Decode(iphex)
		if err != nil {
			return err
		}

		ip := string(ipbz[:])

		if !requested || !strings.EqualFold(ip, ctx.Context.RealIP()) {
			return echo.ErrUnauthorized
		}

		if _, _, isGranted, err := appPermission.GetNetworkAccessResponse(
			&bind.CallOpts{Pending: true},
			big.NewInt(appID),
			common.HexToAddress(ctx.Config().ChainTree.Root.NetworkAccountAddress),
		); err != nil {
			return err
		} else if !isGranted {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
