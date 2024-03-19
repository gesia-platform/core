package main

import "github.com/gesia-platform/core/apps/notary-account/app"

func main() {
	app := app.NewNotaryApplication("config.local.toml")
	app.Run()
}
