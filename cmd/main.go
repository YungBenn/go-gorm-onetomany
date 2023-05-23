package main

import (
	"log"
	"os"

	"github.com/YungBenn/go-gorm-fiber/config"
	"github.com/YungBenn/go-gorm-fiber/internal/routes"
	"github.com/YungBenn/go-gorm-fiber/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func buildServer() error {
	err := config.LoadENV()
	if err != nil {
		return err
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db := storage.Connect(config)

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	routes.SetupUserRoutes(app, db)
	routes.SetupItemRoutes(app, db)

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK!")
	})

	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}

func main() {
	err := buildServer()
	if err != nil {
		log.Fatal(err)
	}
}
