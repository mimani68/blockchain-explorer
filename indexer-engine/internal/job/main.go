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
	blocState := make(chan int)
	s.Every(cfg.Server.CronJobInterval).Seconds().Do(networkscaning.SyncNetwork, cfg, blocState, blockRepo, transactionRepo)
	s.StartAsync()
}
