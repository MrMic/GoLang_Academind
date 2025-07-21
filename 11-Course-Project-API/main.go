package main

import (
	"github.com/gin-gonic/gin"
	"michaelchlon.fr/api/db"
	"michaelchlon.fr/api/routes"
)

func main() {
	db.InitDB() // * NOTE: Initialize the database connection and create necessary tables
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8082") // * NOTE: Start the server on port 8082 (localhost:8082)
}
