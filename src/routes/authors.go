package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/go-olist-challenge/src/controllers"
)

func AuthorsRoutes(eng *gin.Engine) {
	authorRouter := eng.Group("/authors")
	{
		authorRouter.POST("/", controllers.CreateAuthor)
		authorRouter.GET("/", controllers.GetAuthors)
		authorRouter.GET("/:name", controllers.SearchAuthorByName)
		authorRouter.DELETE("/:id", controllers.DeleteAuthor)
	}
}
