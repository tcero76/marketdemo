package main

import (
	"os"

	"github.com/tcero76/marketplace/bff-service/server"
)

func main() {
	e, err := server.StartServer()
	if err != nil {
		panic(err)
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
