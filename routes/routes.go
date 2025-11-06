package routes

import (
	"go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.GET("/get-books", controller.GetBooks)
	server.GET("/get-books/:id", controller.GetBookByID)
	server.POST("/add-book", controller.AddBook)
}
