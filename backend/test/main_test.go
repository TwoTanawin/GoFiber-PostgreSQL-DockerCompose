// backend/test/main_test.go
package test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var app *fiber.App
var db *gorm.DB

type Movie struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Title string `json:"title"`
    Year  int    `json:"year"`
}

func initDB() {
    var err error
    dsn := "user=admin password=password dbname=movie_app host=localhost port=5432 sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }
    db.AutoMigrate(&Movie{})
}

func setupApp() *fiber.App {
    app := fiber.New()
    initDB()
    app.Post("/movies", createMovie)
    app.Get("/movies", getMovies)
    app.Get("/movies/:id", getMovie)
    app.Put("/movies/:id", updateMovie)
    app.Delete("/movies/:id", deleteMovie)
    return app
}

func TestCreateMovie(t *testing.T) {
    app = setupApp()

    movie := map[string]interface{}{
        "title": "Inception",
        "year":  2010,
    }
    movieJSON, _ := json.Marshal(movie)

    req, err := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(movieJSON))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    app.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)

    var response map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &response)
    assert.Equal(t, "Inception", response["title"])
    assert.Equal(t, float64(2010), response["year"])
}

func TestGetMovies(t *testing.T) {
    app = setupApp()

    req, err := http.NewRequest(http.MethodGet, "/movies", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    rr := httptest.NewRecorder()
    app.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var response []map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &response)
    assert.NotEmpty(t, response)
}
