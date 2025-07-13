package main

import (
	"github.com/joho/godotenv"
	"github.com/withzeus/mugi-identity/cmd/app"
)

func main() {
	godotenv.Load()

	server := app.NewApp()
	server.Run()
}
