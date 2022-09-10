package networkscaning

import (
	"testing"

	"app.io/config"
	"app.io/internal/data/db"
	"app.io/internal/data/repository"
)

func TestSyncBlocks(t *testing.T) {

	cfg := config.Config{}
	cfg.Database.DbUri = "tcp://localhost:9001?database=app"
	cfg.Database.Type = "clickhouse"
	cfg.Server.TatumApiToken = "b211febc-6e35-4430-959d-9ae103799f4b"
	cfg.Server.NetworkTitle = "ethereum"
	cfg.Server.NumberOfBlockForCapturing = 2

	dbInstance := db.NewDatabase(cfg)

	blockRepo := repository.CreateBlockRepository(cfg, *dbInstance.Db)
	transactionRepo := repository.CreateTransactionRepository(cfg, *dbInstance.Db)

	blockList, err := SyncBlocks(0, 0, cfg, blockRepo, transactionRepo)

	if len(blockList) <= 0 || err != nil {
		t.Errorf("Having issue here")
	}
}
