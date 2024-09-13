package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Movie struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func initDB() {
	var err error
	// Updated DSN to use PostgreSQL service name from Docker Compose
	dsn := "user=admin password=password dbname=movie_app host=postgres port=5432 sslmode=disable"


	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Movie{})
}

func main() {
	app := fiber.New()

	// Enable CORS for all origins
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowMethods: "GET,POST,PUT,DELETE", // Allow specified methods
		AllowHeaders: "Origin, Content-Type, Accept", // Allow specified headers
	}))

	initDB()

	app.Post("/movies", createMovie)
	app.Get("/movies", getMovies)
	app.Get("/movies/:id", getMovie)
	app.Put("/movies/:id", updateMovie)
	app.Delete("/movies/:id", deleteMovie)

	app.Listen(":3000")
}

func createMovie(c *fiber.Ctx) error {
	var movie Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	result := db.Create(&movie)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(movie)
}

func getMovies(c *fiber.Ctx) error {
	var movies []Movie
	result := db.Find(&movies)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.JSON(movies)
}

func getMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie Movie
	result := db.First(&movie, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}
	return c.JSON(movie)
}

func updateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var existingMovie Movie
	result := db.First(&existingMovie, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}
	db.Model(&existingMovie).Updates(movie)
	return c.JSON(existingMovie)
}

func deleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	result := db.Delete(&Movie{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
