package service

import (
	"context"

	"app.io/internal/data/domain"
	proto_service "app.io/internal/transport/grpc/proto_service"
)

func NewTransactionService(proto_service.UnimplementedTransactionServiceServer) proto_service.TransactionServiceServer {
	return &TransactionService{}
}

type TransactionService struct {
	proto_service.UnimplementedTransactionServiceServer
}

func (blockSrv *TransactionService) GetTransaction(ctx context.Context, request *proto_service.TransactionRequest) (*proto_service.TransactionResponse, error) {
	var transaction *domain.Transaction
	var err error
	if request.TransactionHash == "" {
		transaction, err = trxRepo.GetLastTransaction()
	} else {
		transaction, err = trxRepo.GetTransactionByHash(request.TransactionHash)
	}
	if err != nil || transaction.Hash == "" {
		return &proto_service.TransactionResponse{
			Success: false,
			Message: "Unable to get transaction",
		}, nil
	}

	return &proto_service.TransactionResponse{
		Success: true,
		Transaction: &proto_service.Transaction{
			BlockNumber: transaction.BlockNumber,
			Hash:        transaction.Hash,
			Amount:      transaction.Amount,
			Nonce:       transaction.Nonce,
			From:        transaction.From,
			To:          transaction.To,
		},
	}, nil
}
