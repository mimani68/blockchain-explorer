package service

import (
	"context"

	"app.io/data/entity"
	pb "app.io/service/proto_service"
)

func NewTransactionService(transactionGrpcClient pb.TransactionServiceClient, ctx context.Context) TransactionService {
	return &transactionServiceImpl{
		ctx:                   ctx,
		transactionGrpcClient: transactionGrpcClient,
	}
}

type transactionServiceImpl struct {
	ctx                   context.Context
	transactionGrpcClient pb.TransactionServiceClient
}

func (service *transactionServiceImpl) LastTransaction() (response entity.Transaction, err error) {
	trx, err := service.transactionGrpcClient.GetTransaction(service.ctx, &pb.TransactionRequest{})
	if err != nil {
		return
	}
	response = entity.Transaction{
		BlockNumber: trx.Transaction.BlockNumber,
		Hash:        trx.Transaction.Hash,
		Nonce:       trx.Transaction.Nonce,
		Amount:      trx.Transaction.Amount,
		From:        trx.Transaction.From,
		To:          trx.Transaction.To,
	}
	return
}

func (service *transactionServiceImpl) GettransactionByHash(hash string) (response entity.Transaction, err error) {
	trx, err := service.transactionGrpcClient.GetTransaction(service.ctx, &pb.TransactionRequest{
		TransactionHash: hash,
	})
	if err != nil {
		return
	}
	response = entity.Transaction{
		BlockNumber: trx.Transaction.BlockNumber,
		Hash:        trx.Transaction.Hash,
		Nonce:       trx.Transaction.Nonce,
		Amount:      trx.Transaction.Amount,
		From:        trx.Transaction.From,
		To:          trx.Transaction.To,
	}
	return
}
