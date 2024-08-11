package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/go-olist-challenge/src/db"
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"github.com/joaooliveira247/go-olist-challenge/src/repositories"
	"github.com/joaooliveira247/go-olist-challenge/src/utils"
)

func CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.BindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid body"})
		return
	}

	if err := book.Prepare(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	db, err := db.GetDBConnection()

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"msg": "unexpected error"},
		)
		return
	}

	repository := repositories.NewBookRepository(db)

	if _, err := repository.InsertBook(book); err != nil {
		if errors.Is(err, utils.BookAlreadyExistsError) {
			ctx.JSON(
				http.StatusConflict, gin.H{"msg": err.Error()},
			)
			return
		}
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"msg": "unexpected error"},
		)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": book.ID})
}

func GetBooks(ctx *gin.Context) {
	db, err := db.GetDBConnection()

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"msg": "unexpected error"},
		)
		return
	}

	repository := repositories.NewBookRepository(db)

	var bookQuery utils.BookQuery
	if ctx.ShouldBindQuery(&bookQuery) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid query"})
		return
	}
	books, err := repository.GetBooks()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"msg": "unexpected error"},
		)
		return
	}
	ctx.JSON(http.StatusOK, books)
}
