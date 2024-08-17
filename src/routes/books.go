package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/go-olist-challenge/src/controllers"
)

func BooksRoutes(eng *gin.Engine) {
	BooksRouter := eng.Group("/books")
	{
		BooksRouter.POST("/", controllers.CreateBook)
		BooksRouter.GET("/", controllers.GetBooks)
		BooksRouter.DELETE("/:id", controllers.DeleteBook)
	}
}
