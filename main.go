package main

import (
	"HatPanel/config"
	"HatPanel/fyne"
)

func main() {
	projectConfig := config.ReadConfig()

	app := fyne.NewApp(projectConfig)
	app.Launch()
}
