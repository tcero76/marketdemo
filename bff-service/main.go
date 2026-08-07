package main

import (
	"os"

	"github.com/tcero76/marketplace/bff-service/oauth2"
	"github.com/tcero76/marketplace/bff-service/server"
)

func main() {
	validator := oauth2.NewHydraTokenValidator()
	e := server.StartServer(validator)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
