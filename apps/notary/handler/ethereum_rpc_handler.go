package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gesia-platform/core/apps/notary/notarycontext"
	"github.com/gesia-platform/core/apps/notary/store"
	"github.com/labstack/echo/v4"
)

type ResponseJsonRPC struct {
	Result []string
}

func (handler *Handler) EthereumRPCHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nc := c.(*notarycontext.NotaryContext)

		applicationID, err := strconv.ParseInt(c.Request().Header.Get("x-application-id"), 10, 64)
		if err != nil {
			return err
		}

		// Instantiate notary public contract
		instance, err := store.NewNotaryPublicStore(common.HexToAddress(nc.Config().NotaryPublicAddress), handler.chainClient.Notary)
		if err != nil {
			return err
		}

		// Get domain
		domain, _, err := instance.GetApplicationDetails(
			&bind.CallOpts{Pending: true},
			big.NewInt(applicationID),
		)

		if err != nil {
			return err
		}

		if len([]byte(domain)) == 0 {
			return fmt.Errorf("domain not registered by applicationID. %d", applicationID)
		}

		// Fetch Clique Signers
		var signers []common.Address

		var body ResponseJsonRPC

		if resp, err := http.Post(nc.Config().RPCURL, "application/json", bytes.NewBuffer(
			[]byte(`{"jsonrpc":"2.0","method": "clique_getSigners", "params":[], "id": 2}`),
		)); err != nil {
			return err
		} else {
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				return err
			}

			for _, signer := range body.Result {
				signers = append(signers, common.HexToAddress(signer))
			}
		}

		// Check notrazied from signer
		var signatures int
		for _, signer := range signers {
			if notarized, _, _, err := instance.GetApplicationNotarizationDetails(
				&bind.CallOpts{Pending: true},
				big.NewInt(applicationID),
				nc.Config().Chain,
				signer,
			); err != nil {
				return err
			} else {
				if notarized {
					signatures++
				}
			}
		}

		quorum := int(math.Round(float64(len(signers)) * 2 / 3))
		if quorum > signatures {
			return errors.New("insufficient 2/3 required signature quorum")
		}

		// Check ip within domain
		ips, err := net.LookupIP(domain)
		if err != nil {
			return err
		}

		var whitelistedIP bool
		for _, ip := range ips {
			if strings.EqualFold(ip.To4().String(), nc.RealIP()) {
				whitelistedIP = true
			}
		}

		if !whitelistedIP {
			return echo.ErrForbidden
		}

		return next(c)
	}
}
