package middlewares

import "github.com/gofiber/fiber/v2/middleware/logger"

func LoggerConfig() logger.Config {
	return logger.Config{
		Format:     "${time} ${status} - ${method} ${path} ${query:}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Tokyo",
	}
}
