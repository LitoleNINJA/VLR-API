package main

import (
	server "VLR-API/server"
	"log"
)

func main() {

	app := server.NewApp()

	if err := app.Run("8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
