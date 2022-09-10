package service

import (
	"context"

	"app.io/internal/data/repository"
	proto_service "app.io/internal/transport/grpc/proto_service"
)

func NewBlockService(blockRepo *repository.BlockRepository, proto_service proto_service.UnimplementedBlockServiceServer) proto_service.BlockServiceServer {
	return &BlockService{}
}

type BlockService struct {
	proto_service.UnimplementedBlockServiceServer
}

func (blockSrv *BlockService) GetBlock(ctx context.Context, blockNumber *proto_service.BlockRequest) (*proto_service.BlockResponse, error) {
	if blockNumber.BlockNumber <= 0 {
		blockNumber.BlockNumber = 15507757
	}
	return &proto_service.BlockResponse{
		Success: true,
		Block: &proto_service.Block{
			Number:  blockNumber.BlockNumber,
			Hash:    "934y34ty85",
			TxCount: 14,
			Transactions: []*proto_service.Transaction{
				{
					BlockNumber: blockNumber.BlockNumber,
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
