package main

import (
	"encoding/hex"
	"fmt"
	"os"

	gocontext "context"

	"github.com/gesia-platform/core/api"
	"github.com/gesia-platform/core/api/handler"
	"github.com/gesia-platform/core/chaintree"
	"github.com/gesia-platform/core/config"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/notary"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	"github.com/redis/go-redis/v9"
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
	context.SetKeychain(redis.NewClient(&redis.Options{
		Addr:       config.Keychain.Host,
		ClientName: config.ChainTree.Address,
		Username:   config.ChainTree.Address,
		Password:   config.Keychain.Password,
		DB:         0,
	}))

	if err := context.Keychain().Ping(gocontext.Background()).Err(); err != nil {
		panic("keychain connection unstabled")
	} else {
		fmt.Printf("keychain connected host: %s, user: %s\n", config.Keychain.Host, config.ChainTree.Address)
	}

	notary := notary.NewNotary()

	//notary.SubscribeNetworkAccessRequested(context)
	fmt.Println("root chain network access requested event subscribed")
	//notary.SubscribeNotarizedWithCondition(context)

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

}
