package grpc

import (
	"fmt"
	"net"
	"os"

	"app.io/config"
	"app.io/internal/data/db"
	"app.io/internal/data/repository"
	"app.io/internal/job"
	"app.io/pkg/logHandler"
)

func RunServer() {

	cfg := config.NewConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))

	if err != nil {
		logHandler.Log(logHandler.INFO, logHandler.ERROR, err.Error())
		os.Exit(1)
	} else {
		logHandler.Log(logHandler.INFO, fmt.Sprintf("Application starting"))
		logHandler.Log(logHandler.INFO, "Envirnoment: "+cfg.Server.Env)
	}

	dbInstance := db.NewDatabase(*cfg)
	dbInstance.AutoMigrate()

	blockRepo := repository.CreateBlockRepository(*cfg, *dbInstance.Db)
	transactionRepo := repository.CreateTransactionRepository(*cfg, *dbInstance.Db)

	//
	// Capture network data in background
	//
	go job.BackgroundProcesses(*cfg, blockRepo, transactionRepo)

	//
	// Serve data using gRPC format to other internal systems
	//
	GrpcServerBuilder(listener, *cfg, nil, blockRepo, transactionRepo)

}
