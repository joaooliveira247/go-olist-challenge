package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	
	if err := author.Prepare(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

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

func GetAuthors(ctx *gin.Context) {
	db, err := db.GetDBConnection()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "unexpected error"})
		return
	}

	repository := repositories.NewAuthorRepository(db)

	authors, err := repository.GetAuthors()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func SearchAuthorByName(ctx *gin.Context) {
	name := ctx.Param("name")

	db, err := db.GetDBConnection()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "unexpected error"})
		return
	}

	repository := repositories.NewAuthorRepository(db)

	authors, err := repository.GetAuthorsByName(name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "error when try search author"})
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func DeleteAuthor(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid uuid"})
		return
	}

	db, err := db.GetDBConnection()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "unexpected error"})
		return
	}

	repository := repositories.NewAuthorRepository(db)

	if err = repository.DeleteAuthor(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "error when try delete author"})
	}
	ctx.JSON(http.StatusNoContent, nil)
}
