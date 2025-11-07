package routes

import (
	"go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.GET("/get-books", controller.GetBooks)
	server.GET("/get-books/:id", controller.GetBookByID)
	server.POST("/add-book", controller.AddBook)
	server.POST("/update-book/:id", controller.UpdateBook)
	server.POST("/delete-book/:id", controller.DeleteBook)
}
