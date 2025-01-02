package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/utqinadhif/fibertest/routes"
)

func main() {
	app := fiber.New()

	routes.Router(app)

	log.Fatal(app.Listen(":3000"))
}
