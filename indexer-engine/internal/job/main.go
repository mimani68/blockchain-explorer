package job

import (
	"time"

	"app.io/config"
	"app.io/internal/data/db"
	"github.com/go-co-op/gocron"
)

func BackgroundProcesses(cfg config.Config, db db.Database) {
	s := gocron.NewScheduler(time.Local)
	s.Every(cfg.Server.CronJobInterval).Seconds().Do(SyncNetwork, cfg, db)
	s.StartAsync()
}
