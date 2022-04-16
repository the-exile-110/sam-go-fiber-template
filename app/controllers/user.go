package controllers

import (
	"app/service"
	"app/utils/messages"
	"app/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ListUsers(c *fiber.Ctx) error {
	res, err := service.CreateServiceFactory().ListUsers()

	if err != nil {
		return response.Fail(c, fiber.StatusBadRequest, messages.CurdGetFailMsg, err.Error())
	} else {
		return response.Success(c, messages.CurdStatusOkMsg, res)
	}
}
