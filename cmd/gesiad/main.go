package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/gesia-platform/core/api"
	"github.com/gesia-platform/core/api/handler"
	"github.com/gesia-platform/core/chaintree"
	"github.com/gesia-platform/core/config"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/notary"
	"github.com/gesia-platform/core/types"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gesiad",
	Short: "Gesia PoA Node",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start gesia poa node",
		Run: func(cmd *cobra.Command, args []string) {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			start(configPath)
		},
	}

	genBlsKeyCmd = &cobra.Command{
		Use:   "gen-bls-key",
		Short: `Generate BLS Key Pair (Hex)`,
		Run: func(cmd *cobra.Command, args []string) {
			sk, err := blst.RandKey()
			if err != nil {
				panic(err)
			}

			skHex := hex.EncodeToString(sk.Marshal())
			pkHex := hex.EncodeToString((sk.PublicKey().Marshal()))

			cmd.Printf("Generated BLS SecretKey: %s\n", skHex)
			cmd.Printf("Generated BLS PublicKey: %s\n", pkHex)
		},
	}

	genBlsSignatureCmd = &cobra.Command{
		Use:   "gen-bls-signature",
		Short: `Generate BLS Signatrue given appID, secretKey, messageType`,
		Run: func(cmd *cobra.Command, args []string) {
			secretKey, err := cmd.Flags().GetString("secret-key")
			if err != nil {
				panic(err)
			}

			appID, err := cmd.Flags().GetUint64("app-id")
			if err != nil {
				panic(err)
			}

			messageType, err := cmd.Flags().GetString("message-type")
			if err != nil {
				panic(err)
			}

			skbz, err := hex.DecodeString(secretKey)
			if err != nil {
				panic(err)
			}

			secret, err := blst.SecretKeyFromBytes(skbz)
			if err != nil {
				panic(err)
			}

			var message [32]byte
			switch messageType {
			case string(types.NetworkAccessPermissionKey):
				message = types.GetNetwrokAccessPermissionMessage(*big.NewInt(int64(appID)))
			}

			cmd.Printf("Signed BLS Signature: %s\n", hex.EncodeToString(secret.Sign(message[:]).Marshal()))
		},
	}
)

func start(configPath string) {
	config := config.NewConfig(configPath)
	fmt.Println("config marshalled")

	chainTree := chaintree.NewChainTree(
		config.ChainTree.Root.RPCURL,
		config.ChainTree.Root.WSURL,
		config.ChainTree.Host.RPCURL,
		config.ChainTree.Host.WSURL,
	)
	fmt.Println("chaintree created")

	context := &context.Context{}

	context.SetChainTree(chainTree)
	context.SetConfig(config)

	notary := notary.NewNotary()

	notary.SubscribeNetworkAccessRequested(context)
	fmt.Println("root chain network access requested event subscribed")
	notary.SubscribeNotarizedWithCondition(context)

	apiHandler := handler.NewAPIHandler(context, notary)

	api := api.NewAPI(config.Port, apiHandler)
	fmt.Println("api created")

	api.Start()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	startCmd.Flags().String("config", "", "toml config absolute path")
	rootCmd.AddCommand(startCmd)

	rootCmd.AddCommand(genBlsKeyCmd)

	genBlsSignatureCmd.Flags().String("secret-key", "", "bls secret key hex")
	genBlsSignatureCmd.Flags().Uint64("app-id", 0, "taget app id")
	genBlsSignatureCmd.Flags().String(
		"message-type",
		"",
		`target meessage type
	1. 'network_access_permission'`)

	rootCmd.AddCommand(genBlsSignatureCmd)
}
