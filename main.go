package main

import (
	"border_patrol/api/app"
	"border_patrol/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
