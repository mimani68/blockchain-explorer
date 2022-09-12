package main

import (
	"context"

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

	ctx := context.Background()

	conn := grpcclient.StartClient(cfg.Get("INDEXER_ENGINE"))
	defer conn.Close()

	grpcClientTrxService := pb.NewTransactionServiceClient(conn)
	transactionService := service.NewTransactionService(grpcClientTrxService, ctx)
	transactionController := controller.NewTransactionController(&transactionService)

	grpcClientBlockService := pb.NewBlockServiceClient(conn)
	grpcScanBlockService := pb.NewScanServiceClient(conn)
	blockService := service.NewBlockService(grpcClientBlockService, grpcScanBlockService, ctx)
	blockController := controller.NewBlockController(&blockService)

	grpcClientStatsService := pb.NewStatsServiceClient(conn)
	statsService := service.NewStatsService(grpcClientStatsService, ctx)
	statsController := controller.NewStatsController(&statsService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	transactionController.Route(app)
	blockController.Route(app)
	statsController.Route(app)

	errLunchingApp := app.Listen(cfg.Get("HOST") + ":" + cfg.Get("PORT"))
	exception.PanicIfNeeded(errLunchingApp)
}
