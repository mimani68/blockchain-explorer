package grpc

import (
	"fmt"
	"net"
	"testing"

	"app.io/config"
	"app.io/internal/data/db"
	"app.io/internal/data/repository"
)

func TestGrpcServerBuilder(t *testing.T) {

	cfg := config.Config{}
	cfg.Database.DbUri = "tcp://localhost:9001?database=app"
	cfg.Database.Type = "clickhouse"
	cfg.Server.Host = "0.0.0.0"
	cfg.Server.Port = "3001"
	cfg.Server.TatumApiToken = "b211febc-6e35-4430-959d-9ae103799f4b"
	cfg.Server.NetworkTitle = "ethereum"
	cfg.Server.NumberOfBlockForCapturing = 2

	status := make(chan string)

	listener, _ := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))

	dbInstance := db.NewDatabase(cfg)

	blockRepo := repository.CreateBlockRepository(cfg, *dbInstance.Db)
	transactionRepo := repository.CreateTransactionRepository(cfg, *dbInstance.Db)

	go GrpcServerBuilder(listener, cfg, status, blockRepo, transactionRepo)

	if "SERVER_IS_AVAIALBE" != <-status {
		t.Errorf("Having issue here")
	}
}
