package main

import (
	"go-rest-api/database"
	"go-rest-api/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
