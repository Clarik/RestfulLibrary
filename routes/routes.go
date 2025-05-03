package routes

import (
	"LibrarySystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/books", getBooks)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/books", createBook)
	authenticated.PUT("/books/:id", updateBook)
	authenticated.DELETE("/books/:id", deleteBook)
}
