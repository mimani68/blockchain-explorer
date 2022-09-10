package main

import (
	"app.io/config"
	"app.io/controller"
	"app.io/pkg/exception"
	"app.io/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.New()

	// FIXME: adding grpc client
	// grpcClient := as()

	transactionService := service.NewTransactionService(nil)

	transactionController := controller.NewTransactionController(&transactionService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	transactionController.Route(app)

	err := app.Listen(cfg.Get("Host") + ":" + cfg.Get("Port"))
	exception.PanicIfNeeded(err)
}
