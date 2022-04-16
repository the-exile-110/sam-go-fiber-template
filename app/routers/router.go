package routers

import (
	"app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterMapping(app *fiber.App) {
	app.Get("/", controllers.StatusCheck)
	user := app.Group("/user")
	{
		user.Get("/list-users", controllers.ListUsers)
	}
}
