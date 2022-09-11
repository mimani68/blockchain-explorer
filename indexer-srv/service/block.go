package service

import (
	"context"

	"app.io/data/entity"
	pb "app.io/service/proto_service"
)

func NewBlockService(blockGrpcClient pb.BlockServiceClient, scanGrpcClient pb.ScanServiceClient, ctx context.Context) BlockService {
	return &blockServiceImpl{
		ctx:             ctx,
		blockGrpcClient: blockGrpcClient,
		scanGrpcClient:  scanGrpcClient,
	}
}

type blockServiceImpl struct {
	ctx             context.Context
	blockGrpcClient pb.BlockServiceClient
	scanGrpcClient  pb.ScanServiceClient
}

func (service *blockServiceImpl) ForceScan(startBlock, endBlock int) (response []entity.Block, err error) {
	blockList, err := service.scanGrpcClient.FreshScan(service.ctx, &pb.ScanRequest{
		StartBlock: int64(startBlock),
		EndBlock:   int64(endBlock),
	})
	if err != nil {
		return
	}
	for _, blk := range blockList.Block {
		response = append(response, entity.Block{
			Number:  blk.Number,
			Hash:    blk.Hash,
			TxCount: int64(blk.TxCount),
		})
	}
	return
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
