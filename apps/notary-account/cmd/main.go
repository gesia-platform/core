package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gesia-platform/core/apps/notary-account/app"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hugo",
		Short: "Start notary application",
		Run: func(cmd *cobra.Command, args []string) {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil || strings.EqualFold(configPath, "") {
				panic(fmt.Errorf("invalid config flag. %d", err))
			}

			app.NewNotaryApplication(configPath)
		},
	}

	rootCmd.Flags().String("config", "", ".toml config path")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}