package controllers

import (
	"app/utils/response"
	"github.com/gofiber/fiber/v2"
)

func StatusCheck(c *fiber.Ctx) error {
	return response.Success(c, "OK", "")
}
