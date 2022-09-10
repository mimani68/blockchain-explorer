package service

import (
	"context"

	proto_service "app.io/internal/transport/grpc/proto_service"
)

func NewStatsService(proto_service.UnimplementedStatsServiceServer) proto_service.StatsServiceServer {
	return &StatsService{}
}

type StatsService struct {
	proto_service.UnimplementedStatsServiceServer
}

func (blockSrv *StatsService) Stat(ctx context.Context, blockNumber *proto_service.StatsRequest) (*proto_service.StatsResponse, error) {
	if blockNumber.StartBlock <= 0 {
		blockNumber.StartBlock = 100000
	}
	if blockNumber.EndBlock <= 0 {
		blockNumber.StartBlock = 15507757
	}
	return &proto_service.StatsResponse{
		Success: true,
		Stats: &proto_service.Stats{
			TotalAmount:       "1000",
			TotalTransactions: "10000",
		},
	}, nil
}
