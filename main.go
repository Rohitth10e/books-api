package main

import (
	"fmt"
	"go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(func(c *gin.Context) {
		fmt.Println("morgan: ", c.Request.Method, c.Request.URL)
		c.Next()
	})
	routes.Routes(server)
	server.Run(":8080")
}
