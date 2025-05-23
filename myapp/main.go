package main

import (
	"log"
	"myapp/config"
	"myapp/routes"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDatabase()
	r := routes.SetupRouter(db)
	log.Fatal(r.Run(":8080"))
}
