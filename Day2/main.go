package main

import (
	"day2-agmc/routes"
)

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
