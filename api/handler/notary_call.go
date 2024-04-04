package handler

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"math/big"

	gocontext "context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/store"
	"github.com/labstack/echo/v4"
)

func (handler *APIHandler) NotaryCall(c echo.Context) error {
	body := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return err
	}

	abiData, err := hexutil.Decode(body["data"].(string))
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(body["to"].(string))

	ctx := c.(*context.Context)
	client := ctx.ChainTree().Host.Client()
	chainId := big.NewInt(int64(ctx.Config().ChainTree.Host.ChainID))

	notaryAccount, err := store.NewNotaryAccountStore(
		ctx.NotaryAccount(),
		client,
	)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(ctx.Config().ChainTree.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(gocontext.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(gocontext.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainId,
	)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1000000)
	auth.GasPrice = gasPrice

	tx, err := notaryAccount.NotaryCall(auth, toAddress, abiData)
	if err != nil {
		return err
	}

	c.Logger().Info("notary call transaction: ", tx.Hash())

	if err := c.JSON(200, tx); err != nil {
		return err
	}

	return nil
}
