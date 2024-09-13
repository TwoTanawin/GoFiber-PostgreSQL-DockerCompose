// tests/main_test.go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateMovie(t *testing.T) {
    app := setupApp() // This should initialize your Fiber app

    // Define the movie data
    movie := map[string]interface{}{
        "title": "Inception",
        "year":  2010,
    }
    movieJSON, _ := json.Marshal(movie)

    // Create a request
    req, err := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(movieJSON))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a response recorder
    rr := httptest.NewRecorder()

    // Serve the request
    app.ServeHTTP(rr, req)

    // Check the status code
    assert.Equal(t, http.StatusCreated, rr.Code)

    // Check the response body
    var response map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &response)
    assert.Equal(t, "Inception", response["title"])
    assert.Equal(t, float64(2010), response["year"])
}

func TestGetMovies(t *testing.T) {
    app := setupApp() // This should initialize your Fiber app

    // Create a request
    req, err := http.NewRequest(http.MethodGet, "/movies", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    // Create a response recorder
    rr := httptest.NewRecorder()

    // Serve the request
    app.ServeHTTP(rr, req)

    // Check the status code
    assert.Equal(t, http.StatusOK, rr.Code)

    // Check the response body
    var response []map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &response)
    assert.NotEmpty(t, response)
}
