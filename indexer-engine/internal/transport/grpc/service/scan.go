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
		request.StartBlock = 0
	}
	if request.EndBlock <= 0 {
		request.EndBlock = 0
	}

	blockList, err := networkscaning.SyncBlocks(int(request.StartBlock), int(request.EndBlock), cfg, blockRepo, trxRepo)
	if err != nil {
		return &proto_service.ScanResponse{
			Success: false,
			Message: "Unable sync with network",
		}, nil
	}

	result := &proto_service.ScanResponse{
		Success: true,
		Message: "forcing fresh scan",
	}
	for _, block := range blockList {
		tmp := &proto_service.Block{
			Number:  block.Number,
			Hash:    block.Hash,
			TxCount: int32(block.TxCount),
		}
		for _, trx := range block.Transaction {
			a := &proto_service.Transaction{
				BlockNumber: trx.BlockNumber,
				Hash:        trx.Hash,
				Amount:      trx.Amount,
				Nonce:       trx.Nonce,
				From:        trx.From,
				To:          trx.To,
			}
			tmp.Transactions = append(tmp.Transactions, a)
		}
		result.Block = append(result.Block, tmp)
	}
	return result, nil
}
