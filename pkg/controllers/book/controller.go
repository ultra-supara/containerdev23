package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ultra-supara/containerdev23/pkg/usecases"
)

type BookController struct {
	listBooks *usecases.ListBooks
}

func NewBookController(listBooks *usecases.ListBooks) *BookController {
	return &BookController{listBooks: listBooks}
}

func (c *BookController) List(ctx *gin.Context) {
	books, err := c.listBooks.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    books,
	})
}
