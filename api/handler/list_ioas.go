package handler

import (
	"crypto/ecdsa"
	"strings"

	gocontext "context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gesia-platform/core/context"
	"github.com/labstack/echo/v4"
)

var cahced = map[string]string{}

func (handler *APIHandler) ListIOAs(c echo.Context) error {
	ctx := c.(*context.Context)
	goCtx := gocontext.Background()

	keys, err := ctx.Keychain().Keys(goCtx, "*").Result()
	if err != nil {
		return err
	}

	var addresses []string
	for _, key := range keys {
		if !strings.EqualFold(cahced[key], "") {
			addresses = append(addresses, cahced[key])
		} else {
			value, err := ctx.Keychain().Get(goCtx, key).Result()
			if err != nil {
				return err
			}
			privateKey, err := crypto.HexToECDSA(value)
			if err == nil {
				publicKey := privateKey.Public()
				publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
				if ok {
					address := crypto.PubkeyToAddress(*publicKeyECDSA)
					cahced[key] = address.Hex()
					addresses = append(addresses, cahced[key])
				}

			}
		}
	}

	return c.JSON(200, addresses)
}
