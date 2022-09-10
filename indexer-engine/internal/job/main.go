package job

import (
	"time"

	"app.io/config"
	"app.io/internal/data/repository"
	networkscaning "app.io/internal/pkg/networkScaning"
	"github.com/go-co-op/gocron"
)

func BackgroundProcesses(cfg config.Config, blockRepo *repository.BlockRepository, transactionRepo *repository.TransactionRepository) {
	s := gocron.NewScheduler(time.Local)
	s.Every(cfg.Server.CronJobInterval).Seconds().Do(networkscaning.SyncBlocks, 0, 0, cfg, blockRepo, transactionRepo)
	s.StartAsync()
}
