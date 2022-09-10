package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"app.io/config"
	"app.io/internal/data/db"
	"app.io/internal/data/repository"
	proto_service "app.io/internal/transport/grpc/proto_service"
	"app.io/internal/transport/grpc/service"
	"app.io/pkg/logHandler"
	"google.golang.org/grpc"
)

func grpcServerBuilder(listener net.Listener, cfg config.Config, db db.Database) {
	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)

	blockRepo := repository.CreateBlockRepository(cfg, *db.Db)
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
	fmt.Println("\nStopping the server... ")
	grpcServer.Stop()
	listener.Close()
	fmt.Println("Closing database conncetion")
	fmt.Println("Done")
}
