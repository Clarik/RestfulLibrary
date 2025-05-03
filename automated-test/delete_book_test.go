package routes

import (
	"LibrarySystem/common"
	"LibrarySystem/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBook(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	// Step 1: Create a test book
	bookID := createTestBook(router, "Book to Delete", "Author", "Description")

	// Step 2: Delete the book
	req, _ := http.NewRequest("DELETE", "/books/"+strconv.FormatUint(bookID, 10), nil)
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Step 3: Assert the response
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Book removed successfully")

	// Step 4: Verify the book is deleted from the database
	var count int64
	common.DB.Model(&models.Book{}).Where("id = ?", bookID).Count(&count)
	assert.Equal(t, int64(0), count, "The book should be deleted from the database")
}
