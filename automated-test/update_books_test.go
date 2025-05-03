package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test case 1: No authorization token, no body
func TestUpdateBook_NoAuthNoBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	bookID := createTestBook(router, "Original Title", "Original Author", "Original Description")

	req, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(bookID, 10), nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Token Not Valid")
}

// Test case 2: No authorization token, valid body
func TestUpdateBook_NoAuthValidBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	bookID := createTestBook(router, "Original Title", "Original Author", "Original Description")

	body := map[string]string{
		"title":       "Updated Title",
		"author":      "Updated Author",
		"description": "Updated Description",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(bookID, 10), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Token Not Valid")
}

// Test case 3: No authorization token, invalid JSON body
func TestUpdateBook_NoAuthInvalidJSON(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	bookID := createTestBook(router, "Original Title", "Original Author", "Original Description")

	invalidJSON := `{"title": "Updated Title", "author": "Updated Author", "description":}`
	req, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(bookID, 10), bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Token Not Valid")
}

// Test case 4: Authorization token, no body
func TestUpdateBook_AuthNoBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	bookID := createTestBook(router, "Original Title", "Original Author", "Original Description")

	req, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(bookID, 10), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid request body")
}

// Test case 5: Authorization token, invalid body
func TestUpdateBook_AuthInvalidBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	bookID := createTestBook(router, "Original Title", "Original Author", "Original Description")

	invalidBody := map[string]string{
		"summary": "Incomplete Book",
	}
	jsonInvalidBody, _ := json.Marshal(invalidBody)

	req, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(bookID, 10), bytes.NewBuffer(jsonInvalidBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid field in request body")
}

// Test case 6: Authorization token, valid body
func TestUpdateBook_AuthValidBody(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	bookID := createTestBook(router, "Original Title", "Original Author", "Original Description")

	body := map[string]string{
		"title":       "Updated Title",
		"author":      "Updated Author",
		"description": "Updated Description",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(bookID, 10), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Book Updated Successfully")
}

// Test case 7: Authorization token, valid body, but title already exists
func TestUpdateBook_AuthDuplicateTitle(t *testing.T) {
	router := setupTestRouter()
	defer cleanupTestDB()

	_ = createTestBook(router, "Existing Title", "Author 1", "Description 1")
	book2ID := createTestBook(router, "Another Title", "Author 2", "Description 2")

	body := map[string]string{
		"title":       "Existing Title", // Duplicate title
		"author":      "Updated Author",
		"description": "Updated Description",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(book2ID, 10), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "static-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "UNIQUE constraint failed")
}
