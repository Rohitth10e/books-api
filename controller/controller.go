package controller

import (
	"go-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: "1", Title: "1984", Author: "George Orwell"},
	{ID: "2", Title: "The Hobbit", Author: "J.R.R. Tolkien"},
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"book": book,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Requested book not found",
	})
}

func AddBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}
	newBook.ID = strconv.Itoa(len(books) + 1)
	books = append(books, newBook)

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"book":    newBook,
	})
}
