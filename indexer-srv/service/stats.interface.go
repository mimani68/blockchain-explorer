package service

import "app.io/data/model"

type StatsService interface {
	TotalStatsReport() (response model.Stats, err error)
	CustomStatsReport(startBlock, endBlock int) (response model.Stats, err error)
}
