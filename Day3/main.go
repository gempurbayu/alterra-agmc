package main

import (
	"day2-agmc/config"
	m "day2-agmc/middlewares"
	"day2-agmc/routes"
)

func main() {

	config.InitDB()
	config.InitMigrate()

	e := routes.New()

	m.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":8080"))
}
