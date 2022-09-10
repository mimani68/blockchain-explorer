package service

import (
	"context"

	"app.io/internal/data/repository"
	proto_service "app.io/internal/transport/grpc/proto_service"
)

func NewBlockService(blockRepo *repository.BlockRepository, proto_service proto_service.UnimplementedBlockServiceServer) proto_service.BlockServiceServer {
	// return &BlockService{
	// 	protobuf_service: proto_service,
	// 	blockRepo:        blockRepo,
	// }
	return &BlockService{}
}

type BlockService struct {
	proto_service.UnimplementedBlockServiceServer
	// blockRepo        *repository.BlockRepository
}

func (blockSrv *BlockService) StoreBlockation(ctx context.Context, input *proto_service.BlockRequest) (*proto_service.BlockResponse, error) {
	return &proto_service.BlockResponse{}, nil
}

func (blockSrv *BlockService) Estimate(ctx context.Context, input *proto_service.BlockRequest) (*proto_service.EstimateResponse, error) {
	return &proto_service.EstimateResponse{}, nil
}
