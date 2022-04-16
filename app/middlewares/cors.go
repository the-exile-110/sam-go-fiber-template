package middlewares

import "github.com/gofiber/fiber/v2/middleware/cors"

func Cors() cors.Config {
	return cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Origin, Content-Type, Accept, Accept-Language, Content-Length, Authorization, X-Requested-With, X-CSRF-Token, X-XSRF-Token, *",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	}
}
