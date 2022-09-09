package job

import (
	"time"

	"app.io/config"
	"app.io/internal/data/db"
	"github.com/go-co-op/gocron"
)

func BackgroundProcesses(cfg config.Config, db db.Database) {
	s := gocron.NewScheduler(time.Local)
	// FIXME
	// s.Every(cfg.Server.Interval).Seconds().Do(SystemCheck)
	s.Every(1000).Seconds().Do(SystemCheck)
	// s.Every(60).Seconds().Do(syncNetwork)
	s.StartAsync()
}
