package main

import (
	"github.com/tcero76/marketplace/bff-service/server"
)

func main() {
	e := server.StartServer()
	log.Info("Servidor iniciado en el puerto: ", os.Getenv("PORT"))
	e.Start(":" + os.Getenv("PORT"))
}
