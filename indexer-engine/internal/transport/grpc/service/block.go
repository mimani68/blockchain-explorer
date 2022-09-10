package service

import (
	"context"

	"app.io/internal/data/domain"
	proto_service "app.io/internal/transport/grpc/proto_service"
	"app.io/pkg/logHandler"
)

func NewBlockService(proto_service proto_service.UnimplementedBlockServiceServer) proto_service.BlockServiceServer {
	return &BlockService{}
}

type BlockService struct {
	proto_service.UnimplementedBlockServiceServer
}

func (blockSrv *BlockService) GetBlock(ctx context.Context, blockNumber *proto_service.BlockRequest) (*proto_service.BlockResponse, error) {
	var block *domain.Block
	var err error
	if blockNumber.BlockNumber <= 0 {
		block, err = blockRepo.GetLastBlock()
	} else {
		block, err = blockRepo.GetBlockByNumber(int(blockNumber.BlockNumber))
	}
	if err != nil {
		logHandler.Log(logHandler.ERROR, err.Error())
		return &proto_service.BlockResponse{
			Success: false,
			Message: "retriving block from db occured error",
		}, err
	}
	trxList, err := trxRepo.GetTransactionsByBlockNumber(int(block.Number))
	if err != nil {
		logHandler.Log(logHandler.ERROR, err.Error())
		return &proto_service.BlockResponse{
			Success: false,
			Message: "retriving block from db occured error",
		}, err
	}
	result := &proto_service.BlockResponse{
		Success: true,
		Block: &proto_service.Block{
			Number:  block.Number,
			Hash:    block.Hash,
			TxCount: int32(block.TxCount),
		},
	}
	for _, trx := range trxList {
		a := &proto_service.Transaction{
			BlockNumber: trx.BlockNumber,
			Hash:        trx.Hash,
			Amount:      trx.Amount,
			Nonce:       trx.Nonce,
			From:        trx.From,
			To:          trx.To,
		}
		result.Block.Transactions = append(result.Block.Transactions, a)
	}
	return result, nil
}
