package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"app.io/config"
	"app.io/internal/data/repository"
	proto_service "app.io/internal/transport/grpc/proto_service"
	"app.io/internal/transport/grpc/service"
	"app.io/pkg/logHandler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func GrpcServerBuilder(listener net.Listener, cfg config.Config, status chan string, blockRepo *repository.BlockRepository, transactionRepo *repository.TransactionRepository) {
	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)

	scanService := service.NewScanService(cfg, blockRepo, transactionRepo, proto_service.UnimplementedScanServiceServer{})
	proto_service.RegisterScanServiceServer(grpcServer, scanService)
	blockService := service.NewBlockService(proto_service.UnimplementedBlockServiceServer{})
	proto_service.RegisterBlockServiceServer(grpcServer, blockService)
	statService := service.NewStatsService(proto_service.UnimplementedStatsServiceServer{})
	proto_service.RegisterStatsServiceServer(grpcServer, statService)
	trxService := service.NewTransactionService(proto_service.UnimplementedTransactionServiceServer{})
	proto_service.RegisterTransactionServiceServer(grpcServer, trxService)
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logHandler.Log(logHandler.INFO, err.Error(), "GRPC")
		}
	}()
	logHandler.Log(logHandler.INFO, fmt.Sprintf("server listening at %v", listener.Addr()))
	status <- "SERVER_IS_AVAIALBE"

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	logHandler.Log(logHandler.INFO, "Stopping the server...")
	grpcServer.Stop()
	listener.Close()
	logHandler.Log(logHandler.INFO, "Closing database conncetion")
	logHandler.Log(logHandler.INFO, "Shutdown application")
}
