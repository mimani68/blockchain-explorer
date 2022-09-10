package job

import (
	"testing"

	"app.io/config"
	"app.io/internal/data/db"
	"app.io/internal/data/repository"
)

func TestSyncNetwork(t *testing.T) {

	cfg := config.Config{}
	cfg.Database.DbUri = "tcp://localhost:9001?database=app"
	cfg.Database.Type = "clickhouse"
	cfg.Server.TatumApiToken = "b211febc-6e35-4430-959d-9ae103799f4b"
	cfg.Server.NetworkTitle = "ethereum"
	cfg.Server.NumberOfBlockForCapturing = 2

	blocStateChannel := make(chan int)
	dbInstance := db.NewDatabase(cfg)

	blockRepo := repository.CreateBlockRepository(cfg, *dbInstance.Db)
	transactionRepo := repository.CreateTransactionRepository(cfg, *dbInstance.Db)

	go SyncNetwork(cfg, blocStateChannel, blockRepo, transactionRepo)
	lastBlock := <-blocStateChannel
	t.Logf("Last block value was %d", lastBlock)

	if lastBlock < 1 {
		t.Errorf("Having issue here")
	}
}
