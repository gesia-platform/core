package handler

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/gesia-platform/core/types"
	"github.com/labstack/echo/v4"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
)

func (handler *APIHandler) EthereumRPCHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*context.Context)

		appID, err := strconv.ParseInt(c.Request().Header.Get("x-application-id"), 10, 64)
		if err != nil {
			return err
		}

		appPermission, err := store.NewAppPermissionStore(
			common.HexToAddress(ctx.Config().ChainTree.Root.Address),
			ctx.ChainTree().Root.Client(),
		)
		if err != nil {
			return err
		}

		requested, ip, err := appPermission.GetNetworkAccessRequest(
			&bind.CallOpts{Pending: true},
			big.NewInt(appID),
			common.HexToAddress(ctx.Config().ChainTree.Root.NetworkAccountAddress),
		)
		if err != nil {
			return err
		}

		if !requested || !strings.EqualFold(ip, ctx.Context.RealIP()) {
			return echo.ErrUnauthorized
		}

		notaryPublic, err := store.NewNotaryPublicStore(
			common.HexToAddress(ctx.Config().ChainTree.Host.NotaryPublicAddress),
			ctx.ChainTree().Host.Client(),
		)
		if err != nil {
			return err
		}

		// Fetch Clique Signers
		var signers []common.Address
		var response types.JsonRPCResponse

		if err := ctx.ChainTree().Host.Client().Client().Call(
			&response,
			"clique_getSigners",
		); err != nil {
			return err
		} else {
			for _, signer := range response.Result {
				signers = append(signers, common.HexToAddress(signer))
			}
		}

		var signatures [][]byte
		var pubkeys []blst.PublicKey

		for _, signer := range signers {
			if pubkey, signature, err := notaryPublic.GetNotarization(
				&bind.CallOpts{Pending: true},
				types.NetworkAccessPermissionPrefix,
				signer,
			); err == nil && len(signature) >= 1 {
				signatures = append(signatures, signature)
				if key, err := blst.PublicKeyFromBytes(pubkey); err != nil {
					return err
				} else {
					pubkeys = append(pubkeys, key)
				}
			}
		}

		if len(pubkeys) == 0 {
			return ehco.ErrUnauthorized
		}

		message := types.GetNetwrokAccessPermissionMessage(*big.NewInt(appID))

		aggreatedSig, err := blst.AggregateCompressedSignatures(signatures)
		if err != nil {
			return err
		}

		if verified := aggreatedSig.FastAggregateVerify(pubkeys, message); !verified {
			return ehco.ErrUnauthorized
		}

		return next(c)
	}
}
