package service

import (
	"context"

	proto_service "app.io/internal/transport/grpc/proto_service"
)

func NewTransactionService(proto_service.UnimplementedTransactionServiceServer) proto_service.TransactionServiceServer {
	return &TransactionService{}
}

type TransactionService struct {
	proto_service.UnimplementedTransactionServiceServer
}

func (blockSrv *TransactionService) GetTransaction(ctx context.Context, request *proto_service.TransactionRequest) (*proto_service.TransactionResponse, error) {
	if len(request.TransactionHash) <= 0 {
		// get last transaction of db
		request.TransactionHash = "0x7fg376t65f"
	}
	return &proto_service.TransactionResponse{
		Success: true,
		Transaction: &proto_service.Transaction{
			BlockNumber: 2135451,
			Hash:        "0x7fg376t65f",
			Amount:      1,
			Nonce:       4,
			From:        "0xyyyyy",
			To:          "0xerg45jg945t",
		},
	}, nil
}
