package service

import (
	"context"

	"app.io/internal/data/repository"
	proto_service "app.io/internal/transport/grpc/proto_service"
	"app.io/pkg/logHandler"
)

var repo *repository.BlockRepository

func NewBlockService(blockRepo *repository.BlockRepository, proto_service proto_service.UnimplementedBlockServiceServer) proto_service.BlockServiceServer {
	repo = blockRepo
	return &BlockService{}
}

type BlockService struct {
	proto_service.UnimplementedBlockServiceServer
}

func (blockSrv *BlockService) GetBlock(ctx context.Context, blockNumber *proto_service.BlockRequest) (*proto_service.BlockResponse, error) {
	if blockNumber.BlockNumber <= 0 {
		blockNumber.BlockNumber = 100
	}
	blockList, err := repo.GetBlocks(0, 10)
	if err != nil {
		logHandler.Log(logHandler.ERROR, err.Error())
		return &proto_service.BlockResponse{
			Success: false,
			Message: "retriving block from db occured with error",
		}, err
	}
	if len(blockList) < 1 {
		logHandler.Log(logHandler.ERROR, "list of blocks is empty")
		return &proto_service.BlockResponse{
			Success: false,
			Message: "list of blocks is empty",
		}, err
	}
	result := &proto_service.BlockResponse{
		Success: true,
		Block: &proto_service.Block{
			Number:  blockList[0].Number,
			Hash:    blockList[0].Hash,
			TxCount: int32(blockList[0].TxCount),
			// 	Transactions: []*proto_service.Transaction{
			// 		{
			// 			BlockNumber: blockNumber.BlockNumber,
			// 			Hash:        "ohisi99yf8732f48gb5yby7f",
			// 			Amount:      121550000,
			// 			Nonce:       0,
			// 			From:        "0xigr4y58hg8gh84h5h5",
			// 			To:          "0xmopqj2nfh28hr4",
			// 		},
		},
	}
	// for _, item := range transactionList {
	// 	result.Block.Transactions = append(result.Block.Transactions, proto_service.Transaction{
	// 		BlockNumber: item.BlockNumber,
	// 		Hash:        item.Hash,
	// 		Amount:      0,
	// 		Nonce:       item.,
	// 		From:        item.,
	// 		To:          item.,
	// 	})
	// }
	return result, nil
}
