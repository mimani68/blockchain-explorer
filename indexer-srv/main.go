package main

import (
	"context"
	"time"

	"app.io/config"
	"app.io/controller"
	"app.io/pkg/exception"
	grpcclient "app.io/pkg/grpcClient"
	"app.io/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	pb "app.io/service/proto_service"
)

func main() {
	cfg := config.New()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	conn := grpcclient.StartClient(cfg.Get("INDEXER_ENGINE"))
	defer conn.Close()

	grpcClientTrxService := pb.NewTransactionServiceClient(conn)
	transactionService := service.NewTransactionService(grpcClientTrxService, ctx)
	transactionController := controller.NewTransactionController(&transactionService)

	grpcClientBlockService := pb.NewBlockServiceClient(conn)
	blockService := service.NewBlockService(grpcClientBlockService, ctx)
	blockController := controller.NewBlockController(&blockService)

	// grpcClientService := pb.NewBlockServiceClient(conn)
	// transactionService := service.NewTransactionService(grpcClientService, ctx)
	// transactionController := controller.NewTransactionController(&transactionService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	transactionController.Route(app)
	blockController.Route(app)

	errLunchingApp := app.Listen(cfg.Get("HOST") + ":" + cfg.Get("PORT"))
	exception.PanicIfNeeded(errLunchingApp)
}