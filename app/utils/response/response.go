package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Success[T any](c *fiber.Ctx, msg string, data T) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"msg":        msg,
		"data":       data,
	})
}

// Fail Return fail
func Fail(c *fiber.Ctx, statusCode int, msg string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"statusCode": statusCode,
		"msg":        msg,
		"data":       "",
	})
}
