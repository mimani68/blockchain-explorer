package service

import (
	"context"

	"app.io/data/entity"
	pb "app.io/service/proto_service"
)

func NewBlockService(blockGrpcClient pb.BlockServiceClient, ctx context.Context) BlockService {
	return &blockServiceImpl{
		ctx:             ctx,
		blockGrpcClient: blockGrpcClient,
	}
}

type blockServiceImpl struct {
	ctx             context.Context
	blockGrpcClient pb.BlockServiceClient
}

func (service *blockServiceImpl) LastBlock() (response entity.Block, err error) {
	blk, err := service.blockGrpcClient.GetBlock(service.ctx, &pb.BlockRequest{})
	if err != nil {
		return
	}
	response = entity.Block{
		Number:  blk.Block.Number,
		Hash:    blk.Block.Hash,
		TxCount: int64(blk.Block.TxCount),
	}
	return
}

func (service *blockServiceImpl) GetBlockByNumber(blockNumber int) (response entity.Block, err error) {
	blk, err := service.blockGrpcClient.GetBlock(service.ctx, &pb.BlockRequest{
		BlockNumber: int64(blockNumber),
	})
	if err != nil {
		return
	}
	response = entity.Block{
		Number:  blk.Block.Number,
		Hash:    blk.Block.Hash,
		TxCount: int64(blk.Block.TxCount),
	}
	return
}
