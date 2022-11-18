package main

import (
	"log"
	"os"

	"shorturl/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.Resolve)
	app.Post("/submit", routes.Shorten)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment file.")
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	setupRoutes(app)
	app.Listen(os.Getenv("APP_PORT"))

}
