package service

import (
	"context"

	"app.io/data/model"
	pb "app.io/service/proto_service"
)

func NewStatsService(statsGrpcClient pb.StatsServiceClient, ctx context.Context) StatsService {
	return &statsServiceImpl{
		ctx:             ctx,
		statsGrpcClient: statsGrpcClient,
	}
}

type statsServiceImpl struct {
	ctx             context.Context
	statsGrpcClient pb.StatsServiceClient
}

func (service *statsServiceImpl) TotalStatsReport() (response model.Stats, err error) {
	stats, err := service.statsGrpcClient.Stat(service.ctx, &pb.StatsRequest{})
	if err != nil {
		return
	}
	response = model.Stats{
		TotalAmount:       stats.Stats.TotalAmount,
		TotalTransactions: stats.Stats.TotalTransactions,
	}
	return
}

func (service *statsServiceImpl) CustomStatsReport(startBlock, endBlock int) (response model.Stats, err error) {
	stats, err := service.statsGrpcClient.Stat(service.ctx, &pb.StatsRequest{
		StartBlock: int64(startBlock),
		EndBlock:   int64(endBlock),
	})
	if err != nil {
		return
	}
	response = model.Stats{
		TotalAmount:       stats.Stats.TotalAmount,
		TotalTransactions: stats.Stats.TotalTransactions,
	}
	return
}
