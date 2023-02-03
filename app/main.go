package main

import (
	"github.com/SetGecko/study-microservice/tree/Homework_8/app/app"
	"github.com/SetGecko/study-microservice/tree/Homework_8/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
