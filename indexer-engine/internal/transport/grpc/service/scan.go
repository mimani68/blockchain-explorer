package service

import (
	"context"

	"app.io/config"
	"app.io/internal/data/repository"
	networkscaning "app.io/internal/pkg/networkScaning"
	proto_service "app.io/internal/transport/grpc/proto_service"
)

var cfg config.Config
var blockRepo *repository.BlockRepository
var trxRepo *repository.TransactionRepository

func NewScanService(_cfg config.Config, _blockRepo *repository.BlockRepository, _trxRepo *repository.TransactionRepository, proto_service proto_service.UnimplementedScanServiceServer) proto_service.ScanServiceServer {
	cfg = _cfg
	blockRepo = _blockRepo
	trxRepo = _trxRepo
	return &ScanService{}
}

type ScanService struct {
	proto_service.UnimplementedScanServiceServer
}

func (blockSrv *ScanService) FreshScan(ctx context.Context, request *proto_service.ScanRequest) (*proto_service.ScanResponse, error) {
	if request.StartBlock <= 0 {
		request.StartBlock = 100
	}
	if request.EndBlock <= 0 {
		request.EndBlock = 200
	}

	blockList, err := networkscaning.SyncNetwork(cfg, nil, blockRepo, trxRepo)
	if err != nil {
		return &proto_service.ScanResponse{
			Success: false,
			Message: "Unable sync with network",
		}, nil
	}

	if len(blockList) <= 0 {
		blockList, err = blockRepo.GetBlocks(0, 100)
		if err != nil {
			return &proto_service.ScanResponse{
				Success: false,
				Message: "Unable sync with network",
			}, nil
		}
	}
	return &proto_service.ScanResponse{
		Success: true,
		Message: "forcing fresh scan",
		Block: &proto_service.Block{
			Number:  request.StartBlock,
			Hash:    "934y34ty85",
			TxCount: 14,
			Transactions: []*proto_service.Transaction{
				{
					BlockNumber: request.StartBlock,
					Hash:        "ohisi99yf8732f48gb5yby7f",
					Amount:      121550000,
					Nonce:       0,
					From:        "0xigr4y58hg8gh84h5h5",
					To:          "0xmopqj2nfh28hr4",
				},
			},
		},
	}, nil
}
