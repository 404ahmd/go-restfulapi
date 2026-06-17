package main

import (
	"go-restfulapi/config"
	"go-restfulapi/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRouter()

	router.Run(":8080")
}
