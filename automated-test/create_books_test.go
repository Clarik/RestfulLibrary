package routes

import (
	"LibrarySystem/common"
	"LibrarySystem/db"
	"LibrarySystem/models"
	"LibrarySystem/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	db.InitDummyDB()
	router := gin.Default()
	routes.RegisterRoutes(router)
	return router
}

func cleanupTestDB() {
	db.CloseDummyDB()
}

func createTestBook(router *gin.Engine, title, author, description string) uint64 {
	body := map[string]string{
		"title":       title,
		"author":      author,
		"description": description,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(nil, http.StatusCreated, resp.Code)

	// Parse the response to get the created book's ID
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	book := response["book"].(map[string]interface{})
	return uint64(book["id"].(float64))
}

// Test case 1: No authorization token, no body
func TestCreateBook_NoAuthNoBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	req, _ := http.NewRequest("POST", "/books", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Token Not Valid")
}

// Test case 2: No authorization token, valid body
func TestCreateBook_NoAuthValidBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	body := map[string]string{
		"title":       "Test Book",
		"author":      "Test Author",
		"description": "Test Description",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Token Not Valid")
}

// Test case 3: No authorization token, invalid JSON body
func TestCreateBook_NoAuthInvalidJSON(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	invalidJSON := `{"title": "Test Book", "author": "Test Author", "description":}`
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Token Not Valid")
}

// Test case 4: Authorization token, no body
func TestCreateBook_AuthNoBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	req, _ := http.NewRequest("POST", "/books", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid request body")
}

// Test case 5: Authorization token, invalid body
func TestCreateBook_AuthInvalidBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	invalidBody := map[string]string{
		"title": "Incomplete Book",
	}
	jsonInvalidBody, _ := json.Marshal(invalidBody)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonInvalidBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Title, Author, and Description can't empty.")
}

// Test case 6: Authorization token, valid body
func TestCreateBook_AuthValidBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	body := map[string]string{
		"title":       "Test Book",
		"author":      "Test Author",
		"description": "Test Description",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Book created")

	// Verify the book was created in the database
	var createdBook models.Book
	common.DB.First(&createdBook, "title = ?", "Test Book")
	assert.Equal(t, "Test Book", createdBook.Title)
	assert.Equal(t, "Test Author", createdBook.Author)
	assert.Equal(t, "Test Description", createdBook.Description)
}
