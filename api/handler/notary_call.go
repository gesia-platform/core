package handler

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	gocontext "context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gesia-platform/core/context"
	"github.com/labstack/echo/v4"
)

func (handler *APIHandler) NotaryCall(c echo.Context) error {
	body := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return err
	}

	var toPlain string
	if body["to"] == nil {
		toPlain = ""
	} else {
		toPlain = body["to"].(string)
	}

	dataPlain := body["data"].(string)

	fmt.Printf("notary call requested to: %s, data: %s\n", toPlain, dataPlain)

	abiData, err := hexutil.Decode(dataPlain)
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(toPlain)

	ctx := c.(*context.Context)
	client := ctx.ChainTree().Host.Client()
	chainId := big.NewInt(int64(ctx.Config().ChainTree.Host.ChainID))

	// Keychain에서 불러오기
	pkHex, err := ctx.Keychain().Get(gocontext.Background(), ctx.AppID().String()).Result()
	if err != nil {
		return err
	} else if len([]byte(pkHex)) == 0 {
		return errors.New("empty notary private key")
	}

	privateKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(gocontext.TODO(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(gocontext.TODO())
	if err != nil {
		return err
	}

	var tx *types.Transaction
	if strings.EqualFold(toPlain, "") {
		tx = types.NewContractCreation(nonce, big.NewInt(0), uint64(4700000), gasPrice, abiData)
	} else {
		tx = types.NewTransaction(nonce, toAddress, big.NewInt(0), uint64(4700000), gasPrice, abiData)

	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		fmt.Println("failed to sign notarized transaction")
		return err
	}

	if err := client.SendTransaction(gocontext.Background(), signedTx); err != nil {
		fmt.Printf("failed notary transaction: %s\n", signedTx.Hash())
		return err
	} else {
		fmt.Printf("success notary transaction: %s\n", signedTx.Hash())
	}

	return c.JSON(200, signedTx)
}
