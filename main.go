package main

import (
	"LibrarySystem/db"
	"LibrarySystem/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	ginServer := gin.Default()

	routes.RegisterRoutes(ginServer)

	ginServer.Run(":8080") // localhost:8080
}
