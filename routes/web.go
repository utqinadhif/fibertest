package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/utqinadhif/fibertest/app/controllers"
)

func Router(c *fiber.App) {
	c.Get("/", controllers.Index)

	c.Static("/", "./public")
}
