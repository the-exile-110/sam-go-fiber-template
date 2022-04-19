package main

import (
	"app/middlewares"
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"

	"app/routers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadaptor "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	app = fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: middlewares.ErrorHandler,
	})
	fiberLambda *fiberadaptor.FiberLambda
	ENV         = os.Getenv("ENV")
)

func init() {
	routers.RouterMapping(app)
	fiberLambda = fiberadaptor.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	zap.L().Info("Starting Lambda ...")

	app.Use(logger.New(middlewares.LoggerConfig()), cors.New(middlewares.Cors()), recover.New())

	if ENV == "" {
		err := app.Listen(":3000")
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
	lambda.Start(Handler)
}
