package grpc

import (
	"net"

	"app.io/config"
	"app.io/internal/data/db"
)

func grpcServerBuilder(listener net.Listener, cfg config.Config, db db.Database) {
	// serverOptions := []grpc.ServerOption{}
	// grpcServer := grpc.NewServer(serverOptions...)

	// userRepo := repository.CreateUserRepository(cfg, *db.Db)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	// grpcUserService.RegisterUserServiceServer(grpcServer, userHandler)
	// // pb.RegisterUserServiceServer(grpcServer)

	// go func() {
	// 	if err := grpcServer.Serve(listener); err != nil {
	// 		msg := fmt.Sprintf("failed to serve %v", err)
	// 		logHandler.Log(logHandler.INFO, msg, "GRPC")
	// 	}
	// }()
	// logHandler.Log(logHandler.INFO, fmt.Sprintf("server listening at %v", listener.Addr()))

	// c := make(chan os.Signal)
	// signal.Notify(c, os.Interrupt)
	// <-c
	// fmt.Println("\nStopping the server... ")
	// grpcServer.Stop()
	// listener.Close()
	// fmt.Println("Closing database conncetion")
	// fmt.Println("Done")
}
