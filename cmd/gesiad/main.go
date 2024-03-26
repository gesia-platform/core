package main

import (
	"github.com/gesia-platform/core/api"
	"github.com/gesia-platform/core/api/handler"
	"github.com/gesia-platform/core/chaintree"
	"github.com/gesia-platform/core/config"
	"github.com/gesia-platform/core/context"
	"github.com/gesia-platform/core/notary"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "Start gesia poa node",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			panic(err)
		}

		start(configPath)
	},
}

func start(configPath string) {
	config := config.NewConfig(configPath)

	chainTree := chaintree.NewChainTree(config.ChainTree.Root.RPCURL, config.ChainTree.Host.RPCURL)

	context := &context.Context{}

	context.SetChainTree(chainTree)
	context.SetConfig(config)

	notary := notary.NewNotary()

	notary.SubscribeNetworkAccessRequested(context)
	notary.SubscribeNotarizedWithCondition(context)

	apiHandler := handler.NewAPIHandler(context, notary)

	api := api.NewAPI(config.Port, apiHandler)

	api.Start()
}

func init() {
	rootCmd.Flags().StringP("config", "c", "", "toml config absolute path")
}
