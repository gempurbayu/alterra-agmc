package main

import (
	"day2-agmc/config"
	"day2-agmc/routes"
)

func main() {

	config.InitDB()
	config.InitMigrate()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
