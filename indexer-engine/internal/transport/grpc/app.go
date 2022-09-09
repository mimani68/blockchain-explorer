package grpc

import (
	"app.io/config"
	"app.io/internal/data/db"
	"app.io/internal/job"
)

func RunServer() {

	cfg := config.NewConfig()

	// listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.ServerHost, cfg.Server.ServerPort))

	// if err != nil {
	// 	logHandler.Log(logHandler.INFO, logHandler.ERROR, err.Error())
	// 	os.Exit(1)
	// } else {
	// 	logHandler.Log(logHandler.INFO, fmt.Sprintf("Application started on \"%s:%s\"", cfg.Server.ServerHost, cfg.Server.ServerPort))
	// 	logHandler.Log(logHandler.INFO, "Envirnoment: "+cfg.Server.Env)
	// }

	dbInstance := db.NewDatabase(*cfg)
	dbInstance.Db.AutoMigrate()

	job.SyncNetwork(*cfg)

	//
	// Capture network data in background
	//
	go job.BackgroundProcesses(*cfg, dbInstance)

	//
	// Serve data using gRPC format to other internal systems
	//
	// grpcServerBuilder(listener, *cfg, dbInstance)

}
