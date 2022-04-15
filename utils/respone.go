package utils

import "github.com/gofiber/fiber/v2"

func Result(c *fiber.Ctx, result any) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": true,
		"result":  result,
	})
}

func Error(c *fiber.Ctx, code uint, message string) error {
	return c.JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    code,
			"message": message,
		},
	})
}
