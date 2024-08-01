package routes

import "github.com/gin-gonic/gin"

func RoutesRegistry(eng *gin.Engine) {
	AuthorsRoutes(eng)
	BooksRoutes(eng)
}