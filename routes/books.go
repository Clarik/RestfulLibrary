package routes

import (
	"LibrarySystem/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllBooks(context *gin.Context) {
	books, err := models.GetBooks()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch books",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get Books Data",
		"books":   *books,
	})

}

func getBooks(context *gin.Context) {

	pageParam := context.Query("page")

	if pageParam == "" {
		getAllBooks(context)
		return
	}

	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid page value",
			"error":   err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(context.DefaultQuery("size", "5"))
	if err != nil || limit <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid size value",
			"error":   err.Error(),
		})
		return
	}

	order := context.DefaultQuery("order", "asc")
	if order != "asc" && order != "desc" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Order must be 'asc' for ascending or 'desc' for descending.",
		})
		return
	}

	offset := (page - 1) * limit

	books, err := models.GetBooksWithPagination(limit, offset, "title "+order)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch books",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get Books Data",
		"books":   *books,
	})
}

func createBook(context *gin.Context) {
	var book models.Book

	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if book.Title == "" || book.Author == "" || book.Description == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Title, Author, and Description can't empty.",
		})
		return
	}

	err = book.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't create the book",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Book created",
		"book":    book,
	})
}

func updateBook(context *gin.Context) {
	var updatesField map[string]any
	err := context.ShouldBindJSON(&updatesField)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body. Please provide valid JSON data.",
			"error":   err.Error(),
		})
	}

	if len(updatesField) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Request body can't be empty",
		})
		return
	}

	validFields := map[string]bool{
		"title":       true,
		"author":      true,
		"description": true,
	}

	for key := range updatesField {
		if !validFields[key] {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid field in request body: " + key,
			})
			return
		}
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid book ID, must be an positive integer",
			"error":   err.Error(),
		})
		return
	}

	book, err := models.GetBookByID(id)
	if err != nil || book == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found with specific id",
			"error":   err.Error(),
		})
		return
	}

	err = book.Update(updatesField)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update request data",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Book Updated Successfully",
		"book":    book,
	})
}

func deleteBook(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid book ID",
			"error":   err.Error(),
		})
		return
	}

	book, err := models.GetBookByID(id)
	if err != nil || book == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not fetch book specific id",
			"error":   err.Error(),
		})
		return
	}

	err = book.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not remove the book",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Book removed successfully",
	})
}
