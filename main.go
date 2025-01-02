package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/utqinadhif/fibertest/config"
	"github.com/utqinadhif/fibertest/routes"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	engine := html.New("./resources/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Router(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
