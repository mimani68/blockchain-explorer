package service

import (
	"context"
	"fmt"

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
		blockNumber.StartBlock = 0
	}
	if blockNumber.EndBlock <= 0 {
		blockNumber.EndBlock = 2000000000000000000
	}
	transactionSum, _ := trxRepo.GetTransactionsSum(int(blockNumber.StartBlock), int(blockNumber.EndBlock))
	transactionNumber, _ := trxRepo.GetTransactionsNumber(int(blockNumber.StartBlock), int(blockNumber.EndBlock))

	return &proto_service.StatsResponse{
		Success: true,
		Stats: &proto_service.Stats{
			TotalAmount:       fmt.Sprintf("%d", transactionNumber),
			TotalTransactions: fmt.Sprintf("%d", transactionSum),
		},
	}, nil
}
