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
)

func grpcServerBuilder(listener net.Listener, cfg config.Config, blockRepo *repository.BlockRepository, transactionRepo *repository.TransactionRepository) {
	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)

	blockService := service.NewBlockService(blockRepo, proto_service.UnimplementedBlockServiceServer{})
	proto_service.RegisterBlockServiceServer(grpcServer, blockService)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			msg := fmt.Sprintf("failed to serve %v", err)
			logHandler.Log(logHandler.INFO, msg, "GRPC")
		}
	}()
	logHandler.Log(logHandler.INFO, fmt.Sprintf("server listening at %v", listener.Addr()))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	logHandler.Log(logHandler.INFO, "Stopping the server...")
	grpcServer.Stop()
	listener.Close()
	logHandler.Log(logHandler.INFO, "Closing database conncetion")
	logHandler.Log(logHandler.INFO, "Shutdown application")
}
