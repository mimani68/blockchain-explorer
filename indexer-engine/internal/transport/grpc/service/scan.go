package service

import (
	"context"

	proto_service "app.io/internal/transport/grpc/proto_service"
)

func NewScanService(proto_service.UnimplementedScanServiceServer) proto_service.ScanServiceServer {
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
	return &proto_service.ScanResponse{
		Success: true,
		Message: "forceing fresh scan",
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
