package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/utqinadhif/fibertest/app/model"
	"github.com/utqinadhif/fibertest/config"
)

func Index(c *fiber.Ctx) error {
	var user []model.User

	err := config.DB.Find(&user).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"suceess": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"suceess": true,
		"message": "Successfully get all user",
		"data":    user,
	})
}
