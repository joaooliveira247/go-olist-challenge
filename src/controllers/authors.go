package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/go-olist-challenge/src/db"
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"github.com/joaooliveira247/go-olist-challenge/src/repositories"
)

func CreateAuthor(ctx *gin.Context) {
	var author models.Authors
	if err := ctx.BindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid body"})
		return
	}
	author.GenUUID()
	db, err := db.GetDBConnection()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "unexpected error"})
		return
	}

	repository := repositories.NewAuthorRepository(db)

	if _, err := repository.InsertAuthor(author); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "unexpected error"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": author.ID})
}
