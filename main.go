package main

import (
	"books-api-gin-gorm/database"
	"books-api-gin-gorm/routes"
)

func main() {
	database.StartDB()

	router := routes.SetupRoutes()

	router.Run(":8080")
}
