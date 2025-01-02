package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/utqinadhif/fibertest/config"
	"github.com/utqinadhif/fibertest/routes"
)

func main() {
	config.ConnectDB()

	engine := html.New("./resources/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Router(app)

	log.Fatal(app.Listen(":3000"))
}
